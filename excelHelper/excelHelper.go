package excelHelper

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/access"
	"github.com/tealeg/xlsx"
)

//ConnectToXlsx connects to the excel file through the path
func ConnectToXlsx(xlsxPath string) *xlsx.File {
	sheet, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		log.Fatalln(err)
	}
	return sheet
}

//IdentifyCols returns all the columns in the sheet in a slice of string
func IdentifyCols(sheet *xlsx.File) []string {
	var colNamesSlice []string
	for i := 0; i < sheet.Sheets[0].MaxCol; i++ {
		cellValue := sheet.Sheets[0].Cell(0, i).Value
		cellValue = strings.ToLower(cellValue)
		colNamesSlice = append(colNamesSlice, cellValue)
	}
	return colNamesSlice
}

//ColCompare identifies the common strings in the two string slices
func ColCompare(tghCols []string, twhCols []string) []string {
	var commonColsSlice []string
	for _, twh := range twhCols {
		for _, tgh := range tghCols {
			if tgh == twh {
				commonColsSlice = append(commonColsSlice, tgh)
			}
		}
	}
	return commonColsSlice
}

//NotPresentinSlice identifies the strings that are not common between the two string slices
func NotPresentinSlice(originalCols []string, commonColsSlice []string) []string {
	var unCommonColsSlice []string
	var present bool
	for _, original := range originalCols {
		present = false
		for _, commonCol := range commonColsSlice {
			if commonCol == original {
				present = true
				break
			}
		}
		if !present {
			unCommonColsSlice = append(unCommonColsSlice, original)
		}
	}
	return unCommonColsSlice
}

//PrintSlice prints the slice
func PrintSlice(slice []string) {
	for _, s := range slice {
		fmt.Printf(" %s |", s)
	}
}

//ReadRow reads a row in excel and returns it as a slice of string
func ReadRow(sheet *xlsx.File) []string {
	return nil
}

//WriteStruct writes the struct type to text so it can be copied into peri.go
func WriteStruct(colNameSlice []string) {
	path := "C:\\Users\\raymond chou\\Desktop\\struct.txt"
	accessHelper.CreateFile(path)
	file, _ := accessHelper.ConnectToTxt(path)
	for _, c := range colNameSlice {
		var structPrint string
		lowerC := strings.ToLower(c)
		if strings.Contains(lowerC, "reop") {
			continue
		}
		if strings.Contains(lowerC, "mi") {
			continue
		}
		if strings.Contains(lowerC, "pace") {
			continue
		}
		if strings.Contains(lowerC, "tia") {
			continue
		}
		if strings.Contains(lowerC, "stroke") {
			continue
		}
		if strings.Contains(lowerC, "survival") {
			continue
		}
		upperC := strings.ToUpper(c)
		structPrint += upperC
		structPrint += " " + "string"
		structPrint += "`json:\"" + lowerC + "\"`\n"
		accessHelper.FileWrite(file, structPrint)
	}
}

//PeriOpLiteral prints to text the literal for the stuct
func PeriOpLiteral(colNameSlice []string) {
	path := "C:\\Users\\raymond chou\\Desktop\\periopliteral.txt"
	accessHelper.CreateFile(path)
	file, _ := accessHelper.ConnectToTxt(path)
	for _, c := range colNameSlice {
		var structPrint string
		lowerC := strings.ToLower(c)
		if strings.Contains(lowerC, "reop") {
			continue
		}
		if strings.Contains(lowerC, "mi") {
			continue
		}
		if strings.Contains(lowerC, "pace") {
			continue
		}
		if strings.Contains(lowerC, "tia") {
			continue
		}
		if strings.Contains(lowerC, "stroke") {
			continue
		}
		if strings.Contains(lowerC, "survival") {
			continue
		}
		upperC := strings.ToUpper(c)
		structPrint += upperC
		structPrint += " " + "string"
		structPrint += "`json:\"" + lowerC + "\"`\n"
		accessHelper.FileWrite(file, structPrint)
	}
}

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
