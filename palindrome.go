package main
import "fmt"

func main() {
    s := "helllleh"
    c := 0

    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			c++
		}
	}

	if (c != 0) {
	    fmt.Println(false)
	} else {
	    fmt.Println(true)
	}
}
