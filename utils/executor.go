package utils

import "os/exec"

func SingleOut(one, two string) string {
	cmd := exec.Command(one, two)
	cmdOut, err := cmd.Output()
	if err != nil {
		return string(cmdOut)
	}
	return string(cmdOut)
}

// func SevenOut(one, two, three, four, five, six, seven string) []string {
// 	cmd := exec.Command(one, two)
// 	cmdOut, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println(cmdOut)
// 	}
// 	pipeOne := exec.Command(three, four, five)
// 	pipeTwo := exec.Command(six, seven)
// 	pipeOneOut, err := pipeOne.StdoutPipe()
// 	if err != nil {
// 		fmt.Println(pipeOneOut)
// 	}
// 	pipeOne.Start()
// 	pipeTwo.Stdin = pipeOneOut
// 	pipeTwoOut, err := pipeOne.Output()
// 	if err != nil {
// 		fmt.Println(pipeTwo)
// 	}
// 	cmdStr := string(pipeTwoOut)
// 	cmdSlice := strings.Split(cmdStr, ",")
// 	return cmdSlice
// }
