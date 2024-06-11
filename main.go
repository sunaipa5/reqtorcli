package main

import (
	"flag"
	"os"
	"reqtorcli/dye"
    "encoding/json"
	"github.com/sunaipa5/reqtor"
)

func main() {
	urlFlag := flag.String("u", "", "")
	fileFlag := flag.String("f", "", "")
	consoleFlag := flag.Bool("nc", false, "Console flag")
	headersFlag := flag.String("h", "", "")
	flag.Parse()
	if flag.Lookup("nc").Value.String() == ""{
		*consoleFlag = true
	}

	var requesturl string
	var filename string
	var consoleMod bool
	var headersFile string

	if *urlFlag != "" {
		dye.Magenta("Request Url:", *urlFlag)
		requesturl = *urlFlag
	}

	if *fileFlag != "" {
		dye.Bmagenta("Output file:", *fileFlag)
		filename = *fileFlag
	}

	if *consoleFlag {
		dye.Magenta("Write console disabled")
		consoleMod = true
	}

	if *headersFlag != "" {
		dye.Bmagenta("Headers file:", *headersFlag)
		headersFile = *headersFlag
	}
	if flag.NFlag() == 0 {
		dye.Bmagenta(` ____            _             
|  _ \ ___  ____| |_ ___  _ __ 
| |_) / _ \/ _  | __/ _ \| '__|
|  _ <  __/ (_| | || (_) | |   
|_| \_\___|\__, |\__\___/|_|   
              |_|`)
		dye.Magenta(`+---------+-----------------------+------------------------------+
| Command |   Parameter           |           Example            |
+---------+-----------------------+------------------------------+
| -u      | <request-url>         | https://example.com/         |
| -f      | <file-name>           | output.html                  |
| -nc     | <don't-write-console> | null                         |       
| -h      | <request-headers>     | headers.json                 |
+---------+-----------------------+------------------------------+`)
		return
	} else {
		err := reqtor.Start()
		if err != nil {
			dye.Red(err)
			return
		}
		var requestHeaders map[string]string
		if headersFile != ""{
			file, err := os.Open(headersFile)
			if err != nil {
				dye.Red("Headers file can't read:",err)
				return
			}
			defer file.Close()
		
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&requestHeaders)
			if err != nil {
				dye.Red("Headers file can't decode:",err)
				return
			}
		}

		resBody, err := reqtor.GetBody(requesturl, requestHeaders)
		if err != nil {
			dye.Red(err)
			return
		}
		err = reqtor.Stop()
		if err != nil {
			dye.Red(err)
			return
		}

		if !consoleMod {
			dye.White(string(resBody))
		}
		if filename != "" {
			os.WriteFile(filename, resBody, 0666)
		}
        
     
	}
}
