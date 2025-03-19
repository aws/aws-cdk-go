package awscodepipeline


// Execution mode.
//
// Example:
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   	ExecutionMode: codepipeline.ExecutionMode_PARALLEL,
//   })
//
type ExecutionMode string

const (
	// QUEUED mode.
	//
	// Executions are processed one by one in the order that they are queued.
	//
	// This requires pipeline type V2.
	ExecutionMode_QUEUED ExecutionMode = "QUEUED"
	// SUPERSEDED mode.
	//
	// A more recent execution can overtake an older one.
	//
	// This is the default.
	ExecutionMode_SUPERSEDED ExecutionMode = "SUPERSEDED"
	// PARALLEL mode.
	//
	// In PARALLEL mode, executions run simultaneously and independently of one
	// another. Executions don't wait for other runs to complete before starting
	// or finishing.
	//
	// This requires pipeline type V2.
	ExecutionMode_PARALLEL ExecutionMode = "PARALLEL"
)

