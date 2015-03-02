
BYTE_SIZE = 8

def get_char_from_binary(bits):
    return chr(int(bits, 2))

if __name__ == "__main__":
    data = raw_input("PLEASE INSERT GIRDER: ")
    result = "".join([ get_char_from_binary(data[BYTE_SIZE*i:BYTE_SIZE*(i+1)]) for i in range(len(data)/BYTE_SIZE)])
    print(result)
