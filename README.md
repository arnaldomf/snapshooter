# Snapshooter
[![Build Status](https://img.shields.io/travis/dafiti/snapshooter/master.svg?style=flat-square)](https://travis-ci.org/dafiti/snapshooter)
[![Coverage Status](https://img.shields.io/coveralls/dafiti/snapshooter/master.svg?style=flat-square)](https://coveralls.io/github/dafiti/snapshooter?branch=master)

This programs schedules instances snapshots

# configuration

The credentials are expected to be exported as Environment Variables
**AWS_SECRET_ACCESS_KEY** and **AWS_ACCESS_KEY_ID**

The configuration file uses the json format, and must have the fields:

1. provider name (string)
2. instances (list) - instance list to take snapshots on each provider

Each instance has the fields:

1. name (string) - instance's domain name
2. window_hour (int) - hour in which the snapshot will be taken
3. id (int) - for digitalocean is possible to fetch an instance by ID
4. region (string) - for aws the region name is necessary

## file example

    {
      "droplets": [{"id": 1, "name": "Goku", "window_hour": 10 }],
      "ec2": [{"name": "Vegeta", "region": "us-east-1", "window_hour": 20}]
    }
