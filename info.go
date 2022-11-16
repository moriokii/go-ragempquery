package ragempquery

import (
	"encoding/json"
	"io"
	"net/http"
)

type ServerInfo struct {
	Name       string
	Gamemode   string
	Lang       string
	Players    int
	Peak       int
	MaxPlayers int
}

func (c *Client) Info() (s *ServerInfo, err error) {
	s = &ServerInfo{}
	resp, err := http.Get("https://cdn.rage.mp/master/")

	body, err := io.ReadAll(resp.Body)

	mapper := make(map[string]interface{})

	json.Unmarshal(body, &mapper)

	for i, d := range mapper {
		if i != c.addr {
			continue
		}
		m, _ := d.(map[string]interface{})
		s.Name = m["name"].(string)
		s.Gamemode = m["gamemode"].(string)
		s.Lang = m["lang"].(string)
		s.Players = int(m["players"].(float64))
		s.Peak = int(m["peak"].(float64))
		s.MaxPlayers = int(m["maxplayers"].(float64))
	}

	defer resp.Body.Close()
	return
}
