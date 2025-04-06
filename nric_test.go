package nric

import (
	"errors"
	"testing"
)

func TestNewNRIC(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		// Valid IDs
		{"Valid pre-2000 NRIC", "S6083480F", nil},
		{"Valid post-2000 NRIC", "T5717279C", nil},
		{"Valid pre-2000 FIN", "F6470401W", nil},
		{"Valid post-2000 FIN", "G8877699U", nil},
		{"Valid post-2022 FIN", "M5043078W", nil},
		{"Valid post-2022 FIN with J", "M2424771J", nil},

		// Invalid format IDs
		{"Too long", "G88776991Z", ErrInvalidFormat},
		{"Invalid characters", "GAAAAAAAZ", ErrInvalidFormat},
		{"Invalid prefix", "Z1111111A", ErrInvalidFormat},
		{"Invalid checksum char pre-2000 NRIC", "S6083480K", ErrInvalidFormat},
		{"Invalid checksum char pre-2000 FIN", "F6083480B", ErrInvalidFormat},

		// Invalid checksum IDs
		{"Invalid checksum pre-2000 NRIC", "S1234567E", ErrInvalidChecksum},
		{"Invalid checksum post-2000 NRIC", "T5717279A", ErrInvalidChecksum},
		{"Invalid checksum pre-2000 FIN", "F6470401K", ErrInvalidChecksum},
		{"Invalid checksum post-2000 FIN", "G8877699L", ErrInvalidChecksum},
		{"Invalid checksum post-2022 FIN", "M8877689K", ErrInvalidChecksum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nric, err := NewNRIC(tt.id)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("NewNRIC(%q) expected error %v, got nil", tt.id, tt.wantErr)
					return
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NewNRIC(%q) got error %v, want %v", tt.id, err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewNRIC(%q) unexpected error: %v", tt.id, err)
				return
			}
			if nric.String() != tt.id {
				t.Errorf("NewNRIC(%q) = %q, want %q", tt.id, nric.String(), tt.id)
			}
		})
	}
}

func TestIsForeigner(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		want      bool
		shouldErr bool
	}{
		{"Pre-2000 NRIC", "S6083480F", false, false},
		{"Post-2000 NRIC", "T5717279C", false, false},
		{"Pre-2000 FIN", "F6470401W", true, false},
		{"Post-2000 FIN", "G8877699U", true, false},
		{"Post-2022 FIN", "M5043078W", true, false},
		{"Invalid ID", "S1234567E", false, true}, // Won't reach this due to error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nric, err := NewNRIC(tt.id)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("NewNRIC(%q) expected error, got nil", tt.id)
				}
				return
			}
			if err != nil {
				t.Errorf("NewNRIC(%q) unexpected error: %v", tt.id, err)
				return
			}
			if got := nric.IsForeigner(); got != tt.want {
				t.Errorf("IsForeigner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIs2000(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		want      bool
		shouldErr bool
	}{
		{"Pre-2000 NRIC", "S6083480F", false, false},
		{"Post-2000 NRIC", "T5717279C", true, false},
		{"Pre-2000 FIN", "F6470401W", false, false},
		{"Post-2000 FIN", "G8877699U", true, false},
		{"Post-2022 FIN", "M5043078W", false, false},
		{"Invalid ID", "S1234567E", false, true}, // Won't reach this due to error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nric, err := NewNRIC(tt.id)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("NewNRIC(%q) expected error, got nil", tt.id)
				}
				return
			}
			if err != nil {
				t.Errorf("NewNRIC(%q) unexpected error: %v", tt.id, err)
				return
			}
			if got := nric.Is2000(); got != tt.want {
				t.Errorf("Is2000() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSeriesM(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		want      bool
		shouldErr bool
	}{
		{"Pre-2000 NRIC", "S6083480F", false, false},
		{"Post-2000 NRIC", "T5717279C", false, false},
		{"Pre-2000 FIN", "F6470401W", false, false},
		{"Post-2000 FIN", "G8877699U", false, false},
		{"Post-2022 FIN", "M5043078W", true, false},
		{"Invalid ID", "S1234567E", false, true}, // Won't reach this due to error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nric, err := NewNRIC(tt.id)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("NewNRIC(%q) expected error, got nil", tt.id)
				}
				return
			}
			if err != nil {
				t.Errorf("NewNRIC(%q) unexpected error: %v", tt.id, err)
				return
			}
			if got := nric.IsSeriesM(); got != tt.want {
				t.Errorf("IsSeriesM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		want      string
		shouldErr bool
	}{
		{"Pre-2000 NRIC", "S6083480F", "S6083480F", false},
		{"Post-2000 NRIC", "T5717279C", "T5717279C", false},
		{"Pre-2000 FIN", "F6470401W", "F6470401W", false},
		{"Post-2000 FIN", "G8877699U", "G8877699U", false},
		{"Post-2022 FIN", "M5043078W", "M5043078W", false},
		{"Invalid ID", "S1234567E", "", true}, // Won't reach this due to error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nric, err := NewNRIC(tt.id)
			if tt.shouldErr {
				if err == nil {
					t.Errorf("NewNRIC(%q) expected error, got nil", tt.id)
				}
				return
			}
			if err != nil {
				t.Errorf("NewNRIC(%q) unexpected error: %v", tt.id, err)
				return
			}
			if got := nric.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}
