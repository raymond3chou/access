package access

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//OrderedMap works like a Map but is ordered
type OrderedMap struct {
	Colname string
	Value   string
}

//ConvertToString converts an array of NullString interfaces to an array of string
func ConvertToString(vals []interface{}) []string {
	row := make([]string, len(vals))
	for i, val := range vals {
		value := val.(*sql.NullString)
		row[i] = value.String
	}
	return row
}

//ConvertToText takes in the queried row divided in an array of strings based off of the column
//maincolumns contains the master columns and a flag for which ever one was used
//the function arranges based on
func ConvertToText(maincolumns []string, cols []OrderedMap) string {
	var row string
	found := false
	row = "\n"
	for _, mastercol := range maincolumns {
		found = false
		for i := range cols {
			if strings.Contains(cols[i].Colname, mastercol) {
				row += cols[i].Value + "|"
				found = true
				break
			}
		}
		if !found {
			row += "|"
		}
	}
	row = strings.TrimSuffix(row, "|")

	return row
}

//ConvertToOrderedMap converts a string array to an array of orderedMap
func ConvertToOrderedMap(cols []OrderedMap, rowstring []string) []OrderedMap {
	endindex := len(rowstring)
	i := 0
	for key := range cols {
		if i < endindex {
			cols[key].Value = rowstring[i]
			i++
		} else {
			break
		}

	}

	return cols
}

// ConnectToTxt Connects to Text File
func ConnectToTxt(filedir string) (*os.File, bool) {
	file, err := os.OpenFile(filedir, os.O_APPEND|os.O_RDWR|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Unable to Open Text File: %s", filedir)
		fmt.Print(err)
		return file, false
	}
	return file, true
}

//FileWrite Writes the queried row into a text file
func FileWrite(file *os.File, row string) int {
	_, err := file.WriteString(row)
	if err != nil {
		fmt.Println("Could Not Write String")
		return 0
	}
	file.Sync()
	return 1
}

//ReadFile reads a file
func ReadFile(filename string) string {
	fileoutput, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(fileoutput)
}

//CreateFile creates a file
func CreateFile(path string) bool {
	var f, err = os.Create(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	f.Close()
	return true
}

//ReadPath reads a from User input
func ReadPath(typeofpath string) string {
	fmt.Printf("\n\n------ENTER PATH FOR %s ------", typeofpath)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Path for %s: ", typeofpath)
	path, _ := reader.ReadString('\n')
	return path
}

//CreateErrorLog gets the path and creates the error file
func CreateErrorLog(test bool) *os.File {
	var path string
	if test {
		path = "C:\\Users\\raymond chou\\Desktop\\ErrorLog.log"
	} else {
		path = ReadPath("Error Log")
	}
	CreateFile(path)
	errFile, err := ConnectToTxt(path)
	if !err {
		log.Fatal("Did not connect\n")
	}
	return errFile
}
