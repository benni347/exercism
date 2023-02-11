def response(hey_bob):
    hey_bob = hey_bob.strip()
    if hey_bob == "":
        return "Fine. Be that way!"
    if not hey_bob.isupper():
        if hey_bob[-1] == "?":
            return "Sure."

    if hey_bob.isupper():
        if hey_bob[-1] == "?":
            return "Calm down, I know what I'm doing!"
        return "Whoa, chill out!"
    return "Whatever."
