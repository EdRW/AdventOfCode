from collections import Counter, deque
from dataclasses import dataclass
from enum import IntEnum
from math import lcm
from typing import Callable, Literal, NamedTuple, TypedDict, cast

Direction = Literal["top", "bottom", "left", "right"]
reverse_direction: dict[Direction, Direction] = {
    "top": "bottom",
    "bottom": "top",
    "left": "right",
    "right": "left",
}


class Pos(NamedTuple):
    x: int
    y: int


@dataclass
class Pipe:
    io: dict[Direction, Direction]

    def __init__(self, dir_1: Direction, dir_2: Direction):
        self.io = {dir_1: dir_2, dir_2: dir_1}

    def can_enter_from(self, enter_from: Direction) -> bool:
        return enter_from in self.io

    def traverse(self, pos: Pos, enter_from: Direction) -> tuple[Pos, Direction]:
        exit_direction = self.io[enter_from]
        x, y = pos
        if exit_direction == 'top':
            y -= 1
        elif exit_direction == 'bottom':
            y+=1
        elif exit_direction == 'left':
            x -=1
        else:
            x+= 1
        
        return Pos(x, y), exit_direction


pipe_dict = {
    "|": Pipe("bottom", "top"),
    "-": Pipe("left", "right"),
    "L": Pipe("right", "top"),
    "J": Pipe("left", "top"),
    "7": Pipe("left", "bottom"),
    "F": Pipe("right", "bottom"),
}


def is_pipe(thing: str):
    return thing in pipe_dict


def find_start(pipe_map: list[str]):
    # find S
    for y, row in enumerate(pipe_map):
        for x, thing in enumerate(row):
            if thing == "S":
                return Pos(x, y)
    raise ValueError("Cant find the S")


def search_enterable_pipe(pos: Pos, pipe_map: list[str]):
    x_max = len(pipe_map[0]) - 1
    y_max = len(pipe_map) - 1

    adj_pos_list = adjacent_positions(pos, x_max, y_max)
    for adj_pos, adj_direction in adj_pos_list:
        pipe_symbol = pipe_map[adj_pos.y][adj_pos.x]
        pipe = pipe_dict.get(pipe_symbol)
        if pipe and pipe.can_enter_from(reverse_direction[adj_direction]):
            return adj_pos, reverse_direction[adj_direction], pipe_symbol
    raise ValueError(f"No compatible adjacent pipes found at pos: {pos}")


def adjacent_positions(pos: Pos, x_max: int, y_max: int):
    adjacency_list: list[tuple[Pos, Direction]] = []
    if pos.x > 0:
        adjacency_list.append((Pos(pos.x - 1, pos.y), "left"))
    if pos.y > 0:
        adjacency_list.append((Pos(pos.x, pos.y - 1), "top"))

    if pos.x < x_max:
        adjacency_list.append((Pos(pos.x + 1, pos.y), "right"))
    if pos.y < y_max:
        adjacency_list.append((Pos(pos.x, pos.y + 1), "bottom"))
    return adjacency_list


def part_1(input_str: str):
    pipe_map = input_str.splitlines(keepends=False)

    start_pos = find_start(pipe_map)

    # determine initial orientation
    pos, entry_direction, pipe_symbol = search_enterable_pipe(start_pos, pipe_map)

    steps = 1
    while pipe_symbol != "S":
        # i just came out of a pipe so i have
        # new pos and direction
        # print(f'pos: ({pos.x}, {pos.y}), dir: {entry_direction}, sym: {pipe_symbol}')
        pipe = pipe_dict[pipe_symbol]
        pos, direction = pipe.traverse(pos, entry_direction)
        # print(f'exit pos: ({pos.x}, {pos.y}), exit dir: {direction}')
        entry_direction = reverse_direction[direction]
        pipe_symbol = pipe_map[pos.y][pos.x]
        steps += 1

    print(steps//2)


def part_2(input_str: str):
    pass


def main():
    input_filename = "day10/input.txt"
    # input_filename = "day10/test_input.txt"
    # input_filename = "day10/test_input_2.txt"

    with open(input_filename) as f:
        input_str = f.read()

    part_1(input_str)

    print()

    part_2(input_str)


if __name__ == "__main__":
    main()
