package gmailfunctions

import (
	"fmt"
	"io/ioutil"
	"log"

	gmailtoken "github.com/arctheowl/EmailReports/GmailToken"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

//CheckMail - This function is to check the email account for the requirements to select the correct emails
func CheckMail() {
	fmt.Println("Checking the mail for the correct mail to download")
	user := "me"

	//This section gets a list of messages that can then be serached through an index
	srv := createclient()

	mesamount, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mesamount)

	meslist, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}
	//if len(r.Labels) == 0 {
	messageid := meslist.Messages[6].Id

	mail, err := srv.Users.Messages.Get(user, messageid).Do()
	if err != nil {
		log.Fatalf("unable to get messages")
	}

	fmt.Println(mail.Payload.Headers[4])

}

func mailloop() {

}

//This function returns a gmail.service/Client that can then be used to interact with the gmail api.
func createclient() *gmail.Service {
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

	return srv
}
