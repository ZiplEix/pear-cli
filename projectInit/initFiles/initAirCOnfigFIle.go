package initfiles

import "os"

var AirConfig = `
# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "."
tmp_dir = "tmp"

[build]
	# Just plain old shell command. You could use "make" as well.
	cmd = "go build -o ./tmp/main ./"
	# Binary file yields from "cmd".
	bin = "tmp/main"
	# Customize binary, can setup environment variables when run your app.
	full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
	# Watch these filename extensions.
	include_ext = ["go", "tpl", "tmpl", "html"]
	# Ignore these filename extensions or directories.
	exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
	# Watch these directories if you specified.
	include_dir = []
	# Exclude files.
	exclude_file = []
	# Exclude specific regular expressions.
	exclude_regex = ["_test\\.go"]
	# Exclude unchanged files.
	exclude_unchanged = true
	# Follow symlink for directories
	follow_symlink = true
	# This log file places in your tmp_dir.
	log = "air.log"
	# It's not necessary to trigger build each time file changes if it's too frequent.
	delay = 1000 # ms
	# Stop running old binary when build errors occur.
	stop_on_error = true
	# Send Interrupt signal before killing process (windows does not support this feature)
	send_interrupt = false
	# Delay after sending Interrupt signal
	kill_delay = 500 # ms
	# Add additional arguments when running binary (bin/full_bin). Will run './tmp/main hello world'.
	args_bin = ["hello", "world"]

[log]
	# Show log time
	time = false

[color]
	# Customize each part's color. If no color found, use the raw app log.
	main = "magenta"
	watcher = "cyan"
	build = "yellow"
	runner = "green"

[misc]
	# Delete tmp directory on exit
	clean_on_exit = true
`

func initAirConfigFile() {
	if _, err := os.Stat("./.air.toml"); err == nil {
		return
	}

	file, err := os.Create("./.air.toml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(AirConfig)
	if err != nil {
		panic(err)
	}
}
