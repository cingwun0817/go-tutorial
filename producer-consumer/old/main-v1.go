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

	_ "go.uber.org/automaxprocs"
)

var wg sync.WaitGroup
var someMapMutex = sync.RWMutex{}
var changeMap map[string]TagChange

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
	tagMap := getTagMap()

	changeMap = buildChangeMap(tagMap)

	consumerCount := 2
	parseCh := make(chan []Tag)

	// result := sync.Map{}
	// result := make(map[string]Tag)

	wg.Add(consumerCount)

	go producer(parseCh)

	// for i := 0; i < consumerCount; i++ {
	// 	go consumer(i, parseCh, &result)
	// }

	wg.Wait()

	// tagStat := make(map[string]int)
	// for _, tag := range result {
	// 	tagId := tag.TagId

	// 	tagStat[tagId] += 1
	// }

	// tagFile, err := os.Create("output-tags.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer tagFile.Close()

	// tw := bufio.NewWriter(tagFile)
	// for _, tag := range result {
	// 	tagByte, err := json.Marshal(tag)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Fprintln(tw, string(tagByte))
	// }
	// tw.Flush()

	// statFile, err := os.Create("output-stat.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer statFile.Close()

	// sw := bufio.NewWriter(statFile)
	// for tagId, stat := range tagStat {
	// 	str := fmt.Sprintf("%s: %d", tagId, stat)

	// 	fmt.Fprintln(sw, str)
	// }
	// sw.Flush()

	fmt.Println("End")
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

type Tag struct {
	LineId    string `json:"line_id"`
	TagId     string `json:"tag_id"`
	TagName   string `json:"tag_name"`
	ChannelId string `json:"channel_id"`
	CreatedAt string `json:"created_at"`
}

func producer(parseCh chan []Tag) {
	defer close(parseCh)

	file, err := os.Open("middle-pxmart-tags.log")
	// file, err := os.Open("large-pxmart-tags.log")
	// file, err := os.Open("small-pxmart-tags.log")
	// file, err := os.Open("pxmart-tags.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var chunk []Tag
	count := 0
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		byteContent := scanner.Bytes()

		var tag Tag
		json.Unmarshal(byteContent, &tag)

		chunk = append(chunk, tag)
		count++

		if count >= 10 {
			parseCh <- chunk

			total += count
			chunk = nil
			count = 0

			fmt.Printf("Produce: %d\n", total)
		}
	}

	parseCh <- chunk
	total += count

	fmt.Printf("Produce: %d\n", total)
}

// func consumer(worker int, parseCh <-chan []Tag, result *map[string]Tag) {
// 	defer wg.Done()

// 	for data := range parseCh {
// 		fmt.Printf("Worker %d: process %d data\n", worker, len(data))

// 		for _, row := range data {
// 			hashKey := row.TagId + row.LineId

// 			assignBool := false
// 			someMapMutex.RLock()
// 			if tag, ok := (*result)[hashKey]; ok {
// 				if tag.CreatedAt == "0" {
// 					assignBool = true
// 				} else if tag.CreatedAt < row.CreatedAt {
// 					assignBool = true
// 				}
// 			} else {
// 				if findChangeMap, ok := changeMap[row.TagId]; ok {
// 					row.TagId = findChangeMap.Id
// 					row.TagName = findChangeMap.Name
// 				}

// 				assignBool = true
// 			}
// 			someMapMutex.RUnlock()

// 			if assignBool {
// 				someMapMutex.Lock()
// 				(*result)[hashKey] = row
// 				someMapMutex.Unlock()
// 			}
// 		}
// 	}
// }
