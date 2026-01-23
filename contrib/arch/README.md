# Arch Linux Installation

## Building from PKGBUILD

1. Clone this repo:
   ```bash
   git clone https://github.com/bkuri/ppc.git
   cd ppc
   ```

2. Build package:
   ```bash
   cd contrib/arch
   makepkg -si
   ```

3. Verify installation:
   ```bash
   ppc --version
   ```

## Submitting to AUR (Future)

Once ready for AUR:

1. Create AUR package repo:
   ```bash
   ssh aur@aur.archlinux.org setup ppc
   ```

2. Push PKGBUILD:
   ```bash
   git clone ssh://aur@aur.archlinux.org/ppc.git
   cp contrib/arch/PKGBUILD ppc/
   cd ppc
   git add PKGBUILD
   git commit -m "Update to v0.2.0"
   git push
   ```

3. Users install with:
   ```bash
   yay -S ppc
   # or
   paru -S ppc
   ```
