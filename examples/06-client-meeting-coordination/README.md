# Example 06: Client Meeting Coordination

## Problem

Meetings are chaotic. Decisions get lost, ownership is unclear, and follow-up work is inconsistent. When multiple people (humans and agents) participate without structure, the outcome is often:

- Missed decisions that have to be rediscovered later
- Ambiguous ownership ("Who's supposed to do that?")
- Inconsistent meeting notes that can't be reused as context
- Different teams have different meeting styles, making onboarding hard

Ad-hoc notes taken after the fact are a poor solution. By then, the nuance of what was decided—and why—is already fading. People remember different things, and the conversation becomes a source of confusion rather than clarity.

## Why this exists

A policy can orchestrate a meeting structure deterministically. When the meeting rules are captured as markdown policies:

- Multiple actors (humans + agents) can coordinate via shared rules, not custom code
- Decisions are captured automatically during the meeting, not manually after
- Notes become the source of truth and context for future work
- The same policy structure can be reused for different meeting types

This example shows that meeting discipline is just applied policy. You don't need custom software to orchestrate human-agent teams—you need clear rules that everyone follows.

## What this example demonstrates

This example introduces several new concepts:

- **Role-based expectations:** Clear responsibilities for the client, coordinator, and four specialized agents
- **Phase discipline:** A linear, deterministic flow (refinement → review → Q&A → decision → assignments) that ensures nothing is missed
- **Response format rules:** Every team response must include Perspective | Concern | Recommendation, ensuring substantive feedback
- **Markdown contract:** The meeting notes structure is defined as a contract, guaranteeing consistent output
- **Variable substitution:** Project name, date, participants, and number of topics are injected at compile time
- **Expandability:** The same policy structure works for retrospectives, design reviews, and other meeting types (see `profiles/retrospective.yml` and `profiles/design-review.yml` for templates)

## How to use this example

This example is a reference implementation. To use it for your organization:

1. **Copy modules to your prompts directory:**
   ```bash
   cp -r examples/06-client-meeting-coordination/prompts/* your-project/prompts/
   ```

2. **Create a profile in your profiles directory:**
   ```bash
   mkdir -p your-project/profiles
   cp examples/06-client-meeting-coordination/profiles/kickoff-meeting.yml your-project/profiles/
   ```

3. **Update the profile to reference your modules:**
   ```yaml
   # your-project/profiles/kickoff-meeting.yml
   mode: kickoff
   contract: meeting-notes
   requires:
     - base
     - kickoff
     - role-expectations
     - phase-discipline
     - response-format
     - note-structure
   vars:
     project_name: "Your Project Name"
     meeting_date: "2026-01-23"
   ```

4. **Compile the prompt:**
   ```bash
   cd your-project
   ppc build --profile kickoff-meeting --prompts prompts
   ```

The profiles in this example (`retrospective.yml`, `design-review.yml`) are templates showing how to adapt the policy structure for other meeting types.

## Output

When you compile `kickoff-meeting.yml`, you get a comprehensive meeting notes document:

```markdown
# Kickoff Meeting: SaaS Dashboard Project
**Date:** 2026-01-23
**Attendees:** Client Stakeholder, Project Coordinator, Requirements Clarifier, Technical Advisor, Design Advisor, Meeting Summarizer
**Meeting Duration:** 45 minutes
**Topics:** 3

## Topic 1: Overall Concept

### Phase 1: Requirements Refinement
**Client Initial Requirement:**
> "We need a dashboard for managing team projects. It should show project lists, task details, and team analytics. Updates need to be real-time."

**Refined Requirement:**
After Q&A with the Project Coordinator, the requirement is clarified as:
A web-based dashboard for managing internal team projects. It must support three core views: project list, task detail, and team analytics. Real-time updates should reflect changes within 5 seconds, not 1 second. The dashboard is for internal use by 10-20 team members.

### Phase 2: Team Review
**Requirements Clarifier:**
- Perspective: I understand you want a dashboard for managing projects and tasks with real-time updates.
- Concern: What's the expected data volume? How many projects, tasks, and users will this support at launch vs. 6 months out?
- Recommendation: Clarify the scale so we can design the database and caching strategy appropriately.

[... rest of the meeting notes ...]
```

The output is a single markdown file that captures:

- What was discussed in each topic
- What the team flagged and recommended
- What was decided and why
- What was deferred and what risks were accepted
- Next steps with clear owners

## What to copy

To adapt this example for your organization:

1. **Copy the module structure** from `prompts/` to your project's `prompts/` directory.

2. **Adapt `base.md`** to your organization's context (e.g., change "You are a meeting facilitator system" to match your domain).

3. **Copy and adapt `prompts/policies/role-expectations.md`** to match your team structure. Add roles like "DevOps Advisor" or remove roles you don't need.

4. **Copy and adapt `prompts/policies/phase-discipline.md`** to match your meeting type. For example, retrospectives might have phases like "What went well" → "What didn't" → "Action items."

5. **Create a profile** in your `profiles/` directory that references your modules and sets configuration variables.

The policies are modular. You can swap out `prompts/policies/role-expectations.md` for different team structures without changing the rest of the system. You can swap out `prompts/policies/phase-discipline.md` for different meeting types without changing the role expectations.

## Adapt for your meeting type

This example is built around a kickoff meeting, but the policy structure is adaptable to other meeting types:

### Retrospective

To adapt this for retrospectives:

1. Copy all modules to your prompts directory.
2. Edit `prompts/policies/role-expectations.md` to replace roles with retrospectives-specific roles:
   - "Client Stakeholder" → "Project Lead"
   - "Project Coordinator" → "Facilitator"
   - "Technical Advisor" → "Process Reviewer"
   - "Design Advisor" → "Team Member"
   - Remove: Requirements Clarifier
3. Edit `phase-discipline.md` to replace phases with retrospectives-specific phases:
   - Phase 1: "What went well"
   - Phase 2: "What didn't go well"
   - Phase 3: "What should we do differently"
   - Phase 4: "Action items"

See `profiles/retrospective.yml` for a template of what the profile might look like.

### Design Review

To adapt this for design reviews:

1. Copy all modules to your prompts directory.
2. Edit `prompts/policies/role-expectations.md` to replace roles with design review-specific roles:
   - "Client Stakeholder" → "Product Manager"
   - "Project Coordinator" → "Design Lead"
   - "Technical Advisor" → "Accessibility Reviewer"
   - "Design Advisor" → "UX Designer"
   - Remove: Requirements Clarifier
3. Edit `phase-discipline.md` to replace phases with design review-specific phases:
   - Phase 1: "Design presentation"
   - Phase 2: "Team feedback"
   - Phase 3: "Clarifying questions"
   - Phase 4: "Design decision"

See `profiles/design-review.yml` for a template of what the profile might look like.

The meeting notes structure (`note-structure.md` and `meeting-notes.md`) remains the same—only the roles and phase definitions change.

## Common failure modes

**Skipping a phase:** If you skip the "Followup Questions" phase, you'll miss key requirements that surface during team feedback. The meeting notes will be incomplete because the team's concerns weren't addressed.

**Mixing roles:** If the Client Stakeholder acts like an agent (e.g., suggesting technical implementations), it creates confusion about who's deciding vs. who's advising. The policy prevents this by enforcing role separation.

**Summarizer not recording:** If the Meeting Summarizer doesn't capture decisions with full rationale and tradeoffs, the meeting notes become useless. The meeting becomes a conversation without an artifact.

**Not capturing dissent:** If you only record consensus and not dissent, you lose the nuance of decision-making. Later, when someone asks "Did everyone agree?", you won't have an answer. The policy requires recording dissent and why it was overruled.

## Key takeaway

**Policies orchestrate workflows. Meetings are just an applied workflow.**

This example shows that you don't need complex software to coordinate human-agent teams. You need clear policies that define:

- Who plays what role
- When each phase happens
- What format every response must follow
- What the output must look like

When these policies are captured as markdown, they become inspectable, versionable, and reusable across teams and meeting types. A deterministic meeting is just a compiled prompt policy.

## Validation

To validate the module structure:

```bash
ppc doctor --prompts prompts
```

This checks that all modules are well-formed, dependencies are valid, and there are no circular dependencies.
