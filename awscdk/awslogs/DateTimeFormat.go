package awslogs


// Standard datetime formats for the DateTimeConverter processor.
//
// Provides common format patterns for date/time conversion.
type DateTimeFormat string

const (
	// ISO 8601 format (yyyy-MM-ddTHH:mm:ssZ).
	DateTimeFormat_ISO_8601 DateTimeFormat = "ISO_8601"
	// Unix timestamp (seconds since epoch).
	DateTimeFormat_UNIX_TIMESTAMP DateTimeFormat = "UNIX_TIMESTAMP"
	// Custom format specified by the targetFormat parameter.
	DateTimeFormat_CUSTOM DateTimeFormat = "CUSTOM"
)

