package main

import (
	"fmt"

	gfunc "github.com/arctheowl/EmailReports/GmailFunctions"
)

func main() {

	mail := gfunc.SelectMail()
	CsvData := gfunc.GetAttachmentData(mail)
	fmt.Println(CsvData)
}
