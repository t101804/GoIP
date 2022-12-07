package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/alphadose/itogami"
	"github.com/fatih/color"
)

const runTimes uint32 = 10000

var sum uint32
var clear map[string]func()

func init() {
	rand.Seed(time.Now().UnixNano())
	clear = make(map[string]func()) //In	itialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
func n1() {
	var list string
	var thread uint64
	CallClear()
	banner := `
   ______      ________  
  / ____/___  /  _/ __ \
 / / __/ __ \ / // /_/ / 
/ /_/ / /_/ // // ____/  [ Join : https://t.me/DailyToolz ] 
\____/\____/___/_/  GoIP is Free Domain To IPv4 Made by Golang 
		     
`

	scanner := bufio.NewScanner(strings.NewReader(banner))
	for scanner.Scan() {
		c := []int{int(color.FgCyan), int(color.FgGreen), int(color.FgMagenta), int(color.FgRed), int(color.FgYellow), int(color.FgBlue)}
		index := rand.Intn(len(c))
		color.New(color.Attribute(c[index])).Println(scanner.Text())
	}
	CyB := color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	BlB := color.New(color.FgBlue).Add(color.Bold).SprintFunc()
	white := color.New(color.FgWhite).Add(color.Bold).SprintFunc()
	time := fmt.Sprintf("%s%s%s%s%s%s%s ", white("["), CyB(time.Now().Hour()), CyB(":"), CyB(time.Now().Minute()), CyB(":"), CyB(time.Now().Second()), white("]"))
	input := fmt.Sprintf("%s%s%s", white("["), BlB("INPUT-LIST"), white("] > "))
	threadput := fmt.Sprintf("%s%s%s", white("["), BlB("INPUT-THREAD"), white("] > "))
	fmt.Print(time, input)
	fmt.Scanln(&list)
	fmt.Print(time, threadput)
	fmt.Scanln(&thread)
	Runner(list, uint64(thread))

}

func Changer(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("|-> %s\n", scanner.Text())
		ips, _ := net.LookupIP(scanner.Text())
		for _, ip := range ips {
			if ipv4 := ip.To4(); ipv4 != nil {
				fmt.Println("|--> IPv4: ", ipv4)
				w, _ := os.OpenFile("ipv4.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				data := fmt.Sprintf("%s\n", ipv4.String())
				_, _ = w.WriteString(data)
			}
		}
	}
}

func Runner(list string, thread uint64) {
	var wg sync.WaitGroup
	file, _ := os.Open(list)
	defer file.Close()
	// Use the common pool
	start := time.Now()
	pool := itogami.NewPool(thread)

	syncCalculateSum := func() {
		Changer(file)
		wg.Done()
	}
	for i := uint32(0); i < runTimes; i++ {
		wg.Add(1)
		// Submit task to the pool
		pool.Submit(syncCalculateSum)
	}
	wg.Wait()
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("\nfinished all tasks with %s \ndont forget to support by join channel", elapsed)
}

func main() {
	n1()
}
