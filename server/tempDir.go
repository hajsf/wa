package server

import (
	"fmt"
	"os"
	"os/exec"
)

func createTemp() string {
	dirName := "waTemp"
	//os.TempDir() is => C:\Users\username\AppData\Local\Temp
	tmpDir := os.TempDir() + "\\" + dirName
	if p, err := os.Stat(tmpDir); os.IsNotExist(err) {
		err = os.Mkdir(tmpDir, 0755)
		defer os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Printf("err 2: %v", err)
		} else {
			fmt.Println("temp created at:", p)
			_, exists := os.LookupEnv("waTemp")
			if !exists {
				//
				//err = os.Setenv(`cv`, tmpDir)
				_ = exec.Command(`SETX`, `waTemp`, tmpDir).Run()
				if err != nil {
					fmt.Printf("Error: %s\n", err)
				}
				//	fmt.Println("tmpDir: ", tmpDir) */
			} else {
				fmt.Println("Env exisit")
			}
		}
	} else {
		fmt.Println("checking Env ")
		_, exists := os.LookupEnv("waTemp")
		if !exists {
			//
			//err = os.Setenv(`cv`, tmpDir)
			_ = exec.Command(`SETX`, `waTemp`, tmpDir).Run()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Println("Env created")
			}
			//	fmt.Println("tmpDir: ", tmpDir) */
		} else {
			fmt.Println("Env exisit")
		}
	}

	return tmpDir
}
