package connectors

import (
	"fmt"

	"github.com/dafiti/snapshooter/models"
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
	GetInstancesByName(gibni []*GetInstanceByNameInput) ([]models.Instance, error)
	GetInstanceByName(gibni *GetInstanceByNameInput) (models.Instance, error)
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

// GetInstanceByNameInput packs information needed to fetch an instance on cloud provider
type GetInstanceByNameInput struct {
	Name       string
	Region     string
	WindowHour int
}
