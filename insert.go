package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
	"time"
	"strings"
	"os/exec"

)

func insert() {
	log.SetFlags(log.LUTC)

	_, err := os.Stat("stored_tasks.txt")

	var title, header, description string

	if err == nil {
		file, err := os.OpenFile("stored_tasks.txt", os.O_APPEND | os.O_WRONLY, 0644)

		if err == nil {
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

			file.WriteString(fmt.Sprintf("%s;%s;%s", title, header, description))
			
			log.Println("[+] Task add with success !")
			time.Sleep(2 * time.Second)

			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()

			
		}else {
			log.Fatalf("[-] Can't open the file: %v", err)

		}


	}else {
		log.Fatal("[-] Please create a new file and add a new Task !")

	}
	
}