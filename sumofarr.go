package main
import "fmt"

func main() {
    array := []int{1, 2, 3, 4, 5}
    sum := 0

    for _, value := range array {
        sum += value
    }

    fmt.Println("Sum of array:",sum)
}
