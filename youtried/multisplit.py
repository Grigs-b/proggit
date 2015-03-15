'''
Given a string and a list of separators, return the string split
on all the given separators
Example:
input: 'abcdefgh', ['a', 'de', 'g']
output: ['bc', 'f', 'h']
'''


def multisplit(string, delims):

    for delim in delims:
        string = string.replace(delim, ":")
    return [ item for item in string.split(":") if item ]

''' strings used for testing
result = multisplit('abcdefgh', ['a', 'de', 'g'])
print (result)


result = multisplit('this is a test. this is only a test. do not adjust your television', ['.', ' ', 'is', 'es'])
print (result)
'''

if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser(description='Process some integers.')

    parser.add_argument('--string', type=str, help='the string to split')
    parser.add_argument('--delims', type=str, nargs='+',
                        help='delimiters to split the given string on')

    args = parser.parse_args()
    print(multisplit(args.string, args.delims))
