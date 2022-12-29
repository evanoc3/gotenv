package gotenv

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestParseDotenv(t *testing.T) {
	require := require.New(t)

	input := "TEST1=TEST2\n" +
					 "TEST3 =TEST4\n" +
					 "TEST5 = TEST6\n" +
					 "TEST7  =  TEST8\n"
	
	expected := []envEntry{
		{ Key: "TEST1", Val: "TEST2" },
		{ Key: "TEST3", Val: "TEST4" },
		{ Key: "TEST5", Val: "TEST6" },
		{ Key: "TEST7", Val: "TEST8" },
	}
	actual := parseDotenv(ConfigOptions{}, input)
	require.Equal(expected, actual)
}
