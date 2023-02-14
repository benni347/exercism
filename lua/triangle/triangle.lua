local triangle = {}

function triangle.kind(a, b, c)
  local result
  if a <= 0 or b <= 0 or c <= 0 then error("Input Error") end
  if a + b <= c or b + c <= a or c + a <= b then error("Input Error") end
  if a == b and a == c and b == c then
    result = "equilateral"
  elseif a == b or a == c or b == c then
    result = "isosceles"
  else
    result = "scalene"
  end
  return result
end

return triangle
