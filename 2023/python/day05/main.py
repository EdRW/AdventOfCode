from dataclasses import dataclass
from typing import Callable, TypedDict, cast, Literal, NamedTuple
from collections import deque


class Range(NamedTuple):
    start: int
    len: int

    def __lt__(self, other: "Range") -> bool:
        return self.start < other.start

    def __contains__(self, src: int) -> bool:
        return self.start <= src < (self.start + self.len)

    @property
    def end(self):
        return self.start + self.len - 1

    def __repr__(self) -> str:
        return f"Range: {self.start}, {self.len}"


@dataclass
class MapRange:
    dest_start: int
    src_start: int
    len: int

    @staticmethod
    def parse_from_str(line: str) -> "MapRange":
        vals = [int(val) for val in line.split()]
        return MapRange(dest_start=vals[0], src_start=vals[1], len=vals[2])

    def __getitem__(self, src: int) -> int:
        if src not in self:
            raise KeyError(f"{src} not found")
        return self.dest_start + (src - self.src_start)

    def __contains__(self, src: int) -> bool:
        return self.src_start <= src < (self.src_start + self.len)

    @property
    def src_range(self) -> Range:
        return Range(self.src_start, self.len)

    @property
    def dest_range(self) -> Range:
        return Range(self.dest_start, self.len)

    @property
    def src_end(self):
        return self.src_start + self.len

    @property
    def dest_end(self):
        return self.dest_start + self.len

    def __repr__(self) -> str:
        return f"MapRange: {self.dest_start}, {self.src_start}, {self.len}"


@dataclass
class FarmMap:
    title: str
    ranges: list[MapRange]

    @staticmethod
    def parse_from_str(section: str) -> "FarmMap":
        lines = section.splitlines(keepends=False)
        title = lines[0].split()[0]

        ranges = [MapRange.parse_from_str(line) for line in lines[1:]]

        return FarmMap(title, ranges=ranges)

    def __getitem__(self, src: int) -> int:
        for range in self.ranges:
            if src in range:
                return range[src]
        return src

    def __repr__(self) -> str:
        return f"FarmMap: {self.title}, {self.ranges}"

    def get_dest_ranges(
        self,
        src_range: Range,
    ) -> set[Range]:
        # dest_range (0,6) :    0  1  2  3  4  5    associated narrowed dest ranges
        # src_range1 (10,4):   10 11 12 13          -> (0,4)
        # src_range2 (12,3):         12 13 14       -> (2,3)
        # src_range3 (14,3):               14 15 16 -> (4,2)  + partial src(16,1)
        # src_range4 (9,3) : 9 10 11                -> (0,2) + partial src(9,1)
        # src_range5 (9,8) : 9 10 11 12 13 14 15 16 -> (0,6) + partial src(9,1) + partial(16,1)
        dest_ranges: set[Range] = set()

        for map_range in self.ranges:
            if src_range.start in map_range:
                # the beginning of the source range is in the map range
                dest_range_start = map_range[src_range.start]
                max_supported_len = min(
                    map_range.dest_end - dest_range_start, src_range.len
                )
                dest_range = Range(dest_range_start, max_supported_len)
                dest_ranges.add(dest_range)

                if src_range.len > max_supported_len:
                    # the end of the source range extends beyond the map range
                    # split into a separate src_range and recurse
                    partial_src_range = Range(
                        start=src_range.start + max_supported_len,
                        len=src_range.len - max_supported_len,
                    )
                    partial_dest_range = self.get_dest_ranges(partial_src_range)
                    dest_ranges.update(partial_dest_range)
            elif src_range.end in map_range:
                # the beginning of the source range is not in the map range
                # but the end of the source range is in the map range

                dest_range_end = map_range[src_range.end]
                max_supported_len = dest_range_end - map_range.dest_start + 1
                dest_range = Range(map_range.dest_start, max_supported_len)
                dest_ranges.add(dest_range)

                # split into separate src range and recurse
                partial_src_range = Range(
                    start=src_range.start, len=src_range.len - max_supported_len
                )
                partial_dest_range = self.get_dest_ranges(partial_src_range)
                dest_ranges.update(partial_dest_range)
            elif (
                src_range.start < map_range.src_start
                and src_range.end > map_range.src_end
            ):
                # both ends of the source range lie outside the map range
                # but it is possible that some values inside the source range
                # also lie inside the map range

                dest_ranges.add(map_range.dest_range)

                # split into separate ranges for beginning and end and recurse

                # part before map range
                pre_partial_range = Range(
                    start=src_range.start,
                    len=map_range.src_start - src_range.start,
                )
                pre_partial_dest_range = self.get_dest_ranges(pre_partial_range)
                dest_ranges.update(pre_partial_dest_range)

                # part after map range
                post_partial_range = Range(
                    start=map_range.src_end,
                    len=src_range.end - map_range.src_end,
                )
                post_partial_dest_range = self.get_dest_ranges(post_partial_range)
                dest_ranges.update(post_partial_dest_range)

        if not dest_ranges:
            dest_ranges.add(src_range)

        return dest_ranges


def parse_seeds(line: str) -> list[int]:
    seeds = line.split(":")[1].split()
    return [int(seed) for seed in seeds]


def parse_seeds2(line: str) -> list[Range]:
    range_strs = line.split(":")[1].split()

    seeds: list[Range] = []
    for i in range(0, len(range_strs), 2):
        start = int(range_strs[i])
        length = int(range_strs[i + 1])
        seeds.append(Range(start, length))
    return seeds


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


def seed_location_ranges(farm_maps: list[FarmMap], seed_range: Range) -> set[Range]:
    src_ranges: set[Range] = set([seed_range])

    for farm_map in farm_maps:
        dest_ranges: set[Range] = set()

        for src_range in src_ranges:
            curr_dest_ranges = farm_map.get_dest_ranges(src_range)
            dest_ranges.update(curr_dest_ranges)

        src_ranges = dest_ranges

    return src_ranges


def part2(sections: list[str]):
    seed_ranges = parse_seeds2(sections[0])

    farm_maps: list[FarmMap] = []
    for section in sections[1:]:
        farm_map = FarmMap.parse_from_str(section)
        farm_maps.append(farm_map)

    location_ranges = seed_location_ranges(farm_maps, seed_ranges[0])
    min_location = min(location_ranges).start
    for seed_range in seed_ranges[1:]:
        location_ranges = seed_location_ranges(farm_maps, seed_range)
        min_location = min(min_location, min(location_ranges).start)

    print(min_location)


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
