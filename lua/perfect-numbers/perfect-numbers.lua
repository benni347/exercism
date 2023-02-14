local function aliquot_sum(n)
  local sum = 0
  for i = 1, n / 2 do if n % i == 0 then sum = sum + i end end
  return sum
end

local function classify(n)
  local aliquot_sum_var = aliquot_sum(n)
  local result
  if aliquot_sum_var == n then
    result = "perfect"
  elseif aliquot_sum_var > n then
    result = "abundant"
  else
    result = "deficient"
  end
  return result
end

return { aliquot_sum = aliquot_sum, classify = classify }
