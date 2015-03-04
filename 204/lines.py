

if __name__ == "__main__":
    find = raw_input("Enter line to find in Macbeth: ")
    lines = []
    with open('204/macbeth.txt', 'r') as f:
        lines = f.readlines()

    # we're popping from the list, so reverse it
    lines.reverse()
    toparse = []
    building = ""
    while lines:
        current = lines.pop()
        while lines and current.startswith("    "):
            building = "".join([building, current.lstrip(" ")])
            current = lines.pop()

        toparse.append(building)
        building = ""

    results = [block for block in toparse if find in block]
    for result in results:
        print(result)
