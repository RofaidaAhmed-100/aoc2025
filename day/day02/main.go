package main
import(
	"bufio"
    "fmt"
    "log"
    "os"
    "strings"
)
import "strconv"


func isInvalid(num int) bool {
    s := strconv.Itoa(num)  
    n := len(s)
    if n%2 != 0 {                
        return false
    }
    half := n / 2
    return s[:half] == s[half:] 
}

func main(){
file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    scanner.Scan() 
    line := scanner.Text()

    ranges := strings.Split(line, ",")
    result :=0
	for _, r := range ranges {
    parts := strings.Split(r, "-")
    first := 0
    last := 0
    fmt.Sscanf(parts[0], "%d", &first)
    fmt.Sscanf(parts[1], "%d", &last)
	for num := first; num <= last; num++ {
        if isInvalid(num) {
            //fmt.Println(num) 
			result += num
        }
    }
}
fmt.Println(" sum of invalid IDs:", result)
	
}