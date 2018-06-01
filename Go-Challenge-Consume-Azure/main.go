package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

//Task ...
type Task struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description"`
	Label       string   `json:"label,omitempty"`
	User        string   `json:"user,omitempty"`
	IsComplete  bool     `json:"iscomplete,omitempty"`
	Duedate     *Duedate `json:"duedate,omitempty"`
}

//Duedate ...
type Duedate struct {
	Day   string `json:"day,omitempty"`
	Month string `json:"month,omitempty"`
}

//ResponseObject ...
var ResponseObject []Task

//PrintMenu ...
func PrintMenu() {

	a := color.New(color.FgBlue, color.Bold)
	b := color.New(color.FgCyan)
	c := color.New(color.FgGreen)
	d := color.New(color.FgRed)
	e := color.New(color.FgMagenta)

	//whiteBackground := red.Add(color.BgGreen)
	//whiteBackground.Println("Red text with white background.")

	color.Yellow("OPTIONS-----------------------------------")
	fmt.Printf("\n")
	a.Println("1.- List Tasks")
	b.Println("2.- List Single Task")
	c.Println("3.- Create Task")
	d.Println("4.- Delete Task")
	e.Println("0.- Exit")
	fmt.Printf("\n")
	color.Yellow("------------------------------------------")

}

//Options ...
func Options() {

	var input int
	fmt.Printf("\nSelect an option: ")

	n, err := fmt.Scanln(&input)

	if n < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}

	switch input {
	case 1:
		ListTasks()
		fmt.Print("Press 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		CallClear()
		PrintMenu()
	case 2:
		ListSingleTask()
		time.Sleep(2 * time.Second)
		fmt.Print("Press 'Space' and 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\040')
		CallClear()
		PrintMenu()
	case 3:
		CreateTask()
		time.Sleep(2 * time.Second)
		fmt.Print("Press 'Space' and 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\040')
		CallClear()
		PrintMenu()
	case 4:
		DeleteTask()
		time.Sleep(2 * time.Second)
		fmt.Print("Press 'Space' and 'Enter' to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\040')
		CallClear()
		PrintMenu()
	case 0:
		CallClear()
		os.Exit(2)
	default:
		fmt.Println("def")
	}
}

//*********************Clear Function******************************

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
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

//CallClear ...
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

//************************************************************************************

//ListTasks ...
func ListTasks() {

	response, err := http.Get("https://go-poncho.azurewebsites.net/tasks")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &ResponseObject)

	//fmt.Println(len(ResponseObject))

	fmt.Printf("\n\n************************************\n\n")

	for _, item := range ResponseObject {
		fmt.Println("Id: ", item.ID)
		fmt.Println("Title: ", item.Title)
		fmt.Println("Description: ", item.Description)
		fmt.Println("Label: ", item.Label)
		fmt.Println("User: ", item.User)
		fmt.Println("Completed: ", item.IsComplete)
		fmt.Println("Day: ", item.Duedate.Day)
		fmt.Println("Month: ", item.Duedate.Month)
		fmt.Printf("\n************************************\n")
	}
}

//ListSingleTask ...
func ListSingleTask() {

	response, err := http.Get("https://go-poncho.azurewebsites.net/tasks")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &ResponseObject)

	var opc int

	fmt.Printf("\nPending Tasks: \n\n")
	for _, item := range ResponseObject {
		fmt.Printf("Id: %v  ---  Title: %v\n", item.ID, item.Title)
	}

	fmt.Printf("\n\nSelect a task: ")
	fmt.Scanf("%d", &opc)

	opc = opc - 1

	fmt.Printf("\n************************************\n\n")
	fmt.Println("Id: ", ResponseObject[opc].ID)
	fmt.Println("Title: ", ResponseObject[opc].Title)
	fmt.Println("Description: ", ResponseObject[opc].Description)
	fmt.Println("Label: ", ResponseObject[opc].Label)
	fmt.Println("User: ", ResponseObject[opc].User)
	fmt.Println("Completed: ", ResponseObject[opc].IsComplete)
	fmt.Println("Day: ", ResponseObject[opc].Duedate.Day)
	fmt.Println("Month: ", ResponseObject[opc].Duedate.Month)
	fmt.Printf("\n************************************\n")

}

//CreateTask ...
func CreateTask() {

	response11, err11 := http.Get("https://go-poncho.azurewebsites.net/tasks")

	if err11 != nil {
		fmt.Print(err11.Error())
		os.Exit(1)
	}

	data, err22 := ioutil.ReadAll(response11.Body)

	if err22 != nil {
		log.Fatal(err22)
	}

	json.Unmarshal(data, &ResponseObject)

	fmt.Printf("\nPending Tasks: \n\n")
	for _, item := range ResponseObject {
		fmt.Printf("Id: %v  ---  Title: %v\n", item.ID, item.Title)
	}

	fmt.Printf("\n************************************\n\n")

	index := len(ResponseObject) + 1
	indexString := strconv.Itoa(index)
	values := []string{"https://go-poncho.azurewebsites.net/tasks", indexString}
	url := strings.Join(values, "/")
	fmt.Println(url)

	fmt.Printf("\n************************************\n\n")

	var (
		title       string
		description string
		label       string
		labelInt    int
		user        string
		completed   bool
		day         string
		month       string
		conti       = false
	)

	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Insert title: ")
	//fmt.Scanf("%s", &title)
	title, err0 := in.ReadString('\n')
	title = strings.Replace(title, "\n", "", -1)
	time.Sleep(1 * time.Second)

	fmt.Printf("Insert description: ")
	description, err1 := in.ReadString('\n')
	description = strings.Replace(description, "\n", "", -1)
	time.Sleep(1 * time.Second)

	for conti == false {
		fmt.Printf("Insert label (Family - 1  Work - 2  Personal - 3): ")
		fmt.Scanf("%d", &labelInt)
		time.Sleep(1 * time.Second)
		fmt.Print("")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		switch labelInt {
		case 1:
			label = "Family"
			conti = true
		case 2:
			label = "Work"
			conti = true
		case 3:
			label = "Personal"
			conti = true
		default:
			fmt.Println("Insert a valid option")
		}
	}

	fmt.Printf("Insert user: ")
	user, err2 := in.ReadString('\n')
	user = strings.Replace(user, "\n", "", -1)
	time.Sleep(1 * time.Second)

	completed = false

	fmt.Printf("Insert day: ")
	day, err3 := in.ReadString('\n')
	day = strings.Replace(day, "\n", "", -1)
	time.Sleep(1 * time.Second)

	fmt.Printf("Insert month: ")
	month, err4 := in.ReadString('\n')
	month = strings.Replace(month, "\n", "", -1)
	time.Sleep(1 * time.Second)

	fmt.Printf("\n\n")

	if err0 != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Error en la lectura")
	}

	ind := strconv.Itoa(index)

	task0 := &Task{ID: ind, Title: title, Description: description, Label: label, User: user, IsComplete: completed, Duedate: &Duedate{Day: day, Month: month}}
	jsonData, err := json.Marshal(task0)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post(url, "appication/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
	} else {
		data, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			log.Fatal(err)
		}
		json.Unmarshal(data, &ResponseObject)
		fmt.Println(string(data))
	}
}

//DeleteTask ...
func DeleteTask() {

	response, err := http.Get("https://go-poncho.azurewebsites.net/tasks")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err2 := ioutil.ReadAll(response.Body)

	if err2 != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &ResponseObject)

	fmt.Println(len(ResponseObject))

	var opc int

	fmt.Printf("\nPending Tasks: \n\n")
	for _, item := range ResponseObject {
		fmt.Printf("Id: %v  ---  Title: %v\n", item.ID, item.Title)
	}

	fmt.Printf("\n\nSelect a task for delete: ")
	fmt.Scanf("%d", &opc)

	indexString := strconv.Itoa(opc)
	values := []string{"https://go-poncho.azurewebsites.net/tasks", indexString}
	url := strings.Join(values, "/")
	fmt.Println(url)

	fmt.Printf("\n************************************\n\n")

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Print("Error trying to Delete")
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

}

//************************************************************************************

func main() {

	PrintMenu()

	for {
		Options()
	}

}
