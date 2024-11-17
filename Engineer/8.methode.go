package main

import "fmt"

type Avatar struct {
 URL  string
 Size int64
}

type Client struct {
 ID   int64
 Name string
 Age  int
 IMG  *Avatar
}

func (c Client) HasAvatar() bool {
 if c.IMG != nil && c.IMG.URL != "" {
  return true
 }
 return false
}

func (c *Client) UpdateAvatar(newURL string) {
 if c.IMG == nil {
  c.IMG = &Avatar{} // Создаем новый Avatar, если он nil
 }
 c.IMG.URL = newURL // Присваиваем новое значение URL
}

// func main() {
 client := Client{
  ID:   7,
  Name: "Андрей",
  Age:  35,
  IMG: &Avatar{
   URL:  "some URL",
   Size: 10,
  },
 }

 fmt.Println(client)
 fmt.Println("Has Avatar:", client.HasAvatar())

 // Обновляем аватар
 client.UpdateAvatar("New URL")
 fmt.Println("Updated Client:", client)
 fmt.Println("Has Avatar:", client.HasAvatar())
}
