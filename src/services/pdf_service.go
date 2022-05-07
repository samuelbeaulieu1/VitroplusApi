package services

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/classes"
)

func renderTemplate(branchReport *classes.BranchReport) (string, responses.Error) {
	data, err := os.ReadFile("template/branchReport.html")
	logoSrc, _ := filepath.Abs("template/images/vitroplusZiebart.png")
	tpl := ""
	if err == nil {
		if tmpl, err := template.New("Branch report").Parse(string(data)); err == nil {
			buff := new(bytes.Buffer)
			branchReport.Logo = logoSrc
			if err = tmpl.Execute(buff, branchReport); err == nil {
				tpl = buff.String()
			}
		}
	}

	if err != nil {
		return "", responses.NewError("Impossible de créer le rapport")
	}
	return tpl, nil
}

func CreateReport(branchReport *classes.BranchReport) ([]byte, responses.Error) {
	tpl, err := renderTemplate(branchReport)
	if err != nil {
		return nil, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, responses.NewError("Impossible de créer le rapport")
	}

	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	page := wkhtmltopdf.NewPageReader(strings.NewReader(tpl))
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)
	err = pdfg.Create()
	if err != nil {
		return nil, responses.NewError("Impossible de créer le rapport")
	}

	return pdfg.Bytes(), nil
}
