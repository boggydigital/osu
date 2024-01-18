package main

import (
	"bufio"
	"bytes"
	"os/exec"
)

const (
	gitBin           = "git"
	gitTagCmd        = "tag"
	gitTagDeleteFlag = "-d"
	gitPushCmd       = "push"
	originArg        = "origin"
)

func execGitCmd(args ...string) (*bytes.Buffer, error) {

	gitPath := ""

	if path, err := exec.LookPath(gitBin); err == nil {
		gitPath = path
	} else {
		return nil, err
	}

	var bts []byte
	buf := bytes.NewBuffer(bts)
	cmd := exec.Command(gitPath, args...)
	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return buf, nil
}

func GitLatestTag() (*SemVer, error) {

	latestSv := &SemVer{}

	buf, err := execGitCmd(gitTagCmd)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		if sv, err := ParseSemVerTag(scanner.Text()); err == nil {
			if latestSv.Less(sv) {
				latestSv = sv
			}
		} else {
			return nil, err
		}
	}

	return latestSv, nil
}

func GitDeleteTag(sv *SemVer) error {
	_, err := execGitCmd(gitTagCmd, gitTagDeleteFlag, sv.String())
	return err
}

func GitTag(sv *SemVer) error {
	_, err := execGitCmd(gitTagCmd, sv.String())
	return err
}

func GitIncrementLatestTag() (*SemVer, error) {
	lsvt, err := GitLatestTag()
	if err != nil {
		return nil, err
	}

	lsvt.Increment()

	return lsvt, GitTag(lsvt)
}

func GitPushOrigin(sv *SemVer) error {
	_, err := execGitCmd(gitPushCmd, originArg, sv.String())
	return err
}
