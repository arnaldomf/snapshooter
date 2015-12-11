package models

import "github.com/aws/aws-sdk-go/service/ec2"
import "strings"

// BlockMapping has the Name (xvda/xvdb...) and the EBS ID
type BlockMapping struct {
	Name string
	ID   string
}

// EC2Instance holds all information needed to create a snapshot
type EC2Instance struct {
	Name          string
	Domain        string
	FQDN          string
	ID            string
	WindowHour    int
	BlockMappings []*BlockMapping
}

// GetEC2Instance returns a new EC2Inctance
func GetEC2Instance(instance *ec2.Instance) *EC2Instance {
	ec2Instance := new(EC2Instance)
	for _, tag := range instance.Tags {
		if *tag.Key == "Name" {
			values := strings.Split(*tag.Value, ".")
			ec2Instance.Name = values[0]
			ec2Instance.Domain = strings.Join(values[1:], ".")
			ec2Instance.FQDN = strings.Join(values, ".")
			break
		}
	}
	ec2Instance.ID = *instance.InstanceId
	for _, v := range instance.BlockDeviceMappings {
		bm := new(BlockMapping)
		bm.Name = *v.DeviceName
		bm.ID = *v.Ebs.VolumeId
		ec2Instance.BlockMappings = append(ec2Instance.BlockMappings, bm)
	}
	return ec2Instance
}

func (instance *EC2Instance) windowHour(hh int) {
	instance.WindowHour = hh
}
