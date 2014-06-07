// jimmyr.go
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	_ "log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Author string `json:"author"`
	Score  int    `json:"score"`
	URL    string `json:"url"`
	Title  string `json:"title"`
}

func (i Item) String() string {
	return fmt.Sprintf(
		"Author: %s\nScore: %d\nURL: %s\nTitle: %s\n\n",
		i.Author,
		i.Score,
		i.URL,
		i.Title)
}

type Response struct {
	PageData struct {
		Children []struct {
			Entry Item `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type SubInfo struct {
	Weight int
	Take   int
	Name   string
}

func ReadListFile(filename string) []SubInfo {
	result := make([]SubInfo, 0, 20)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, " ")
		i, err := strconv.Atoi(arr[0])
		if err == nil {
			entry := SubInfo{Weight: i, Name: arr[1]}
			result = append(result, entry)
		}
	}
	return result
}

func EvalTakes(subs []SubInfo) {
	var totalWeight int = 0

	for _, v := range subs {
		totalWeight += v.Weight
	}

	for i, v := range subs {
		d := v.Weight * 50 / totalWeight
		fmt.Println(d)
		if d == 0 {
			d = 1
		}
		subs[i].Take = d
	}

}

func Get(sub string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", sub)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.PageData.Children))
	for i, child := range r.PageData.Children {
		items[i] = child.Entry
	}
	return items, nil
}

func main() {

	subs := ReadListFile("./filelist.txt")
	EvalTakes(subs)

	for _, s := range subs {
		fmt.Println(s.Name, s.Take, s.Weight)
	}
	//x := Sum(&array)  // Note the explicit address-of operator

	/*items, err := Get("newzealand")

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}*/
}
