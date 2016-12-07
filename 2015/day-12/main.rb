require 'json'

input = JSON.parse(File.read('input.json'))

def sum_item(input)
  if input.is_a?(Array)
    sum = 0
    input.each do |i|
      sum += sum_item(i)
    end
    sum
  elsif input.is_a?(Hash)
    sum = 0
    input.values.each do |v|
      return 0 if v == 'red'
      sum += sum_item(v)
    end
    sum
  else
    return input.to_i
  end
end

puts sum_item(input)
