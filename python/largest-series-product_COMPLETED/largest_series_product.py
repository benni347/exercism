def largest_product(series, size):
    if size == 0 or series == "":
        raise ValueError("span must be smaller than string length")
    if len(series) < size:
        raise ValueError("span must be smaller than string length")
    if size < 0:
        raise ValueError("span must not be negative")
    if not series.isdigit():
        raise ValueError("digits input must only contain digits")

    products = []

    for i in range(len(series) - size + 1):
        product = 1
        for j in range(size):
            product *= int(series[i + j])
        products.append(product)

    return max(products)
