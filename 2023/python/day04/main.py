from dataclasses import dataclass
from typing import Callable, TypedDict, cast
from collections import deque


@dataclass
class Card:
    card_num: int
    winning_numbers: set[int]
    my_numbers: list[int]

    @staticmethod
    def parse_from_str(line: str) -> "Card":
        sections = line.strip().split(":")
        card_num = int(sections[0].split()[1])

        body_sections = sections[1].split("|")
        winning_nums = {int(num) for num in body_sections[0].split()}
        my_nums = [int(num) for num in body_sections[1].split()]

        return Card(card_num, winning_nums, my_nums)

    @property
    def value(self):
        num_matches = self.num_matches
        if num_matches == 0:
            return num_matches
        return 2 ** (num_matches - 1)

    @property
    def num_matches(self):
        nums = []
        for my_num in self.my_numbers:
            if my_num in self.winning_numbers:
                nums.append(my_num)

        return len(nums)


def part1(lines: list[str]):
    total_value = 0
    for line in lines:
        card = Card.parse_from_str(line)
        total_value += card.value

    print(total_value)


def part2(lines: list[str]):
    card_dict: dict[int, Card] = {}
    card_queue: deque[Card] = deque()

    for line in lines:
        card = Card.parse_from_str(line)
        card_dict[card.card_num] = card
        card_queue.append(card)

    num_cards = 0
    while card_queue:
        num_cards += 1
        card = card_queue.popleft()

        for i in range(1, card.num_matches + 1):
            card_queue.append(card_dict[card.card_num + i])

    print(num_cards)


def main():
    input_filename = "day04/input.txt"
    # input_filename = "day04/test_input.txt"

    with open(input_filename) as f:
        lines = f.readlines()
        part1(lines)
        part2(lines)


if __name__ == "__main__":
    main()
