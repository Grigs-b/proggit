'''
Goal: Take a list of integers of positive and negative numbers and return
the pair of numbers who's summation is closest to 0

Example:
input: [-12, 6, 8, -10, -5, 14, -2]
output: 6, -5
'''

#sorted [-12, -10, -5, -2, 6, 8, 14]
import sys

def magnitude(a, b):
    return abs(abs(a) - abs(b))

def list_pairs(dataset):
    dataset.sort()
    end = len(dataset) - 1
    start = 0
    result = (start, end)
    smallest = sys.maxint
    check = smallest
    #shortcuts if we have lists of all pos/neg numbers. not required, but fewer checks
    if dataset[start] > 0:
        return (dataset[start], dataset[start+1])
    elif dataset[end] < 0:
        return (dataset[end-1], dataset[end])

    while start < end and check != 0:

        check = dataset[start] + dataset[end]

        if abs(check) < abs(smallest):
            smallest = check
            result = (dataset[start], dataset[end])
        if abs(dataset[start]) > abs(dataset[end]):
            start += 1
        else:
            end -= 1



    return result



if __name__ == "__main__":
    data = [-12, 6, 8, -10, -5, 14, -2]
    print('Pairs for {}: {}'.format(data, list_pairs(data)))
    pos = [12, 6, 8, 10, 5, 14, 2]
    print('Pairs for {}: {}'.format(pos, list_pairs(pos)))
    neg = [-12, -6, -8, -10, -5, -14, -2]
    print('Pairs for {}: {}'.format(neg, list_pairs(neg)))
    #local testing /repeatability
    import random
    #random.seed(9001)
    ran = [ random.randint(-100, 100) for x in xrange(15)]
    print('Pairs for {}: {}'.format(ran, list_pairs(ran)))
