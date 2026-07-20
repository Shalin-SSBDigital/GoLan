# 🛑 STOP — Read This First

**Every AI agent touching this project MUST read this file and the master [[Second Brain AI_CONTEXT.md]] before creating, editing, or moving any documentation.**

This project uses an **external Second Brain vault** as its single source of truth for all documentation. **Nothing stays in this repo except code and this file.**

---

## Quick Reference

| Resource | Path |
|---|---|
| Second Brain vault root | `G:\My Drive\Obsi\Second-Brain` (Windows) / `/Users/shalinshah/Library/CloudStorage/GoogleDrive-2002shalin@gmail.com/My Drive/Obsi/Second-Brain` (macOS) |
| Master AI_CONTEXT.md | `00_System/AI_CONTEXT.md` in the vault |
| This project's docs | `01_Projects/go-lan/` in the vault |
| Project templates | `00_System/Templates/Project.md` in the vault |
| Tags reference | `00_System/Tags.md` in the vault |

---

## Agent Workflow

1. **READ** the master `00_System/AI_CONTEXT.md` in the Second Brain vault — it defines the full operating system for content placement, naming, tags, and templates
2. **READ** `01_Projects/go-lan/README.md` in the vault to understand project context
3. **FIND** existing docs in `01_Projects/go-lan/` before creating new ones
4. **CREATE** new doc files in the vault's `01_Projects/go-lan/` folder following the Project.md template
5. **UPDATE** `Tasks.md` in the vault when completing work
6. **LINK** new docs using `[[wikilinks]]` to related notes in the vault

---

## The Rule

> **Everything goes to the brain, nothing stays in the project folder.**
>
> - Code examples → this repo (`.go` files)
> - Explanations, architecture, tasks, notes, decisions, bugs → `01_Projects/go-lan/` in the Second Brain
> - Reusable Go snippets → `06_Resources/Snippets/Go/` in the Second Brain
> - Go knowledge articles → `02_Knowledge/Go/` in the Second Brain
>
> **The only `.md` file allowed in this repo is this one.**

---

## Project Context

This is **GoLan** — a personal Go learning repository. It contains runnable Go examples covering fundamentals: variables, constants, conditionals, loops, arrays, slices, functions, and more. Each concept has a dedicated folder with annotated source files.

- **Repo location:** `/Users/shalinshah/Developer-Shalin/GoLan`
- **Project docs:** [[01_Projects/go-lan/README]] in the Second Brain
- **Active tasks:** [[01_Projects/go-lan/Tasks]] in the Second Brain
- **Notes & gotchas:** [[01_Projects/go-lan/Notes]] in the Second Brain
