from collections import Counter, deque
from dataclasses import dataclass
from enum import IntEnum
from math import lcm
from typing import Callable, Literal, NamedTuple, TypedDict, cast


class Children(TypedDict):
    L: str
    R: str


Network = dict[str, Children]


def network_from_str(lines: list[str]):
    # print(lines)
    network: Network = {}
    for line in lines:
        parts = line.split("=")
        node = parts[0].strip()
        children_parts = parts[1].split(",")
        left_child = children_parts[0].strip()[1:]
        right_child = children_parts[1].strip()[:-1]
        children = Children(L=left_child, R=right_child)
        network[node] = children
    return network


def part_1(input_str: str):
    lines = input_str.splitlines(keepends=False)
    instr_str = lines[0]

    network = network_from_str(lines[2:])

    step = 0
    node = "AAA"
    while node != "ZZZ":
        instruction = instr_str[step % len(instr_str)]
        node = network[node][instruction]
        step += 1
    print(step)


def part_2(input_str: str):
    lines = input_str.splitlines(keepends=False)
    instr_str = lines[0]

    network = network_from_str(lines[2:])

    nodes = [node for node in network if node.endswith("A")]
    
    steps: list[int] =[]
    for node in nodes:
        step = 0
        while not node.endswith("Z"):
            instruction = instr_str[step % len(instr_str)]
            node = network[node][instruction]
            step += 1
        print(step)
        steps.append(step)

    print(lcm(*steps))


            
            


def main():
    input_filename = "day08/input.txt"
    # input_filename = "day08/test_input.txt"

    with open(input_filename) as f:
        input_str = f.read()

    part_1(input_str)
    print()

    # input_filename = "day08/test_input_2.txt"

    with open(input_filename) as f:
        input_str = f.read()

    part_2(input_str)


if __name__ == "__main__":
    main()
