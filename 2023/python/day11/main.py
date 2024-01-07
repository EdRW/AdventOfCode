from collections import Counter, deque
from collections.abc import Iterator
from dataclasses import dataclass
from enum import IntEnum, StrEnum, Enum
from math import lcm
import os
import subprocess
import time
from typing import Callable, Iterable, Literal, NamedTuple, Self, TypedDict, cast


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


class Universe:
    size: Size
    image_data: list[str]

    def __init__(self, input: str) -> None:
        self.print_str = input
        self.image_data = input.splitlines(keepends=False)
        self.size = Size(len(self.image_data), len(self.image_data[0]))

        # iterate over rows and cols
        # keep track of which indices are empty
        # [x] create 2 lists of booleans. one the size of the rows, one the size of the cols
        # [ ] This will be used later to determine which rows and cols to expand
        self.rows_to_expand: list[int] = []
        self.galaxies: list[Galaxy] = []

        cols_to_expand: set[int] = set(range(self.size.cols - 1))
        print(cols_to_expand)

        for i, row in enumerate(self.image_data):
            galaxy_found = False
            for j, col in enumerate(row):
                if col == "#":
                    galaxy_found = True
                    if j in cols_to_expand:
                        cols_to_expand.remove(j)
                    self.galaxies.append(Galaxy(i, j))
            if not galaxy_found:
                self.rows_to_expand.append(i)

        self.cols_to_expand = list(cols_to_expand)

        self.rows_to_expand.sort()
        self.cols_to_expand.sort()

        print(f"rows to expand {self.rows_to_expand}")
        print(f"cols to expand {self.cols_to_expand}")

    def expand(self) -> "Universe":
        expanded_image_data = [*self.image_data]
        # iterate over rows_to_expand and cols_to_expand and expand if necessary
        for offset, row_index in enumerate(self.rows_to_expand):
            row = self.image_data[row_index]
            expanded_image_data.insert(row_index + offset, row)

        for offset, col_index in enumerate(self.cols_to_expand):
            # print(f"\nCOL: {col_index}")
            for row_index, row in enumerate(expanded_image_data):
                # print(f"row: {row_index}, col: {col_index}")

                # print(" " * col_index + "v")
                # print(row)
                expanded_row = (
                    row[: col_index + offset] + "." + row[col_index + offset :]
                )
                # print(expanded_row)
                # print()
                expanded_image_data[row_index] = expanded_row

        return Universe("\n".join(expanded_image_data))

    def __str__(self) -> str:
        return f"Universe ({self.size.rows} , {self.size.cols})\n{self.print_str}\n"

def distance(galaxy_a: Galaxy, galaxy_b: Galaxy) -> int:
    row_dist = abs(galaxy_a.row - galaxy_b.row)
    col_dist = abs(galaxy_a.col - galaxy_b.col)
    return row_dist + col_dist

def part_1(input: str):
    universe = Universe(input)
    print(universe)

    expanded_Universe = universe.expand()
    print(expanded_Universe)

    distance_sum = 0
    # iterate over the list of Galaxies
    for i, galaxy in enumerate(expanded_Universe.galaxies[:-1]):
        for j, otherGalaxy in enumerate(expanded_Universe.galaxies[i+1:]):
            #   use their positions to determine their distances
            dist = distance(galaxy, otherGalaxy) 
            #   sum the distances
            distance_sum += dist

    print(distance_sum)
    


def part_2():
    pass


def main():
    DAY = os.getenv("DAY")
    if DAY is None:
        raise EnvironmentError("Env var `$DAY` is not set.")
    DAY = int(DAY)
    print(f"Running for day {DAY}\n")

    filenames = [
        # f"day{DAY}/test_input.txt",
        f"day{DAY}/input.txt",
    ]
    for input_filename in filenames:
        with open(input_filename) as f:
            input_str = f.read()

        print()
        part_1(input_str)
        print()

        part_2()
        print("\n#########################\n")


if __name__ == "__main__":
    main()
