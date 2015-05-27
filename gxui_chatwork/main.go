package main

import (
	"bufio"
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const roomID = "33408194"

var apiToken = ""

func appMain(driver gxui.Driver) {
	const width = 320
	theme := dark.CreateTheme(driver)

	window := theme.CreateWindow(width, 240, "Hello")
	window.OnClose(driver.Terminate)

	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)
	window.AddChild(layout)

	textBox := theme.CreateTextBox()
	textBox.SetMultiline(true)
	textBox.SetDesiredWidth(width)
	layout.AddChild(textBox)

	button := theme.CreateButton()
	button.SetText("Post")
	button.OnClick(func(gxui.MouseEvent) {
		postMessage(textBox.Text())
		textBox.SetText("")
	})
	button.SetHorizontalAlignment(gxui.AlignCenter)
	layout.AddChild(button)

	fmt.Println(apiToken)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("API Token: ")
	token, _ := reader.ReadString('\n')
	apiToken = token

	gl.StartDriver(appMain)
}

func postMessage(message string) {
	urlString := "https://api.chatwork.com/v1/rooms/" + roomID + "/messages"

	values := url.Values{}
	values.Add("body", message)

	req, _ := http.NewRequest("POST", urlString, strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-ChatWorkToken", apiToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	contentString := string(content)

	fmt.Println(contentString)
}
