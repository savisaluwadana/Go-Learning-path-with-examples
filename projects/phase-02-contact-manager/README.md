# Phase 02 Mini Project: Contact Manager CLI

An in-memory CRUD contact manager using structs and methods.

## Run

```bash
go run ./projects/phase-02-contact-manager
```

## Commands

- `add <name> <email> <phone>`
- `list`
- `update <id> <name> <email> <phone>`
- `delete <id>`
- `help`
- `exit`

## Example Session

```text
> add Alice alice@example.com +94770000000
> add Bob bob@example.com +94771111111
> list
> update 1 Alicia alicia@example.com +94772222222
> delete 2
> exit
```

## Test

```bash
go test ./projects/phase-02-contact-manager/...
```

