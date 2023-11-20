module Hamming (distance) where

-- | Calculate the Hamming Distance between two DNA strands.
distance :: String -> String -> Maybe Int
distance xs ys
    | length xs /= length ys = Nothing
    | otherwise = Just $ length $ filter id $ zipWith (/=) xs ys
