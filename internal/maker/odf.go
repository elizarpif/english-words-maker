package maker

import (
	"math/rand"
	"os"

	"github.com/kpmy/odf/generators"
	"github.com/kpmy/odf/mappers"
	"github.com/kpmy/odf/mappers/attr"
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns"
)

const timesNewRoman = "Times New Roman"

//func toPdf(fileName string, words []string) error {
//	cfg := pdfcpu.NewDefaultConfiguration()
//	reftable := pdfcpu.NewFreeHeadXRefTableEntry()
//
//	//pdfcpu.CreateContext(reftable, cfg)
//}

func toOdf(fileName string, words []string) error {
	//output, err := os.Create(fileName)
	//if err != nil {
	//	return err
	//}

	output, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	m := model.ModelFactory()

	//we need an empty model
	//	m := model.ModelFactory()
	//standard formatter
	fm := &mappers.Formatter{}
	//couple them
	fm.ConnectTo(m)
	//we want text
	fm.MimeType = xmlns.MimeText
	//yes we can
	fm.Init()
	//pretty simple

	fm.RegisterFont(timesNewRoman, timesNewRoman)

	for _, v := range words {
		size := rand.Intn(30) + 12
		fm.SetAttr(new(attr.TextAttributes).FontFace(timesNewRoman).Size(size))
		fm.WriteString(v)
	}

	//store file
	generators.GeneratePackage(m, nil, output, fm.MimeType)
	//cleanup
	return output.Close()
}
