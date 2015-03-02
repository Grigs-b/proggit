import re

def format(isbn):
    return [c for c in isbn if re.match('[0-9X]', c)]

def validate(isbn):
    if len(isbn) != 10:
        return False
    if isbn[-1] == "X":
        isbn[-1] = 10

    return sum([ int(x)*y for x, y in zip(isbn, range(10, 0, -1))]) % 11 == 0



if __name__ == "__main__":
    raw_isbn = raw_input("Input ISBN to check: ")
    isbn = format(raw_isbn)
    print(isbn)
    print("Validation check on {} : {}".format(raw_isbn, validate(isbn)))
