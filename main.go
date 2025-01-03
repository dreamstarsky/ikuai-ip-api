package main

import (
	"encoding/json"
	"fmt"
	"ikuai-ip-api/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	MaxGetipTryTime = 2
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	r := gin.Default()

	url := os.Getenv("IKUAI_URL")
	username := os.Getenv("IKUAI_USERNAME")
	password := os.Getenv("IKUAI_PASSWORD")

	if url == "" || username == "" || password == "" {
		log.Println("$IKUAI_URL, $IKUAI_USERNAME or $IKUAI_PASSWORD not found!")
		os.Exit(1)
	}

	ikuai := api.NewIkuai(url, username, password)

	err := ikuai.Login()
	if err != nil {
		log.Println("Initial login failed:", err)
	} else {
		log.Println("Initial login successful")
	}

	r.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		for i := 0; i < MaxGetipTryTime; i++ {
			ip, err := getip(id, ikuai)
			if err == nil {
				c.String(200, ip)
				return
			} else if i+1 < MaxGetipTryTime {
				log.Println("Getip failed:", err, "Try again...")
			}
		}

		c.String(404, "Not find\n")
	})

	r.Run(":8080")
}

type IfaceCheck struct {
	Id              int    `json:"id"`
	Interface       string `json:"interface"`
	ParentInterface string `json:"parent_interface"`
	IpAddr          string `json:"ip_addr"`
	Gateway         string `json:"gateway"`
	Internet        string `json:"internet"`
	Updatetime      string `json:"updatetime"`
	AutoSwitch      string `json:"auto_switch"`
	Result          string `json:"result"`
	Errmsg          string `json:"errmsg"`
	Comment         string `json:"comment"`
}

type IfaceCheckData struct {
	IfaceCheck []IfaceCheck `json:"iface_check"`
}

func getip(id string, ikuai api.Ikuai) (string, error) {

	resp, err := ikuai.Call("show", "monitor_iface", "iface_check")
	if err != nil {
		return "", err
	}

	if resp.Result != 30000 {
		return "", fmt.Errorf("API Error: %s", resp.ErrMsg)
	}

	var data IfaceCheckData
	err = json.Unmarshal(resp.Data, &data)
	if err != nil {
		return "", err
	}

	for _, iface := range data.IfaceCheck {
		if iface.Interface == id {
			return iface.IpAddr, nil
		}
	}

	return "", fmt.Errorf("interface %s not found", id)
}
