package test

import "testing"

func Test_Email(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		want    bool
		wantErr string
	}{
		{"Valid Email", "krishna@info.com", true, ""}, {"valid email 2", "vamsik@info.com", true, ""},
		{"missed @", "krishninfo.com", false, ""},
		{"missed domainname", "krishna@.com", false, ""},
		{"missed domain", "krishna@", false, ""},
		{"missed extention", "krishna@info", false, ""},
		{"missing username", "@info.com", false, ""},
		{"Invalid email", "krishna@info..com", false, ""},

		{"empty", "", false, "cannot be empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateEmail(tt.email)

			if got != tt.want {
				t.Errorf("ValidateEmail(%q)=%v, want=%v", tt.email, got, tt.want)
			}

			if err != nil && err.Error() != tt.wantErr {
				t.Errorf("ValidateEmail(%v)=error %v, wantErr=%v", tt.email, err, tt.wantErr)
			} else if err == nil && tt.wantErr != "" {
				t.Errorf("ValidateEmail(%q)=error %v, wantErr=%v", tt.email, err, tt.wantErr)
			}
		})
	}
}
