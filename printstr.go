package main
import "fmt"

func main() {
    var in string

    fmt.Print("Enter something: ")
    fmt.Scanln(&in)

    fmt.Printf("You entered: %s\n", in)
}

