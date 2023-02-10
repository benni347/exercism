def is_triangle_inequality_satisfied(a, b, c):
    """
    Check if the triangle inequality holds for the given triangle sides.
    
    Args:
    a (int): Length of side 1.
    b (int): Length of side 2.
    c (int): Length of side 3.
    
    Returns:
    bool: True if the triangle inequality holds, False otherwise.
    
    Example:
    >>> is_triangle_inequality_satisfied(3, 4, 5)
    True
    >>> is_triangle_inequality_satisfied(1, 1, 3)
    False
    """
    if a + b > c and a + c > b and b + c > a:
        return True
    else:
        return False


def equilateral(sides):
    """
    Check if a triangle is equilateral.

    Args:
    sides (list[int]): List of length 3 representing the sides of the triangle.

    Returns:
    bool: True if the triangle is equilateral, False otherwise.

    Example:
    >>> equilateral([3, 3, 3])
    True
    >>> equilateral([3, 4, 4])
    False
    """
    return sides[0] > 0 and sides[1] > 0 and sides[2] > 0 and sides[0] == sides[1] == sides[2]



def isosceles(sides):
    """
    Check if a triangle is isosceles.

    Args:
    sides (list[int]): List of length 3 representing the sides of the triangle.

    Returns:
    bool: True if the triangle is isosceles, False otherwise.

    Example:
    >>> is_isosceles_triangle([3, 4, 4])
    True
    >>> is_isosceles_triangle([3, 4, 5])
    False
    """
    if sides[0] > 0 and sides[1] > 0 and sides[2] > 0:
        if is_triangle_inequality_satisfied(sides[0], sides[1], sides[2]):
            return True if sides[0] == sides[1] or sides[0] == sides[2] or sides[1] == sides[2] else False
        return False
    return False


def scalene(sides):
    """
    Check if a triangle is scalene.

    Args:
    sides (list[int]): List of length 3 representing the sides of the triangle.

    Returns:
    bool: True if the triangle is scalene, False otherwise.

    Example:
    >>> scalene([3, 4, 5])
    True
    >>> scalene([3, 3, 4])
    False
    """
    if is_triangle_inequality_satisfied(sides[0], sides[1], sides[2]):
        return sides[0] > 0 and sides[1] > 0 and sides[2] > 0 and sides[0] != sides[1] != sides[2] and sides[0] != sides[2]
    return False
