https://powcoder.com
代写代考加微信 powcoder
Assignment Project Exam Help
Add WeChat powcoder
package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected triangleType
	}{
		{1001, 5, 6, UnknownTriangle},
		// TODO add more tests for 100% test coverage
	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
