"Square the sum of the first `n` positive integers"
function square_of_sum(n)
    sum_var = 0
    sum_var = sum(1:n)
    result = sum_var ^ 2
    result
end

"Sum the squares of the first `n` positive integers"
function sum_of_squares(n)
    sum_var = 0
    for i in 1:n
        sum_var += i ^ 2
    end
    sum_var
end

"Subtract the sum of squares from square of the sum of the first `n` positive ints"
function difference(n)
    square_of_sum(n) - sum_of_squares(n)
end
