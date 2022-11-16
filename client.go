package ragempquery

type Client struct {
	addr string
}

func NewClient(addr string) (c *Client, err error) {
	c = &Client{
		addr: addr,
	}
	err = nil
	return
}
