package mycontainer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var ErrNoContainerID = errors.New("no container ID detected")

// ContainerID returns current container ID or [ErrNoContainerID].
// It uses several internal approaches (cpuset/mountinfo)
// in order to calculate self container ID. It doesn't use hostname as container id (which is common approach)
// since hostname can be easily modified.
func ContainerID() (string, error) {
	if id, err := containerIDv1(); err == nil {
		return id, nil
	}
	if id, err := containerIDv2(); err == nil {
		return id, nil
	}

	return "", ErrNoContainerID
}

// looks for a "legacy" mount point in container. Was stable till 2020.
func containerIDv1() (string, error) {
	const path = `/proc/1/cpuset`

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("detect container ID: %w", err)
	}
	id := filepath.Base(strings.TrimSpace(string(data)))
	if id == "/" {
		return "", ErrNoContainerID
	}

	return id, nil
}

// v2 on demand regexp parsers.
//
//nolint:gochecknoglobals
var (
	v2Init   sync.Once
	v2Regexp *regexp.Regexp
)

// looks for container id in CGroup "modern" approach.
func containerIDv2() (string, error) {
	const path = `/proc/self/mountinfo`

	v2Init.Do(func() {
		v2Regexp = regexp.MustCompile(`containers/([[:alnum:]]{64})/`)
	})
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read mountinfo: %w", err)
	}
	matches := v2Regexp.FindStringSubmatch(string(content))
	if len(matches) == 0 {
		return "", ErrNoContainerID
	}

	return matches[len(matches)-1], nil
}
