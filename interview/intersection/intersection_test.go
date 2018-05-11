package intersection

import "testing"

func TestIntersect(t *testing.T) {
	type args struct {
		a, b []Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"should return the intersection of the point slices",
			args{
				[]Point{Point{0.0, 0.0}, Point{1.0, 1.0}, Point{2.0, 2.0}},
				[]Point{Point{0.0, 0.0}, Point{3.0, 3.0}, Point{4.0, 4.0}},
			},
			1,
		},
		{
			"should return the intersection of the point slices",
			args{
				[]Point{Point{0.0, 0.0}, Point{1.0, 1.0}, Point{2.0, 2.0}},
				[]Point{Point{0.0, 0.0}, Point{3.0, 3.0}, Point{1.0, 1.0}},
			},
			2,
		},
		{
			"should return the intersection of the point slices",
			args{
				[]Point{Point{5.0, 3.0}, Point{1.0, 1.0}, Point{2.0, 2.0}},
				[]Point{Point{0.0, 0.0}, Point{3.0, 3.0}, Point{4.0, 4.0}},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Intersect(%T, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
