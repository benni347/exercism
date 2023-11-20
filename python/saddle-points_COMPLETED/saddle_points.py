def saddle_points(matrix):
    """Find the saddle points of a matrix."""
    # if the matrix is iregular
    if any(len(row) != len(matrix[0]) for row in matrix):
        raise ValueError("irregular matrix")
    saddle_points = []
    for i in range(len(matrix)):
        for j in range(len(matrix[i])):
            value = matrix[i][j]
            if value == max(matrix[i]) and value == min([row[j] for row in matrix]):
                saddle_points.append({"row": i + 1, "column": j + 1})
    return saddle_points
