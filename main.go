package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func getNoteData()(string, string){
	title :=getUserInput("Title  : ")
	
  
	content :=getUserInput("\n Content : ")
	return title,content
	
}
func main(){
   title, content:=getNoteData()
   userNote,err :=note.New(title,content)
   if err!=nil{
	fmt.Println(err)
	return
   }
   userNote.Display()
   
   err = userNote.Save()

   if err !=nil{
	fmt.Println("Saving the note is failed")
	return
   }
   fmt.Println("Saving the note is succeeded \n")
}

func getUserInput(prompt string)(string){
	fmt.Print(prompt)
	// var value string
	// fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)
	text, err :=reader.ReadString('\n')

	if err != nil{
		return ""
	}
	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}