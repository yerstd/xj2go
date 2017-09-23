package xj2go

import (
	"testing"
)

func Test_xmlToMap(t *testing.T) {
	fs := []string{
		"./testxml/[Content_Types].xml",
		"./testxml/xl/workbook.xml",
		"./testxml/xl/styles.xml",
		"./testxml/xl/sharedStrings.xml",
		"./testxml/xl/_rels/workbook.xml.rels",
		"./testxml/xl/theme/theme1.xml",
		"./testxml/xl/worksheets/sheet1.xml",
		"./testxml/docProps/app.xml",
		"./testxml/docProps/core.xml",
	}

	pkgname := "excel"
	for k, v := range fs {
		t.Run("xml to map"+string(k), func(t *testing.T) {
			xj := New(v, pkgname)
			// paths, err := xj.xmlToPaths()
			// fmt.Println(paths)
			_, err := xj.xmlToPaths()
			if err != nil {
				t.Errorf("xmlToMap() error = %v", err)
				return
			}
		})
	}
}

func Test_XMLToStruct(t *testing.T) {
	fs := []string{
		"./testxml/[Content_Types].xml",
		"./testxml/xl/workbook.xml",
		"./testxml/xl/styles.xml",
		"./testxml/xl/sharedStrings.xml",
		"./testxml/xl/_rels/workbook.xml.rels",
		"./testxml/xl/theme/theme1.xml",
		"./testxml/xl/worksheets/sheet1.xml",
		"./testxml/docProps/app.xml",
		"./testxml/docProps/core.xml",
	}

	pkgname := "excel"
	for k, v := range fs {
		t.Run("leaf paths "+string(k), func(t *testing.T) {
			// filename := pkgname + "/" + path.Base(v) + ".go"
			xj := New(v, pkgname)
			err := xj.XMLToStruct()
			if err != nil {
				t.Errorf("XMLToStruct() error = %v", err)
				return
			}
		})
	}
}
