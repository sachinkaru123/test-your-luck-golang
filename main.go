package main
import "fmt"
import "math/rand"
import "github.com/common-nighthawk/go-figure"
import "time"


var completed bool = false
var start time.Time


func main() {
    var name string = ""
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    start = time.Now()

    figure.NewFigure("Try Your Luck", "doom" , true).Print()
    fmt.Println("Hello, " + name + "!")

    inits()
    fmt.Println("Goodbye, " + name + "!")
}

func inits() {
    runLoop()
}


func checkIsNumber(number int) bool {
    var randNumber int = rand.Intn(10)
    fmt.Println("Random number generated:", randNumber)
    if (number == randNumber) {
        return true
    }
    return false
}

func runLoop() {
    for i := 1; i < 6; i++ {
        if (completed) {
            fmt.Println("exiting game ....")
            fmt.Println("")
            fmt.Println("")

            break
        }
        fmt.Println("Iteration number:", i)
        if (checkIsNumber(i)) {

            end:=time.Now()
            elapsed:=end.Sub(start)
            
            fmt.Println("")
            fmt.Println("------------congratulations---------------")
            fmt.Println("")
            fmt.Println( "==== You have won.====")
            fmt.Println("Your Lucky number is:", i )
            fmt.Println("Time taken to win:", elapsed)
            completed = true
        } else {
            fmt.Println(i, "is not a valid number.")
        }
    }
    if(!completed) {
            inits()
    }

}