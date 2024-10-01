# watchdog

Basic usage instructions:

```bash
Usage:
  watchdog [flags]

Flags:
  -q, --banner            Disables banner
  -d, --delay int         Time in milliseconds for polling the filesystem (default 2000)
  -h, --help              help for watchdog
  -l, --loglevel string   Include verbose messages from program execution [error, warn, info, debug] (default "info")
  -m, --monitor string    Location to monitor for cached IPA files (default "/Library/Group Containers/K36BKF7T3D.group.com.apple.configurator/Library/Caches/")
  -o, --out string        Location in the user's home directory for dumping files (default "/Desktop/watchdog")
  -v, --version           version for watchdog
```

## Purpose

watchdog is a utility which assists in pulling IPA files from an OSX filesystem when being cached by Configurator.

## Sample Execution

```bash
> watchdog -q
[+] Initalizing watchdog instance
[+] Polling Rate set to 2s
[+] Watching directory: /Users/pentester/Library/Group Containers/K36BKF7T3D.group.com.apple.configurator/Library/Caches
[+] Outputting IPAs to: /Users/pentester/Desktop/watchdog
[+] Press CTRL+C to exit the program
[+] Found new IPA. Copying...
[+] Successfully copied IPA: MyApp.ipa
[+] Found new IPA. Copying...
[+] Successfully copied IPA: Customer.ipa
```
