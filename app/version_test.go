package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	ver := NewVersion("")

	assert.NotNil(t, ver)
	assert.Equal(t, "", ver.raw)
	assert.Equal(t, uint64(0), ver.Major)
	assert.Equal(t, uint64(0), ver.Minor)
	assert.Equal(t, uint64(0), ver.Patch)
}

func TestVersionParse(t *testing.T) {
	ver := NewVersion("1.0.0-alpha.1+sha.001").Parse()

	assert.NotNil(t, ver)
	assert.Equal(t, uint64(1), ver.Major)
	assert.Equal(t, uint64(0), ver.Minor)
	assert.Equal(t, uint64(0), ver.Patch)
	assert.Equal(t, "alpha.1", ver.Prerelease)
	assert.Equal(t, "sha.001", ver.Build)

	ver1 := NewVersion("2.1.2-0.3.7").Parse()
	assert.NotNil(t, ver1)
	assert.Equal(t, uint64(2), ver1.Major)
	assert.Equal(t, uint64(1), ver1.Minor)
	assert.Equal(t, uint64(2), ver1.Patch)

	ver2 := NewVersion("10.3.4-x.7.z.94").Parse()
	assert.NotNil(t, ver2)
	assert.Equal(t, uint64(10), ver2.Major)
	assert.Equal(t, uint64(3), ver2.Minor)
	assert.Equal(t, uint64(4), ver2.Patch)

	ver3 := NewVersion("5.20.60+sha.94").Parse()
	assert.NotNil(t, ver3)
	assert.Equal(t, uint64(5), ver3.Major)
	assert.Equal(t, uint64(20), ver3.Minor)
	assert.Equal(t, uint64(60), ver3.Patch)

	ver4 := NewVersion("").Parse()
	assert.Empty(t, ver4.String())
}

func TestVersionCompare(t *testing.T) {
	ver1 := NewVersion("1.0.0-alpha").Parse()
	ver2 := NewVersion("1.0.0-alpha.1").Parse()
	ver3 := NewVersion("1.0.0-alpha.beta").Parse()
	ver4 := NewVersion("1.0.0-beta").Parse()
	ver5 := NewVersion("1.0.0-beta.2").Parse()
	ver6 := NewVersion("1.0.0-beta.11").Parse()
	ver7 := NewVersion("1.0.0-rc.1").Parse()
	ver8 := NewVersion("1.0.0").Parse()

	assert.True(t, ver1.LessThan(ver2))
	assert.True(t, ver2.LessThan(ver3))
	assert.True(t, ver3.LessThan(ver4))
	assert.True(t, ver4.LessThan(ver5))
	assert.True(t, ver5.LessThan(ver6))
	assert.True(t, ver6.LessThan(ver7))
	assert.True(t, ver7.LessThan(ver8))

	assert.True(t, ver8.GreaterThan(ver7))
	assert.True(t, ver7.GreaterThan(ver6))
	assert.True(t, ver6.GreaterThan(ver5))
	assert.True(t, ver5.GreaterThan(ver4))
	assert.True(t, ver4.GreaterThan(ver3))
	assert.True(t, ver3.GreaterThan(ver2))
	assert.True(t, ver2.GreaterThan(ver1))
}

func TestVersionCompareRelease(t *testing.T) {
	versions := []struct {
		first  string
		second string
		expect int
	}{
		{"1.0.0-alpha", "1.0.0-beta", -1},
		{"1.0.0-alpha.3.1", "1.0.0-alpha.3.2", -1},
		{"1.0.0", "1.0.0", 0},
		{"1.0.0", "1.0.0-beta", 1},
		{"1.0.0-alpha", "1.0.0", -1},
		{"1.0.0-beta.2", "1.0.0-beta.11", -1},
		{"1.0.0-beta.4", "1.0.0-beta.-2", -1},
		{"1.0.0-beta.-2", "1.0.0-beta.-3", -1},
		{"1.0.0-beta.foo", "1.0.0-beta", 1},
		{"1.2.3", "1.5.1", -1},
		{"2.2.3", "1.5.1", 1},
		{"2.2.3", "2.2.2", 1},
		{"3.2-beta", "3.2-beta", 0},
		{"1.3", "1.1.4", 1},
		{"4.2", "4.2-beta", 1},
		{"4.2-beta", "4.2", -1},
		{"4.2-alpha", "4.2-beta", -1},
		{"4.2-alpha", "4.2-alpha", 0},
		{"4.2-beta.2", "4.2-beta.1", 1},
		{"4.2-beta2", "4.2-beta1", 1},
		{"4.2-beta", "4.2-beta.2", -1},
		{"4.2-beta", "4.2-beta.foo", -1},
		{"4.2-beta.2", "4.2-beta", 1},
		{"4.2-beta.foo", "4.2-beta", 1},
		{"1.2+bar", "1.2+baz", 0},
		{"1.0.0-beta.4", "1.0.0-beta.-2", -1},
		{"1.0.0-beta.-2", "1.0.0-beta.-3", -1},
		{"1.0.0-beta.-3", "1.0.0-beta.5", 1},
		{"1.2.3", "1.5.1", -1},
		{"2.2.3", "1.5.1", 1},
		{"3.2-beta", "3.2-beta", 0},
		{"3.2.0-beta.1", "3.2.0-beta.5", -1},
		{"3.2-beta.4", "3.2-beta.2", 1},
		{"7.43.0-SNAPSHOT.99", "7.43.0-SNAPSHOT.103", -1},
		{"7.43.0-SNAPSHOT.FOO", "7.43.0-SNAPSHOT.103", 1},
		{"7.43.0-SNAPSHOT.99", "7.43.0-SNAPSHOT.BAR", -1},
	}

	for _, item := range versions {
		ver1 := NewVersion(item.first).Parse()
		ver2 := NewVersion(item.second).Parse()

		assert.Equal(t, item.expect, ver1.Compare(ver2))
	}
}
func TestVersionCompareEqual(t *testing.T) {

	ver1 := NewVersion("1.0.0-alpha.1").Parse()
	ver2 := NewVersion("1.0.0-alpha.1").Parse()

	assert.True(t, ver1.Equal(ver2))
	assert.True(t, ver2.Equal(ver1))
}

func TestVersionString(t *testing.T) {
	ver := NewVersion("1.0.0-alpha.1")
	assert.Equal(t, "1.0.0-alpha.1", ver.String())
}

func TestVersionsSort(t *testing.T) {
	versions := []*Version{
		NewVersion("3.0.0-alpha").Parse(),
		NewVersion("3.0.0").Parse(),
		NewVersion("2.0.0").Parse(),
		NewVersion("1.0.0").Parse(),
	}

	NewVersion("").Sort(versions)

	assert.Equal(t, "1.0.0", versions[0].String())
	assert.Equal(t, "2.0.0", versions[1].String())
	assert.Equal(t, "3.0.0-alpha", versions[2].String())
	assert.Equal(t, "3.0.0", versions[3].String())
}

func TestVersionsLatest(t *testing.T) {
	versions := []*Version{
		NewVersion("3.0.0-alpha").Parse(),
		NewVersion("3.0.0").Parse(),
		NewVersion("2.0.0").Parse(),
		NewVersion("1.0.0").Parse(),
	}
	actual := NewVersion("").Latest(versions)

	assert.NotNil(t, actual)
	assert.Equal(t, "3.0.0", actual.String())
}
