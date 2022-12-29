package gotenv

import (
	"os"
	"regexp"
	"strings"

	"github.com/evanoc3/gotenv/internal/utils/slices"
)


type ConfigOptions struct {
	Path string
	AllowComments bool
}


func Config(options ConfigOptions) error {
	options = normalizeOptions(options)

	fileContents, err := os.ReadFile(options.Path)
	if err != nil {
		return err
	}

	envEntries := parseDotenv(options, string(fileContents))

	for _, envEntry := range envEntries {
		if err := os.Setenv(envEntry.Key, envEntry.Val); err != nil {
			return err
		}
	}

	return nil
}


type envEntry struct {
	Key string
	Val string
}


func parseDotenv(options ConfigOptions, dotenvContent string) []envEntry {
	dotenvLines := strings.Split(dotenvContent, "\n")
	dotenvLines = slices.Transform(dotenvLines, func (line string) string { return strings.TrimSpace(line) })
	dotenvLines = slices.Filter( dotenvLines, func(line string) bool { return !strings.HasPrefix(line, "#"); } )

	envEntries := make([]envEntry, 0, len(dotenvLines))

	for _, line := range dotenvLines {
		var re *regexp.Regexp
		if options.AllowComments {
			re = regexp.MustCompile(`(\S+)\s*=\s*(.*)\s?#.*?$`)
		} else {
			re = regexp.MustCompile(`(\S+)\s*=\s*(.+)`)
		}

		submatches := re.FindStringSubmatch(line)

		if len(submatches) != 3 {
			continue
		}

		envEntries = append(envEntries, envEntry{
			Key: submatches[1],
			Val: submatches[2],
		})
	}

	return envEntries
}


func normalizeOptions(options ConfigOptions) ConfigOptions {
	if options.Path == "" {
		options.Path = ".env"
	}

	return options
}
