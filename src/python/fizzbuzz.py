#!/usr/bin/env python


def main(i: int) -> None:
    for _i in range(0, i):
        print(get_msg(_i + 1))


def get_msg(i: int) -> str:
    """Get the message for integer"""
    msg = str(i)
    if i % 3 == 0:
        msg = "Fizz"
    if i % 5 == 0:
        if msg != str(i):
            msg += "Buzz"
        else:
            msg = "Buzz"

    return msg


if __name__ == "__main__":
    main(15)
