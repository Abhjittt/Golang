package main
import "fmt"

func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    var num = 5

    result := factorial(num)

    fmt.Println("Factorial:",result)
}
