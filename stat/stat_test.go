package stat

import (
	"testing"

	"github.com/nathanfrench/procfs/util"
)

func TestParsingStat(t *testing.T) {
	// set the GLOBAL_SYSTEM_START
	util.GLOBAL_SYSTEM_START = 1388417200
	s, err := New("./testfiles/stat")

	if err != nil {
		t.Fatal("Got error", err)
	}

	if s == nil {
		t.Fatal("stat is missing")
	}

	// if s.Starttime.seconds() != 1388604586 {
	// t.Fatal("Start time is wrong, expected: 1388604586", s.Starttime.EpochSeconds)
	// }

}
