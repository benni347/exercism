def steps(number):
    if number <= 0:
        raise ValueError("Only positive integers are allowed")
    amount_of_steps = 0
    while number != 1:
        if number % 2 == 0:
            number /= 2
        else:
            number = 3 * number + 1
        amount_of_steps += 1
    return amount_of_steps
