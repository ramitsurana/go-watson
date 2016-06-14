package main

import (
	"fmt"
	"runtime"
	"log"
	"os/exec"
	"bufio"
    	"os"		
)


func Install(){

	if runtime.GOOS == "windows" {
    	fmt.Println("Cool Good to know you are Windows user")
	}
	if runtime.GOOS == "linux" {
        fmt.Println("Cool so you are a Linux User")
        }
	if runtime.GOOS == "osx" {
        fmt.Println("Cool so you are a OSX User")
        }
	//Installing watson Generator
	cmd := exec.Command("npm", "install", "-g", "generator-watson")
	err := cmd.Start()
	if err != nil {
                log.Fatal(err)
        }
        log.Printf("Installing Watson awesomness..")
        err= cmd.Wait()
	log.Printf("You are all set to go.Love it.Fork it.Try running $yo watson")
	
	//Installing Watson utilities
	fmt.Println("You can install the following watson utilites (Kale/Runner/Ally)")
   	scan := bufio.NewScanner(os.Stdin)
   	scan.Scan()   	

	//Installing Kale
	if scan.Text() == "kale" { 
	cmd1 := exec.Command("wget","https://github.com/IBM-Watson/kale/releases/download/v1.5.0/kale-1.5.0-standalone.jar")
	cmd2 := exec.Command("alias", "kale='java -jar /full/path/to/kale-1.5.0-standalone.jar'")	

	err1 := cmd1.Start()
	err2 := cmd2.Start()
	if err1 != nil {
                log.Fatal(err1)
        }
	if err2 != nil {
                log.Fatal(err2)
        }
        log.Printf("Installing Kale..")
        err1= cmd1.Wait()
	err2= cmd2.Wait()
	log.Printf("You have installed kale successfully.")
	}               		

	//Installing Runner
	if scan.Text() == "runner" {                
	cmd3 := exec.Command("npm", "install", "--global", "gulp-cli")
	cmd4 := exec.Command("npm", "install", "--save-dev", "gulp")	

	err3 := cmd3.Start()
	err4 := cmd4.Start()
	if err3 != nil {
                log.Fatal(err)
        }
	if err4 != nil {
                log.Fatal(err)
        }
        log.Printf("Installing Runner..")
        err3= cmd3.Wait()
	err4= cmd4.Wait()
	log.Printf("You have installed Runner successfully.")
	}

	//Installing ally.js
	if scan.Text() == "ally" {
	//Checking Bower is installed or not
	log.Printf("Checking if you have bower installed or not")
	cmd5 := exec.Command("bower", "--version")
	err5 := cmd5.Start()
	if err5 != nil { 
		log.Fatal(err)               
		log.Printf("Sorry you do not have bower installed.Please install Bower manually to kickastart.")				
        } else {
		log.Printf("You have Bower installed.That's Great.:)")			
	}	                 
	
	//Installing Ally
	//fmt.Println("Installing Ally")	
	cmd7 := exec.Command("bower", "install", "a11y.js", "--save")	
	err7 := cmd7.Start()	
	if err7 != nil {
                log.Fatal(err)
        }	
        log.Printf("Installing Ally.js..")
        err7= cmd7.Wait()	
	log.Printf("You have installed Ally.js successfully.")
	}
}	
