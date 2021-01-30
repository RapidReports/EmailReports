package getattachmentdata

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	gmailtoken "github.com/arctheowl/EmailReports/GmailToken"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

//GetAttachmentData :
//This function accesses the gmail api with a token and downloads the attachment file data.
func GetAttachmentData() string {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := gmailtoken.GetClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	r, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return "No labels found."
	}
	fmt.Println("Labels:")
	for _, l := range r.Labels {
		fmt.Printf("- %s\n", l.Name)
	}

	profile, err := srv.Users.GetProfile(user).Do()
	if err != nil {
		log.Fatalf("unable to get profile")
	}

	fmt.Println("Profile")
	fmt.Println(profile.EmailAddress)
	fmt.Println(profile.MessagesTotal)

	inbox, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatalf("unable to get messages")
	}

	fmt.Println("Messages:")

	fmt.Println(inbox.Messages[13].Id)

	messageid := inbox.Messages[1].Id
	mail, err := srv.Users.Messages.Get(user, messageid).Do()
	if err != nil {
		log.Fatalf("unable to get messages")
	}

	attachmentID := mail.Payload.Parts[1].Body.AttachmentId
	fmt.Println("Messages:")
	fmt.Println("Sent to: ", mail.Payload.Headers[0].Value)
	// /fmt.Println(attachmentID)

	attachment, err := srv.Users.Messages.Attachments.Get(user, messageid, attachmentID).Do()
	if err != nil {
		log.Fatalf("unable to get messages")
	}

	fmt.Println(attachment.Data)
	sd, e := base64.StdEncoding.DecodeString(attachment.Data)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(sd))

	return string(sd)
}
