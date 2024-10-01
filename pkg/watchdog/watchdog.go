package watchdog

import (
	"os"
	"io"
	"fmt"
	"os/user"
	"path/filepath"
	"time"
	"strings"

	// 3rd Party
	"github.com/brutalgg/cli"
)

type watcher struct {
	MonitorFolder string
	OutputFolder  string
	PollRate      time.Duration
	User          *user.User
}

func New(monitorfolder string, outputfolder string, pollrate int) *watcher {
	u, _ := user.Current()
	cli.Debug("Automatically detected username: %v", u.Username)

	p := time.Duration(pollrate) * time.Millisecond
	cli.Info("Polling Rate set to %v", p)

	w := filepath.Join("/Users", u.Username, monitorfolder)
	cli.Info("Watching directory: %v", w)

	o := filepath.Join("/Users", u.Username, outputfolder)
	cli.Info("Outputting IPAs to: %v", o)

	return &watcher{
		MonitorFolder: w,
		OutputFolder:  o,
		PollRate:      p,
		User:          u,
	}
}

func (w watcher) Watch() error {
	exists, err := checkDir(w.MonitorFolder)
	if err != nil {
		return err
	}
	if !exists{
		return fmt.Errorf("CRITICAL: Monitor folder not found")
	}

	err = ensureDir(w.OutputFolder)
	if err !=nil{
		return err
	}

	cli.Infoln("Press CTRL+C to exit the program")

	for {
		if err := filepath.WalkDir(w.MonitorFolder, w.grab); err != nil {
			cli.Errorln(err)
		}
		time.Sleep(w.PollRate)
	}
	return nil
}

func (w watcher) grab(path string, d os.DirEntry, err error) error {
	return nil
}

func checkDir(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// Directory does not exist
		return false, nil
	}
	if err != nil {
		// Some other error occurred
		return false, err
	}
	return info.IsDir(), nil // Return true if it's a directory
}

func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		cli.Debugln("Output directory not found. Trying to create...")
		err = os.MkdirAll(dir, os.ModePerm) // Use os.ModePerm for default permissions
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
		cli.Debugln("Output directory successfully created")
	}
	return nil
}