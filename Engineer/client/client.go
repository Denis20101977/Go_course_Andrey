package client


type Person struct{
	Name string
	Age int
}

type Avatar struct{
	URL string
	Size int64
}

type Client struct{
	Person
	Id int64
	Img Avatar
}

func(c Client) HasAvatar() bool{
	if c.Img.URL != "" {
		return true
	}
	return false
}

func (c *Client) UpdateAvatar(){
	c.Img.URL = "New URL"
}

