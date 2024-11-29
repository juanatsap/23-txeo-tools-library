package models

type Client struct {
	ID        int
	LongName  string
	ShortName string
}

type Clients []Client

func (c Client) String() string {
	return c.LongName
}

func (c Client) Int() int {
	return c.ID
}

func (c Client) Next() Client {
	return Client{ID: c.ID + 1}
}

func (c Client) Prev() Client {
	return Client{ID: c.ID - 1}
}
