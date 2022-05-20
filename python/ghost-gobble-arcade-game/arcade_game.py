"""Functions for implementing the rules of the classic arcade game Pac-Man."""


def eat_ghost(power_pellet_active, touching_ghost):
    """Verify that Pac-Man can eat a ghost if he is empowered by a power pellet.

    :param power_pellet_active: bool - does the player have an active power pellet?
    :param touching_ghost: bool - is the player touching a ghost?
    :return: bool - can the ghost be eaten?
    """
    can_eat = False
    if power_pellet_active and touching_ghost:
        can_eat = True
    return can_eat


def score(touching_power_pellet, touching_dot):
    """Verify that Pac-Man has scored when a power pellet or dot has been eaten.

    :param touching_power_pellet: bool - does the player have an active power pellet?
    :param touching_dot: bool - is the player touching a dot?
    :return: bool - has the player scored or not?
    """
    score_bool = False
    if touching_dot or touching_power_pellet:
        score_bool = True
    return score_bool


def lose(power_pellet_active, touching_ghost):
    """
    Trigger the game loop to end (GAME OVER) when Pac-Man touches a ghost without his power pellet.

    :param power_pellet_active: bool - does the player have an active power pellet?
    :param touching_ghost: bool - is the player touching a ghost?
    :return: bool - has the player lost the game?
    """
    lose_b = False
    if touching_ghost:
        if power_pellet_active is False:
            lose_b = True
    return lose_b


def win(has_eaten_all_dots, power_pellet_active, touching_ghost):
    """Trigger the victory event when all dots have been eaten.

    :param has_eaten_all_dots: bool - has the player "eaten" all the dots?
    :param power_pellet_active: bool - does the player have an active power pellet?
    :param touching_ghost: bool - is the player touching a ghost?
    :return: bool - has the player won the game?
    """
    win_b = False
    lose_b = lose(power_pellet_active, touching_ghost)

    if has_eaten_all_dots and not lose_b:
        if has_eaten_all_dots:
            win_b = True

    return win_b
