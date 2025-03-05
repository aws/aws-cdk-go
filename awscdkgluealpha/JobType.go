package awscdkgluealpha


// The job type.
//
// If you need to use a JobType that doesn't exist as a static member, you
// can instantiate a `JobType` object, e.g: `JobType.of('other name')`.
// Experimental.
type JobType string

const (
	// Command for running a Glue Spark job.
	// Experimental.
	JobType_ETL JobType = "ETL"
	// Command for running a Glue Spark streaming job.
	// Experimental.
	JobType_STREAMING JobType = "STREAMING"
	// Command for running a Glue python shell job.
	// Experimental.
	JobType_PYTHON_SHELL JobType = "PYTHON_SHELL"
	// Command for running a Glue Ray job.
	// Experimental.
	JobType_RAY JobType = "RAY"
)

