package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}
// type displayer interface{
// 	Display()
// }

type outputtable interface{
	saver
	Display() 
}

func main(){


   title, content:=getNoteData()

    todoText :=getUserInput(" Todo Text :")

	todo, err := todo.New(todoText)

	if err!= nil{
		fmt.Println(err)
		return 
	}
	userNote,err :=note.New(title,content)
	if err!=nil{
	 fmt.Println(err)
	 return
	}

	err = outputData(todo)
	if err !=nil{
		return
	  }
  
   outputData(userNote)
   
  anyValueAllowed(1)
  anyValueAllowed(1.4)
  anyValueAllowed("test")
  
}

func anyValueAllowed( value interface{}){
	stringVal, ok := value.(string)
	if ok{
		fmt.Print("string",stringVal)
		return
	}
	switch value.(type){
	case int:
		fmt.Println("its an integer",value)
	case float32:
		fmt.Println("its an integer",value)
	case string:
		fmt.Println("its an integer",value)
	}
	fmt.Println(value)
}
func outputData(data outputtable)error{
	//to display and to save
	data.Display()
	return saveData(data)

}
func saveData(data saver) error{
	err := data.Save()

   if err !=nil{
	fmt.Println("Saving the note is failed")
	return err
   }
   fmt.Println("Saving the note is succeeded ")
   return nil
}

func getNoteData()(string, string){
	title :=getUserInput("Title  : ")
	
  
	content :=getUserInput("\n Content : ")
	return title,content
	
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

