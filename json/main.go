package main

import (
	"encoding/json"
	"fmt"
)

type payload struct {
	Region string `json:"region"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func main() {
	bodyStr := `{"region":"ap-northeast-1","key":"AKIAVLGKSUQHKQAGPLKL","secret":"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","from":{"bucket":"leooooo-mpc","key":"1529888140439425\/small-2020-01-01.txt"},"to":{"bucket":"leooooo-mpc","key":"1529888140439425\/output.parquet"},"schema":{"Tag":"name=transform-example","Fields":[{"Tag":"name=d_1, type=BYTE_ARRAY, convertedtype=UTF8"},{"Tag":"name=d_2, type=BYTE_ARRAY, convertedtype=UTF8"},{"Tag":"name=d_3, type=BYTE_ARRAY, convertedtype=UTF8"},{"Tag":"name=d_4, type=BYTE_ARRAY, convertedtype=UTF8"},{"Tag":"name=day, type=BYTE_ARRAY, convertedtype=UTF8"},{"Tag":"name=m_1, type=FLOAT"},{"Tag":"name=m_2, type=FLOAT"},{"Tag":"name=m_3, type=FLOAT"},{"Tag":"name=m_4, type=FLOAT"}]},"columns":[{"source_column":"media_account_uid","column":"d_1","column_type":"dim"},{"source_column":"campaign_id","column":"d_2","column_type":"dim"},{"source_column":"adset_id","column":"d_3","column_type":"dim"},{"source_column":"ad_id","column":"d_4","column_type":"dim"},{"source_column":"day","column":"day","column_type":"dim"},{"source_column":"clicks","column":"m_1","column_type":"met"},{"source_column":"impressions","column":"m_2","column_type":"met"},{"source_column":"cost","column":"m_3","column_type":"met"},{"source_column":"results","column":"m_4","column_type":"met"}]}`

	body := []byte(bodyStr)

	var payload payload
	json.Unmarshal(body, &payload)

	fmt.Println(payload)
}
