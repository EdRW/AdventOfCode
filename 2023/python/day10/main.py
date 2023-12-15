from collections import Counter, deque
from collections.abc import Iterator
from dataclasses import dataclass
from enum import IntEnum, StrEnum, Enum
from math import lcm
import os
import subprocess
import time
from typing import Callable, Iterable, Literal, NamedTuple, TypedDict, cast


class Direction(Enum):
    N = 1
    S = 2
    E = 3
    W = 4


class Turn(Enum):
    CLOCKWISE = 1
    COUNTER_CLOCKWISE = 2


reverse_direction: dict[Direction, Direction] = {
    Direction.N: Direction.S,
    Direction.S: Direction.N,
    Direction.E: Direction.W,
    Direction.W: Direction.E,
}


def directed_arrow(symbol: str, direction: Direction):
    directed_arrow_dict: dict[str, dict[Direction, str]] = {
        "|": {Direction.N: "↑", Direction.S: "↓"},
        "-": {Direction.W: "←", Direction.E: "→"},
        "L": {Direction.N: "⬑", Direction.E: "↳"},
        "J": {Direction.N: "⬏", Direction.W: "↲"},
        "7": {Direction.S: "↴", Direction.W: "↰"},
        "F": {Direction.S: "⬐", Direction.E: "↱"},
    }
    return directed_arrow_dict[symbol][direction]


direction_turn_map: dict[tuple[Direction, Direction], Turn] = {
    (Direction.N, Direction.E): Turn.COUNTER_CLOCKWISE,
    (Direction.E, Direction.N): Turn.CLOCKWISE,
    (Direction.E, Direction.S): Turn.COUNTER_CLOCKWISE,
    (Direction.S, Direction.E): Turn.CLOCKWISE,
    (Direction.W, Direction.N): Turn.COUNTER_CLOCKWISE,
    (Direction.N, Direction.W): Turn.CLOCKWISE,
    (Direction.S, Direction.W): Turn.COUNTER_CLOCKWISE,
    (Direction.W, Direction.S): Turn.CLOCKWISE,
}

def get_turn_dirs(symbol: str, inside_turn: Turn) -> tuple[Direction, Direction]:
    symbol_turn_map = {
        "L": (Direction.N, Direction.E),
        "J": (Direction.W, Direction.N),
        "7": (Direction.S, Direction.W),
        "F": (Direction.E, Direction.S),
    }
    clockwise_directions = symbol_turn_map[symbol]
    if inside_turn == Turn.CLOCKWISE:
        return clockwise_directions

    return (
        reverse_direction[clockwise_directions[0]],
        reverse_direction[clockwise_directions[1]],
    )


class Pos(NamedTuple):
    x: int
    y: int


class Pipe:
    io: dict[Direction, Direction]
    symbol: str

    def __init__(self, symbol: str, *dirs: Direction):
        self.io = {}
        for i, entry_dir in enumerate(dirs):
            for j, exit_dir in enumerate(dirs):
                if j != i:
                    self.io[entry_dir] = exit_dir

        self.symbol = symbol

    def can_enter_from(self, enter_from: Direction) -> bool:
        return enter_from in self.io

    def traverse(self, pos: Pos, enter_from: Direction) -> tuple[Pos, Direction]:
        exit_direction = self.io[enter_from]
        x, y = pos
        dx, dy = displacement(exit_direction)

        return Pos(x + dx, y + dy), exit_direction


class PipeMap:
    path_map: list[list[str]]
    start_pos: Pos

    def __init__(self, __iterable: Iterable[str]) -> None:
        self.path_map = [[char for char in row] for row in __iterable]
        self.start_pos = find_start(__iterable)

    def size(self):
        x_max = len(self.path_map[0])
        y_max = len(self.path_map)
        return x_max, y_max

    def at(self, pos: Pos) -> str:
        return self.path_map[pos.y][pos.x]

    @staticmethod
    def from_str(input_str: str):
        pipe_map = input_str.splitlines(keepends=False)
        return PipeMap(pipe_map)

    def path(self):
        pos, entry_direction, pipe_symbol = search_enterable_pipe(
            self.start_pos, self.path_map
        )
        yield pipe_dict["S"], reverse_direction[entry_direction], reverse_direction[
            entry_direction
        ], self.start_pos

        while pipe_symbol != "S":
            # i just came out of a pipe so i have
            # new pos and direction
            # print(f'pos: ({pos.x}, {pos.y}), dir: {entry_direction}, sym: {pipe_symbol}')
            pipe = pipe_dict[pipe_symbol]
            next_pos, exit_direction = pipe.traverse(pos, entry_direction)
            yield pipe, entry_direction, exit_direction, pos,
            # print(f'exit pos: ({pos.x}, {pos.y}), exit dir: {direction}')
            entry_direction = reverse_direction[exit_direction]
            pipe_symbol = self.at(next_pos)
            pos = next_pos

    def __iter__(self):
        for row in self.path_map:
            yield row

    def __str__(self):
        path_map = [["." for _ in row] for row in self.path_map]

        path_map[self.start_pos.y][self.start_pos.x] = Color.BLUE_BG.color_str("S")

        halfway_point = furthest_dist(self)

        path = self.path()
        next(path)
        step = 1
        for pipe, _, exit_direction, pos in path:
            if pipe.symbol == "S":
                break
            if step < halfway_point:
                color = Color.GREEN
            elif step > halfway_point:
                color = Color.YELLOW
            else:
                color = Color.RED
            path_map[pos.y][pos.x] = color.color_str(
                directed_arrow(pipe.symbol, exit_direction)
            )

            step += 1
        map_str = ""
        for row in path_map:
            map_str += "".join(row) + "\n"

        return map_str

    def animated_print(self, fps=1):
        def print_map(path: list[list[str]]):
            map_str = ""
            for row in path:
                map_str += "".join(row) + "\n"
            os.system("clear")
            print(map_str, end="\r", flush=True)

        path_map = [["." for _ in row] for row in self.path_map]

        path_map[self.start_pos.y][self.start_pos.x] = Color.BLUE.color_str("S")

        halfway_point = furthest_dist(self)

        sleep_secs = 1 / fps
        print_map(path_map)

        path = self.path()
        next(path)
        step = 1
        for pipe, _, exit_direction, pos in path:
            print(pipe)
            if pipe.symbol == "S":
                break
            if step < halfway_point:
                color = Color.GREEN
            elif step > halfway_point:
                color = Color.YELLOW
            else:
                color = Color.RED
            path_map[pos.y][pos.x] = color.color_str(
                directed_arrow(pipe.symbol, exit_direction)
            )
            time.sleep(sleep_secs)
            print_map(path_map)

            step += 1

        path_map[self.start_pos.y][self.start_pos.x] = Color.BLUE_BG.color_str("S")
        time.sleep(sleep_secs)
        print_map(path_map)


class Color(StrEnum):
    BLINK = "\33[6m"
    GREEN = "\033[92m"
    YELLOW = "\033[93m"
    BLUE = "\033[94m"
    BLUE_BG = "\033[104m"
    VIOLET = "\033[95m"
    RED = "\033[91m"
    ENDC = "\033[0m"

    def color_str(self, txt: str):
        return f"{self}{txt}{Color.ENDC}"

    def print(self, txt: str):
        print(self.color_str(txt))


def displacement(direction: Direction) -> Pos:
    x, y = 0, 0

    if direction == Direction.N:
        y -= 1
    elif direction == Direction.S:
        y += 1
    elif direction == Direction.E:
        x += 1
    else:
        x -= 1

    return Pos(x, y)


pipe_dict = {
    "S": Pipe("S", Direction.S, Direction.N, Direction.E, Direction.W),
    "|": Pipe("|", Direction.S, Direction.N),
    "-": Pipe("-", Direction.E, Direction.W),
    "L": Pipe("L", Direction.E, Direction.N),
    "J": Pipe("J", Direction.W, Direction.N),
    "7": Pipe("7", Direction.W, Direction.S),
    "F": Pipe("F", Direction.E, Direction.S),
}


def is_inside_turn(
    entry_direction: Direction, exit_direction: Direction, inside_turn: Turn
) -> bool:
    turn = direction_turn_map[(entry_direction, exit_direction)]
    return turn == is_inside_turn




def is_pipe(thing: str):
    return thing in pipe_dict


def find_start(pipe_map: Iterable[str]):
    # find S
    for y, row in enumerate(pipe_map):
        for x, thing in enumerate(row):
            if thing == "S":
                return Pos(x, y)
    raise ValueError("Cant find the S")


def search_enterable_pipe(pos: Pos, pipe_map: list[list[str]]):
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
        adjacency_list.append((Pos(pos.x - 1, pos.y), Direction.W))
    if pos.y > 0:
        adjacency_list.append((Pos(pos.x, pos.y - 1), Direction.N))

    if pos.x < x_max:
        adjacency_list.append((Pos(pos.x + 1, pos.y), Direction.E))
    if pos.y < y_max:
        adjacency_list.append((Pos(pos.x, pos.y + 1), Direction.S))
    return adjacency_list


def calc_enclosed_area():
    pass


def print_stats(pipe_map: PipeMap):
    symbol_counter = Counter()
    direction_counter = Counter()
    turn_counter = Counter()
    for pipe, entry_direction, exit_direction, pos in pipe_map.path():
        symbol_counter.update(pipe.symbol)
        direction_counter.update({exit_direction: 1})
        if pipe.symbol in "LJF7":
            turn = {direction_turn_map[(entry_direction, exit_direction)]}
            # print(f"{pipe.symbol}, {entry_direction}, {exit_direction} -> {turn}")
            turn_counter.update(turn)
    print(f"stats:")
    for key, val in symbol_counter.items():
        print(f"{key} : {val}")
    for key, val in direction_counter.items():
        print(f"{key} : {val}")
    for key, val in turn_counter.items():
        print(f"{key} : {val}")


def furthest_dist(pipe_map: PipeMap):
    return len(list(pipe_map.path())) // 2


def part_1(pipe_map: PipeMap):
    print(furthest_dist(pipe_map))


def in_bounds(pos: Pos, x_max: int, y_max: int):
    return 0 <= pos.x < x_max and 0 <= pos.y < y_max

def paint_adj_insides(pos: Pos, pipe_map: list[list[str]], memo: set[Pos]):
    x_max = len(pipe_map[0])
    y_max = len(pipe_map)
    adj = adjacent_positions(pos, x_max=x_max, y_max=y_max)
    for adj_pos, _ in adj:
        symbol = pipe_map[adj_pos.y][adj_pos.x]
        if symbol != '.':
            continue
        pipe_map[adj_pos.y][adj_pos.x] = '*'
        memo.add(adj_pos)
        paint_adj_insides(adj_pos, pipe_map, memo)
        
    
    


def part_2(pipe_map: PipeMap):
    pipe_set: set[Pos] = set()
    in_set: set[Pos] = set()
    # out_set: set[Pos] = set()
    clean_map = [["." for _ in row] for row in pipe_map]
    turn_counter = Counter()
    for pipe, entry_direction, exit_direction, pos in pipe_map.path():
        pipe_set.add(pos)
        clean_map[pos.y][pos.x] = pipe.symbol
        if pipe.symbol in "LJF7":
            turn_counter.update({direction_turn_map[(entry_direction, exit_direction)]})

    most_turns, _ = turn_counter.most_common()[0]
    print(f"{most_turns} *most")

    for i in range(len(clean_map)):
        for j in range(len(clean_map[i])):
            pos = Pos(j, i)
            if pos in pipe_set:
                continue

    inside = "*"
    outside = "O"
    rev_turn = {
        Turn.CLOCKWISE: Turn.COUNTER_CLOCKWISE,
        Turn.COUNTER_CLOCKWISE: Turn.CLOCKWISE,
    }
    inside_directions =None
    for pipe, entry_direction, exit_direction, pos in pipe_map.path():
        # print(f"{pipe.symbol}, {pos}, {entry_direction}, {exit_direction}")
        if pipe.symbol not in "LJF7":
            if inside_directions:
                for adj_direction in inside_directions:
                    offset = displacement(adj_direction)
                    adj_pos = Pos(pos.x + offset.x, pos.y + offset.y)
                    # print(f"inner adj: {adj_direction}, {adj_pos}")
                    if adj_pos in pipe_set or not in_bounds(adj_pos, *pipe_map.size()):
                        continue
                    clean_map[adj_pos.y][adj_pos.x] = inside
                    # print(f"{pipe.symbol}, {pos}")
                    in_set.add(adj_pos)
                
            continue

        # print(f"{most_turns} *most")
        turn = direction_turn_map[entry_direction, exit_direction]
        # print(turn)
        inside_turn = (
            Turn.CLOCKWISE
            if most_turns == turn
            else Turn.COUNTER_CLOCKWISE
        )
        # print(inside_turn)
        inside_directions = get_turn_dirs(pipe.symbol, inside_turn)
        for adj_direction in inside_directions:
            offset = displacement(adj_direction)
            adj_pos = Pos(pos.x + offset.x, pos.y + offset.y)
            # print(f"inner adj: {adj_direction}, {adj_pos}")
            if adj_pos in pipe_set or not in_bounds(adj_pos, *pipe_map.size()):
                continue
            clean_map[adj_pos.y][adj_pos.x] = inside
            # print(f"{pipe.symbol}, {pos}")
            in_set.add(adj_pos)

        # outside_directions = get_turn_dirs(pipe.symbol, rev_turn[inside_turn])
        # for adj_direction in outside_directions:
        #     offset = displacement(adj_direction)
        #     adj_pos = Pos(pos.x + offset.x, pos.y + offset.y)
        #     # print(f"outer adj: {adj_direction}, {adj_pos}")
        #     if adj_pos in pipe_set or not in_bounds(adj_pos, *pipe_map.size()):
        #         continue
        #     clean_map[adj_pos.y][adj_pos.x] = outside
        
    # recursively color points adjacency to inside points
    # print(in_set)
    for pos in [key for key in in_set]:
        paint_adj_insides(pos, clean_map, in_set)


    map_str = ""
    for row in clean_map:
        map_str += "".join(row) + "\n"

    
        

    print(map_str)
    # print(in_set)
    print (len(in_set))

    # for symbol in "LJF7":
    #     print(f"Symbol: {symbol}: {get_inside_dirs(symbol,Turn.CLOCKWISE )}")
    #     print(f"Symbol: {symbol}: {get_inside_dirs(symbol,Turn.COUNTER_CLOCKWISE )}")


def main():
    filenames = [
        # "day10/test_input.txt",
        # "day10/test_input_2.txt",
        # "day10/test_input_3.txt",
        # "day10/test_input_4.txt",
        # "day10/test_input_5.txt",
        # "day10/test_input_6.txt",
        "day10/input.txt",
    ]
    for input_filename in filenames:
        with open(input_filename) as f:
            input_str = f.read()

        pipe_map = PipeMap.from_str(input_str)
        print(pipe_map)
        # pipe_map.animated_print(fps=240)
        # print_stats(pipe_map)
        print()
        part_1(pipe_map)
        print()

        part_2(pipe_map)
        print('\n#########################\n')


if __name__ == "__main__":
    main()
