package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

func main() {

	args := os.Args
	fmt.Print(`
--------------------------------------------------------------------------------
--------------------------------------------------------------------------------
--------------------------------------------------------------------------------
Atlased to Phaser Parser
Version: 0.0
Build Date: 2024-01-24
Author: Gabriel A. Zorrilla - @GabrielZorrilla - zorrilla.me
Site: https://github.com/GAZ082/atlased_phaserjs_parser

FREE TO USE FOR ANY ENDEVOUR BUT PLEASE ASK FOR AUTHORIZATION AND MENTION THE
                        ORIGINAL AUTHOR IN YOUR PROJECT.
--------------------------------------------------------------------------------
--------------------------------------------------------------------------------
--------------------------------------------------------------------------------



`)
	fmt.Println("")
	if len(args) != 3 {
		fmt.Print(`ERROR: Must provide two files, the input (Atlased JSON) and the output destination JSON.`)
		os.Exit(1)
	}
	input := readAtlasedJson(args[1])
	output, _ := os.Create(args[2])
	write_buffer := bufio.NewWriter(output)
	parseAtlasedJson(input, write_buffer)
	write_buffer.Flush()
	println(`DONE!`)
}

func readAtlasedJson(inputFile string) string {
	b, err := os.ReadFile(inputFile)
	if err != nil {
		print(err.Error())
		return "Something wrong with the input file."
	}
	return string(b)
}
func parseAtlasedJson(json string, wbuffer *bufio.Writer) {
	wbuffer.WriteString(`{ "frames": {`)
	results := gjson.Get(json, "regions").Array()
	lastItem := len(results)
	for counter, i := range results {
		value := fmt.Sprintf(`
		"%v": {
			"frame": {
				"x":%v,
				"y":%v,
				"w":%v,
				"h":%v
			},
			"pivot": {
				"x":%.2f,
				"y":%.2f
			}
		},`,
			i.Get("name").String(),
			i.Get("rect.0").Int(),
			i.Get("rect.1").Int(),
			i.Get("rect.2").Int(),
			i.Get("rect.3").Int(),
			i.Get("origin.0").Float()/i.Get("rect.2").Float(),
			i.Get("origin.1").Float()/i.Get("rect.3").Float(),
		)
		if counter == lastItem-1 {
			value = value[:len(value)-1]
		}
		wbuffer.WriteString(value)
		counter++
	}
	wbuffer.WriteString(`
	}
}`)
}
