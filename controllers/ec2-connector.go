package controllers

import (
	"errors"

	"bitbucket.org/dafiti/snap-shooter/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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
