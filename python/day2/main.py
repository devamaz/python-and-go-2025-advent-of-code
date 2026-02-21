def read_input(filename: str) -> str:
    with open(filename, "r", encoding="utf-8") as f:
        return f.read()


inputPath = "../../input/day2.txt"


def part1() -> int:
    input = read_input(inputPath)
    invalid_ids_sum: int = 0

    for range_str in input.split(","):
        start_str, end_str = range_str.split("-")
        start, end = int(start_str), int(end_str)

        for r in range(start, end + 1):
            id = str(r)
            mid = len(id) // 2
            if id[:mid] == id[mid:]:
                invalid_ids_sum += r

    return invalid_ids_sum


def part2() -> int:
    input = read_input(inputPath)
    invalid_ids_sum: int = 0

    for range_str in input.split(","):
        start_str, end_str = range_str.split("-")
        start, end = int(start_str), int(end_str)

        for r in range(start, end + 1):
            id = str(r)
            if id in (id + id)[1:-1]:
                invalid_ids_sum += r

    return invalid_ids_sum


if __name__ == "__main__":
    print(part1())
    print(part2())
