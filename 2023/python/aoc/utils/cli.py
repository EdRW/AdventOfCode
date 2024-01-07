from argparse import ArgumentParser, Namespace
import os
from typing import TypedDict

DAY = os.getenv("DAY")


class Options(TypedDict):
    test: bool | None
    day: int 


def parse_args(args: list[str]) -> Options:
    parser = ArgumentParser()
    parser.add_argument(
        "-t",
        "--test",
        action="store_true",
    )
    parser.add_argument("-d", "--day", type=int, default=DAY, required=(DAY is None))

    parsed_args = parser.parse_args(args)
    options = vars(parsed_args)
    return Options({"test": options["test"], "day": options["day"]})

def day_dir():
    return  f'day{DAY}'   