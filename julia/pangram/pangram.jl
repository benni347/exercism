"""
    ispangram(input)

Return `true` if `input` contains every alphabetic character (case insensitive).

"""
ispangram(input) = 'a':'z' ⊆ lowercase(input)

