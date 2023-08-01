package main

import (
	"fmt"
	"github.com/ives22/stool/cmd"
)

func main() {
	//sshRsaFile := "/Users/liyanjie/.ssh/id_rsa"
	//_, err := os.Stat(sshRsaFile)
	//if os.IsNotExist(err) {
	//	fmt.Printf("File %s does not exist\n", sshRsaFile)
	//} else {
	//	cmd := exec.Command("cat", "/Users/liyanjie/.ssh/id_rsa.pub")
	//	output, err := cmd.Output()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Print(string(output))
	//}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}