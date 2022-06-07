package nexus_test

import "github.com/stretchr/testify/assert"

func in(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}

func (suite *NexusClientSuite) TestRealmsListAndUpdate() {
	realms, err := suite.client.Realms.List()
	assert.NoError(suite.T(), err)

	newRealm := "DockerToken"

	if in(newRealm, realms) == false {
		realms = append(realms, newRealm)
	}

	err = suite.client.Realms.Update(realms)
	assert.NoError(suite.T(), err)
}
