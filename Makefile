BIN=ppc
PLATFORMS=linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILDDIR=dist

all: build

build:
	go build -ldflags="-s -w" -o $(BIN) ./cmd/build-prompt

test:
	go test ./...

smoke: build
	./scripts/smoke.sh

clean:
	rm -f $(BIN)

# Extract OS/ARCH from platform
os = $(word 1,$(subst /, ,$@))
arch = $(word 2,$(subst /, ,$@))

# Build single platform binary (non-windows)
linux/amd64 linux/arm64 darwin/amd64 darwin/arm64:
	@echo "Building $@..."
	@mkdir -p $(BUILDDIR)
	GOOS=$(os) GOARCH=$(arch) go build \
		-ldflags="-s -w -X main.Version=$(VERSION)" \
		-o $(BUILDDIR)/ppc_$(VERSION)_$(subst /,_,$@) \
		./cmd/build-prompt

# Windows binary gets .exe extension
windows/amd64:
	@echo "Building $@..."
	@mkdir -p $(BUILDDIR)
	GOOS=windows GOARCH=amd64 go build \
		-ldflags="-s -w -X main.Version=$(VERSION)" \
		-o $(BUILDDIR)/ppc_$(VERSION)_windows_amd64.exe \
		./cmd/build-prompt

# Build all platforms
release-all: $(PLATFORMS)

# Create release archives
archive: release-all
	@echo "Creating archives..."
	@for platform in linux_amd64 linux_arm64 darwin_amd64 darwin_arm64; do \
		mkdir -p $(BUILDDIR)/$$platform; \
		cp $(BUILDDIR)/ppc_$(VERSION)_$$platform $(BUILDDIR)/$$platform/ppc; \
		chmod +x $(BUILDDIR)/$$platform/ppc; \
		cp LICENSE $(BUILDDIR)/$$platform/; \
		cp README.md $(BUILDDIR)/$$platform/; \
		tar -czf $(BUILDDIR)/ppc_$(VERSION)_$$platform.tar.gz -C $(BUILDDIR) $$platform; \
	done
	@mkdir -p $(BUILDDIR)/windows_amd64; \
		cp $(BUILDDIR)/ppc_$(VERSION)_windows_amd64.exe $(BUILDDIR)/windows_amd64/ppc.exe; \
		cp LICENSE $(BUILDDIR)/windows_amd64/; \
		cp README.md $(BUILDDIR)/windows_amd64/; \
		tar -czf $(BUILDDIR)/ppc_$(VERSION)_windows_amd64.tar.gz -C $(BUILDDIR) windows_amd64
	@echo "Archives created in $(BUILDDIR)/"

# Generate checksums
checksums: archive
	@./scripts/generate-checksums.sh

# Clean build artifacts
clean-release:
	rm -rf $(BUILDDIR)

# Clean everything
clean-all: clean clean-release
	rm -f $(BUILDDIR)/*.tar.gz

.PHONY: all build test smoke clean release-all clean-release clean-all checksums archive
