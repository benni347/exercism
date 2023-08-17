def sum_of_multiples(level: int, factors: list[int]) -> int:
    """
    Calculates the sum of factors below level.

    @level: int which is the highest the factors will check for multiples of themself.
    @factors: list[int] which contains the possible factors.
    @return: int of the sum
    """
    if 0 in factors:
        factors.remove(0)

    multiples = set()

    for factor in factors:
        for multiple in range(factor, level, factor):
            multiples.add(multiple)

    return sum(multiples)
