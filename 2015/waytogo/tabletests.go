var flagtests = []struct {
	in  string
	out string
}{
	{"%-1.2a", "[%-1.2a]"},
	{"%+1.2a", "[%+1.2a]"},
}

func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		s := Sprintf(tt.in, &flagprinter)
		if s != tt.out {
			t.Errorf("Sprintf(%q, &flagprinter) => %q, want %q", tt.in, s, tt.out)
		}
	}
}