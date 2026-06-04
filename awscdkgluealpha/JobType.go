package awscdkgluealpha


// The job type.
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
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	JobType_RAY JobType = "RAY"
)

