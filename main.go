package main

import (
	"fmt"

	getattachmentdata "github.com/arctheowl/EmailReports/GetAttachmentData"
)

func main() {
	CsvData := getattachmentdata.GetAttachmentData()
	fmt.Println(CsvData)
}
