package gmailfunctions

import (
	"fmt"
	"log"

	"google.golang.org/api/gmail/v1"
)

//Mail is the gathering of all the data needed to download the attachments from an email
type Mail struct {
	MessageID      string
	From           string
	AttachmentID   string
	AttachmentData string
}

//SelectMail will select the appropriate mail from the gmail account based on set requirements.
//This function will then be passed onto getattachment to download the file data.
func SelectMail() []Mail {

	fmt.Println("Selecting the mail that this program will run on...")
	mail := CheckMail()
	fmt.Println(mail[0])

	client := Createclient()

	user := "me"
	profile, err := client.Users.GetProfile(user).Do()
	if err != nil {
		log.Fatalf("unable to get profile")
	}

	fmt.Println("Profile you are connected to is:")
	fmt.Println(profile.EmailAddress)

	fmt.Println("Mail Info:")
	selectedmail := mailinfo(mail, client)

	//fmt.Println(selectedmail)

	return selectedmail
}

//This function iterates the mail(map of messageID's and client) and returns
func mailinfo(m map[int]string, client *gmail.Service) []Mail {
	selectedmail := []Mail{}

	for i := 0; i < len(m); i++ {
		//fmt.Println(m[i])
		messageinfo, err := client.Users.Messages.Get("me", m[i]).Do()
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(messageinfo.Payload.Headers[7])
		from := fromheader(messageinfo.Payload.Headers)

		if messageinfo.Payload.Headers[from].Value == "Liam Devlin <lidevlin95@gmail.com>" {
			selectedmail = append(selectedmail, Mail{MessageID: messageinfo.Id, From: messageinfo.Payload.Headers[from].Value})
		}

	}
	return selectedmail
}

//this function is to iterate through the message headers to find the "from" header
func fromheader(headers []*gmail.MessagePartHeader) int {
	for i := 0; i < len(headers); i++ {
		//fmt.Println(headers[i].Name)
		if headers[i].Name == "From" {
			//fmt.Println("Found IT HERE:")
			//fmt.Println("From:", headers[i].Value)
			return i
		}
	}
	return 0
}
