package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("ipconfig")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	r := bytes.NewReader(buf)
	r2 := transform.NewReader(r, simplifiedchinese.GBK.newDecoder())
	buf2, _ := ioutil.ReadAll(r2)
	fmt.Println(string(buf2))

}
