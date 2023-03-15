package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// DISTROS SCRIPTS

func DistroScript(Distro, ScriptPath string) {
	var answer string
	fmt.Printf("Are you sure you want to install %s script? [Y/n]: ", Distro)
	fmt.Scan(&answer)

	switch answer {
	case "Y", "y", "Yes", "yes", "yeah", "ДА", "д", "да", "Да":
		c := exec.Command("sh", ScriptPath)
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
		c.Run()
	default:
		Clear()
		Logo()
		DistroMenu()
	}
}

// Russian

func RuDistroScript(Distro, ScriptPath string) {
	var answer string
	fmt.Printf("Вы уверены, что хотите выполнить скрипт установки %s ? [Да/нет]: ", Distro)
	fmt.Scan(&answer)

	switch answer {
	case "Y", "y", "Yes", "yes", "yeah", "ДА", "д", "да", "Да":
		c := exec.Command("sh", ScriptPath)
		c.Stdout = os.Stdout
		c.Stdin = os.Stdin
		c.Stderr = os.Stderr
		c.Run()
	default:
		Clear()
		Logo()
		DistroMenu()
	}
}

// manjaro
// func Manjaro_script() {
// 	var answer string
// 	fmt.Print("Are you sure you want to install Manjaro script? [Y/n]: ")
// 	fmt.Scan(&answer)

// 	switch answer {
// 	case "Y", "y", "Yes", "yes":
// 		c := exec.Command("sh", "src/distros/Manjaro/manjaro.sh")
// 		c.Stdout = os.Stdout
// 		c.Stdin = os.Stdin
// 		c.Stderr = os.Stderr
// 		c.Run()
// 	default:
// 		Clear()
// 		Logo()
// 		DistroMenu()
// 	}
// }

// solus
// func Solus_script() {
// 	var answer string
// 	fmt.Print("Are you sure you want to install Solus script? [Y/n]: ")
// 	fmt.Scan(&answer)

// 	switch answer {
// 	case "Y", "y", "Yes", "yes":
// 		c := exec.Command("sh", "src/distros/Solus/solus.sh")
// 		c.Stdout = os.Stdout
// 		c.Stdin = os.Stdin
// 		c.Stderr = os.Stderr
// 		c.Run()
// 	default:
// 		Clear()
// 		Logo()
// 		DistroMenu()
// 	}
// }

// fedora
// func Fedora_script() {
// 	var answer string
// 	fmt.Print("Are you sure you want to install Fedora script? [Y/n]: ")
// 	fmt.Scan(&answer)

// 	switch answer {
// 	case "Y", "y", "Yes", "yes", "yeah":
// 		c := exec.Command("sh", "src/distros/Fedora/fedora.sh")
// 		c.Stdout = os.Stdout
// 		c.Stdin = os.Stdin
// 		c.Stderr = os.Stderr
// 		c.Run()
// 	default:
// 		Clear()
// 		Logo()
// 		DistroMenu()
// 	}
// }
