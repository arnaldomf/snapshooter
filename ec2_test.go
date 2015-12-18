package main

import (
	"testing"

	"github.com/bitly/go-simplejson"
)

func TestIFEC2ClientIsNotNil(t *testing.T) {
	awsConn := new(AWS)
	client := awsConn.Client("us-east-1")
	if client == nil {
		t.Error("Expected an object, got nil")
	}
}

func TestEC2GetInstancesWithNoInstance(t *testing.T) {
	awsConn := new(AWS)
	config, _ := simplejson.NewJson([]byte(`{}`))
	awsConn.config = config
	if awsConn.GetInstances() != nil {
		t.Error("Expected nil")
	}
}

func TestEC2GetInstancesWithInstanceNotFound(t *testing.T) {
	awsConn := new(AWS)
	config, _ := simplejson.NewJson([]byte(`{
    "ec2": [{"name": "test.aws.teste.com.br", "window_hour": 10, "region": "sa-east-1"}]
    }`))
	awsConn.config = config
	instances := awsConn.GetInstances()
	if instances == nil {
		t.Error("Expected [], received", instances)
	}
}
