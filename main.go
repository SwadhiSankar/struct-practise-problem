package main

import (
	"fmt"

	"example.com/note/note"
)

func getNoteData()(string, string){
	title :=getUserInput("Title")
	
  
	content :=getUserInput("Content")
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
   
   

}

func getUserInput(prompt string)(string){
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	
	return value
}