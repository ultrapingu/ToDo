package mymath

import (
	"fmt"
	"math"
	"testing"
)

const epsilon = 0.000000001

func approxEqualVec4(a, b Vec4) bool {
	return math.Abs(a.X-b.X) < epsilon &&
		math.Abs(a.Y-b.Y) < epsilon &&
		math.Abs(a.Z-b.Z) < epsilon &&
		math.Abs(a.W-b.W) < epsilon
}

func approxEqualFloat(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestAdd(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		arg2     Vec4
		expected Vec4
	}{
		{NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, 0.0)},
		{Up(), NewPoint(0.0, 0.0, 0.0), Up()},
		{Right(), NewPoint(0.0, 0.0, 0.0), Right()},
		{Front(), NewPoint(0.0, 0.0, 0.0), Front()},
		{NewPoint(0.0, 0.0, 0.0), Up(), NewPoint(0.0, 1.0, 0.0)},
		{NewPoint(0.0, 0.0, 0.0), Right(), NewPoint(1.0, 0.0, 0.0)},
		{NewPoint(0.0, 0.0, 0.0), Front(), NewPoint(0.0, 0.0, 1.0)},
		{NewPoint(1.3, 3.4, 12.001), NewPoint(0.1, 0.555, 0.001), NewPoint(1.4, 3.955, 12.002)},
		{NewVector(1.3, 3.4, 12.001), NewVector(0.1, 0.555, 0.001), NewVector(1.4, 3.955, 12.002)},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test Add(%s, %s) == %s", data.arg1.String(), data.arg2.String(), data.expected.String())
		t.Run(testName, func(t *testing.T) {
			if output := Add(data.arg1, data.arg2); !approxEqualVec4(output, data.expected) {
				t.Errorf("Output %s not equal to expected %s", output.String(), data.expected.String())
			}
		})
	}
}

func TestSub(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		arg2     Vec4
		expected Vec4
	}{
		{NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, 0.0)},
		{Up(), NewPoint(0.0, 0.0, 0.0), Up()},
		{Right(), NewPoint(0.0, 0.0, 0.0), Right()},
		{Front(), NewPoint(0.0, 0.0, 0.0), Front()},
		{NewPoint(0.0, 0.0, 0.0), Up(), NewPoint(0.0, -1.0, 0.0)},
		{NewPoint(0.0, 0.0, 0.0), Right(), NewPoint(-1.0, 0.0, 0.0)},
		{NewPoint(0.0, 0.0, 0.0), Front(), NewPoint(0.0, 0.0, -1.0)},
		{NewPoint(1.3, 3.4, 12.001), NewPoint(0.1, 0.555, 0.001), NewPoint(1.2, 2.845, 12.0)},
		{NewVector(1.3, 3.4, 12.001), NewVector(0.1, 0.555, 0.001), NewVector(1.2, 2.845, 12.0)},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test Sub(%s, %s) == %s", data.arg1.String(), data.arg2.String(), data.expected.String())
		t.Run(testName, func(t *testing.T) {
			if output := Sub(data.arg1, data.arg2); !approxEqualVec4(output, data.expected) {
				t.Errorf("Output %s not equal to expected %s", output.String(), data.expected.String())
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		arg2     float64
		expected Vec4
	}{
		{NewVector(0.0, 0.0, 0.0), 0.0, NewVector(0.0, 0.0, 0.0)},
		{NewVector(0.0, 0.0, 0.0), 1.0, NewVector(0.0, 0.0, 0.0)},
		{Up(), 0.0, NewVector(0.0, 0.0, 0.0)},
		{Up(), 1.0, NewVector(0.0, 1.0, 0.0)},
		{Up(), 2.0, NewVector(0.0, 2.0, 0.0)},
		{Up(), -1.0, NewVector(0.0, -1.0, 0.0)},
		{Right(), 0.0, NewVector(0.0, 0.0, 0.0)},
		{Right(), 1.0, NewVector(1.0, 0.0, 0.0)},
		{Right(), 2.0, NewVector(2.0, 0.0, 0.0)},
		{Right(), -1.0, NewVector(-1.0, 0.0, 0.0)},
		{Front(), 0.0, NewVector(0.0, 0.0, 0.0)},
		{Front(), 1.0, NewVector(0.0, 0.0, 1.0)},
		{Front(), 2.0, NewVector(0.0, 0.0, 2.0)},
		{Front(), -1.0, NewVector(0.0, 0.0, -1.0)},
		{NewVector(-1.3, -3.4, 12.001), -5.0, NewVector(6.5, 17.0, -60.005)},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test Multiply(%s, %f) == %s", data.arg1.String(), data.arg2, data.expected.String())
		t.Run(testName, func(t *testing.T) {
			if output := Multiply(data.arg1, data.arg2); !approxEqualVec4(output, data.expected) {
				t.Errorf("Output %s not equal to expected %s", output.String(), data.expected.String())
			}
		})
	}
}

func TestDot(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		arg2     Vec4
		expected float64
	}{
		{NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, 0.0), 0.0},
		{Up(), NewPoint(0.0, 0.0, 0.0), 0.0},
		{Up(), Up(), 1.0},
		{Up(), NewPoint(0.0, -1.0, 0.0), -1.0},
		{Up(), Right(), 0.0},
		{Up(), Front(), 0.0},
		{Right(), NewPoint(0.0, 0.0, 0.0), 0.0},
		{Right(), Up(), 0.0},
		{Right(), NewPoint(-1.0, 0.0, 0.0), -1.0},
		{Right(), Right(), 1.0},
		{Right(), Front(), 0.0},
		{Front(), NewPoint(0.0, 0.0, 0.0), 0.0},
		{Front(), Up(), 0.0},
		{Front(), NewPoint(0.0, 0.0, -1.0), -1.0},
		{Front(), Right(), 0.0},
		{Front(), Front(), 1.0},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test Dot(%s, %s) == %f", data.arg1.String(), data.arg2.String(), data.expected)
		t.Run(testName, func(t *testing.T) {
			if output := Dot(data.arg1, data.arg2); !approxEqualFloat(output, data.expected) {
				t.Errorf("Output %f not equal to expected %f", output, data.expected)
			}
		})
	}
}

func TestCross(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		arg2     Vec4
		expected Vec4
	}{
		{NewVector(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0)},
		{Up(), NewVector(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0)},
		{NewVector(0.0, 0.0, 0.0), Up(), NewVector(0.0, 0.0, 0.0)},
		{Right(), Up(), Front()},
		{Up(), Right(), NewVector(0.0, 0.0, -1.0)},
		{Front(), Right(), Up()},
		{Right(), Front(), NewVector(0.0, -1.0, 0.0)},
		{Front(), NewVector(0.0, -1.0, 0.0), Right()},
		{NewVector(0.0, -1.0, 0.0), Front(), NewVector(-1.0, 0.0, 0.0)},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test Cross(%s, %s) == %s", data.arg1.String(), data.arg2.String(), data.expected.String())
		t.Run(testName, func(t *testing.T) {
			if output := Cross(data.arg1, data.arg2); !approxEqualVec4(output, data.expected) {
				t.Errorf("Output %s not equal to expected %s", output.String(), data.expected.String())
				t.Failed()
			}
		})
	}
}

func TestMagnitudeSq(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		expected float64
	}{
		{NewVector(0.0, 0.0, 0.0), 0.0},
		{NewVector(1.0, 0.0, 0.0), 1.0},
		{NewVector(0.0, 1.0, 0.0), 1.0},
		{NewVector(0.0, 0.0, 1.0), 1.0},
		{NewVector(-1.0, 0.0, 0.0), 1.0},
		{NewVector(0.0, -1.0, 0.0), 1.0},
		{NewVector(0.0, 0.0, -1.0), 1.0},
		{NewVector(1.0, 1.0, 1.0), 3.0},
		{NewVector(-1.0, 1.0, 1.0), 3.0},
		{NewVector(1.0, -1.0, 1.0), 3.0},
		{NewVector(1.0, 1.0, -1.0), 3.0},
		{NewVector(-1.0, -1.0, -1.0), 3.0},
		{NewVector(15.3, -6.45, 56.3244), 3448.13053536},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test %s.MagnitudeSq() == %f", data.arg1.String(), data.expected)
		t.Run(testName, func(t *testing.T) {
			if output := data.arg1.MagnitudeSq(); !approxEqualFloat(output, data.expected) {
				t.Errorf("Output %f not equal to expected %f", output, data.expected)
				t.Failed()
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		expected float64
	}{
		{NewVector(0.0, 0.0, 0.0), 0.0},
		{NewVector(1.0, 0.0, 0.0), 1.0},
		{NewVector(0.0, 1.0, 0.0), 1.0},
		{NewVector(0.0, 0.0, 1.0), 1.0},
		{NewVector(-1.0, 0.0, 0.0), 1.0},
		{NewVector(0.0, -1.0, 0.0), 1.0},
		{NewVector(0.0, 0.0, -1.0), 1.0},
		{NewVector(1.0, 1.0, 1.0), 1.7320508076},
		{NewVector(-1.0, 1.0, 1.0), 1.7320508076},
		{NewVector(1.0, -1.0, 1.0), 1.7320508076},
		{NewVector(1.0, 1.0, -1.0), 1.7320508076},
		{NewVector(-1.0, -1.0, -1.0), 1.7320508076},
		{NewVector(15.3, -6.45, 56.3244), 58.720784526},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test %s.Magnitude() == %f", data.arg1.String(), data.expected)
		t.Run(testName, func(t *testing.T) {
			if output := data.arg1.Magnitude(); !approxEqualFloat(output, data.expected) {
				t.Errorf("Output %f not equal to expected %f", output, data.expected)
				t.Failed()
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	var testData = []struct {
		arg1     Vec4
		expected Vec4
	}{
		{NewVector(1.0, 0.0, 0.0), NewVector(1.0, 0.0, 0.0)},
		{NewVector(0.0, 1.0, 0.0), NewVector(0.0, 1.0, 0.0)},
		{NewVector(0.0, 0.0, 1.0), NewVector(0.0, 0.0, 1.0)},
		{NewVector(-1.0, 0.0, 0.0), NewVector(-1.0, 0.0, 0.0)},
		{NewVector(0.0, -1.0, 0.0), NewVector(0.0, -1.0, 0.0)},
		{NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0)},
		{NewVector(1.0, 1.0, 1.0), NewVector(0.577350269, 0.577350269, 0.577350269)},
		{NewVector(-1.0, 1.0, 1.0), NewVector(-0.577350269, 0.577350269, 0.577350269)},
		{NewVector(1.0, -1.0, 1.0), NewVector(0.577350269, -0.577350269, 0.577350269)},
		{NewVector(1.0, 1.0, -1.0), NewVector(0.577350269, 0.577350269, -0.577350269)},
		{NewVector(-1.0, -1.0, -1.0), NewVector(-0.577350269, -0.577350269, -0.577350269)},
		{NewVector(15.3, -6.45, 56.3244), NewVector(0.260555102, -0.109841857, 0.959190182)},
	}

	for _, data := range testData {
		testName := fmt.Sprintf("Test %s.Normalize() == %s", data.arg1.String(), data.expected.String())
		t.Run(testName, func(t *testing.T) {
			data.arg1.Normalize()
			if !approxEqualVec4(data.arg1, data.expected) {
				t.Errorf("Output %s not equal to expected %s", data.arg1.String(), data.expected.String())
				t.Failed()
			}
		})
	}
}
