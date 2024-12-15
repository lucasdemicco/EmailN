package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	name := "Campaign X"
	content := "Body"
	contacts := []string{"teste1@teste.com", "teste2@teste.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.NotNil(campaign.Name)
	assert.NotEmpty(campaign.Name)
	assert.Equal(campaign.Content, content)
	assert.NotNil(campaign.Content)
	assert.NotEmpty(campaign.Content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotEmpty(campaign.Contacts)
}
