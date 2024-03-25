package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

var testObj = []byte(`{
"id" : "b33dc8e4-f1ed-436d-b570-96ca10b4dd36",
"latitude": 83.231,
"longitude" : 85.232123,
"tags": {
	"tagName1" : 123,
  	"tagName2" : "test",
  	"tagName3" : 123.2323,
  	"tagName4" : true,
  	"tagName5" : "2023-02-01 12:23:12"
	},
"tags_types": {
	"tagName1": "float",
  	"tagName2": "timestamp",
  	"tagName3": "integer",
  	"tagName4": "float",
  	"tagName5": "timestamp"
	}
}`)

type Obj struct {
	ID        string         `json:"id"`
	Latitude  float32        `json:"latitude"`
	Longitude float32        `json:"longitude"`
	Tags      map[string]any `json:"tags"`
	TagsTypes map[string]any `json:"tags_types"`
}

func get(data []byte) ([]byte, error) {
	o := new(Obj)
	if err := json.Unmarshal(data, o); err != nil {
		return nil, fmt.Errorf("unmarshal err: %w", err)
	}

	for tagName, tagValue := range o.Tags {
		if tagType, ok := o.TagsTypes[tagName]; ok {
			switch tagType {
			case "timestamp":
				t, _ := time.Parse(time.DateTime, tagValue.(string))
				o.Tags[tagName] = t.Format(time.DateTime)
				continue
			case "integer":
				if v, ok := tagValue.(bool); ok {
					res := 0
					if v == true {
						res = 1
					}
					o.Tags[tagName] = res
					continue
				}
				if v, ok := tagValue.(string); ok {
					o.Tags[tagName], _ = strconv.Atoi(v)
					continue
				}
				if v, ok := tagValue.(float64); ok {
					o.Tags[tagName] = int(v)
					continue
				}
			case "float":
				if v, ok := tagValue.(bool); ok {
					res := 0
					if v == true {
						res = 1
					}
					o.Tags[tagName] = float64(res)
					continue
				}
				if v, ok := tagValue.(string); ok {
					o.Tags[tagName], _ = strconv.ParseFloat(v, 64)
					continue
				}

			}
		}
	}

	res, err := json.MarshalIndent(o, "", "	")
	if err != nil {
		return nil, fmt.Errorf("unmarshal err at the and: %w", err)
	}
	return res, nil
}

func main() {
	res, err := get(testObj)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(res))
}
