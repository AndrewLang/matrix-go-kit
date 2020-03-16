package app

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	log "github.com/andrewlang/matrix-go-kit/log"
)

const (
	// VersionParseRegex regex for parse version string
	VersionParseRegex = `v?([0-9]+)(\.[0-9]+)?(\.[0-9]+)?` +
		`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
		`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?`

	// Dot char
	Dot = "."
)

/*
Version represent a version number/string
see semantic version definition https://semver.org/
*/
type Version struct {
	Major            uint64
	Minor            uint64
	Patch            uint64
	PrereleaseNumber uint64
	BuildNumber      uint64
	Prerelease       string
	Build            string
	raw              string
	regExpression    *regexp.Regexp
	logger           log.ILogger
}

// NewVersion create new instance
func NewVersion(versionString string) *Version {
	version := &Version{
		raw:           versionString,
		Major:         0,
		Minor:         0,
		Patch:         0,
		regExpression: regexp.MustCompile("^" + VersionParseRegex + "$"),
		logger:        CreateLogger("Version"),
	}

	return version
}

// Parse parse from the raw version string
func (ver *Version) Parse() *Version {
	parts := ver.regExpression.FindStringSubmatch(ver.raw)
	if parts == nil || len(parts) == 0 {
		return ver
	}

	ver.Major = ver.parseUnit(parts[1], "")
	ver.Minor = ver.parseUnit(parts[2], Dot)
	ver.Patch = ver.parseUnit(parts[3], Dot)
	ver.Prerelease = parts[5]
	ver.Build = parts[8]

	// ver.logger.Debug("Parse version string: ", ver.raw, parts)

	return ver
}

// String to string content
func (ver *Version) String() string {
	return ver.raw
}

// Compare with dest, value -1, 0, 1 means less, equal large than the dest version
func (ver *Version) Compare(dest *Version) int {

	if result := ver.compareNumber(ver.Major, dest.Major); result != 0 {
		return result
	}
	if result := ver.compareNumber(ver.Minor, dest.Minor); result != 0 {
		return result
	}
	if result := ver.compareNumber(ver.Patch, dest.Patch); result != 0 {
		return result
	}

	return ver.comparePrerelease(dest)
}

// LessThan dest
func (ver *Version) LessThan(dest *Version) bool {
	return ver.Compare(dest) < 0
}

// Equal dest
func (ver *Version) Equal(dest *Version) bool {
	return ver.Compare(dest) == 0
}

// GreaterThan dest
func (ver *Version) GreaterThan(dest *Version) bool {
	return ver.Compare(dest) > 0
}

// parseUnit string to uint64
func (ver *Version) parseUnit(value string, prefix string) uint64 {
	if value == "" {
		return 0
	}

	value = strings.TrimPrefix(value, prefix)
	result, _ := strconv.ParseUint(value, 10, 64)

	return result
}

// Sort versions with ascending order
func (ver *Version) Sort(versions []*Version) {
	sort.Slice(versions, func(i, j int) bool {
		return versions[i].LessThan(versions[j])
	})
}

// Latest get latest version in give versions
func (ver *Version) Latest(versions []*Version) *Version {
	ver.Sort(versions)
	length := len(versions)
	return versions[length-1]
}

// compareNumber format compare result to -1, 0, 1
func (ver *Version) compareNumber(first, second uint64) int {
	if first < second {
		return -1
	}
	if first > second {
		return 1
	}
	return 0
}

// comparePrerelease part
func (ver *Version) comparePrerelease(dest *Version) int {
	if ver.Prerelease == "" && dest.Prerelease == "" {
		return 0
	}

	if ver.Prerelease == "" {
		return 1
	}

	if dest.Prerelease == "" {
		return -1
	}

	// split the prelease versions by their part. The separator, per the spec is a .
	srcParts := strings.Split(ver.Prerelease, Dot)
	oparts := strings.Split(dest.Prerelease, Dot)

	// Find the longer length of the parts to know how many loop iterations to
	// go through.
	slen := len(srcParts)
	olen := len(oparts)

	length := slen
	if olen > slen {
		length = olen
	}

	// Iterate over each part of the prereleases to compare the differences.
	for i := 0; i < length; i++ {
		// Since the lentgh of the parts can be different we need to create
		// a placeholder. This is to avoid out of bounds issues.
		stemp := ""
		if i < slen {
			stemp = srcParts[i]
		}

		otemp := ""
		if i < olen {
			otemp = oparts[i]
		}

		d := ver.comparePrePart(stemp, otemp)
		if d != 0 {
			return d
		}
	}

	return 0
}

// comparePrePart prerelease part
func (ver *Version) comparePrePart(first, second string) int {
	// Fastpath if they are equal
	if first == second {
		return 0
	}

	// When s or o are empty we can use the other in an attempt to determine the response.
	if first == "" && second != "" {
		return -1
	}

	if second == "" && first != "" {
		return 1
	}

	/*
	 When comparing strings "99" is greater than "103". To handle
	 cases like this we need to detect numbers and compare them. According
	 to the semver spec, numbers are always positive. If there is a - at the
	 start like -99 this is to be evaluated as an alphanum. numbers always
	 have precedence over alphanum. Parsing as Uints because negative numbers
	 are ignored.
	*/

	oi, n1 := strconv.ParseUint(second, 10, 64)
	si, n2 := strconv.ParseUint(first, 10, 64)

	// The case where both are strings compare the strings
	if n1 != nil && n2 != nil {
		if first > second {
			return 1
		}
		return -1
	} else if n1 != nil {
		// o is a string and s is a number
		return -1
	} else if n2 != nil {
		// s is a string and o is a number
		return 1
	}
	// Both are numbers
	if si > oi {
		return 1
	}
	return -1

}
