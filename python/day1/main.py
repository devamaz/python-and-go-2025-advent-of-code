from collections.abc import Generator


# read file with generator expression
def read_file(filename: str) -> Generator[str, None, None]:
    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            yield line.strip()


def part1() -> int:
    moves = read_file("input2.txt")
    length: int = 100
    curr: int = 50
    count: int = 0

    for move in moves:
        value: int = int(move[1:])
        if move[0] == "L":
            curr = (curr - value) % length
        else:
            curr = (curr + value) % length

        if curr == 0:
            count += 1

    return count


def part2() -> int:
    moves = read_file("../input/day1.txt")

    length: int = 100
    curr: int = 50
    count: int = 0

    for move in moves:
        direction = -1 if move[0] == "L" else 1
        value = int(move[1:])

        count += value // length
        value %= length
        for _ in range(value):
            curr = (curr + direction) % length
            count += curr == 0

    return count


if __name__ == "__main__":
    # print(part1())
    print(part2())
