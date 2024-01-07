import sys
from typing import cast

from aoc.utils import cli


def main():
    args = sys.argv
    options = cli.parse_args(args[1:])

    DAY = options["day"]
    if DAY is None:
        raise EnvironmentError("Env var `$DAY` is not set.")
    DAY = int(DAY)
    print(f"Running for day {DAY}\n")

    from importlib import import_module

    module = import_module(f"aoc.day{DAY:02}.main")

    part_1_output, part_2_output = cast(tuple[str, str], module.main())

    print("\nPart 1 Output:\n")
    print(part_1_output)


    print("\nPart 2 Output:\n")
    print(part_2_output)

if __name__ == "__main__":
    main()
