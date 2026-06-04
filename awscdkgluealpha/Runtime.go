package awscdkgluealpha


// AWS Glue runtime determines the runtime engine of the job.
// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
// Migrate to Amazon EKS with KubeRay Operator. See
// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
type Runtime string

const (
	// Runtime for a Glue for Ray 2.4.
	// Deprecated: AWS Glue for Ray is closed to new customers as of April 30, 2026.
	// Migrate to Amazon EKS with KubeRay Operator. See
	// https://docs.aws.amazon.com/glue/latest/dg/awsglue-ray-jobs-availability-change.html
	Runtime_RAY_TWO_FOUR Runtime = "RAY_TWO_FOUR"
)

