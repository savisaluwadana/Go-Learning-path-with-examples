package converter

import (
	"math"
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		name    string
		kind    string
		from    string
		to      string
		value   float64
		want    float64
		wantErr bool
	}{
		{name: "c to f", kind: "temperature", from: "c", to: "f", value: 100, want: 212},
		{name: "f to c", kind: "temperature", from: "f", to: "c", value: 32, want: 0},
		{name: "km to mi", kind: "distance", from: "km", to: "mi", value: 5, want: 3.1068559611866697},
		{name: "mi to m", kind: "distance", from: "mi", to: "m", value: 1, want: 1609.344},
		{name: "kg to lb", kind: "weight", from: "kg", to: "lb", value: 2, want: 4.409245243697551},
		{name: "invalid kind", kind: "speed", from: "km", to: "mi", value: 10, wantErr: true},
		{name: "invalid unit", kind: "temperature", from: "c", to: "m", value: 10, wantErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Convert(tc.kind, tc.from, tc.to, tc.value)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tc.want) > 1e-9 {
				t.Fatalf("got %f, want %f", got, tc.want)
			}
		})
	}
}
