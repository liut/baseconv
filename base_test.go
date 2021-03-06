package baseconv

import (
	"testing"
)

type testpair struct {
	n, s string
	f, t int
}

var pairs = []testpair{
	{"", "", 1, 2},
	{"12", "12", 10, 10},
	{"297eb4fe7f4993", "37051e7sxw3", 16, 36},
	{"49262201ced289", "5mqek115cax", 16, 36},
	{"12cfe6fab9b713e", "n67j3ox4qpq", 16, 36},
	{"dd0274386e26030c", "3czs1w1ul6cr0", 16, 36},
	{"dd0274386e26030c0301", "4pyem417r86neold", 16, 36},
	{"5cc163b92ab9b482b4486999d354f91e", "5hos8aw6atq7kcpvn1gweaf4u", 16, 36},
	{"5cc163b92ab9b482b4486999d354f91e", "2P1FKmvE5PjCN4PhpocjBs", 16, 62},
	{"d1fb1bc11d2e992b4be5f770f35e345aa75a1d11", "oj0fiwpthr2v1ecf3p27ktokmh1a51t", 16, 36},
	{"d1fb1bc11d2e992b4be5f770f35e345aa75a1d11", "tXzMTsS55hMM9DEUf4BXSlq509b", 16, 62},
	{"d1fb1bc11d2e992b4be5f770f35e345aa75a1d1104", "1ZHZLT1diZProPNKUef2LeorGEBTu", 16, 62},
}

func testEqual(t *testing.T, msg string, args ...interface{}) bool {
	if args[len(args)-2] != args[len(args)-1] {
		t.Errorf(msg, args...)
		return false
	}
	return true
}

func TestEncode(t *testing.T) {
	for _, p := range pairs {
		got, err := Convert(p.n, p.f, p.t)
		if err != nil {
			t.Error(err)
		} else {
			testEqual(t, "convert(%q, %d, %d) = %q, want %q", p.n, p.f, p.t, got, p.s)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, p := range pairs {
		got, err := Convert(p.s, p.t, p.f)
		if err != nil {
			t.Error(err)
		} else {
			testEqual(t, "baseconvert(%q, %d, %d) = %q, want %q", p.s, p.t, p.f, got, p.n)
		}
	}
}
