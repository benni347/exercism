def transform(legacy_data):
    d = {}
    for score, letters in legacy_data.items():
        for letter in letters:
            d[letter.lower()] = score
    return d
