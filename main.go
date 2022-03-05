package main

import (
	"fmt" // standart fmt
	"os"
	"os/exec" // to execute shell commands

	"github.com/mbndr/figlet4go" // figlet logo
)

const (
	version = "1.0"
	// menu translate
	// ru
	selectru = "ВЫБРАТЬ"
	backru   = "НАЗАД"
	quitru   = "ВЫХОД"
	// en
	selecten = "SELECT"
)

// colors
// colorReset := "\033[0m"
// ColorYellow = "\033[33m"
// ColorCyan = "\033[36m"
// colorWhite := "\033[37m"

func main() {
	c()    // clear screen
	logo() // output logotype
	mm()   // start main menu (english version)
}

// GENERAL FUNCS
//  logotype output
func logo() {
	ascii := figlet4go.NewAsciiRender()

	// Adding the colors to RenderOptions
	options := figlet4go.NewRenderOptions()
	// options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorYellow,
	}

	// ascii.LoadFont("src/font/")

	renderStr, _ := ascii.RenderOpts("PIN3APPLE", options)
	fmt.Print(renderStr)
}

// clear screen
func c() {
	out, _ := exec.Command("clear").Output()
	fmt.Println(string(out))
}

// Menus
// main menu
func mm() {
	fmt.Printf(" \033[36m1)\033[37m DISTROS  \033[36m2)\033[37m MYOS\n\n \033[36m3)\033[37m SETTINGS \033[36m\033[36m0)\033[31m EXIT\n\n\033[0m")

	var scan string

	fmt.Print("\033[33m SELECT: \033[0m")
	fmt.Scan(&scan)

	switch scan {
	case "0":
		println("")
	case "1":
		c()
		logo()
		dm()
	case "2":
		c()
		myos()
		mm()
	case "3":
		c()
		logo()
		settings()
	default:
		c()
		logo()
		mm()
	}
}

// Distros menu
func dm() {
	fmt.Printf(" \033[36m1)\033[37m MANJARO  \033[36m2)\033[37m SOLUS\n\n \033[36m3)\033[37m FEDORA \033[36m\033[36m  \n\n \033[36m5)\033[33m BACK \033[36m0)\033[31m EXIT\n\n\033[0m")

	var scan string

	fmt.Printf("\033[33m %s: \033[0m", selecten)
	fmt.Scan(&scan)

	switch scan {
	case "0":
		println("")
	case "1":
		manjaro_script()
	case "2":
		solus_script()
	case "3":
		fedora_script()
	case "5":
		c()
		logo()
		mm()
	default:
		c()
		logo()
		dm()
	}
}

// os info
func myos() {
	out, _ := exec.Command("screenfetch").Output()
	fmt.Println(string(out))
}

//  settings
func settings() {
	fmt.Printf("\n \033[33mAUTHOR:\033[0m https://github.com/Avdushin\n \033[33mVERSION: \033[0m%sgo\033[36m \n\n \033[33mLANGUAGE:\n\n \033[36m1)\033[0m РУССКИЙ \n \033[36m2)\033[0m ENGLISH \n\n \033[36m5) \033[33mBACK \033[36m0)\033[31m EXIT\n\n", version)

	var scan string

	fmt.Print("\n\033[33m ACTION: \033[0m")
	fmt.Scan(&scan)

	switch scan {
	case "0":
		print("")
	case "1":
		c()
		logo()
		ruApp()
	case "2":
		c()
		logo()
		mm()
	case "5":
		c()
		logo()
		mm()
	default:
		c()
		logo()
		settings()
	}
}

//  LANGUAGES

//  ru
func ruApp() {
	fmt.Printf("\n \033[36m1)\033[37m ДИСТРИБУТИВЫ  \033[36m2)\033[37m МОЯ ОС\n\n \033[36m3)\033[37m НАСТРОЙКИ \033[36m\033[36m    0)\033[31m %s\n\n\033[0m", quitru)

	var scan string

	fmt.Printf("\033[33m %s: \033[0m", selectru)
	fmt.Scan(&scan)

	switch scan {
	case "0":
		println("")
	case "1":
		c()
		logo()
		rudm()
	case "2":
		c()
		myos()
		ruApp()
	case "3":
		c()
		logo()
		rusettings()
	default:
		c()
		logo()
		ruApp()
	}
}

func rudm() {
	fmt.Printf("\n \033[36m1)\033[37m MANJARO  \033[36m2)\033[37m SOLUS\n\n \033[36m3)\033[37m FEDORA \033[36m\033[36m  \n\n \033[36m5)\033[33m %s \033[36m0)\033[31m %s\n\n\033[0m", backru, quitru)

	var scan string

	fmt.Printf("\033[33m %s: \033[0m", selectru)
	fmt.Scan(&scan)

	switch scan {
	case "0":
		println("")
	case "1":
		manjaro_script()
	case "2":
		solus_script()
	case "3":
		fedora_script()
	case "5":
		c()
		logo()
		ruApp()
	default:
		c()
		logo()
		rudm()
	}
}

func rusettings() {
	fmt.Printf("\n \033[33mАВТОР:\033[0m https://github.com/Avdushin\n \033[33mВЕРСИЯ: \033[0m%sgo\033[36m \n\n \033[33mЯЗЫК:\n\n \033[36m1)\033[0m РУССКИЙ \n \033[36m2)\033[0m ENGLISH \n\n \033[36m5) \033[33m%s \033[36m0)\033[31m %s\n\n", version, backru, quitru)

	var scan string

	fmt.Printf("\n\033[33m %s: \033[0m", selectru)
	fmt.Scan(&scan)

	switch scan {
	case "0":
		print("")
	case "1":
		c()
		logo()
		ruApp()
	case "2":
		c()
		logo()
		mm()
	case "5":
		c()
		logo()
		ruApp()
	default:
		c()
		logo()
		rusettings()
	}
}

// SCRIPTS

// manjaro
func manjaro_script() {
	c := exec.Command("sh", "src/distros/Manjaro/manjaro.sh")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}

// solus
func solus_script() {
	c := exec.Command("sh", "src/distros/Solus/solus.sh")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}

// fedora
func fedora_script() {
	c := exec.Command("sh", "src/distros/Fedora/fedora.sh")
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	c.Run()
}
