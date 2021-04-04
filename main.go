package main // files that share same package name can access each others functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("")                      // Go SDK method names are capitalized
	fmt.Println("Welcome", "Datou", "!") // Print accepts multiple variables and prints with whitespace
	fmt.Println("")

	// Basics
	explainFunctions()
	explainVariables()
	explainStrings()
	explainLoops()
	explainArraysAndSlices()
	explainMaps()
	explainPointers()
	explainStructs()
	explainInterfaces()
	explainReceivers()
	explainTests()

	// Practical
	explainFiles()
	explainHTTP()
	explainTime()

	// Advanced
	explainChannels()

}

func explainFunctions() {
	fmt.Println("=== FUNCTIONS ===")
	fmt.Println("")

	fmt.Println("Functions can have different signatures")
	fmt.Println(" ", functionWithSingleReturn(""))
	fmt.Print("  ")
	fmt.Println(functionWithMultipleReturns())
	fmt.Println("    s1, s2 := functionWithMultipleReturns()")

	fmt.Println("Go has anonymous functions")
	fmt.Println("  go func(echo string) { fmt.Println(echo) }(\"echo\")")
	func(echo string) { fmt.Println(echo) }("  echo")

	fmt.Println("")
}

func functionWithMultipleReturns() (string, string) {
	return "func functionWithMultipleReturns() (string,int) { return \"\",\"\" }", ""
}

func functionWithSingleReturn(echo string) string {
	return "func functionWithSingleReturn(echo string) string { return echo }"
}

func explainVariables() {

	fmt.Println("=== VARIABLES ===")
	fmt.Println()

	fmt.Println("Default variables values")
	// declare without initiating
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string

	fmt.Println("  defaultInt", defaultInt)
	fmt.Println("  defaultFloat32", defaultFloat32)
	fmt.Println("  defaultFloat64", defaultFloat64)
	fmt.Println("  defaultBool", defaultBool)
	fmt.Println("  defaultString", defaultString)

	fmt.Println()
	fmt.Println("Constants and variables can be declared as")
	const x1 = "const x = \"\" // constant"
	x2 := "x := \"y\" // variable"
	fmt.Println(" ", x1)
	fmt.Println(" ", x2)

	fmt.Println()

	fmt.Println("nil is the zero value in Go representing an uninitialized value")

	fmt.Println()

}

func explainStrings() {
	fmt.Println("=== STRINGS ===")
	fmt.Println()

	str1 := fmt.Sprintf("%s", "String interpolation in Go - fmt.Sprintf(\"%s\", \"\")")
	fmt.Println(str1)

	fmt.Println()
}

func explainLoops() {
	fmt.Println("=== LOOPS ===")
	fmt.Println()

	fmt.Println("Different styles of loops")
	fmt.Println("  Classic loop - for i := 0; i < 10; i++ {}")
	fmt.Print("  ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, "")
	}
	fmt.Println()

	x := 0
	fmt.Println("  While loop - for x < 5 {}")
	fmt.Print("  ")
	for x < 5 {
		fmt.Print(x, "")
		x++
	}
	fmt.Println()

	fmt.Println("  Infinite loop - for {}")

	fmt.Println()
}

func explainArraysAndSlices() {
	fmt.Println("=== ARRAYS ===")
	fmt.Println()

	fmt.Println("Arrays in Go are fixed length")
	strArray := [2]string{"one", "two"}
	fmt.Println("  strArray := [2]string{\"one\", \"two\"} =", strArray)
	fmt.Println()

	fmt.Println("To add elements, you use a slice instead of an array and append to the slice")
	strSlice := []string{"one", "two"}
	strSlice = append(strSlice, "three")
	fmt.Println("  strArray := []string{\"one\", \"two\"}")
	fmt.Println("  strSlice = append(strSlice,\"three\")=", strSlice)
	fmt.Println()

	fmt.Println("Looping through arrays/slices")
	fmt.Println("  for i, value := range strSlice")
	for i, value := range strSlice {
		fmt.Println(" ", i, value)
	}

	fmt.Println()
}

func explainMaps() {
	fmt.Println("=== MAPS ===")
	fmt.Println()

	fmt.Println("Maps in Go are as you would expect")
	m := map[string]string{
		"name":  "value",
		"other": "value2",
	}
	fmt.Println("  m := map[string]string{ \"name\":\"value\", \"other\":\"value2\" }")
	fmt.Println("  yields:", m)
	fmt.Println()

	fmt.Println("Add entry to map")
	m["another"] = "value3"
	fmt.Println("  m[\"another\"] = \"value3\"")
	fmt.Println("  yields:", m)
	fmt.Println()

	fmt.Println("Delete entry to map")
	delete(m, "another")
	fmt.Println("  delete(m,\"another\")")
	fmt.Println("  yields:", m)
	fmt.Println()
}

func explainPointers() {
	fmt.Println("=== POINTERS ===")
	fmt.Println()

	fmt.Println("Go is a pass-by-value language - pointers help to directly access the reference in memory")
	str := "one"
	strPtr := &str
	fmt.Println("  the value of &str:", strPtr)
	fmt.Println("  the value of *ptr:", *strPtr)
	fmt.Println("  the value of *&str:", *&str)

	fmt.Println()
}

func explainStructs() {
	fmt.Println("=== STRUCTS ===")
	fmt.Println()

	fmt.Println("In this example, a new type Person struct is used with properties name (string) and age (int)")
	fmt.Println("person{\"Bob\", 10} = ", person{"Bob", 10})
	fmt.Println("person{name:\"George\"} = ", person{name: "George"})
	fmt.Println("p := newPerson(\"Joe\", 30) = ", newPerson("Joe", 30))

	fmt.Println()
}

func explainInterfaces() {
	fmt.Println("=== INTERFACES ===")
	fmt.Println()

	fmt.Println("To implement an interface in Go, we just need to implement all the methods in the interface.")
	fmt.Println("  type personInterface interface{ printName() }")
	fmt.Println("  type person struct")
	fmt.Println("  person will conform to personInterface as long as it implements printName()")
	fmt.Println("    func (person p) printName() {}")

	fmt.Println()
}

func explainReceivers() {
	fmt.Println("=== RECEIVERS ===")
	fmt.Println()

	fmt.Println("Receivers allow functions to be called on all entities of a type")
	fmt.Println("  func (p person) printName() {} can be called to all person types")
	fmt.Println("  p := person{name: \"Jose\"}")
	fmt.Print("  p.printName() yields ")
	p := person{name: "Jose"}
	p.printName()

	fmt.Println()
}

func explainTests() {
	fmt.Println("=== TESTING ===")
	fmt.Println()

	fmt.Println("Run command 'go test' - func TestPersonAdultCheck(t *testing.T)")

	fmt.Println()
}

func explainFiles() {
	fmt.Println("=== FILES ===")
	fmt.Println()

	filename := "/tmp/hello.txt"

	fmt.Println("Writing to a file")
	fmt.Println("  ioutil.WriteFile(\"/tmp/hello.txt\", []byte(\"hello\"), 0666)")
	ioutil.WriteFile(filename, []byte("hello"), 0666)
	fmt.Println()

	fmt.Println("Fetch value from file")
	fmt.Println("  ioutil.ReadFile(\"/tmp/hello.txt\")")
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("  Fetched value from file - ", string(bs))
	fmt.Println()

	fmt.Println("Delete file value from file")
	fmt.Println("  os.Remove(\"/tmp/hello.txt\")")
	os.Remove(filename)
	fmt.Println()

	fmt.Println()
}

func explainHTTP() {
	fmt.Println("=== HTTP ===")
	fmt.Println()

	fmt.Println("Http call to google")
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("  Status:", resp.Status)
	//fmt.Println(resp)

	fmt.Println()
}

func explainTime() {
	fmt.Println("=== TIME ===")
	fmt.Println()

	fmt.Println("Example of sleep 5 millisecond")
	fmt.Println("  time.Sleep(5 * time.Millisecond)")
	time.Sleep(5 * time.Millisecond)

	fmt.Println()

}

func explainChannels() {
	fmt.Println("=== CHANNELS ===")
	fmt.Println()

	fmt.Println("Channels are used to send and receive values with the channel operator")
	fmt.Println("  ch := make(chan string) // create channel")
	fmt.Println("  ch <- v // send message")
	fmt.Println("  v := <- ch // receive message")

	ch := make(chan string)                       // make is used to instantiate slices, maps, channels
	go func(ch chan string) { ch <- "hello" }(ch) // run as a concurrent goroutine
	v := <-ch
	fmt.Println("  v received -", v)
	fmt.Println()

	fmt.Println("Channels can be kept open through a range loop")
	fmt.Println("  for msg := range ch")
	ch2 := make(chan string)
	go func(ch2 chan string) { ch2 <- "hello" }(ch2)
	for msg := range ch2 {
		fmt.Println(" ch2 received -", msg)
		go close(ch2)
	}
	fmt.Println()
}
