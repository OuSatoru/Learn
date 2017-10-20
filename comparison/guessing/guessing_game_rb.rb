puts 'Guess the number!'
secret_number = rand(101)
puts secret_number
loop do
    puts 'Please input your guess'
    guess_str = gets.chomp
    if guess_str.to_i.to_s != guess_str
        next
    end
    guess = guess_str.to_i
    puts "You guessed #{guess}"
    case
    when guess > secret_number
        puts 'Too big!'
    when guess < secret_number
        puts 'Too small!'
    when guess == secret_number
        puts 'You win!'
        break
    end
end

