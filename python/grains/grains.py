def square(num):
    if num < 1 or num > 64:
        raise ValueError("square must be between 1 and 64")
    return 2**(num - 1)


def total():
    result = 0
    for i in range(1, 65):
        result += square(i)
    return result
