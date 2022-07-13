package main

import (
	"log"

	"github.com/xitongsys/parquet-go/schema"
)

func main() {
	schemaStr := `
		{
			"Tag": "name=transform-example",
			"Fields": [
				{"Tag": "name=d_1, type=BYTE_ARRAY, convertedtype=UTF8"},
				{"Tag": "name=d_2, type=BYTE_ARRAY, convertedtype=UTF8"},
				{"Tag": "name=d_3, type=BYTE_ARRAY, convertedtype=UTF8"},
				{"Tag": "name=d_4, type=BYTE_ARRAY, convertedtype=UTF8"},
				{"Tag": "name=day, type=BYTE_ARRAY, convertedtype=UTF8"},
				{"Tag": "name=m_1, type=FLOAT"},
				{"Tag": "name=m_2, type=FLOAT"},
				{"Tag": "name=m_3, type=FLOAT"}
			]
		}
	`

	_, err := schema.NewSchemaHandlerFromJSON(schemaStr)
	if err != nil {
		log.Fatal("schema.NewSchemaHandlerFromJSON(): ", err)
	}

	// fw, err := local.NewLocalFileWriter("output.parquet")
	// if err != nil {
	// 	log.Fatal("NewLocalFileWriter(): ", err)
	// }

	// pw, err := writer.NewJSONWriter(schemaStr, fw, 4)
	// if err != nil {
	// 	log.Fatal("NewJSONWriter(): ", err)
	// }

	// pw.RowGroupSize = 128 * 1024 * 1024
	// pw.PageSize = 8 * 1024
	// pw.CompressionType = parquet.CompressionCodec_SNAPPY

	// fr, err := os.Open("data.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fr.Close()

	// scanner := bufio.NewScanner(fr)

	// for scanner.Scan() {
	// 	line := scanner.Bytes()

	// 	dimensions := map[string]string{}
	// 	json.Unmarshal(line, &dimensions)

	// 	metrics := map[string]float32{}
	// 	json.Unmarshal(line, &metrics)

	// 	rowData := make(map[string]interface{})

	// 	rowData["d_1"] = dimensions["media_account_uid"]
	// 	rowData["d_2"] = dimensions["campaign_id"]
	// 	rowData["d_3"] = dimensions["adset_id"]
	// 	rowData["d_4"] = dimensions["ad_id"]
	// 	rowData["day"] = dimensions["day"]
	// 	rowData["m_1"] = metrics["clicks"]
	// 	rowData["m_2"] = metrics["impressions"]
	// 	rowData["m_3"] = metrics["cost"]

	// 	json, err := json.Marshal(rowData)
	// 	if err != nil {
	// 		log.Fatal("Marshal(): ", err)
	// 	}

	// 	if err = pw.Write(string(json)); err != nil {
	// 		log.Fatal("Write(): ", err)
	// 	}
	// }

	// if err = pw.WriteStop(); err != nil {
	// 	log.Fatal("WriteStop(): ", err)
	// }

	// log.Println("Write Finished")
	// fw.Close()
}
