# snapshooter

This programs schedules instances snapshots

# configuration

The credentials are expected to be exported as Environment Variables
**AWS_SECRET_ACCESS_KEY** and **AWS_ACCESS_KEY_ID**

The configuration file uses the toml format, and must have the fields:

1. region (string) - aws region where the program will find the instances
2. instances (list) - instance list to take snapshots

Each instance has the fields:

1. domain (string) - instance's domain name
2. window_hour (string) - hour in which the snapshot will be taken

**It is expected that each instance have a Name tag in the format
instanceName.domainName**

## file example

    region = "sa-east-1"
    [instances]
        [instances.instance1]
            domain = "dafiti.com.br"
            window_hour = "23"
        [instances.instance2]
            domain = "dafiti.com.br"
            window_hour = "00"
