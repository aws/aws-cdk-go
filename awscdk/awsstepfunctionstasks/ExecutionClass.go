package awsstepfunctionstasks


// The excecution class of the job.
//
// Example:
//   tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
//   	GlueJobName: jsii.String("my-glue-job"),
//   	ExecutionClass: tasks.ExecutionClass_FLEX,
//   })
//
type ExecutionClass string

const (
	// The flexible execution class is appropriate for time-insensitive jobs whose start and completion times may vary.
	//
	// Only jobs with AWS Glue version 3.0 and above and command type `glueetl` will be allowed to set `ExecutionClass` to `FLEX`.
	// The flexible execution class is available for Spark jobs.
	ExecutionClass_FLEX ExecutionClass = "FLEX"
	// The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources.
	ExecutionClass_STANDARD ExecutionClass = "STANDARD"
)

