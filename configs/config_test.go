package configs_test

import (
	"fmt"

	"github.com/dafiti/snapshooter/configs"
)

func Example() {
	config1, err := configs.CreateConfig("./example1.toml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config1.Region)
	fmt.Printf("%s.%s\n", "instance1", config1.Instances["instance1"].Domain)
	fmt.Printf("%s.%s\n", "instance2", config1.Instances["instance1"].Domain)

	configs.ClearConfig()

	config1, err = configs.CreateConfig("./example2.toml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output:
	// sa-east-1
	// instance1.dafiti.com.br
	// instance2.dafiti.com.br
	// instance1: failed to parse window_hour
}
