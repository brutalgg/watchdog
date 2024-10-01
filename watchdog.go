package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	uName string
)

// main
// TODO Command line flags and associated parsing
// 	Allow user to specify location (default: /tmp/ipawatchdog)
//  Specify log debug or not
//
// TODO Useful Help and program about
func main() {

	fmt.Printf(`
._____________  _____     __      __         __         .__         .___             
|   \______   \/  _  \   /  \    /  \_____ _/  |_  ____ |  |__    __| _/____   ____  
|   ||     ___/  /_\  \  \   \/\/   /\__  \\   __\/ ___\|  |  \  / __ |/  _ \ / ___\ 
|   ||    |  /    |    \  \        /  / __ \|  | \  \___|   Y  \/ /_/ (  <_> ) /_/  >
|___||____|  \____|__  /   \__/\  /  (____  /__|  \___  >___|  /\____ |\____/\___  / 
v0.1		     \/         \/        \/          \/     \/      \/     /_____/  `)
	fmt.Println()
	u, _ := user.Current()
	uName = u.Username
	source := filepath.Join("/Users", uName, `/Library/Group Containers/K36BKF7T3D.group.com.apple.configurator/Library/Caches/`)
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})
	log.SetLevel(log.DebugLevel)

	log.Debugf("Detecting username: %s", uName)
	log.Debugf("Using Source Directory: %s", source)
	for {
		if err := filepath.Walk(source, moveIpa); err != nil {
			log.Debug(err)
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

func moveIpa(path string, fi os.FileInfo, err error) error {

	if filepath.Ext(strings.TrimSpace(path)) == ".ipa" {
		dest := filepath.Join("/Users", uName, "/Desktop", filepath.Base(strings.TrimSpace(path)))
		if _, err := os.Stat(dest); !os.IsNotExist(err) {
			// the destination already exists
			return nil
		}
		log.Info("Detected Uncaptured IPA. Capturing...")
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}
		log.Infof("Captured IPA: %s", filepath.Base(strings.TrimSpace(path)))
		return out.Close()

	}
	return nil
}
