package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)


func read() {
	log.SetFlags(log.LUTC)

	var index int = 0

	if file, err := os.Open("stored_tasks.txt"); err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			items := strings.Split(scanner.Text(), ";")
			
			log.Printf("==================== Index: %d ====================\n", index)
			log.Printf("[+] Title: %s\n", items[0])			
			log.Printf("[+] Header: %s\n", items[1])
			log.Printf("[+] Description: %s\n", items[2])
			log.Printf("===================================================\n")

			index++

		}

	}else {
		log.Fatal("[-] Please, create a new task file with some task inside the file !")

	}

}