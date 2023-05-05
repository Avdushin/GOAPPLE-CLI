package utils

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	version = "2.0"
	// menu translate
	// ru
	selectru = "ВЫБРАТЬ"
	backru   = "НАЗАД"
	quitru   = "ВЫХОД"
	// en
	Selecten = "SELECT"
)

// ru app
func RuApp() {
	fmt.Printf("\n \033[36m1)\033[37m ДИСТРИБУТИВЫ  \033[36m2)\033[37m МОЯ ОС\n\n \033[36m3)\033[37m НАСТРОЙКИ \033[36m\033[36m    0)\033[31m %s\n\n\033[0m", quitru)

	var scan string

	fmt.Printf("\033[33m %s: \033[0m", selectru)
	fmt.Scan(&scan)

	switch scan {
	case "0", "exit", "EXIT", "q", "quit", "QUIT", "выход", "ВЫХОД":
		os.Exit(0)
	case "1":
		Clear()
		Logo()
		RuDm()
	case "2":
		Clear()
		MyOsRu()
		RuApp()
	case "3":
		Clear()
		Logo()
		Rusettings()
	case "en", "eng":
		Clear()
		Logo()
		MainMenu()
	default:
		Clear()
		Logo()
		RuApp()
	}
}

// Russian distro menu
func RuDm() {
	fmt.Printf("\n \033[36m1)\033[37m MANJARO  \033[36m2)\033[37m SOLUS\n\n \033[36m3)\033[37m FEDORA \033[36m\033[36m \033[36m4)\033[37m Default \033[36m\033[36m \n\n \033[36m5)\033[33m %s \033[36m0)\033[31m %s\n\n\033[0m", backru, quitru)

	var scan string

	fmt.Printf("\033[33m %s: \033[0m", selectru)
	fmt.Scan(&scan)

	switch scan {
	case "0", "exit", "EXIT", "q", "quit", "QUIT", "выход", "ВЫХОД":
		os.Exit(0)
	case "1":
		RuDistroScript(
			"Manjaro",
			"src/distros/Manjaro/manjaro.sh",
		)
	case "2":
		RuDistroScript(
			"Solus",
			"src/distros/Solus/solus.sh",
		)
	case "3":
		RuDistroScript(
			"Fedora",
			"src/distros/Fedora/fedora.sh",
		)
	case "4":
		RuDistroScript(
			"GOAPPLE",
			"src/distros/setup.sh",
		)
	case "5":
		Clear()
		Logo()
		RuApp()
	default:
		Clear()
		Logo()
		RuDm()
	}
}

// Russsian settings menu
func Rusettings() {
	fmt.Printf("\n \033[33mАВТОР:\033[0m https://github.com/Avdushin\n \033[33mВЕРСИЯ: \033[0m%sgo\033[36m \n\n \033[33mЯЗЫК:\n\n \033[36m1)\033[0m РУССКИЙ \n \033[36m2)\033[0m ENGLISH \n\n \033[36m5) \033[33m%s \033[36m0)\033[31m %s\n\n", version, backru, quitru)

	var scan string

	fmt.Printf("\n\033[33m %s: \033[0m", selectru)
	fmt.Scan(&scan)

	switch scan {
	case "0", "exit", "EXIT", "q", "quit", "QUIT", "выход", "ВЫХОД":
		os.Exit(0)
	case "1":
		Clear()
		Logo()
		RuApp()
	case "2":
		Clear()
		Logo()
		MainMenu()
	case "5":
		Clear()
		Logo()
		RuApp()
	default:
		Clear()
		Logo()
		Rusettings()
	}
}

// Info about current OS Russian edition
func MyOsRu() {
	cmd := exec.Command("neofetch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Не удалось запустить Neofetch. %s\nУстановка neofetch...\n", err.Error())
		NeofetchRu()
	}
}

// neofetch menu Russian edition
func NeofetchRu() {
	cmd := exec.Command("sh", "src/neofetch.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
	Clear()
	MyOs()
}
