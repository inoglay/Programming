for i in 0..100
  if i == 0
     puts i 
  elsif i % 15 == 0 then 
     puts "fizzbuzz" 
  elsif i % 3 == 0 then 
     puts "fizz" 
  elsif i % 5 == 0 then 
     puts "buzz" 
  else
     puts i 
  end
end
