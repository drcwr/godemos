package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

var pngpath string = "/root/go/src/github.com/apeiron/shell/profile/defaultKeys/assetKeys"
// This example demonsrates the generation of headers, footers and page breaks.
func ExampleFpdf_AddPage(file_list []string) {
	pngtotalnum := len(file_list)
	fmt.Printf("ExampleFpdf_AddPage %d\n",pngtotalnum)
	//fmt.Printf("ExampleFpdf_AddPage %d\n%v\n",pngnum,file_list)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTopMargin(30)
	/*pdf.SetHeaderFuncMode(func() {
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 15)
		pdf.Cell(80, 0, "")
		pdf.CellFormat(30, 10, "png list", "1", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)*/

	pdf.SetFooterFunc(func() {
		pdf.SetY(-35)
		pdf.SetFont("Arial", "I", 16)
		pdf.CellFormat(0, 10, fmt.Sprintf("500 apeiron each QR code",),
			"", 0, "L", false, 0, "")
	})
	pdf.AliasNbPages("")
	var ti = 0
	qrnum := 6

	pages := (pngtotalnum+qrnum-1)/qrnum
	for j:=0;j<pages;j++ {
		pdf.AddPage()
		for i:=0;i<qrnum;i++ {
			var x float64 = float64(20 +(i%2)*100)
			var y float64 = float64(10+80*(i/2))
			pdf.Image(filepath.Join(pngpath,file_list[ti]), x, y, 90, 0, false, "", 0, "")
			ti++
			if ti >= pngtotalnum {
				break
			}
		}
	}

	/*
	pdf.SetY(-65)
	pdf.SetFont("Times", "", 12)
	for j := 1; j <= 2; j++ {
		pdf.CellFormat(0, 10, fmt.Sprintf("Printing line number %d", j),
			"", 1, "", false, 0, "")
	}
	*/
	fileStr := "Fpdf_AddPage_cwr"
	err := pdf.OutputFileAndClose(fileStr)
	Summary(err, fileStr)
	// Output:
	// Successfully generated pdf/Fpdf_AddPage.pdf
}

// Summary generates a predictable report for use by test examples. If the
// specified error is nil, the filename delimiters are normalized and the
// filename printed to standard output with a success message. If the specified
// error is not nil, its String() value is printed to standard output.
func Summary(err error, fileStr string) {
	if err == nil {
		fileStr = filepath.ToSlash(fileStr)
		fmt.Printf("Successfully generated %s\n", fileStr)
	} else {
		fmt.Println(err)
	}
}

func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}
 
			return nil
		})
	return dir_list, dir_err
}

func GetAllFile(pathname string,filetype string) ([]string, error) {
	var file_list []string
    rd, _ := ioutil.ReadDir(pathname)
    for _, fi := range rd {
        if fi.IsDir() {
            fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
            GetAllFile(pathname + fi.Name() + "\\",filetype)
        } else {
            //fmt.Println(fi.Name())
            ok := strings.HasSuffix(fi.Name(), filetype)
            if ok {
            	file_list = append(file_list, fi.Name())
            }
        }
    }
    //fmt.Printf("%d\n%v\n",len(file_list),file_list)
    return file_list,nil
}

func main() {
	//list,_ := filepath.Glob("*.png")
	//for _,i := range(list) {
	//	fmt.Println(list)
	//}
	pngfilels,_ := GetAllFile(pngpath,".png")
	//fmt.Printf("pngfilels %d\n%v\n",len(pngfilels),pngfilels)
	
	ExampleFpdf_AddPage(pngfilels)
}