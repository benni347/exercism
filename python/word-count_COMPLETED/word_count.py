import re
from typing import Counter


def count_words(sentence):
    sentence = sentence.lower()

    # Tokenize the sentence: Split on whitespace and all punctuation except apostrophes
    words = re.findall(r"\b\w(?:[\w']*\w)?\b", re.sub(r"[_]", " ", sentence))

    # Count the frequency of each word
    word_count = Counter(words)

    return dict(word_count)
