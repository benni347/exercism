# Score categories.
# Change the values as you see fit.
YACHT = 50
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6
FULL_HOUSE = 7
FOUR_OF_A_KIND = 8
LITTLE_STRAIGHT = 30
BIG_STRAIGHT = 31
CHOICE = 0


def score(dice, category):
    if category == YACHT:
        if all(x == dice[0] for x in dice):
            return YACHT
        else:
            return int("0")
    if category == ONES:
        amount = 0
        for i, data in enumerate(dice):
            if data == ONES:
                amount += 1
        return amount * ONES
    if category == TWOS:
        amount = 0
        for i, data in enumerate(dice):
            if data == TWOS:
                amount += 1
        return amount * TWOS
    if category == THREES:
        amount = 0
        for i, data in enumerate(dice):
            if data == THREES:
                amount += 1
        return amount * THREES
    if category == FOURS:
        amount = 0
        for i, data in enumerate(dice):
            if data == FOURS:
                amount += 1
        return amount * FOURS
    if category == FIVES:
        amount = 0
        for i, data in enumerate(dice):
            if data == FIVES:
                amount += 1
        return amount * FIVES
    if category == SIXES:
        amount = 0
        for i, data in enumerate(dice):
            if data == SIXES:
                amount += 1
        return amount * SIXES

    if category == FULL_HOUSE:
        counts = [0] * 7
        for card in dice:
            counts[card] += 1
        if (2 in counts) and (3 in counts):
            three_of_a_kind = [i for i, count in enumerate(counts) if count == 3][0]
            two_of_a_kind = [i for i, count in enumerate(counts) if count == 2][0]
            return (three_of_a_kind * 3) + (two_of_a_kind * 2)
        return 0

    if category == FOUR_OF_A_KIND:
        element_counts = {}
        for element in dice:
            if element in element_counts:
                element_counts[element] += 1
            else:
                element_counts[element] = 1
        for element, count in element_counts.items():
            if count >= 4:
                return 4 * element
        return int("0")

    if category == LITTLE_STRAIGHT:
        if sorted(dice) == [1, 2, 3, 4, 5]:
            return 30
        else:
            return int("0")

    if category == BIG_STRAIGHT:
        if sorted(dice) == [2, 3, 4, 5, 6]:
            return 30
        else:
            return int("0")

    if category == CHOICE:
        value = 0
        for i in range(len(dice)):
            value += dice[i]
        return value
