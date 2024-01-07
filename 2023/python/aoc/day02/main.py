from dataclasses import dataclass
from typing import Callable, TypedDict, cast

COLOR_LIMITS = {"red": 12, "green": 13, "blue": 14}


class CubeSet(TypedDict, total=False):
    red: int
    green: int
    blue: int


@dataclass
class Game:
    id: int
    sets: list[CubeSet]


def load_game(line: str) -> Game:
    split_line = line.split(": ")
    game_id = int(split_line[0].split()[-1])

    cube_sets_strs = split_line[1].split("; ")
    sets = [parse_cube_set(cube_set_str) for cube_set_str in cube_sets_strs]

    return Game(game_id, sets)


def parse_cube_set(cube_set_str: str):
    color_sets = cube_set_str.split(", ")

    cube_set = CubeSet()
    for color_set in color_sets:
        count_and_color = color_set.split(" ")
        cube_set[count_and_color[1]] = int(count_and_color[0])

    return cube_set


def is_valid_game(game: Game):
    for cube_set in game.sets:
        if not is_valid_set(cube_set):
            return False
    return True


def is_valid_set(cube_set: CubeSet):
    for color in cube_set:
        if cube_set[color] > COLOR_LIMITS[color]:
            return False
    return True


def get_min_cubes(game: Game):
    min_cube_set = CubeSet(red=0, green=0, blue=0)
    for cube_set in game.sets:
        for color in cube_set:
            min_cube_set[color] = max(cube_set[color], min_cube_set[color])

    return min_cube_set


def get_cube_set_power(cube_set: CubeSet):
    power = 1
    for color in cube_set:
        power *= cube_set[color]
    return power


def get_game_power(game: Game):
    min_cubes = get_min_cubes(game)
    power = get_cube_set_power(min_cubes)
    return power


def main():
    input_filename = "day02/input.txt"

    with open(input_filename) as f:
        lines = f.readlines()

    valid_id_sum = 0
    power_sum = 0
    for line in lines:
        game = load_game(line.strip())
        if is_valid_game(game):
            valid_id_sum += game.id

        power_sum += get_game_power(game)

    print(f"part 1: {valid_id_sum}")
    print(f"part 2: {power_sum}")


if __name__ == "__main__":
    main()
