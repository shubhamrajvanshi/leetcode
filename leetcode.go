//Package leetcode  has useful utils for leetcode
package leetcode

import (
	"os"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	)	

// PrintTable prints leetcode table from the json input
func PrintTable(in *string){
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte (*in), &jsonMap)

	for lHeader , lHeaderValue  := range  jsonMap {
		switch vv := lHeaderValue.(type) {
		case map[string]interface{}:
			if lHeader=="headers"{
				for tablename, tcname  := range vv{
					fmt.Println("TableName:", tablename )
					th := make([]string, len(tcname.([]interface{})))
					for i, v1 := range tcname.([]interface{}) {
						th[i] = v1.(string)
					}
				
					table := tablewriter.NewWriter(os.Stdout)
					table.SetHeader(th)
					
					for _,row := range (jsonMap["rows"].(map[string]interface{})[tablename]).([]interface{}){
						tc := make([]string, len(row.([]interface{})))
						for i,rc := range row.([]interface{}){
							tc[i] = fmt.Sprintf("%v",rc)
						}
						table.Append(tc)
					}
					
					table.Render();	
					table = nil
					fmt.Println("")
				}
			}
		}
	}
}

// PrintRaw prints leetcode table from the json input
func PrintRaw(in *string){
	var m map[string]interface{}
	json.Unmarshal([]byte (*in), &m)

	for k , v  := range  m {
		switch vv := v.(type) {
		case map[string]interface{}:
			if k=="headers"{
				for tablename, v  := range vv{
					fmt.Println("TableName:", tablename )
					for _, v1 := range v.([]interface{}) {
						fmt.Print(v1, " ")
					}
					fmt.Println("")	
					
					for _,row := range (m["rows"].(map[string]interface{})[tablename]).([]interface{}){
						for _,rc := range row.([]interface{}){
							fmt.Print(rc," ")
						}
						fmt.Println("")	
					}
					fmt.Println("")	
					fmt.Println("")	
				}
			}
		}
	}
}