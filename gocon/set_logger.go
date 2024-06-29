package gocon

import "path"

var log_dir = path.Join("logs")

func setLogName(logName string) string {
	return path.Join(log_dir, logName)
}
