module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz n
  | n <= 0 = Nothing
  | otherwise = Just (go n 0)
  where
    go :: Integer -> Integer -> Integer
    go num acc
      | num == 1 = acc
      | even num = go (num `div` 2) (acc + 1)
      | otherwise = go (3 * num + 1) (acc + 1)
