from dataclasses import dataclass
from enum import IntEnum
from typing import Callable, TypedDict, cast, Literal, NamedTuple
from collections import deque

Card = {
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "J": 11,
    "Q": 12,
    "K": 13,
    "A": 14,
}


class HandType(IntEnum):
    HIGH_CARD = 1
    ONE_PAIR = 2
    TWO_PAIR = 3
    THREE_OF_KIND = 4
    FULL_HOUSE = 5
    FOUR_OF_KIND = 6
    FIVE_OF_KIND = 7


@dataclass
class Hand:
    bid: int
    cards: list[int]
    type: HandType

    def __init__(self, bid: int, cards: list[int]) -> None:
        self.bid = bid
        self.cards = cards
        self.type = self.determine_type()

    @staticmethod
    def parse_hands_races_str(input_str: str) -> "Hand":
        cards_bid = input_str.split()
        cards = [Card[char] for char in cards_bid[0]]
        bid = int(cards_bid[1])
        return Hand(bid, cards)

    def __lt__(self, other: "Hand") -> bool:
        if self.type == other.type:
            # compare individual cards
            for my_card, other_card in zip(self.cards, other.cards):
                if my_card == other_card:
                    continue

                return my_card < other_card

        return self.type < other.type

    def determine_type(self):
        card_set = {}
        for card in self.cards:
            card_set[card] = card_set.setdefault(card, 0) + 1
        set_len = len(card_set)

        if set_len == 1:
            # five of a kind
            return HandType.FIVE_OF_KIND

        if set_len == 2:
            # four of a kind or full house

            if 4 in card_set.values():
                # four of a kind
                return HandType.FOUR_OF_KIND

            # full house
            return HandType.FULL_HOUSE

        if set_len == 3:
            # three of a kind or two pair
            if 3 in card_set.values():
                return HandType.THREE_OF_KIND

            # two pair
            return HandType.TWO_PAIR

        if set_len == 4:
            # one pair
            return HandType.ONE_PAIR

        # high card
        return HandType.HIGH_CARD


def part_1(input_str: str):
    hand_inputs = input_str.splitlines(keepends=False)

    hands = [Hand.parse_hands_races_str(hand_input) for hand_input in hand_inputs]

    ranked_hands = sorted(hands)
    # [print(hand) for hand in ranked_hands]

    total_winnings = 0
    for i, hand in enumerate(ranked_hands):
        rank = i + 1
        total_winnings += rank * hand.bid
    print(total_winnings)


def part_2(input: str):
    pass


def main():
    input_filename = "day07/input.txt"
    # input_filename = "day07/test_input.txt"

    input_str: str
    with open(input_filename) as f:
        input_str = f.read()

    part_1(input_str)
    part_2(input_str)


if __name__ == "__main__":
    main()
