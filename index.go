package gomaskari

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func LetsGetShitDone(ctx context.Context, files ...string) {
	fmt.Println("number of files got to run")

	// start creating a build for every file
	folderName := "maskari_binaries/"

	for _, file := range files {
		buildName := folderName + file + "_build"
		cmd := exec.CommandContext(ctx, "go", "build", "-o", buildName, file)

		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		if err != nil {
			panic("error in building the file " + file)
		}
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error)

	for _, buildFile := range files {
		buildName := folderName + buildFile + "_build"

		wg.Add(1)

		go func(build string) {
			defer wg.Done()

			buildExecCommand := exec.CommandContext(ctx, build)

			buildExecCommand.Stdout = os.Stdout
			buildExecCommand.Stdin = os.Stdin

			err := buildExecCommand.Start()
			if err != nil {
				errChan <- err
			}

			err = buildExecCommand.Wait()
			if err != nil {
				errChan <- err
			}

		}(buildName)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			// if any script fails, here we can log accordingly
			fmt.Println(err)
		}
	}
}
