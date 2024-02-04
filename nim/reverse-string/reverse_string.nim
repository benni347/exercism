proc reverse*(s: string): string =
  var reversed: string = ""
  for i in countdown(s.len - 1, 0):
    reversed.add(s[i])
  return reversed
