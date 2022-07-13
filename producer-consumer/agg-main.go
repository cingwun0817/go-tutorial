package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type tag struct {
	LineId    string `json:"line_id"`
	TagId     string `json:"tag_id"`
	TagName   string `json:"tag_name"`
	ChannelId string `json:"channel_id"`
	CreatedAt string `json:"created_at"`
}

type uniqueTag struct {
	lineId   string
	tagId    string
	createAt int64
	content  string
}

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	/** init */
	os.RemoveAll("output")
	os.Mkdir("output", 0700)

	of, err := os.OpenFile("output/result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer of.Close()

	fs, err := ioutil.ReadDir("chunk")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fs {
		aggChannel := make(chan []string)

		/** read */
		go aggRead("chunk/"+f.Name(), aggChannel)

		/** write */
		for i := 0; i < 16; i++ {
			wg.Add(1)
			go aggWrite(i, of, aggChannel, wg, mutex)
		}
	}

	wg.Wait()
}

func aggRead(fn string, c chan<- []string) {
	defer close(c)

	data := make(map[string]uniqueTag)

	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := string(s.Bytes())

		endPos := strings.Index(line, "\"}")

		linePos := strings.Index(line, "line_id")
		tagIdPos := strings.Index(line, "tag_id")
		tagNamePos := strings.Index(line, "tag_name")
		createdAtPos := strings.Index(line, "created_at")

		lineId := line[linePos+10 : tagIdPos-3]
		tagId := line[tagIdPos+9 : tagNamePos-3]
		createdAtStr := line[createdAtPos+13 : endPos]

		createdAt, err := strconv.ParseInt(createdAtStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		hashKey := lineId + tagId

		assignBool := false
		if tag, ok := data[hashKey]; ok {
			fmt.Println(tag)
			if tag.createAt == 0 {
				assignBool = true
			} else if tag.createAt < createdAt {
				assignBool = true
			}
		} else {
			assignBool = true
		}

		if assignBool {
			data[hashKey] = uniqueTag{
				lineId:   lineId,
				tagId:    tagId,
				createAt: createdAt,
				content:  line,
			}
		}
	}

	var limit int = 1000
	var total int = 0
	var chunk []string
	for _, row := range data {
		total++

		chunk = append(chunk, row.content)

		if len(chunk) == limit {
			c <- chunk

			chunk = nil
		} else if total == len(data) {
			c <- chunk

			chunk = nil
		}
	}

	fmt.Println(total, len(data))
}

func aggWrite(worker int, f *os.File, c <-chan []string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	content := ""
	for data := range c {
		fmt.Printf("agg %d worker consume: %d\n", worker, len(data))

		for _, row := range data {
			content += fmt.Sprintf("%s\n", row)
		}
	}

	mutex.Lock()
	if _, err := f.WriteString(content); err != nil {
		log.Println(err)
	}
	mutex.Unlock()
}
