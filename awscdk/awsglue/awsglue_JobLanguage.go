package awsglue


// Runtime language of the Glue job.
// Experimental.
type JobLanguage string

const (
	// Scala.
	// Experimental.
	JobLanguage_SCALA JobLanguage = "SCALA"
	// Python.
	// Experimental.
	JobLanguage_PYTHON JobLanguage = "PYTHON"
)

