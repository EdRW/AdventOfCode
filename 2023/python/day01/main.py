from dataclasses import dataclass
from typing import Callable


number_dict = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

word_nums_lens = [len(key) for key in number_dict.keys()]
MIN_NUM_WORD_SIZE = min(word_nums_lens)  # 3
MAX_NUM_WORD_SIZE = max(word_nums_lens)  # 5


@dataclass
class DigitStr:
    """
    Holds a calibration value (represented as digit string value)
    and the index it was found at. DigitStr are sortable by index by default.
    """

    index: int
    value: str

    def __str__(self) -> str:
        return f"(i: {self.index}, val: {self.value})"

    def __repr__(self) -> str:
        return self.__str__()

    def __lt__(self, other: "DigitStr") -> bool:
        return self.index < other.index

    @staticmethod
    def from_word_num(index: int, word: str):
        return DigitStr(index, word_to_digit(word))


def word_to_digit(substring: str) -> str:
    return number_dict[substring]


def is_word_num(substring: str) -> bool:
    return substring in number_dict


def find_word_num(line: str, start: int):
    """Search the next few characters for a word num like 'six'"""

    max_end = min(start + MAX_NUM_WORD_SIZE + 1, len(line))

    for end in range(start + MIN_NUM_WORD_SIZE, max_end):
        substring = line[start:end]
        if is_word_num(substring):
            # substring is a value like "one" or "seven"
            return DigitStr.from_word_num(start, substring)
    return None


def find_numbers(line: str, include_word_nums=False):
    digits: list[DigitStr] = []

    for i, char in enumerate(line):
        # part 1
        if char.isnumeric():
            digits.append(DigitStr(i, char))

        # part 2
        elif include_word_nums and (char_num := find_word_num(line, i)):
            digits.append(char_num)

    # sort
    digits.sort()
    return digits


def calculate_calibration(lines: list[str], find_nums: Callable[[str], list[DigitStr]]):
    calibration_sum = 0
    for line in lines:
        digits = find_nums(line)

        first_num = digits[0]
        last_num = digits[len(digits) - 1]

        calibration_value = first_num.value + last_num.value
        calibration_sum += int(calibration_value)

        # values = [digit.value for digit in digits]
        # print(f"{values} = {calibration_value}")
    return calibration_sum


def main():
    input_filename = "day01/input.txt"

    with open(input_filename) as f:
        lines = f.readlines()

    calibration_sum= calculate_calibration(lines, lambda line : find_numbers(line))
    print(f'part 1: {  calibration_sum}')

    calibration_sum=  calculate_calibration(lines, lambda line : find_numbers(line, True))
    print(f'part 2: {  calibration_sum}')


if __name__ == "__main__":
    main()
