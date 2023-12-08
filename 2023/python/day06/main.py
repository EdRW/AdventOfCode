from dataclasses import dataclass
from typing import Callable, TypedDict, cast, Literal, NamedTuple
from collections import deque


class Race(NamedTuple):
    time: int
    distance: int


def parse_from_races_str(input: str) -> list["Race"]:
    lines = input.splitlines(keepends=False)
    times = lines[0].split(":")[1].split()
    distances = lines[1].split(":")[1].split()
    return [Race(int(time), int(dist)) for time, dist in zip(times, distances)]


def simulate_race_dist(btn_press_ms: int, race_time_ms: int) -> int:
    move_time_ms=max(0, race_time_ms - btn_press_ms)
    #      speed * time = distance
    return btn_press_ms * move_time_ms

def count_ways_to_win(race: Race)-> int:
    ways_to_win = 0
    for btn_press_ms in range(1,race.time + 1):
        simulated_dist = simulate_race_dist(btn_press_ms, race.time)
        if simulated_dist > race.distance:
            ways_to_win += 1
    return ways_to_win

def part_1(input: str):
    races = parse_from_races_str(input)

    error_margin = 1
    for race in races:
        num_ways_to_win = count_ways_to_win(race)
        print(num_ways_to_win)
        error_margin *= num_ways_to_win
    print(error_margin)
       


def part_2(input: str):
    pass


def main():
    input_filename = "day06/input.txt"
    # input_filename = "day06/test_input.txt"

    sections: str
    with open(input_filename) as f:
        sections = f.read()

    part_1(sections)
    part_2(sections)


if __name__ == "__main__":
    main()
