from collections.abc import Generator


# read file with generator expression
def read_file(filename: str) -> Generator[str, None, None]:
    with open(filename, "r", encoding="utf-8") as file:
        for line in file:
            yield line.strip()


inputPath = "../../input/day3.txt"


def part1() -> int:
    joltages = read_file(inputPath)
    count: int = 0

    for joltage in joltages:
        # largest: int = 0

        # for i in range(len(joltage)):
        #     for j in range(i + 1, len(joltage)):
        #         largest = max(largest, int(joltage[i] + joltage[j]))

        # print(largest)
        # count += largest
        #
        best_tens = int(joltage[0])
        largest = 0

        for i in range(1, len(joltage)):
            current = int(joltage[i])
            largest = max(largest, best_tens * 10 + current)
            best_tens = max(best_tens, current)

        count += largest

    return count


def part2() -> int:
    joltages = read_file(inputPath)
    count: int = 0
    turns = 12

    for joltage in joltages:
        best = [-1] * turns

        for digit in joltage:
            current = int(digit)
            for i in range(turns - 1, -1, -1):
                if i == 0:
                    best[0] = max(best[0], current)
                elif best[i - 1] != -1:
                    best[i] = max(best[i], best[i - 1] * 10 + current)

        count += best[-1]
    return count


if __name__ == "__main__":
    print(part1())
    print(part2())
