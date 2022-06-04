import datetime


def add(moment):
    """
    Moment should look like this: "YYYY-MM-DD HH:MIN:SEC"
    :param moment:
    :return:
    """
    gigasec: int = 10 ** 9
    # add one gigasecond to the moment
    return moment + datetime.timedelta(seconds=gigasec)
