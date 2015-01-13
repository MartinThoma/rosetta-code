require 'benchmark'

number_of_values = 10_000_000

puts 'Creating random integers'

array = Array.new(number_of_values)

number_of_values.times do |i|
  array[i] = rand(0..100_000_000)
end

puts 'Start sorting'

time = Benchmark.measure do
  array.sort!
end

puts "Time to sort: #{time}"
