def is_armstrong_number(number):
    power = len(str(number))
    sum_var = sum(int(digit) ** power for digit in str(number))
    return True if sum_var == number else False
