return function(n)
	if n <= 0 then
		error("n must be a positive integer greater than zero")
	end

	local S = { 2 } -- initialize the set S to {2}
	local p = 2 -- initialize the largest prime number in S to 2

	while #S < n do -- repeat until |S| = n
		p = p + 1 -- increment p by 1
		local is_prime = true -- flag to check if p is prime

		-- iterate through each prime number q in S that is less than or equal to the square root of p
		for i = 1, #S do
			local q = S[i]
			if q > math.sqrt(p) then
				break -- all possible factors checked
			end
			if p % q == 0 then -- p is not prime if it is evenly divisible by q
				is_prime = false
				break
			end
		end

		if is_prime then
			table.insert(S, p) -- add p to S if it is prime
		end
	end

	return p -- return the nth prime number
end
