function rotate(key::Int, input::Char)
    if input in 'a':'z'
        'a' + (input - 'a' + key) % 26
    elseif input in 'A':'Z'
        'A' + (input - 'A' + key) % 26
    else
        input
    end
end

rotate(key::Int, input::String) = map(input) do x
        rotate(key, x)
    end

for i = 0:26
    @eval begin
        macro $(Symbol("R$(i)_str"))(input)
            rotate($i, input)
        end
    end
end
