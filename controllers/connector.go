package controllers

import (
	"fmt"

	"bitbucket.org/dafiti/snap-shooter/models"
)

const (
	errorTag = "Controller"
)

// CreateConnectorInput encapsulates information to create a new connection
type CreateConnectorInput struct {
	CloudType string
	Region    string
}

// Connector abstract all other connectors
type Connector interface {
	Connect() error
	GetInstancesByName(names []*string) ([]models.Instance, error)
	CreateSnapshot(instance models.Instance) error
}

func createAWSConnector(
	createConnectorInput *CreateConnectorInput) *EC2Connector {
	connector := &EC2Connector{
		Region: createConnectorInput.Region,
	}
	return connector
}

// CreateConnector decides which connector to use (only EC2Connector supported)
func CreateConnector(createConnectorInput *CreateConnectorInput) (Connector, error) {
	var connector Connector
	switch createConnectorInput.CloudType {
	case "aws":
		connector = createAWSConnector(createConnectorInput)
	default:
		return nil, fmt.Errorf("%s.CreateConnector: %s is not a valid connector\n",
			errorTag, createConnectorInput.CloudType)
	}
	return connector, nil
}
