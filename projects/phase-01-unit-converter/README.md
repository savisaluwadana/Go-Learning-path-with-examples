# Phase 01 Mini Project: Unit Converter CLI

Convert values across temperature, distance, and weight units.

## Run

```bash
go run ./projects/phase-01-unit-converter -kind temperature -from c -to f -value 37
go run ./projects/phase-01-unit-converter -kind distance -from km -to mi -value 5
go run ./projects/phase-01-unit-converter -kind weight -from kg -to lb -value 2
```

## Supported Units

- `temperature`: `c`, `f`, `k`
- `distance`: `m`, `km`, `mi`
- `weight`: `g`, `kg`, `lb`

## Test

```bash
go test ./projects/phase-01-unit-converter/...
```

