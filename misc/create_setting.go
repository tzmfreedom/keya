package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	members := [][]string{}
	for i := 0; i < 33; i++ {
		if i == 15 {
			continue
		}
		member := []string{}
		doc, _ := goquery.NewDocumentFromResponse(GetPage("http://www.keyakizaka46.com/s/k46o/artist/" + fmt.Sprintf("%02d", i+1)))
		member = append(member, strings.TrimSpace(doc.Find(".box-profile_text .name").Text()))
		member = append(member, strings.TrimSpace(doc.Find(".box-profile_text .furigana").Text()))
		doc.Find(".box-profile_text dt").Each(func(i int, s *goquery.Selection) {
			member = append(member, strings.TrimSpace(s.Text()))
		})
		members = append(members, member)
	}
	for _, member := range members {
		fmt.Println("[[members]]")
		fmt.Println("name = \"" + member[0] + "\"")
		fmt.Println("hiragana = \"" + member[1] + "\"")
		fmt.Println("birthday = \"" + member[2] + "\"")
		fmt.Println("constellation = \"" + member[3] + "\"")
		fmt.Println("height = \"" + member[4] + "\"")
		fmt.Println("home = \"" + member[5] + "\"")
		fmt.Println("blood_type = \"" + member[6] + "\"")
		fmt.Println("")
	}
}

func GetPage(url string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1)")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//	b, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println(bytes.NewBuffer(b).String())
	return resp
}
