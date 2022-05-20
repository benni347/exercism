"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido,
The creator of the Python language: https://en.wikipedia.org/wiki/Guido_van_Rossum
"""

EXPECTED_BAKE_TIME = 40


def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """
    time = EXPECTED_BAKE_TIME - elapsed_bake_time
    # print(time)
    return time


def preparation_time_in_minutes(layer):
    """
    Each layer should take 2 min.
    :param layer:
    :return:
    """
    preparation_time = 2
    time = layer * preparation_time
    # print(time)
    return time


def elapsed_time_in_minutes(layers, elapsed):
    """
    Total elapsed time.
    :param layers:
    :param elapsed:
    :return:
    """
    time = preparation_time_in_minutes(layers)
    time += elapsed
    print(time)
    return time
