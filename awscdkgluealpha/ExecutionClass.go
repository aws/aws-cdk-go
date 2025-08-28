package awscdkgluealpha


// The ExecutionClass whether the job is run with a standard or flexible execution class.
// See: https://docs.aws.amazon.com/glue/latest/dg/add-job.html
//
// Experimental.
type ExecutionClass string

const (
	// The flexible execution class is appropriate for time-insensitive jobs whose start and completion times may vary.
	// Experimental.
	ExecutionClass_FLEX ExecutionClass = "FLEX"
	// The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources.
	// Experimental.
	ExecutionClass_STANDARD ExecutionClass = "STANDARD"
)

