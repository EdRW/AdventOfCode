from dataclasses import dataclass
from enum import IntEnum
from typing import Callable, TypedDict, cast, Literal, NamedTuple
from collections import Counter, deque
from functools import cmp_to_key

CARD_VAL = {
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

WILD_CARD_VAL = 1


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
    lt_func: Callable[["Hand", "Hand"], bool]

    def __lt__(self, other: "Hand")-> bool:
        return self.lt_func(self, other)
        


@dataclass
class Game:
    card_val: dict[str, int]

    def __init__(self, wild_card: str|None = None):
        self.card_val = {**CARD_VAL}
        if wild_card:
            self.card_val[wild_card] = WILD_CARD_VAL

    def play(self, input_str: str):
        hand_inputs = input_str.splitlines(keepends=False)

        hands: list[Hand] =[]
        for hand_input in hand_inputs:
            cards_bid = hand_input.split()
            cards = [self.card_val[char] for char in cards_bid[0]]
            bid = int(cards_bid[1])
            hands.append(Hand(bid=bid, cards=cards, lt_func=self._compare_hands))
        
        ranked_hands = sorted(hands)

        total_winnings = 0
        for i, hand in enumerate(ranked_hands):
            rank = i + 1
            total_winnings += rank * hand.bid
        return total_winnings

    def _compare_hands(self, hand_a: Hand, hand_b: Hand) -> bool:
        hand_a_type = self._hand_type(hand_a)
        hand_b_type = self._hand_type(hand_b)

        if hand_a_type == hand_b_type:
            # compare individual cards
            for my_card, other_card in zip(hand_a.cards, hand_b.cards):
                if my_card == other_card:
                    continue

                return  my_card < other_card 
            raise ValueError('Hands cannot be the same')

        return hand_a_type < hand_b_type

    def _hand_type(self, hand: Hand):
        card_set = Counter(hand.cards)

        num_wild_cards =card_set.get(WILD_CARD_VAL)
        if num_wild_cards and num_wild_cards != 5:
            card_set.pop(WILD_CARD_VAL)
            most_common = card_set.most_common()
            most_common_key,most_common_count = most_common[0]
            for key,count in most_common[1:]:
                if count == most_common_count and key> most_common_key:
                    most_common_key = key

            card_set.update({most_common_key: num_wild_cards})

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
    game = Game()
    total_winnings = game.play(input_str)
    print(total_winnings)


def part_2(input_str: str):
    game = Game(wild_card='J')
    total_winnings = game.play(input_str)
    print(total_winnings)


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
