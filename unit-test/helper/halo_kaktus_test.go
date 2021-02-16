package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Linux Manjaro")
	}

	result := HaloKaktus("Bay")
	require.Equal(t, "Halo Bay", result, "Result must be 'Halo bay'")

}

func TestSubTest(t *testing.T) {
	t.Run("Bay", func(t *testing.T) {
		result := HaloKaktus("Bay")
		require.Equal(t, "Halo Bay", result, "Result must be 'Halo bay'")
	})
	t.Run("kaks", func(t *testing.T) {
		result := HaloKaktus("kaks")
		require.Equal(t, "Halo kaks", result, "Result must be 'Halo kaks'")
	})
}

func TestHaloRequire(t *testing.T) {
	result := HaloKaktus("Bay")
	require.Equal(t, "Halo Bay", result, "Result must be 'Halo bay'")
	fmt.Println("TestHalo with Require Done")

}
func TestHaloAssert(t *testing.T) {
	result := HaloKaktus("Bay")
	assert.Equal(t, "Halo Bay", result, "Result must be 'Halo bay'")
	fmt.Println("TestHalo with Assert Done")

}

func TestHaloKaktusBy(t *testing.T) {
	result := HaloKaktus("Bay")

	if result != "Halo Bay" {
		//error
		t.Error("Result must be 'Halo Bay'")
		// t.Fail()
		// panic("Result is not 'Halo Bay'")
	}
	fmt.Println("TestHaloKaktusBy Done")
}

func TestHaloJon(t *testing.T) {
	result := HaloKaktus("Kaktus")

	if result != "Halo Kaktus" {
		//error
		t.Fatal("Result must be 'Helo Kaktus'")
		// t.FailNow()
		// panic("Result is not 'Halo Kaktus'")
	}
	fmt.Println("TestHaloKaktusJon Done")

}

func TestTableHai(t *testing.T) {
	tests := []struct {
		name     string
		require  string
		expected string
	}{
		{
			name:     "Bay",
			require:  "Bay",
			expected: "Halo Bay",
		},
		{
			name:     "Kaks",
			require:  "Kaks",
			expected: "Halo Kaks",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HaloKaktus(test.require)
			require.Equal(t, test.expected, result)
		})
	}
}
