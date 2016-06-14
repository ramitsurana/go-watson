package main

import (
	"log"
	"os/exec"	
)

type Error struct {
        Name string
        Err  error
}

func main() {
	cmd := exec.Command("npm", "--version")
	cmd1 := exec.Command("yo", "-v")
	cmd2 := exec.Command("docker", "version")
	err := cmd.Start()
	err1 := cmd1.Start()
	err2 := cmd2.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Testing if npm is installed..")
	err= cmd.Wait()
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Printf("Testing if Yeo installed..")
	err1 = cmd1.Wait()
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Printf("Testing if Docker is installed..")
	err2 = cmd2.Wait()
	//log.Printf("Oops there's a gotcha with npm : %v", err)
	//log.Printf("Oops there's a gotcha with Yeo: %v", err1)
	//log.Printf("Oops there's a gotcha with docker : %v", err2)
	log.Printf("Voila ! You are good to go.Enjoy the amazing watson.")		
}
