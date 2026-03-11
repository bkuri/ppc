# PRD: PPC Variable Substitution for Goal Integration

**Project:** MaksiTrader Prompt Policy Compiler Enhancement  
**Priority:** High  
**Status:** Draft  
**Created:** 2026-03-10  
**Target Completion:** 1-2 sessions  

---

## Executive Summary

Add Jinja2-style variable substitution to the Prompt Policy Compiler (PPC) to enable dynamic prompt compilation from user-defined goals stored in configuration. This allows agents to understand and optimize toward user-specific aspirations rather than just constraints.

---

## Problem Statement

### Current State
- PPC compiles static prompts from modular `.md` files
- Agents know **constraints** (max drawdown, position limits) from config
- Agents do NOT know **aspirations** (target returns, time horizon, objectives)
- Goals exist as user intent, not as agent-consumable data

### Pain Point
Agent C (Controller) can validate "is this safe?" but cannot evaluate "does this help achieve 15% annual return?" or "is this aligned with a 2-year horizon?"

### Desired State
- User defines goals via `maxitrader settings wizard --goals`
- Goals stored in `UserConfig.goals`
- PPC substitutes `{{variable}}` placeholders with goal values during compilation
- Agents receive personalized prompts with explicit outcome targets

---

## Solution Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    User Workflow                             │
├─────────────────────────────────────────────────────────────┤
│  maxitrader settings wizard --goals                          │
│  → "Target annual return? [15%]: 20"                         │
│  → "Time horizon? [12 months]: 24"                           │
│  → Saved to UserConfig.goals                                 │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│                 Prompt Compilation                           │
├─────────────────────────────────────────────────────────────┤
│  prompts/policies/outcome_targets.md:                        │
│    "Target return: {{goals.target_return}}%"                 │
│    "Horizon: {{goals.time_horizon}} months"                  │
│                                                              │
│  PromptCompiler.compile():                                   │
│    1. Load UserConfig.goals                                  │
│    2. Substitute {{goals.*}} with values                     │
│    3. Cache with config hash                                 │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│              Compiled Agent Prompt                           │
├─────────────────────────────────────────────────────────────┤
│  "Target return: 20%"                                        │
│  "Horizon: 24 months"                                        │
│  "When evaluating trades, optimize for 20% annualized..."   │
└─────────────────────────────────────────────────────────────┘
```

---

## Technical Specifications

### 1. Variable Syntax

Use Jinja2-compatible syntax for familiarity:

| Syntax | Description | Example |
|--------|-------------|---------|
| `{{variable}}` | Simple substitution | `{{goals.target_return}}` |
| `{{object.property}}` | Nested access | `{{user.risk_style}}` |
| `{{var \| default}}` | Default if None | `{{goals.horizon \| 12}}` |
| `{% if var %}...{% endif %}` | Conditional blocks | See below |

**Minimal implementation:** Only `{{variable}}` and `{{object.property}}` required initially. Conditionals can be Phase 2.

### 2. Available Variables

Variables are loaded from merged config sources:

```python
VARIABLE_SOURCES = {
    "goals": "UserConfig.goals",        # User-defined aspirations
    "user": "UserConfig",               # User settings (risk_style, etc.)
    "safety": "SafetyConfig",           # Safety thresholds
    "paper": "PaperConfig",             # Paper trading settings
}
```

**Example access:**
```
{{goals.target_return}}      → 15.0
{{goals.primary_objective}}  → "growth"
{{user.risk_style}}          → "balanced"
{{safety.max_drawdown}}      → 0.20
{{paper.starting_balance}}   → 10000.0
```

### 3. GoalsConfig Dataclass

New dataclass in `maxitrader/core/config/user.py`:

```python
@dataclass
class GoalsConfig:
    """User-defined trading goals and aspirations."""
    
    primary_objective: str = "growth"
    # Options: "growth" (maximize returns), "income" (steady yield), 
    #          "preservation" (minimize losses), "balanced" (risk-adjusted)
    
    target_annual_return: float = 15.0
    # Target annualized return percentage (e.g., 15.0 = 15%)
    
    time_horizon_months: int = 12
    # Investment horizon in months
    
    max_acceptable_drawdown: Optional[float] = None
    # Override safety.max_drawdown if set (as decimal, e.g., 0.20)
    
    min_acceptable_sharpe: Optional[float] = None
    # Override safety thresholds if set
    
    target_win_rate: Optional[float] = None
    # Target win rate percentage (e.g., 55.0)
    
    priority_metric: str = "sharpe"
    # Options: "sharpe", "return", "win_rate", "sortino"
    # Which metric to prioritize when making trade-offs
    
    custom_notes: str = ""
    # Free-form notes about user's goals/context
```

### 4. PromptCompiler Modifications

**File:** `maxitrader/chat/build_agent_prompt.py`

**New method:**
```python
def _substitute_variables(self, content: str, config: dict) -> str:
    """
    Substitute {{variable}} placeholders with config values.
    
    Args:
        content: Raw module content with placeholders
        config: Dictionary of variable sources (goals, user, safety, etc.)
    
    Returns:
        Content with placeholders replaced by values
    
    Example:
        "{{goals.target_return}}%" → "15.0%"
    """
```

**Modified `compile()` method:**
```python
def compile(self, agent_type, risk_style, skip_cache=False, validate_only=False) -> str:
    # ... existing code ...
    
    # NEW: Load variable sources
    var_sources = self._load_variable_sources()
    
    # NEW: Substitute variables in each module
    for module_id in resolved:
        module = self._load_module(module_id)
        if module:
            # Substitute variables before appending
            substituted_content = self._substitute_variables(
                module["content"], 
                var_sources
            )
            prompt_parts.append(substituted_content)
    
    # ... rest of compilation ...
```

**Cache key must include config hash:**
```python
def _get_cache_key(self, agent_type, risk_style) -> str:
    key_data = f"{agent_type.value}:{risk_style.value}"
    
    # Include module content (existing)
    for module_id in modules:
        module = self._load_module(module_id)
        if module:
            key_data += f":{module['content']}"
    
    # NEW: Include config hash for variable invalidation
    config_hash = self._get_config_hash()
    key_data += f":{config_hash}"
    
    return hashlib.sha256(key_data.encode()).hexdigest()
```

### 5. Settings Wizard Integration

**File:** `maxitrader/cli/commands/settings.py`

**New function:**
```python
def _wizard_goals(config: UserConfig) -> None:
    """Wizard section for goals configuration."""
    console.print("\n[bold]Trading Goals Configuration[/bold]")
    console.print("[dim]These goals inform agent decision-making[/dim]\n")
    
    # Primary objective
    console.print("\nObjectives:")
    console.print("  growth       - Maximize returns (accept higher volatility)")
    console.print("  income       - Generate steady yield (lower volatility)")
    console.print("  preservation - Minimize losses (conservative)")
    console.print("  balanced     - Risk-adjusted returns (moderate)")
    
    objective = Prompt.ask(
        "Primary objective",
        choices=["growth", "income", "preservation", "balanced"],
        default=config.goals.primary_objective
    )
    config.goals.primary_objective = objective
    
    # Target return
    target_return = Prompt.ask(
        "Target annual return %",
        default=str(config.goals.target_annual_return)
    )
    try:
        config.goals.target_annual_return = float(target_return)
    except ValueError:
        console.print("[yellow]Invalid number, using default[/yellow]")
    
    # Time horizon
    horizon = Prompt.ask(
        "Time horizon (months)",
        default=str(config.goals.time_horizon_months)
    )
    try:
        config.goals.time_horizon_months = int(horizon)
    except ValueError:
        console.print("[yellow]Invalid number, using default[/yellow]")
    
    # Priority metric
    console.print("\nPriority metrics (used for trade-off decisions):")
    console.print("  sharpe   - Risk-adjusted returns")
    console.print("  return   - Absolute returns")
    console.print("  win_rate - Percentage of profitable trades")
    console.print("  sortino  - Downside-risk adjusted returns")
    
    priority = Prompt.ask(
        "Priority metric",
        choices=["sharpe", "return", "win_rate", "sortino"],
        default=config.goals.priority_metric
    )
    config.goals.priority_metric = priority
```

**Wire into main wizard:**
```python
def settings_wizard(args: argparse.Namespace) -> None:
    # ... existing sections ...
    
    if section_all or section_goals:  # NEW
        _wizard_goals(config)
```

---

## Implementation Checklist

### Phase 1: Core Substitution (Required)

- [ ] Add `GoalsConfig` dataclass to `maxitrader/core/config/user.py`
- [ ] Add `goals: GoalsConfig` field to `UserConfig` dataclass
- [ ] Implement `_substitute_variables()` in `PromptCompiler`
- [ ] Modify `compile()` to call substitution on each module
- [ ] Update `_get_cache_key()` to include config hash
- [ ] Add `_load_variable_sources()` to gather config for substitution
- [ ] Write unit tests for variable substitution

### Phase 2: Wizard Integration (Required)

- [ ] Add `_wizard_goals()` function to `settings.py`
- [ ] Wire `--goals` flag into wizard argument parser
- [ ] Add `section_goals` to wizard flow
- [ ] Ensure goals persist via `save_user_config()`
- [ ] Test wizard → config → compilation flow

### Phase 3: Goal Prompt Module (Required)

- [ ] Create `prompts/policies/outcome_targets.md` with `{{goals.*}}` placeholders
- [ ] Add to `default_policies` in `prompts/rules.yml`
- [ ] Remove duplicate content from `prompts/base.md` (streamline)
- [ ] Test compiled prompt contains substituted values

### Phase 4: Edge Cases (Required)

- [ ] Handle missing/None values gracefully (use defaults)
- [ ] Handle type conversion (float → string with formatting)
- [ ] Handle missing config file (use default GoalsConfig)
- [ ] Handle circular references (should not happen, but validate)
- [ ] Log warnings for unresolved variables (don't fail compilation)

### Phase 5: Optional Enhancements (Future)

- [ ] Support `{% if %}` conditional blocks
- [ ] Support filters like `{{value | round(2)}}`
- [ ] Support formatting like `{{value | percent}}` → "15%"
- [ ] Add `maxitrader goals show` command
- [ ] Add `maxitrader goals edit` command

---

## Testing Requirements

### Unit Tests

**File:** `tests/test_ppc_variable_substitution.py`

```python
def test_simple_variable_substitution():
    """{{goals.target_return}} is replaced with value"""
    compiler = PromptCompiler()
    compiler.config = {"goals": {"target_return": 15.0}}
    result = compiler._substitute_variables(
        "Target: {{goals.target_return}}%",
        compiler.config
    )
    assert result == "Target: 15.0%"

def test_nested_variable_substitution():
    """{{user.risk_style}} accesses nested property"""
    compiler = PromptCompiler()
    compiler.config = {"user": {"risk_style": "balanced"}}
    result = compiler._substitute_variables(
        "Risk: {{user.risk_style}}",
        compiler.config
    )
    assert result == "Risk: balanced"

def test_missing_variable_uses_placeholder():
    """Unresolved variables show warning and keep placeholder"""
    compiler = PromptCompiler()
    compiler.config = {}
    result = compiler._substitute_variables(
        "Value: {{unknown.variable}}",
        compiler.config
    )
    # Should log warning but not crash
    assert "{{unknown.variable}}" in result or result == "Value: "

def test_cache_invalidation_on_config_change():
    """Changing config invalidates prompt cache"""
    compiler = PromptCompiler()
    
    # Compile with initial config
    key1 = compiler._get_cache_key(AgentType.A, RiskStyle.BALANCED)
    
    # Change config
    compiler._load_variable_source()  # Reload
    
    # Cache key should be different
    key2 = compiler._get_cache_key(AgentType.A, RiskStyle.BALANCED)
    assert key1 != key2

def test_full_compilation_with_variables():
    """End-to-end: wizard → config → compilation → substituted prompt"""
    # This test would require more setup, mock config, etc.
    pass
```

### Integration Tests

```python
def test_goals_wizard_persists():
    """Running wizard with --goals saves to config"""
    # Mock user input, run wizard, verify config saved

def test_compiled_prompt_contains_goals():
    """Compiled prompt for Agent C contains goal values"""
    prompt = build_agent_prompt("C", "balanced")
    # Assuming goals.target_return = 20
    assert "20" in prompt or "20.0" in prompt
```

---

## Example: outcome_targets.md

```markdown
---
id: policies/outcome_targets
desc: User-defined outcome targets and aspirations.
priority: 2
tags: [policy:outcomes]
requires: [base]
---

## Outcome Targets

You are optimizing for the following user-defined outcomes:

### Primary Objective: {{goals.primary_objective}}

{% if goals.primary_objective == "growth" %}
Maximize risk-adjusted returns while respecting safety constraints.
Accept higher volatility in pursuit of returns.
{% elif goals.primary_objective == "income" %}
Generate consistent yield with minimal drawdowns.
Prioritize stability over maximum returns.
{% elif goals.primary_objective == "preservation" %}
Protect capital above all else. Minimize losses.
Only recommend trades with very high confidence.
{% else %} // balanced
Optimize risk-adjusted returns with moderate volatility.
Balance growth potential with downside protection.
{% endif %}

### Target Metrics

| Metric | Target | Notes |
|--------|--------|-------|
| Annual Return | {{goals.target_annual_return}}% | Primary success criterion |
| Time Horizon | {{goals.time_horizon_months}} months | Duration for achieving targets |
| Priority Metric | {{goals.priority_metric}} | Optimize this when trade-offs required |

{% if goals.target_win_rate %}
| Win Rate | {{goals.target_win_rate}}% | Minimum acceptable |
{% endif %}

### Decision Framework

When evaluating trades or strategies:

1. **Does this align with {{goals.primary_objective}} objective?**
   - Growth: Favor higher return potential
   - Income: Favor consistent, lower-volatility returns
   - Preservation: Favor capital protection

2. **Does this contribute to {{goals.target_annual_return}}% annual target?**
   - Calculate expected contribution
   - Consider time horizon ({{goals.time_horizon_months}} months)

3. **When trade-offs exist, prioritize {{goals.priority_metric}}**
   - If Sharpe vs Return: Choose based on priority_metric
   - If Win Rate vs Return: Choose based on priority_metric

4. **Respect hard constraints from safety_rules.md**
   - Goals inform optimization direction
   - Safety rules are non-negotiable limits

### Communication Style

When discussing results with the user:
- Frame performance relative to their {{goals.target_annual_return}}% target
- Note progress toward goals over {{goals.time_horizon_months}} month horizon
- Highlight alignment with {{goals.primary_objective}} objective
- If goals seem unrealistic, suggest adjustments with reasoning
```

---

## Acceptance Criteria

1. ✅ User can run `maxitrader settings wizard --goals` and set targets
2. ✅ Goals persist in `~/.config/maxitrader/config.json`
3. ✅ Compiled prompts contain substituted goal values
4. ✅ Cache invalidates when goals change
5. ✅ Agent prompts reference specific user targets (not generic values)
6. ✅ All existing tests pass
7. ✅ New tests cover substitution logic
8. ✅ Missing variables don't crash compilation (log warning)

---

## Dependencies

- Existing: `maxitrader/chat/build_agent_prompt.py` (PromptCompiler)
- Existing: `maxitrader/core/config/user.py` (UserConfig)
- Existing: `maxitrader/cli/commands/settings.py` (wizard)
- New: `prompts/policies/outcome_targets.md`

---

## Risks & Mitigations

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Jinja2 syntax too complex | Low | Medium | Start with simple `{{var}}` only |
| Cache invalidation issues | Medium | High | Include full config hash in key |
| Missing variables break prompts | Medium | High | Graceful fallback + logging |
| Config file format changes | Low | Medium | Version config schema |

---

## Questions for Implementation

1. **Variable source priority:** If same key exists in multiple sources (e.g., `safety.max_drawdown` vs `goals.max_acceptable_drawdown`), which wins?
   - Recommendation: Goals override safety (user intent > defaults)

2. **Error handling:** What happens if `{{goals.unknown}}` is used?
   - Recommendation: Log warning, leave placeholder or use empty string

3. **Type formatting:** Should `{{goals.target_return}}` automatically add "%" suffix?
   - Recommendation: No, let prompt author control formatting with `{{goals.target_return}}%`

---

## File Locations

| File | Purpose |
|------|---------|
| `maxitrader/chat/build_agent_prompt.py` | Add substitution logic |
| `maxitrader/core/config/user.py` | Add GoalsConfig |
| `maxitrader/cli/commands/settings.py` | Add wizard section |
| `prompts/policies/outcome_targets.md` | Goal template module |
| `prompts/rules.yml` | Add to default_policies |
| `tests/test_ppc_variable_substitution.py` | Unit tests |

---

**End of PRD**
