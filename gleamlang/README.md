# Advent of Code - Gleam ⭐

1. set the AOC_COOKIE environment variable with your Advent of Code session cookie
2. to work a solution for a new day (or set of days) run the following command:

```bash
gleam run new --fetch --example X Y Z ...
```

3. add your input to `input/<YEAR>/X.txt`
4. add your code to `src/aoc_<YEAR>/day_X.gleam`
5. to run the solutions of a specific day (or set of days) run the following command:

```bash
gleam run run X Y Z ... --example
```

Where X Y Z are day number, e.g. 1, 2, ..., 25
and the `--example` flag indicates to run the example input instead of the real input.
