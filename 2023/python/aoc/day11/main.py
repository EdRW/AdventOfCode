from collections import Counter, deque
from collections.abc import Iterator
from dataclasses import dataclass, InitVar, field
from enum import IntEnum, StrEnum, Enum
from math import lcm
import os
import subprocess
import sys
import time
from typing import (
    Callable,
    Iterable,
    Literal,
    NamedTuple,
    Self,
    Sequence,
    TypedDict,
    cast,
)

from aoc.utils import cli


class Direction(Enum):
    N = 1
    S = 2
    E = 3
    W = 4


class Galaxy(NamedTuple):
    row: int
    col: int


class Size(NamedTuple):
    rows: int
    cols: int


@dataclass
class Universe:
    size: Size
    galaxies: list[Galaxy]
    rows_to_expand: list[int]
    cols_to_expand: list[int]
    image_data: list[str] | None = None
    expansion_rate: int = 0

    def __post_init__(self) -> None:
        self.rows_to_expand.sort()
        self.cols_to_expand.sort()

    def expand(self, expansion_rate: int):
        if expansion_rate < 0:
            raise ValueError("expansion rate must be 0 or greater")
        self.expansion_rate = expansion_rate

    @staticmethod
    def from_str(input: str) -> "Universe":
        image_data = input.splitlines(keepends=False)
        size = Size(len(image_data), len(image_data[0]))

        # iterate over rows and cols
        # keep track of which indices are empty
        rows_to_expand: list[int] = []
        galaxies: list[Galaxy] = []

        cols_to_expand_set: set[int] = set(range(size.cols - 1))
        # print(cols_to_expand_set)

        for i, row in enumerate(image_data):
            galaxy_found = False
            for j, col in enumerate(row):
                if col == "#":
                    galaxy_found = True
                    if j in cols_to_expand_set:
                        cols_to_expand_set.remove(j)
                    galaxies.append(Galaxy(i, j))
            if not galaxy_found:
                rows_to_expand.append(i)

        cols_to_expand = list(cols_to_expand_set)

        return Universe(
            galaxies=galaxies,
            size=size,
            rows_to_expand=rows_to_expand,
            cols_to_expand=cols_to_expand,
            image_data=image_data,
        )

    def __str__(self) -> str:
        image_data = ["." * self.size.cols] * self.size.rows
        for i, galaxy in enumerate(self.galaxies):
            row = image_data[galaxy.row]
            image_data[galaxy.row] = (
                row[: galaxy.col] + str(i + 1) + row[galaxy.col + 1 :]
            )

        expanded_image_data = [*image_data]

        if self.expansion_rate >= 1:
            for offset, row_index in enumerate(self.rows_to_expand):
                row = image_data[row_index]
                for _ in range(self.expansion_rate):
                    expanded_image_data.insert(row_index + offset, row)

            for offset, col_index in enumerate(self.cols_to_expand):
                for row_index, row in enumerate(expanded_image_data):
                    expanded_row = (
                        row[: col_index + offset]
                        + "." * (self.expansion_rate)
                        + row[col_index + offset :]
                    )
                    expanded_image_data[row_index] = expanded_row

        print_str = "\n".join(expanded_image_data)
        return (
            f"Universe ({len(expanded_image_data)} , {len(expanded_image_data[0])})\n"
            + f"expansion rate: {self.expansion_rate}\n"
            + f"{print_str}\n"
        )

    def distance(self, galaxy_a: Galaxy, galaxy_b: Galaxy) -> int:
        def calc_dist(galaxy_a: Galaxy, galaxy_b: Galaxy) -> int:
            row_dist = abs(galaxy_a.row - galaxy_b.row)
            col_dist = abs(galaxy_a.col - galaxy_b.col)
            return row_dist + col_dist

        base_dist = calc_dist(galaxy_a, galaxy_b)
        if self.expansion_rate == 0:
            return base_dist

        def num_expansions_between(
            expansion_list: list[int], val_a: int, val_b: int
        ) -> int:
            if val_a > val_b:
                higher = val_a
                lower = val_b
            else:
                higher = val_b
                lower = val_a

            count = 0
            for val in expansion_list:
                if lower <= val <= higher:
                    count += 1
            return count

        num_row_expansions = num_expansions_between(
            self.rows_to_expand, galaxy_a.row, galaxy_b.row
        )
        num_col_expansions = num_expansions_between(
            self.cols_to_expand, galaxy_a.col, galaxy_b.col
        )


        return base_dist + self.expansion_rate * (
            num_row_expansions + num_col_expansions
        )


def part_1(input: str):
    universe = Universe.from_str(input)
    universe.expand(2-1)

    distance_sum = 0
    # iterate over the list of Galaxies
    for i, galaxy in enumerate(universe.galaxies[:-1]):
        for j, otherGalaxy in enumerate(universe.galaxies[i + 1 :]):
            #   use their positions to determine their distances
            dist = universe.distance(galaxy, otherGalaxy)
            # print(f"Galaxy {i+1} -> Galaxy {i + 2 + j} = {dist} distance")
            #   sum the distances
            distance_sum += dist

    return distance_sum


def part_2(input: str):
    universe = Universe.from_str(input)
    universe.expand(1000000 - 1)

    # print(universe)
    distance_sum = 0
    # iterate over the list of Galaxies
    for i, galaxy in enumerate(universe.galaxies[:-1]):
        for j, otherGalaxy in enumerate(universe.galaxies[i + 1 :]):
            #   use their positions to determine their distances
            dist = universe.distance(galaxy, otherGalaxy)
            # print(f"Galaxy {i+1} -> Galaxy {i + 2 + j} = {dist} distance")
            #   sum the distances
            distance_sum += dist

    return distance_sum


def main():
    args = sys.argv
    options = cli.parse_args(args[1:])

    DAY = options["day"]
    if DAY is None:
        raise EnvironmentError("Env var `$DAY` is not set.")
    DAY = int(DAY)
    print(f"Running for day {DAY}\n")

    input_filename = (
        f"aoc/day{DAY}/test_input.txt" if options["test"] else f"aoc/day{DAY}/input.txt"
    )
    with open(input_filename) as f:
        input_str = f.read()

    print()
    part_1_out = part_1(input_str)
    print()

    part_2_out = part_2(input_str)
    print("\n#########################\n")
    return (str(part_1_out), str(part_2_out))


if __name__ == "__main__":
    main()
