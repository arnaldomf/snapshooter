package controllers

import (
	"errors"

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
