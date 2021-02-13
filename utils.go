package main

import (
	"encoding/json"
)

type ValueType struct {
	Value string
	Type  string
}

var md = make(map[string]string)

func converter(d map[string]string) []byte {

	md["ev"] = "event"
	md["et"] = "event_type"
	md["id"] = "app_id"
	md["uid"] = "user_id"
	md["mid"] = "message_id"
	md["t"] = "page_title"
	md["p"] = "page_url"
	md["l"] = "browser_language"
	md["sc"] = "screen_size"

	keys := make([]string, 0, len(d))

	for k := range d {
		keys = append(keys, k)
	}
	result := make(map[string]interface{})

	attr := map[string]interface{}{}
	trai := map[string]interface{}{}

	for k := range keys {

		if len(keys[k]) > 4 {
			l := keys[k][:4]
			if l == "atrk" {

				key := keys[k]
				val := "atrv" + keys[k][4:]
				ty := "atrt" + keys[k][4:]

				attr[d[key]] = ValueType{d[val], d[ty]}

			} else if l == "uatr" {
				if keys[k][:5] == "uatrk" {
					key := keys[k]
					val := "uatrv" + keys[k][5:]
					ty := "uatrt" + keys[k][5:]
					trai[d[key]] = ValueType{d[val], d[ty]}
				}
			} else {
				key := keys[k]

				result[md[key]] = d[key]
			}
		} else {

			key := keys[k]

			result[md[key]] = d[key]
		}
	}
	result["attributes"] = attr
	result["traits"] = trai
	json1, _ := json.Marshal(result)

	return json1
}
