package reporter

import (
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

const (
	borderString = "1"
	alignString  = "CM"
	linkString   = ""
)

func Report(gurvic, wald, laplas, bayerLaplas map[int]float32, input [3][3]int) error {
	file := "report.pdf"
	pdf := gofpdf.New("P", "mm", "A4", linkString)

	pdf.SetFont("Arial", linkString, 10)
	basicTable := func() {
		pdf.SetX(30)

		pdf.CellFormat(30, 50, "Possible alternatives", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(60, 25, "Possible environment states", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(80, 25, "Criteria", borderString, 0, alignString, false, 0, linkString)

		pdf.Ln(25)
		pdf.SetX(60)

		pdf.CellFormat(20, 25, "S1", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(20, 25, "S2", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(20, 25, "S3", borderString, 0, alignString, false, 0, linkString)

		pdf.CellFormat(20, 25, "Wald", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(20, 25, "Bayes-Laplas", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(20, 25, "Gurvic", borderString, 0, alignString, false, 0, linkString)
		pdf.CellFormat(20, 25, "Laplas", borderString, 0, alignString, false, 0, linkString)

		pdf.Ln(-1)
		for i := 0; i < 3; i++ {
			pdf.SetX(30)
			pdf.CellFormat(30, 25, "A"+strconv.Itoa(i+1), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, strconv.Itoa(input[i][0]), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, strconv.Itoa(input[i][1]), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, strconv.Itoa(input[i][2]), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, fmt.Sprintf("%.2f", wald[i]), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, fmt.Sprintf("%.2f", bayerLaplas[i]), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, fmt.Sprintf("%.2f", gurvic[i]), borderString, 0, alignString, false, 0, linkString)
			pdf.CellFormat(20, 25, fmt.Sprintf("%.2f", laplas[i]), borderString, 0, alignString, false, 0, linkString)
			pdf.Ln(-1)
		}
	}
	explanation := func() {
		pdf.Text(30, 150, "Explanations:")
		pdf.Text(30, 155, "S1 - Competition on the same level")
		pdf.Text(30, 160, "S2 - Competition	a little intensified")
		pdf.Text(30, 165, "S3 - Competition intensified much")

		pdf.Text(30, 180, "A1 - Continue work in a habitual mode")
		pdf.Text(30, 185, "A2 - Activate advertising activity")
		pdf.Text(30, 190, "A3 - Activate advertising activity and lower price")
	}

	pdf.AddPage()
	basicTable()
	explanation()

	return pdf.OutputFileAndClose(file)
}
