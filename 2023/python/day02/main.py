from dataclasses import dataclass
from typing import Callable, TypedDict, cast

COLOR_LIMITS = {"red": 12, "green": 13, "blue": 14}


class CubeSet(TypedDict):
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
    print(split_line)
    print(game_id)

    cube_sets_strs = split_line[1].split("; ")
    sets = [parse_cube_set(cube_set_str) for cube_set_str in cube_sets_strs]
    return Game(game_id, sets)


def parse_cube_set(cube_set_str: str):
    color_sets = cube_set_str.split(", ")
    print(color_sets)

    cube_set = {}
    for color_set in color_sets:
        count_and_color = color_set.split(" ")
        print(count_and_color)
        cube_set[count_and_color[1]] = int(count_and_color[0])

    return cast(CubeSet, cube_set)


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


def main():
    input_filename = "day02/input.txt"

    with open(input_filename) as f:
        lines = f.readlines()

    result = 0
    for line in lines:
        game = load_game(line.strip())
        if is_valid_game(game):
            result += game.id

    print(f"part 1: {  result}")


if __name__ == "__main__":
    main()
