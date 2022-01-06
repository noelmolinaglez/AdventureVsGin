package utils

import (
	"github.com/jung-kurt/gofpdf"
	"time"
)

func CreatePdf(data [][]string) {
	pdf := newReport("test")

	pdf = header(pdf, data[0])
	pdf = table(pdf, data[1:])

	if pdf.Err() {
		//log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
	}

	err := savePDF(pdf)
	if err != nil {
		//log.Fatalf("Cannot save PDF: %s|n", err)
	}
}

func newReport(title string) *gofpdf.Fpdf {
	pdf := gofpdf.New("L", "mm", "Letter", "")

	pdf.AddPage()

	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, title)

	pdf.Ln(12)

	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Mon Jan 2, 2006"))
	pdf.Ln(20)

	return pdf
}

func header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for _, str := range hdr {
		pdf.CellFormat(40, 7, str, "1", 0, "", true, 0, "")
	}
	pdf.Ln(-1)
	return pdf
}

func table(pdf *gofpdf.Fpdf, tbl [][]string) *gofpdf.Fpdf {
	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)

	align := []string{"L", "C", "L", "R", "R", "R"}
	for _, line := range tbl {
		for i, str := range line {
			pdf.CellFormat(40, 7, str, "1", 0, align[i], false, 0, "")
		}
		pdf.Ln(-1)
	}
	return pdf

}

func savePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("report.pdf")
}
