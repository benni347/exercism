def factors(n):
    return set(
        x
        for tup in ([i, n // i] for i in range(1, int(n**0.5) + 1) if n % i == 0)
        for x in tup
    )


def classify(number):
    """A perfect number equals the sum of its positive divisors.

    :param number: int a positive integer
    :return: str the classification of the input integer
    """
    if number <= 0:
        raise ValueError(
            "Classification is only possible for positive integers.")
    factors_of_number = factors(number)
    factors_of_number.discard(number)
    sum_of_factors = 0
    for i, data in enumerate(factors_of_number):
        sum_of_factors += data
    solution = ""
    if sum_of_factors == number:
        solution = "perfect"
    elif sum_of_factors > number:
        solution = "abundant"
    else:
        solution = "deficient"
    return solution
