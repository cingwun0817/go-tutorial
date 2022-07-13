package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var (
	wg            sync.WaitGroup
	changeMap     map[string]TagChange
	consumerCount int
	blockLimit    int
)

type Tag struct {
	LineId    string `json:"line_id"`
	TagId     string `json:"tag_id"`
	TagName   string `json:"tag_name"`
	ChannelId string `json:"channel_id"`
	CreatedAt string `json:"created_at"`
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

func main() {
	consumerCount = 8
	blockLimit = 100000
	// sourceFile := "pxmart-tags.log"
	sourceFile := "large-pxmart-tags.log"
	// sourceFile := "small-pxmart-tags.log"
	// sourceFile := "middle-pxmart-tags.log"

	/** init */
	fmt.Println("Init")

	os.RemoveAll("chunk")
	os.Mkdir("chunk", 0700)

	/** meta */
	fmt.Println("Meta")

	tagMap := getTagMap()

	changeMap = buildChangeMap(tagMap)

	/** read */
	channel := make(chan []Tag)

	go producer(sourceFile, channel)

	/** chunk */
	wg.Add(consumerCount)

	for c := 0; c < consumerCount; c++ {
		go consumer(c, channel)
	}

	wg.Wait()

	fmt.Println("Done")
}

func producer(fileName string, channel chan<- []Tag) {
	defer close(channel)

	sf, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer sf.Close()

	var chunk []Tag
	rowNum := 0
	total := 0
	scanner := bufio.NewScanner(sf)
	for scanner.Scan() {
		rowNum++

		byteContent := scanner.Bytes()

		var tag Tag
		json.Unmarshal(byteContent, &tag)

		chunk = append(chunk, tag)

		if rowNum >= blockLimit {
			channel <- chunk

			total += rowNum
			chunk = nil
			rowNum = 0

			fmt.Printf("Produce: %d\n", total)
		}
	}

	channel <- chunk
	total += rowNum

	fmt.Printf("Produce: %d\n", total)
}

func consumer(worker int, channel <-chan []Tag) {
	defer wg.Done()

	for data := range channel {
		fmt.Printf("Worker %d: process %d data\n", worker, len(data))

		chMap := make(map[string]*os.File)
		for _, tag := range data {
			lineId := tag.LineId

			// chnage
			if findChangeMap, ok := changeMap[tag.TagId]; ok {
				tag.TagId = findChangeMap.Id
				tag.TagName = findChangeMap.Name
			}

			var cf *os.File
			if _, ok := chMap[lineId[0:2]]; !ok {
				chunkFile := fmt.Sprintf("chunk/%s.log", lineId[0:2])

				nweCf, err := os.OpenFile(chunkFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Fatal(err)
				}
				defer cf.Close()

				chMap[lineId[0:2]] = nweCf
			}

			cf = chMap[lineId[0:2]]

			tagByte, err := json.Marshal(tag)
			if err != err {
				log.Fatal(err)
			}

			if _, err := cf.WriteString(fmt.Sprintf("%s\n", string(tagByte))); err != nil {
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
