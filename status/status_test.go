package status

import (
	"testing"

	"github.com/nathanfrench/procfs/util"
)

func TestParsingStatus(t *testing.T) {
	// set the GLOBAL_SYSTEM_START
	util.GLOBAL_SYSTEM_START = 1388417200
	s, err := New("./testfiles/status")

	if err != nil {
		t.Fatal("Got error", err)
	}

	if s == nil {
		t.Fatal("stat is missing")
	}

	if s.Uid != 0 {
		t.Fatal("Uid should be 0, instead:", s.Uid)
	}

	if s.Gid != 1001 {
		t.Fatal("Gid should be 1001, instead:", s.Gid)
	}
	if s.Vcswitch != 72058 {
		t.Fatal("Vcswitch should be 75958, instead,", s.Vcswitch)
	}
	if s.NVcswitch != 3109 {
		t.Fatal("NVcswitch should be 75958, instead,", s.NVcswitch)
	}
}
