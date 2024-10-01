package watchdog

import (
	// "os"
	// "io"
	"fmt"
	"io/fs"
	"os/user"
	"path/filepath"
	"time"
	// "strings"

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

func (w watcher) Watch() {
	cli.Infoln("Press CTRL+C to exit the program")

	for {
		if err := filepath.WalkDir(w.MonitorFolder, w.grab); err != nil {
			cli.Error("%s", err)
		}
		time.Sleep(w.PollRate)
	}
}

func (w watcher) grab(path string, d fs.DirEntry, err error) error {
	return nil
}
