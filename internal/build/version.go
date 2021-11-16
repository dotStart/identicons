package build

import (
	"strconv"
	"time"
)

const shortCommitHashLength = 7

var commitHash = ""
var version = "0.0.0"

var timestampRaw = "0"
var timestamp, _ = strconv.ParseInt(timestampRaw, 10, 64)
var timestampParsed = time.Unix(timestamp, 0)

func HasCommitHash() bool {
	return len(commitHash) != 0
}

func CommitHash() string {
	return commitHash
}

func ShortCommitHash() string {
	if len(commitHash) < shortCommitHashLength {
		return commitHash
	}

	return commitHash[0:shortCommitHashLength]
}

func Version() string {
	return version
}

func FullVersion() string {
	ver := Version()

	if HasCommitHash() {
		ver += "+git-" + ShortCommitHash()
	} else {
		ver += "+dev"
	}

	return ver
}

func BuildTimestamp() time.Time {
	return timestampParsed
}
