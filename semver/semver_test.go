package semver

import (
	"math/rand"
	"testing"
	"time"
)

type greaterThanTest struct {
	greaterVersion string
	lesserVersion  string
}

var greaterThanTests = []greaterThanTest{
	{"0.0.0", "0.0.0-foo"},
	{"0.0.1", "0.0.0"},
	{"1.0.0", "0.9.9"},
	{"0.10.0", "0.9.0"},
	{"0.99.0", "0.10.0"},
	{"2.0.0", "1.2.3"},
	{"0.0.0", "0.0.0-foo"},
	{"0.0.1", "0.0.0"},
	{"1.0.0", "0.9.9"},
	{"0.10.0", "0.9.0"},
	{"0.99.0", "0.10.0"},
	{"2.0.0", "1.2.3"},
	{"0.0.0", "0.0.0-foo"},
	{"0.0.1", "0.0.0"},
	{"1.0.0", "0.9.9"},
	{"0.10.0", "0.9.0"},
	{"0.99.0", "0.10.0"},
	{"2.0.0", "1.2.3"},
	{"1.2.3", "1.2.3-asdf"},
	{"1.2.3", "1.2.3-4"},
	{"1.2.3", "1.2.3-4-foo"},
	{"1.2.3-5-foo", "1.2.3-5"},
	{"1.2.3-5", "1.2.3-4"},
	{"1.2.3-5-foo", "1.2.3-5-Foo"},
	{"3.0.0", "2.7.2+asdf"},
	{"3.0.0+foobar", "2.7.2"},
	{"1.2.3-a.10", "1.2.3-a.5"},
	{"1.2.3-a.b", "1.2.3-a.5"},
	{"1.2.3-a.b", "1.2.3-a"},
	{"1.2.3-a.b.c.10.d.5", "1.2.3-a.b.c.5.d.100"},
}

type invalidTest string

var invalidTests = []invalidTest{
	"1",
	"1.0",
	"1.0.01",
	// TODO: 
	// "1.0.0-a.01", 
	// "1.0.0-1.",
}

func TestCompare(t *testing.T) {
	for _, v := range greaterThanTests {
		gt, err := NewVersion(v.greaterVersion)
		if err != nil {
			t.Error(err)
		}

		lt, err := NewVersion(v.lesserVersion)
		if err != nil {
			t.Error(err)
		}

		if gt.LessThan(*lt) == true {
			t.Errorf("%s should not be less than %s", gt, lt)
		}
	}
}

func testString(t *testing.T, orig string, version *Version) {
	if orig != version.String() {
		t.Errorf("%s != %s", orig, version)
	}
}

func TestString(t *testing.T) {
	for _, v := range greaterThanTests {
		gt, err := NewVersion(v.greaterVersion)
		if err != nil {
			t.Error(err)
		}
		testString(t, v.greaterVersion, gt)

		lt, err := NewVersion(v.lesserVersion)
		if err != nil {
			t.Error(err)
		}
		testString(t, v.lesserVersion, lt)
	}
}

func shuffleStringSlice(src []string) []string {
	dest := make([]string, len(src))
	rand.Seed(time.Now().Unix())
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func TestSort(t *testing.T) {
	sortedVersions := []string{"1.0.0", "1.0.2", "1.2.0", "3.1.1"}
	unsortedVersions := shuffleStringSlice(sortedVersions)

	semvers := []*Version{}
	for _, v := range unsortedVersions {
		sv, err := NewVersion(v)
		if err != nil {
			t.Fatal(err)
		}
		semvers = append(semvers, sv)
	}

	Sort(semvers)

	for idx, sv := range semvers {
		if sv.String() != sortedVersions[idx] {
			t.Fatalf("incorrect sort at index %v", idx)
		}
	}
}

func TestInvalid (t *testing.T) {
	var err error
	for _, versStr := range invalidTests {
		_, err = NewVersion(string(versStr))
		if err == nil {
			t.Errorf("Invalid version string %v did not return an error.", versStr)
		}
	}
}
