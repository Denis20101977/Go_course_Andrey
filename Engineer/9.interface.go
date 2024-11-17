package main

import "fmt"

type Sender interface{
	Send(msg string) error
}

type Email struct{
	Adress string
}

func (e *Email)Send(msg string) error{
	fmt.Printf("Сообщение отправлено по почте", msg, e.Adress)
	return nil
}

func Notyfy(s Sender){
	err := s.Send("Notify message")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Susses")
}

func main(){

}
