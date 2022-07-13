package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type blockContent struct {
	lineId  string
	content string
}

type TagM struct {
	Id            string
	Name          string
	ChangeTagName string
}

type TagChange struct {
	Id   string
	Name string
}

type Tag struct {
	LineId    string `json:"line_id"`
	TagId     string `json:"tag_id"`
	TagName   string `json:"tag_name"`
	ChannelId string `json:"channel_id"`
	CreatedAt string `json:"created_at"`
}

func main() {
	fs := handleBlock()

	for _, f := range fs {
		f.Close()
	}
}

func handleBlock() map[string]*os.File {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	blockChannel := make(chan []blockContent)

	/** init */
	os.RemoveAll("chunk")
	os.Mkdir("chunk", 0700)

	/** meta */
	fmt.Println("Meta")

	tagMap := getTagMap()

	changeMap := buildChangeMap(tagMap)

	/** read */
	go blockRead(changeMap, blockChannel)

	/** write */
	fs := make(map[string]*os.File)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go blockWrite(i, fs, blockChannel, wg, mutex)
	}

	wg.Wait()

	return fs
}

func blockRead(changeTags map[string]TagChange, c chan<- []blockContent) {
	defer close(c)

	var total int = 0
	var chunk []blockContent
	var limit int = 10000

	// f, err := os.Open("small-pxmart-tags.log")
	// f, err := os.Open("large-pxmart-tags.log")
	f, err := os.Open("pxmart-tags.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := string(s.Bytes())

		linePos := strings.Index(line, "line_id")
		tagIdPos := strings.Index(line, "tag_id")
		tagNamePos := strings.Index(line, "tag_name")

		lineId := line[linePos+10 : tagIdPos-3]
		tagId := line[tagIdPos+9 : tagNamePos-3]

		// chnage tag
		if findChangeMap, ok := changeTags[tagId]; ok {
			var tag Tag
			json.Unmarshal(s.Bytes(), &tag)

			tag.TagId = findChangeMap.Id
			tag.TagName = findChangeMap.Name

			tagByte, err := json.Marshal(tag)
			if err != nil {
				log.Fatal(err)
			}
			line = string(tagByte)
		}

		chunk = append(chunk, blockContent{
			lineId:  lineId,
			content: line,
		})
		if len(chunk) == limit {
			c <- chunk
			fmt.Printf("block produce: %d\n", len(chunk))

			total += len(chunk)

			chunk = nil
		}
	}

	c <- chunk
	total += len(chunk)
	fmt.Printf("block produce: %d\n", len(chunk))

	fmt.Printf("total: %d\n", total)
}

func blockWrite(worker int, fs map[string]*os.File, c <-chan []blockContent, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for data := range c {
		fmt.Printf("block %d worker consume: %d\n", worker, len(data))

		for _, row := range data {
			lineId := row.lineId
			lineIdChar := lineId[0:2]

			var f *os.File
			mutex.Lock()
			if _, ok := fs[lineIdChar]; !ok {
				wfName := fmt.Sprintf("chunk/%s.log", lineIdChar)

				nf, err := os.OpenFile(wfName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}

				fs[lineIdChar] = nf
			}
			mutex.Unlock()

			f = fs[lineIdChar]

			if _, err := f.WriteString(fmt.Sprintf("%s\n", string(row.content))); err != nil {
				log.Println(err)
			}
		}
	}
}

func getTagMap() map[string]TagM {
	file, err := os.Open("tag-map.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := make(map[string]TagM)

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		tagId := record[0]
		tagName := record[1]
		changeTagName := record[4]

		result[tagName] = TagM{Id: tagId, Name: tagName, ChangeTagName: changeTagName}
	}

	return result
}

func buildChangeMap(tagMap map[string]TagM) map[string]TagChange {
	result := make(map[string]TagChange)

	for _, tag := range tagMap {
		if tag.ChangeTagName != "" {
			changeTag := tagMap[tag.ChangeTagName]

			result[tag.Id] = TagChange{Id: changeTag.Id, Name: changeTag.Name}
		}
	}

	return result
}
