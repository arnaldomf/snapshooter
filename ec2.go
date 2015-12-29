package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	j "github.com/bitly/go-simplejson"
)

// AWS holds a connector to aws ec2 service per region
type AWS struct {
	config       *j.Json
	clientRegion map[string]*ec2.EC2
}

// GetEC2InstanceByName will gather information from ec2Instances to fetch it from aws
func (a *AWS) GetEC2InstanceByName(idx int) *EC2Instance {
	name, _ := a.config.GetIndex(idx).Get("name").String()
	windowHour, _ := a.config.GetIndex(idx).Get("window_hour").Int()
	region, _ := a.config.GetIndex(idx).Get("region").String()
	describeInstancesInput := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name:   aws.String("tag:Name"),
				Values: []*string{aws.String(name)},
			},
		},
	}

	describeOutput, err := a.Client(region).DescribeInstances(describeInstancesInput)

	if err != nil {
		return nil
	}
	reservations := describeOutput.Reservations
	if len(reservations) == 0 {
		return nil
	}

	ec2Instance := new(EC2Instance)
	ec2Instance.name = name
	ec2Instance.windowHour = windowHour
	ec2Instance.id = *reservations[0].Instances[0].InstanceId
	ec2Instance.svc = a.Client(region)
	ec2Instance.setBlockDevices(reservations[0].Instances[0])
	return ec2Instance
}

// Client creates a new connection to ec2 if one does not exist yet
func (a *AWS) Client(region string) *ec2.EC2 {
	if a.clientRegion == nil {
		a.clientRegion = make(map[string]*ec2.EC2)
	}
	if _, ok := a.clientRegion[region]; !ok {
		a.clientRegion[region] = ec2.New(session.New(), &aws.Config{Region: aws.String(region)})
	}
	return a.clientRegion[region]
}

// GetInstances returns a slice of EC2Instances found on aws based on the config file
func (a *AWS) GetInstances() []Instance {
	var instances []Instance

	for idx := range a.config.MustArray() {
		instances = append(instances, a.GetEC2InstanceByName(idx))
	}

	return instances
}

// EC2BlockDevice is a blockdevice for EC2Instance
type EC2BlockDevice struct {
	ID   string
	Name string
}

// EC2Instance is an aws instance
type EC2Instance struct {
	name         string
	id           string
	windowHour   int
	blockdevices []*EC2BlockDevice
	svc          *ec2.EC2
}

func (ec2Instance *EC2Instance) setBlockDevices(instance *ec2.Instance) {
	for _, blockDevice := range instance.BlockDeviceMappings {
		ec2Block := new(EC2BlockDevice)
		ec2Block.Name = *blockDevice.DeviceName
		ec2Block.ID = *blockDevice.Ebs.VolumeId
		ec2Instance.blockdevices = append(ec2Instance.blockdevices, ec2Block)
	}
}

// Snapshot creates an snapshot from the receiver instance. It will create a
// new snapshot for each volume attached to the instance
func (ec2Instance *EC2Instance) Snapshot() bool {
	if len(ec2Instance.blockdevices) == 0 {
		return false
	}
	now := time.Now().Format("2006-01-02")
	for _, volume := range ec2Instance.blockdevices {
		description := fmt.Sprintf("%s:(%s) @%s", ec2Instance.name, volume.Name, now)
		params := &ec2.CreateSnapshotInput{
			VolumeId:    aws.String(volume.ID),
			Description: aws.String(description),
		}
		_, err := ec2Instance.svc.CreateSnapshot(params)
		if err != nil {
			return false
		}
	}
	return true
}
