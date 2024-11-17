package main

import "fmt"

type Avatar struct{
	URL string
	Size int64
}

type Client struct{
	ID int64
	Name string
	Age int
	IMG *Avatar
}

// func main() {
	i := new(int)
	_ = i

client := Client{}
client.IMG = new(Avatar)
fmt.Printf("%v\n", client)

updateAvatar(&client)
fmt.Printf("%#v\n", client)
fmt.Printf("%#v\n", client.IMG)

updateClient(&client)
fmt.Printf("%#v\n", client)

}

func updateAvatar(client *Client){
	client.IMG.URL = "update URL"
}

func updateClient(client *Client){
	client.Name = "Artem"
}