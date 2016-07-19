package periopcheck

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/access"
)

//CheckIDDuplicates checks for duplicates in chart or ptid
func CheckIDDuplicates(id string, ptid bool) bool {
	var path string
	if ptid {
		path = "C:\\Users\\raymond chou\\Desktop\\PeriOp\\ptid.txt"
	} else {
		path = "C:\\Users\\raymond chou\\Desktop\\PeriOp\\chart.txt"
	}

	accessHelper.CreateFile(path)
	file, _ := accessHelper.ConnectToTxt(path)
	idDup := compareLine(file, id)
	if !idDup {
		accessHelper.FileWrite(file, id)
		return false
	}
	return true
}

//compareLine reads a line from text then compares if it matches cLine
func compareLine(file *os.File, cLine string) bool {
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	for err == nil {
		if cLine == line {
			return true
		}
		fmt.Print(line)
		line, err = reader.ReadString('\n')
	}
	if err != io.EOF {
		log.Println(err)
		return false
	}
	return false
}

//CheckValidNumber converts string to int and checks if the input is within the valid bounds
func CheckValidNumber(min int, max int, input string) bool {
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
		return false
	}
	if i <= max && i >= min {
		return true
	}
	return false
}

//CheckNonNegative coverts input to int and checks if input is non negative
func CheckNonNegative(input string) bool {
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Println(err)
		return false
	}
	if i >= 0 {
		return true
	}
	return false
}

//CheckNonNegativeFloat converts input to float and check if input is non negative
func CheckNonNegativeFloat(input string) bool {
	i, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Println(err)
		return false
	}
	if i >= 0 {
		return true
	}
	return false
}

//ErrorHandler handles the error
func ErrorHandler(rowCheck bool, row int, col string, value string) {
	if rowCheck {
		log.Printf("Row %d  Col %s = %s is not a valid integer", row, col, value)
	}
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return i
}

//CheckPVD checks if PVD defaults to 1 when CORATID is >0
func CheckPVD(pvd string, coratID string) bool {
	p := stringToInt(pvd)
	c := stringToInt(coratID)
	if c > 0 && p == 1 {
		return true
	}
	return false
}
