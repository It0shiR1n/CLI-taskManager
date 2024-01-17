package main

import (
	"log"
	"fmt"
	"strings"
	"os"
	"os/exec"
	"bufio"

)

func update() {
	log.SetFlags(log.LUTC)

	var name, header, description string
	var option int 

	_, err := os.Stat("stored_tasks.txt")

	if err == nil {
		file, _ := os.OpenFile("stored_tasks.txt", os.O_RDWR, 0644)
		content, _ := os.ReadFile("stored_tasks.txt")

		tasksList := strings.Split(string(content), "\n")

		if string(content) == "" {
			log.Println("[-] The file is empty, please use the insert option and add some tasks.")

		}else {
			read()

			log.Print("[-] Please select the item ID on you wan't update: ")
			fmt.Scan(&option)

			for id, _ := range tasksList {
				if option == id {
					log.Println("[+] Type the name: ")
					fmt.Scan(&name)

					log.Println("[+] Type the header: ")
					fmt.Scan(&header)

					log.Println("[+] Type the description: ")
					reader := bufio.NewReader(os.Stdin)
					description, _ = reader.ReadString('\n')

					cmd := exec.Command("clear")
					cmd.Stdout = os.Stdout
					cmd.Run()

					tasksList[id] = fmt.Sprintf("%s;%s;%s\n", name, header, description)

				}

			}


			file.Truncate(0)
			
			for id, _ := range tasksList {
				if tasksList[id] == "" {
					continue

				}else {
					file.WriteString(tasksList[id]+"\n")

				}

			}

		}

	}else {
		log.Fatal("[-] Please use a create option to create a new storage file !")

	}

}