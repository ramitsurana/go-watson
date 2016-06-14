package main

import (
   "bufio"
   "fmt"
   "os"
)


func sample(){

   fmt.Println("You can install the following watson utilites (Kale/Runner)")
   	scan := bufio.NewScanner(os.Stdin)
   	scan.Scan()
   	fmt.Println(scan.Text())

	
	//if scan.Text() = kale {                
	//cmd := exec.Command("wget","https://github.com/IBM-Watson/kale/releases/download/v1.5.0/kale-1.5.0-standalone.jar")
	//cmd := exec.Command("alias", "kale="java -jar /full/path/to/kale-1.5.0-standalone.jar"")	

	//err := cmd.Start()
	//if err != nil {
        //        log.Fatal(err)
        //}
        //log.Printf("Installing Kale..")
        //err= cmd.Wait()
	//log.Printf("You have installed kale successfully.")
}
	
	
