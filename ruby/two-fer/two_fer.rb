def two_fer(name = 'you')
  name = 'you' if name.nil? || name.empty?
  "One for #{name}, one for me."
end

puts two_fer(ARGV[0])

