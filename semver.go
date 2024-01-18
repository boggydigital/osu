package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	semVerTagPfx        = "v"
	incrementMinorPatch = 100
	incrementMajorMinor = 100
)

var ErrInvalidSemVerTag = errors.New("invalid semver tag")

// SemVer https://semver.org
type SemVer struct {
	Major int
	Minor int
	Patch int
	Label string
}

func ParseSemVerTag(tag string) (*SemVer, error) {
	if !strings.HasPrefix(tag, semVerTagPfx) {
		return nil, ErrInvalidSemVerTag
	}

	sv := &SemVer{}

	tagSansPfx := strings.TrimPrefix(tag, semVerTagPfx)

	semVerStr := tagSansPfx
	semVerStr, sv.Label, _ = strings.Cut(semVerStr, "-")

	if parts := strings.Split(semVerStr, "."); len(parts) == 3 {

		if major64, err := strconv.ParseInt(parts[0], 10, 32); err == nil {
			sv.Major = int(major64)
		} else {
			return nil, err
		}

		if minor64, err := strconv.ParseInt(parts[1], 10, 32); err == nil {
			sv.Minor = int(minor64)
		} else {
			return nil, err
		}

		if patch64, err := strconv.ParseInt(parts[2], 10, 32); err == nil {
			sv.Patch = int(patch64)
		} else {
			return nil, err
		}
	}

	return sv, nil
}

func (sv *SemVer) String() string {
	svt := fmt.Sprintf("v%d.%d.%d", sv.Major, sv.Minor, sv.Patch)
	if sv.Label != "" {
		svt = fmt.Sprintf("%s-%s", svt, sv.Label)
	}
	return svt
}

func (sv *SemVer) Less(asv *SemVer) bool {
	if sv.Major < asv.Major {
		return true
	} else if sv.Minor < asv.Minor {
		return true
	} else if sv.Patch < asv.Patch {
		return true
	} else if sv.Label != "" && asv.Label == "" {
		return true
	}
	return false
}

func (sv *SemVer) Increment() {
	sv.Patch += 1
	if sv.Patch == incrementMinorPatch {
		sv.Minor += 1
		sv.Patch = 0
	}
	if sv.Minor == incrementMajorMinor {
		sv.Major += 1
		sv.Minor = 0
	}
}
