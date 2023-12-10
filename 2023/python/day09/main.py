from collections import Counter, deque
from dataclasses import dataclass
from enum import IntEnum
from math import lcm
from typing import Callable, Literal, NamedTuple, TypedDict, cast


def part_1(input_str: str):
    lines = input_str.splitlines(keepends=False)

    prediction_sums = 0
    for line in lines:
        nums = [int(char) for char in line.split()]

        prediction = predict_next(nums)

        prediction_sums += prediction

    print(prediction_sums)


def part_2(input_str: str):
    lines = input_str.splitlines(keepends=False)

    prediction_sums = 0
    for line in lines:
        nums = [int(char) for char in line.split()]

        prediction = predict_prev(nums)

        prediction_sums += prediction

    print(prediction_sums)


def predict_next(nums: list[int]) -> int:
    if sum(nums) == 0:
        return 0

    diffs = calc_diffs(nums)
    return nums[-1] + predict_next(diffs)


def predict_prev(nums: list[int]) -> int:
    if sum(nums) == 0:
        return 0

    diffs = calc_diffs(nums)
    return nums[0] - predict_prev(diffs)


def calc_diffs(nums: list[int]) -> list[int]:
    steps = []
    for i, num in enumerate(nums[1:]):
        steps.append(num - nums[i])
    return steps


def main():
    input_filename = "day09/input.txt"
    # input_filename = "day09/test_input.txt"

    with open(input_filename) as f:
        input_str = f.read()

    part_1(input_str)

    print()

    part_2(input_str)


if __name__ == "__main__":
    main()
