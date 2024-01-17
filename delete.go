package main

import (
	"fmt"
	"log"
	"os"
	"strings"

)

func delete() {
	log.SetFlags(log.LUTC)

	var option int 
	//var newValue string

	_, err := os.Stat("stored_tasks.txt")
	
	if err == nil {
		file, _ := os.OpenFile("stored_tasks.txt", os.O_RDWR, 0644)
		content, _ := os.ReadFile("stored_tasks.txt")

		tasksList := strings.Split(strings.Trim(string(content), " "), "\n")

		if string(content) == "" {
			log.Println("[-] The file is empty, please use the insert option and add some tasks.")

		}else {
			read()

			log.Print("[-] Please select the item ID:")
			fmt.Scan(&option)

			for id, _ := range tasksList {
				if option == id {
					tasksList[id] = ""

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
		log.Fatal("[-] Please create a new file and add a new Task !")

	}

}