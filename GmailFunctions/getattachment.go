package gmailfunctions

import (
	"encoding/base64"
	"fmt"
	"log"

	"google.golang.org/api/gmail/v1"
)

//GetAttachmentData :
//This function accesses the gmail api with a token and downloads the attachment file data.
func GetAttachmentData(selectedmail []Mail) []Mail {

	Srv := Createclient()
	user := "me"

	for i := 0; i < len(selectedmail); i++ {

		attachmentID := getattachmentID(selectedmail[i], Srv)

		selectedmail[i].AttachmentID = attachmentID
		attachment, err := Srv.Users.Messages.Attachments.Get(user, selectedmail[i].MessageID, attachmentID).Do()
		if err != nil {
			log.Fatalf("unable to get messages")
		}

		sd, e := base64.StdEncoding.DecodeString(attachment.Data)
		if e != nil {
			fmt.Println(e)
		}
		selectedmail[i].AttachmentData = string(sd)

	}

	return selectedmail
}

func getattachmentID(m Mail, c *gmail.Service) string {

	mes, err := c.Users.Messages.Get("me", m.MessageID).Do()
	if err != nil {
		log.Fatalf("unable to get messages")
	}
	//fmt.Println(mes.Payload.Parts[1].Body.AttachmentId)
	//acttachmentheader(mes.Payload.PartId)
	return mes.Payload.Parts[1].Body.AttachmentId
}
