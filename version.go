package version

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const SemVerRegExpRaw string = `v?([0-9]+(\.[0-9]+)*?)` +
	`(-([0-9]+[0-9A-Za-z\-~]*(\.[0-9A-Za-z\-~]+)*)|(-([A-Za-z\-~]+[0-9A-Za-z\-~]*(\.[0-9A-Za-z\-~]+)*)))?` +
	`(\+([0-9A-Za-z\-~]+(\.[0-9A-Za-z\-~]+)*))?` +
	`?`

var versionRegExp *regexp.Regexp

func init() {
	versionRegExp = regexp.MustCompile("^" + SemVerRegExpRaw + "$")
}

type Version struct {
	major       int64
	minor       int64
	patch       int64
	identifier  string
	metadata    string
	PrettyPrint bool
}

func (v *Version) String() string {
	if v.PrettyPrint {
		return v.prettyString()
	} else {
		return v.string()
	}
}

func (v *Version) prettyString() (version string) {
	version = fmt.Sprintf("v%v.%v.%v", v.major, v.minor, v.patch)

	if v.identifier != "" {
		version = fmt.Sprintf("%s (%s)", version, v.identifier)
	}

	if v.metadata != "" {
		meta := strings.Split(v.metadata, "-")

		if len(meta) == 2 {
			// try to parse two components for date and git commit hash
			if t, err := time.Parse("20060102", meta[0]); err == nil {
				version = fmt.Sprintf("%s (%s %s)", version, t.Format("2006-01-02"), meta[1])
				return
			} else if t, err := time.Parse("20060102", meta[1]); err == nil {
				version = fmt.Sprintf("%s (%s %s)", version, t.Format("2006-01-02"), meta[0])
				return
			}
		}

		// print meta as it is
		version = fmt.Sprintf("%s (%s)", version, v.metadata)
	}

	return
}

func (v *Version) string() (version string) {
	version = fmt.Sprintf("v%v.%v.%v", v.major, v.minor, v.patch)

	if v.identifier != "" {
		version = fmt.Sprintf("%s-%s", version, v.identifier)
	}

	if v.metadata != "" {
		version = fmt.Sprintf("%s+%s", version, v.metadata)
	}

	return
}

func NewVersion(v string) (*Version, error) {
	matches := versionRegExp.FindStringSubmatch(v)
	if matches == nil {
		return nil, errors.New("malformed version string")
	}

	// parse the version number
	versionSegment := strings.Split(matches[1], ".")
	// integer slice to hold version numbers
	versionSegmentInt := make([]int64, len(versionSegment))

	for i, vs := range versionSegment {
		val, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return nil, errors.New("unable to parse version")
		}

		versionSegmentInt[i] = val
	}

	// if we have less than 3 segments, pad empty segments with 0
	for i := len(versionSegmentInt); i < 3; i++ {
		versionSegmentInt = append(versionSegmentInt, 0)
	}

	pre := matches[7]
	if pre == "" {
		pre = matches[4]
	}

	return &Version{
		major:      versionSegmentInt[0],
		minor:      versionSegmentInt[1],
		patch:      versionSegmentInt[2],
		identifier: pre,
		metadata:   matches[10],
	}, nil
}
