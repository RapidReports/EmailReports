package gmailfunctions

import (
	"fmt"
	"io/ioutil"
	"log"

	gmailtoken "github.com/arctheowl/EmailReports/GmailToken"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

//CheckMail - This function is to check the email account for emails in the inbox
//this function will return a map/list of the messsages in the inbox and their ID for use in the selection process
func CheckMail() map[int]string {
	fmt.Println("Checking the mail for the correct mail to download")
	user := "me"

	//This section gets a list of messages that can then be serached through an index
	srv := Createclient()

	meslist, err := srv.Users.Messages.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels: %v", err)
	}

	messageID := getID(meslist)
	fmt.Println("You have", len(meslist.Messages), "messages in your mailbox")
	//fmt.Println(messageID[0])

	return messageID
}

//This function returns a list of message ID's in an array that can then be searched through easier.
//It needs to be provided with a gmail.listmessages but you need the message ID to get the details on a specific
//mail, like to,from, attachments etc
func getID(meslist *gmail.ListMessagesResponse) map[int]string {

	list := make(map[int]string)

	for i := 0; i < 30; /*len(meslist.Messages)*/ i++ {
		//this section creates the map "list" by iterating over the messages getting the
		//message.id string and mapping it to the iteration number
		list[i] = meslist.Messages[i].Id
	}

	return list
}

//Createclient :
//This function returns a gmail.service/Client that can then be used to interact with the gmail api.
func Createclient() *gmail.Service {
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
