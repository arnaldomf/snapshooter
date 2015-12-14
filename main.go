package main

import (
	"fmt"

	"bitbucket.org/dafiti/snap-shooter/connectors"
	"bitbucket.org/dafiti/snap-shooter/models"
)

func main() {
	createInput := &connectors.CreateConnectorInput{CloudType: "aws", Region: "sa-east-1"}
	conn, _ := connectors.CreateConnector(createInput)
	conn.Connect()
	fmt.Println(conn)
	name := "dft-sa-deploy01.aws.dafiticorp.com.br"
	insts, _ := conn.GetInstancesByName([]*string{&name})
	for _, inst := range insts {
		models.SetWindowHour(inst, "18")
		conn.CreateSnapshot(inst)
	}
}
