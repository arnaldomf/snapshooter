package connectors

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/dafiti/snapshooter/models"
)

// EC2Connector handles the communication with AWS
type EC2Connector struct {
	ServiceConn *ec2.EC2
	Region      string
}

// ConnectionExists checks if a ServiceConn was previously created
func (ec2Connector *EC2Connector) ConnectionExists() bool {
	return ec2Connector.ServiceConn != nil
}

// Connect to the aws
func (ec2Connector *EC2Connector) Connect() error {
	if ec2Connector.ConnectionExists() {
		return errors.New("EC2Connector.Connect: Can't recreate a connection")
	}
	ec2Connector.ServiceConn = ec2.New(session.New(), &aws.Config{
		Region: aws.String(ec2Connector.Region),
	})
	return nil
}

// GetInstancesByName returns EC2Instance with tag name
func (ec2Connector *EC2Connector) GetInstancesByName(names []*string) ([]models.Instance, error) {
	if !ec2Connector.ConnectionExists() {
		return nil, errors.New("EC2Connector.GetInstanceByName: Didnt acquired connection")
	}
	describeInput := new(ec2.DescribeInstancesInput)
	filter := new(ec2.Filter)
	filter.Name = aws.String("tag:Name")
	filter.Values = names
	describeInput.Filters = append(describeInput.Filters, filter)
	describeOutput, err := ec2Connector.ServiceConn.DescribeInstances(describeInput)
	if err != nil {
		return nil, err
	}
	if len(describeOutput.Reservations) == 0 {
		return nil, errors.New("Instance not found")
	}
	var instances []models.Instance
	for _, reservation := range describeOutput.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, models.GetEC2Instance(instance))
		}
	}

	return instances, nil
}

// CreateSnapshot creates a snapshot from an models.EC2Instance
func (ec2Connector *EC2Connector) CreateSnapshot(instance models.Instance) error {
	ec2inst, ok := instance.(*models.EC2Instance)
	if !ok {
		return errors.New("Instance is not from type *models.EC2Instance")
	}
	now := time.Now()
	timeStr := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	for _, volume := range ec2inst.BlockMappings {
		name := fmt.Sprintf("%s:(%s)", ec2inst.Name, volume.Name)
		description := fmt.Sprintf("%s: volume %s @%s", ec2inst.Name, volume.Name, timeStr)
		snapshotInput := new(ec2.CreateSnapshotInput)
		snapshotInput.VolumeId = aws.String(volume.ID)
		snapshotInput.Description = aws.String(description)
		snap, err := ec2Connector.ServiceConn.CreateSnapshot(snapshotInput)
		if err != nil {
			return err
		}
		createTagsInput := &ec2.CreateTagsInput{
			Resources: []*string{snap.SnapshotId},
			Tags: []*ec2.Tag{
				&ec2.Tag{
					Key:   aws.String("Name"),
					Value: aws.String(name),
				},
			},
		}
		if _, err := ec2Connector.ServiceConn.CreateTags(createTagsInput); err != nil {
			return err
		}
	}
	return nil
}
