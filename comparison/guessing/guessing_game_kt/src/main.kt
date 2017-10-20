import java.util.Random

fun main(args: Array<String>) {
    println("Guess the number!")
    val secretNumber = (0..101).random()
//    println(secretNumber)
    loop@ while (true) {
        println("Please input your guess")
        val guessStr = readLine()
        val guess = guessStr?.toIntOrNull() ?: continue
        println("You guessed: $guess")
        when (guess.compareTo(secretNumber)) {
            -1 -> println("Too small!")
            1 -> println("Too big!")
            0 -> {
                println("You win!")
                break@loop
            }
        }
    }
}

fun ClosedRange<Int>.random() = Random().nextInt(endInclusive - start) + start