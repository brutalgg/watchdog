package watchdog

import (
	"os"
	"fmt"
	"os/user"
	"path/filepath"
	"errors"
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
	cli.Debug("Detected username: %v", u.Username)

	p := time.Duration(pollrate) * time.Millisecond
	cli.Debug("Polling Rate set to %v", p)

	w := filepath.Join("/Users", u.Username, monitorfolder)
	cli.Debug("Using watch directory: %v", w)

	o := filepath.Join("/Users", u.Username, outputfolder)
	cli.Debug("Output captured IPAs to the folder: %v", o)

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
		return errors.New("CRITICAL: Monitor folder not found")
	}

	cli.Infoln("Press CTRL+C to exit the program")

	for {
		if err := filepath.WalkDir(w.MonitorFolder, w.grab); err != nil {
			cli.Error("%s", err)
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