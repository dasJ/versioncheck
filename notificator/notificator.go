package notificator

import (
	"fmt"
	"github.com/dasJ/versioncheck/config"
	"github.com/dasJ/versioncheck/moduleRunner"
	"os"
	"os/exec"
	"sync"
)

// Notify runs the notification script asynchronously for each changed version
func Notify(cfg config.VersioncheckConfig, res moduleRunner.RunnerResult) {
	wait := new(sync.WaitGroup)

	for _, item := range res.Changed {
		wait.Add(1)
		go func() {
			args := []string{
				item.Name,
				item.Module,
			}
			args = append(args, item.Tags...)
			cmd := exec.Command(cfg.Notificator, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				fmt.Printf("%s\n", err)
				wait.Done()
				return
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Printf("%s\n", err)
			}
			wait.Done()
		}()
	}
	wait.Wait()
}
