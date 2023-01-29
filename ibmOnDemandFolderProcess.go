package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
### Given a folder
	## Stage 1
    Find all *.ind files
    for each 'ind' file
        check it there is a matching *.out
        if not found 'log an Error message'
    end-for
*/ /*
	## Stage 2
    for each 'ind' file
        substring '8579-3-0-1000FAA-15498-15498-8581.0.'MTF_IN_20120925_HIST.MTF_IN'.ind'
        create/open a new file MTF_IN_20120925_HIST.MTF_IN.ind2
        for loop through lines
            write out all 'GROUP_' lines to 'ind2' file
        end-for [lines loop]
    end-for [file loop]
*/ /*
	## Stage 3
    for each 'ind2' file
        for each 'ind2' file
        read each line
            if it is 'GROUP_FIELD_NAME'
                read next line and use the pair
                & insert the key/value pair into database
            else if it is 'GROUP_OFFSET'
                read next two line (GROUP_LENGTH and GROUP_FILENAME)
            end-if
        end-for [each file]
    end-for all 'ind2' files
*/

var indFileList [512]string

func main() {

	//read parameter for folder name
	var folderName = "sample"
	stageOneProcessing(folderName)
	createIND2Files()

}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func createIND2Files() {
	/*	## Stage 2
		for each 'ind' file
		    substring '8579-3-0-1000FAA-15498-15498-8581.0.'MTF_IN_20120925_HIST.MTF_IN'.ind'
		    create/open a new file MTF_IN_20120925_HIST.MTF_IN.ind2
		    for loop through lines
		        write out all 'GROUP_' lines to 'ind2' file
		    end-for [lines loop]
		end-for [file loop]
	*/
	for i := range indFileList {
		if indFileList[i] != "" {
			//fmt.Println(indFileList[i])
			//fmt.Println(string(indFileList[i][0:(len(indFileList[i]) - 4)]))
			var fileBase = string(indFileList[i][0:(len(indFileList[i]) - 4)])
			fileBase = fileBase + ".ind2"
			fileInd2, err := os.Create(fileBase)
			errorCheck(err)
			defer fileInd2.Close()
			readIndFile, err := os.Open("sample/" + indFileList[i])
			errorCheck(err)
			defer readIndFile.Close()
			fileScanner := bufio.NewScanner(readIndFile)
			fileScanner.Split(bufio.ScanLines)

			for fileScanner.Scan() {
				//fmt.Println(fileScanner.Text())
				var line = fileScanner.Text()
				if strings.Contains(line, "GROUP_") {
					fileInd2.WriteString(line)
				}
			} // end-of line loop
		}
	} // end-of ind files loop
}
func stageOneProcessing(folderName string) {
	println(folderName)
	var filePattern = folderName + "/*.ind"
	files, err := filepath.Glob(filePattern)
	errorCheck(err)

	var i int = 0
	for _, file := range files {
		var printValue = strings.Split(file, "\\")
		indFileList[i] = printValue[len(printValue)-1]
		//fmt.Println(printValue[len(printValue)-1])
		i++
	}
	for i := range indFileList {
		if indFileList[i] != "" {
			//fmt.Println(indFileList[i])
			//fmt.Println(string(indFileList[i][0:(len(indFileList[i]) - 4)]))
			var fileBase = string(indFileList[i][0:(len(indFileList[i]) - 4)])
			// ***** Check if '.out' file exsists
			path := folderName + "/" + fileBase + ".out"
			//fmt.Printf("Out file is : \t %s", path)
			_, err := os.Stat(path)
			errorCheck(err)
			var outFileStatus = !errors.Is(err, os.ErrNotExist)
			if outFileStatus {
				//fmt.Printf("%s exists: %t\n", path, outFileStatus)
			} else {
				fmt.Printf("************ %s File NOT Found", path)
			}
			//fmt.Printf("\t length : \t %d\n", len(indFileList[i]))
			//fmt.Println(i)
		}
	}
}
