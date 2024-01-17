package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"time"

)


func create() {
	log.SetFlags(log.LUTC)
	
	var title, header, description string
	var file *os.File
	
	var yes_or_not []byte = make([]byte, 1)


	
	file, err := os.Create("stored_tasks.txt")

	if err != nil {
		log.Println("[+] Impossible to create the file")
		log.Fatal(err)
		
	}
	

	for {
		for {
			log.Print("[*] Type the title: ")
			fmt.Scan(&title)

			fileCheck, _ := os.ReadFile("stored_tasks.txt")

			if strings.Contains(string(fileCheck), title) == true {
				log.Println("[-] You already have a task with this name, please think another name !")

				time.Sleep(5 * time.Second)
				
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()

				continue

			}else {
				break

			}

		}

		log.Print("[*] Type the header: ")
		fmt.Scan(&header)

		log.Print("[*] Description: ")
		reader := bufio.NewReader(os.Stdin)
		description, _ = reader.ReadString('\n')

		
		log.SetOutput(file)
		log.Printf("%s;%s;%s", title, header, description)
		log.SetOutput(os.Stdout)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
			
		log.Print("[+] Do you want continue to add tasks ? [y/n]: ")
		os.Stdin.Read(yes_or_not)

		if string(yes_or_not) == "y" {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
				
			continue

		}else {
			file.Close()
			break

		}
	}
}