package escape

import "testing"

func TestHTML_Escape(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    string
		wantErr bool
	}{
		{
			"pass through",
			"potato",
			"potato",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HTML{}.Escape(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTML.Escape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTML.Escape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTML_UnEscape(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    string
		wantErr bool
	}{
		{
			"pass through",
			"potato",
			"potato",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HTML{}.UnEscape(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTML.UnEscape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTML.UnEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
