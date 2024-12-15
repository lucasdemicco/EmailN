package campaign

import "time"

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

type Contact struct {
	Email string
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contactsSlice := make([]Contact, len(emails))
	for index, email := range emails {
		contactsSlice[index].Email = email
	}

	return &Campaign{
		ID:        "1",
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contactsSlice,
	}
}
