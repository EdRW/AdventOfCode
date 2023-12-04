from dataclasses import dataclass
from typing import Callable, TypedDict, cast

EngineSchematic = list[list[str]]


def parts_1_and_2(lines: list[str]):
    # parse into grid
    # iterate over grid to find numbers
    #   that are adjacent to a non-period symbol
    # sum all the numbers found in this way

    schematic = get_engine_schematic(lines)
    # [print(line) for line in schematic]

    part_num_sum = 0
    stars_dict: dict[tuple[int, int], list[int]] = {}
    for i, row in enumerate(schematic):
        num_str: str = ""
        for j in range(len(row)):
            char = row[j]
            if char.isnumeric():
                num_str += char
            elif num_str:
                if is_part_num(
                    i, start=j - len(num_str) - 1, end=j + 1, schematic=schematic
                ):
                    part_num_sum += int(num_str)

                # does it have a star?
                star_locs = get_adjacent_stars(
                    i, start=j - len(num_str) - 1, end=j + 1, schematic=schematic
                )
                for star_loc in star_locs:
                    stars_dict.setdefault(star_loc, []).append(int(num_str))

                # reset num_str
                num_str = ""

        if num_str and is_part_num(
            i=i, start=len(row) - len(num_str) - 1, end=len(row), schematic=schematic
        ):
            part_num_sum += int(num_str)

            # does it have a star?
            star_locs = get_adjacent_stars(
                i=i,
                start=len(row) - len(num_str) - 1,
                end=len(row),
                schematic=schematic,
            )
            for star_loc in star_locs:
                stars_dict.setdefault(star_loc, []).append(int(num_str))

    gear_ratio = 0
    for pns_near_star in stars_dict.values():
        if len(pns_near_star) == 2:
            gear_ratio += pns_near_star[0] * pns_near_star[1]

    print(f'part 1: {part_num_sum}')
    print(f'part 2: {gear_ratio}')


def is_part_num(i: int, start: int, end: int, schematic: EngineSchematic):
    # check the area around the found number
    adjacent_chars = get_adjacent_chars(i, start, end, schematic=schematic)
    for adjacent_char in adjacent_chars:
        if is_symbol(adjacent_char):
            return True
    return False


def get_adjacent_chars(
    i: int, start: int, end: int, schematic: EngineSchematic
) -> list[str]:
    start = max(start, 0)
    end = min(end, len(schematic[0]))

    preceding_char = schematic[i][start]
    following_char = schematic[i][end - 1]
    adjacent_chars = [preceding_char, following_char]

    for j in range(start, end):
        # check row above
        if i > 0:
            adjacent_chars.append(schematic[i - 1][j])
        # check row below
        if i < len(schematic) - 1:
            adjacent_chars.append(schematic[i + 1][j])
    return adjacent_chars


def is_symbol(char: str):
    return (not char.isnumeric()) and char != "."


def get_adjacent_stars(
    i: int, start: int, end: int, schematic: EngineSchematic
) -> set[tuple[int, int]]:
    # special cases
    # # . #
    # . * .

    adjacent_stars: set[tuple[int, int]] = set()

    start = max(start, 0)
    end = min(end, len(schematic[0]))

    preceding_char = schematic[i][start]
    if preceding_char == "*":
        adjacent_stars.add((i, start))

    following_char = schematic[i][end - 1]
    if following_char == "*":
        adjacent_stars.add((i, end - 1))

    for j in range(start, end):
        # check row above
        if i > 0 and schematic[i - 1][j] == "*":
            adjacent_stars.add((i - 1, j))
        # check row below
        if i < len(schematic) - 1 and schematic[i + 1][j] == "*":
            adjacent_stars.add((i + 1, j))
    return adjacent_stars


def get_engine_schematic(lines: list[str]) -> EngineSchematic:
    schematic: EngineSchematic = []
    for line in lines:
        parsed_line = parse_line(line.strip())
        schematic.append(parsed_line)
    return schematic


def parse_line(line: str) -> list[str]:
    return [char for char in line]


def main():
    input_filename = "day03/input.txt"
    # input_filename = "day03/test_input.txt"

    with open(input_filename) as f:
        lines = f.readlines()
        parts_1_and_2(lines)


if __name__ == "__main__":
    main()
