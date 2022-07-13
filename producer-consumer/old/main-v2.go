package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
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

var (
	changeMap    map[string]TagChange
	wg           sync.WaitGroup
	resMap       map[string]Tag
	someMapMutex sync.RWMutex
)

func main() {
	consumerCount := 16
	blockLimit := 10000

	resMap = make(map[string]Tag)

	/** init */
	fmt.Println("Init")

	os.RemoveAll("output-pro")
	os.Mkdir("output-pro", 0700)

	channel := make(chan []Tag)
	oChannel := make(chan []Tag)

	/** producer */
	files, err := ioutil.ReadDir("chunk-pro")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer close(channel)

		for _, f := range files {
			fmt.Printf("File: %s\n", f.Name())

			f, err := os.Open(fmt.Sprintf("chunk-pro/%s", f.Name()))
			if err != nil {
				log.Fatal(err)
			}

			var chunk []Tag
			rowNum := 0
			total := 0
			scanner := bufio.NewScanner(f)
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

			var oChunk []Tag
			oRowNum := 0
			oTotal := 0
			someMapMutex.RLock()
			for _, tag := range resMap {
				oRowNum++

				oChunk = append(oChunk, tag)

				if oRowNum >= 500 {
					oChannel <- oChunk

					oTotal += oRowNum
					oChunk = nil
					oRowNum = 0

					fmt.Printf("Output: %d\n", oTotal)
				}
			}

			oChannel <- oChunk
			oTotal += oRowNum

			fmt.Printf("Output: %d\n", oTotal)

			someMapMutex.RUnlock()

			resMap = make(map[string]Tag)
		}
	}()

	/** consumer */
	wg.Add(consumerCount)
	for c := 0; c < consumerCount; c++ {
		go consumer(c, channel)
	}

	/** output consumer */
	// wg.Add(1)
	for c := 0; c < 8; c++ {
		go outputConsumer(c, oChannel)
	}

	wg.Wait()

	fmt.Println("End")
}

func consumer(worker int, channel <-chan []Tag) {
	defer wg.Done()

	for data := range channel {
		fmt.Printf("Worker %d: process %d data\n", worker, len(data))

		for _, row := range data {
			hashKey := row.TagId + row.LineId

			assignBool := false
			someMapMutex.RLock()
			if tag, ok := resMap[hashKey]; ok {
				if tag.CreatedAt == "0" {
					assignBool = true
				} else if tag.CreatedAt < row.CreatedAt {
					assignBool = true
				}
			} else {
				assignBool = true
			}
			someMapMutex.RUnlock()

			if assignBool {
				someMapMutex.Lock()
				resMap[hashKey] = row
				someMapMutex.Unlock()
			}
		}
	}
}

func outputConsumer(worker int, oChannel <-chan []Tag) {
	defer wg.Done()

	of, err := os.OpenFile("output-pro/result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	for data := range oChannel {
		fmt.Printf("Worker %d (Output): process %d data\n", worker, len(data))

		wContent := ""
		for _, tag := range data {
			contentByte, err := json.Marshal(tag)
			if err != nil {
				log.Fatal(err)
			}

			wContent += fmt.Sprintf("%s\n", string(contentByte))
		}

		if _, err := of.WriteString(wContent); err != nil {
			log.Println(err)
		}
	}
}
