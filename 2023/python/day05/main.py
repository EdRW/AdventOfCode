from dataclasses import dataclass
from typing import Callable, TypedDict, cast, Literal
from collections import deque


SectionTitle = Literal[
    "seeds",
    "seed-to-soil",
    "soil-to-fertilizer",
    "fertilizer-to-water",
    "water-to-light",
    "light-to-temperature",
    "temperature-to-humidity",
    "humidity-to-location",
]


@dataclass
class Range:
    dest_start: int
    src_start: int
    len: int

    @staticmethod
    def parse_from_str(line: str) -> "Range":
        vals = [int(val) for val in line.split()]
        return Range(dest_start=vals[0], src_start=vals[1], len=vals[2])

    def __getitem__(self, src: int) -> int:
        if src not in self:
            raise KeyError(f"{src} not found")
        return self.dest_start + (src - self.src_start)

    def __contains__(self, src: int) -> bool:
        return self.src_start <= src < (self.src_start + self.len)


@dataclass
class FarmMap:
    title: SectionTitle
    ranges: list[Range]

    @staticmethod
    def parse_from_str(section: str) -> "FarmMap":
        lines = section.splitlines(keepends=False)
        title = cast(SectionTitle, lines[0].split()[0])

        ranges = [Range.parse_from_str(line) for line in lines[1:]]

        return FarmMap(title, ranges=ranges)

    def __getitem__(self, src: int) -> int:
        for range in self.ranges:
            if src in range:
                return range[src]
        return src


def parse_seeds(line: str) -> list[int]:
    seeds = line.split(":")[1].split()
    return [int(seed) for seed in seeds]


def part1(sections: list[str]):
    seeds = parse_seeds(sections[0])

    farm_maps: list[FarmMap] = []
    for section in sections[1:]:
        farm_map = FarmMap.parse_from_str(section)
        farm_maps.append(farm_map)

    min_location = seed_location(farm_maps, seeds[0])
    for seed in seeds[1:]:
        dest = seed_location(farm_maps, seed)
        min_location = min(min_location, dest)

    print(min_location)


def seed_location(farm_maps: list[FarmMap], seed: int) -> int:
    dest = seed
    for farm_map in farm_maps:
        dest = farm_map[dest]
    return dest


def part2(sections: list[str]):
    pass


def main():
    input_filename = "day05/input.txt"
    # input_filename = "day05/test_input.txt"

    sections: list[str]
    with open(input_filename) as f:
        sections = f.read().split("\n\n")

    part1(sections)
    part2(sections)


if __name__ == "__main__":
    main()
