package periopcheck

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/access"
	"github.com/access/excelHelper"
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

//CheckPVD checks if PVD defaults to 1 when CORATID is >0
func CheckPVD(pvd string, coratID string) bool {
	p := excelHelper.StringToInt(pvd)
	c := excelHelper.StringToInt(coratID)
	if c > 0 && p == 1 {
		return true
	}
	return false
}

//CheckVPROS checks if the type is one of the listed in the codebook
func CheckVPROS(value string) bool {
	listedType := []string{"BP", "CA", "CF", "CP", "EC", "FS", "HO", "HT", "MC", "MP", "PA", "TO", "TF", "TR", "CM", "SJ", "AS", "BS", "CE", "DH", "ED", "EP", "FL", "HA", "HE", "HK", "IO", "LK", "MF", "MH", "MI", "MO", "MS", "OS", "SC", "SE", "SP", "SU", "TD", "D1", "D2", "DV", "PD", "PR", "R", "RC", "RD", "RS", "ST"}
	for _, t := range listedType {
		if t == value {
			return true
		}
	}
	return false
}

//CheckValid check if the value is valid
func CheckValid(value string) bool {
	valid := []string{"1", "2", "3", "4", "4a", "4b", "4c", "4d"}
	for _, v := range valid {
		if v == value {
			return true
		}
	}
	return false
}
