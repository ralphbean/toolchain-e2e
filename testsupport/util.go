package testsupport

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
)

// GenerateName appends generated UUID to the given string
func GenerateName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, uuid.Must(uuid.NewV4()).String())
}

var notAllowedChars = regexp.MustCompile("[^-a-z0-9]")

// NewObjectNamePrefix creates a namePrefix to be used as .ObjectMeta.GenerateName field.
// The name prefix is based on the name of the test using this function.
func NewObjectNamePrefix(t *testing.T) string {
	namePrefix := strings.ToLower(t.Name())
	// Remove all invalid characters
	namePrefix = notAllowedChars.ReplaceAllString(namePrefix, "")

	// Trim if the length exceeds 60 chars (63 is the max)
	if len(namePrefix) > 40 {
		namePrefix = namePrefix[0:40]
	}
	return namePrefix
}
