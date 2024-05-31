package awsstepfunctionstasks


// Properties for the worker configuration.
//
// Example:
//   tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
//   	GlueJobName: jsii.String("my-glue-job"),
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		WorkerType: tasks.WorkerType_G_1X,
//   		 // Worker type
//   		NumberOfWorkers: jsii.Number(2),
//   	},
//   })
//
type WorkerConfigurationProperty struct {
	// The number of workers of a defined `WorkerType` that are allocated when a job runs.
	NumberOfWorkers *float64 `field:"required" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The type of predefined worker that is allocated when a job runs.
	WorkerType WorkerType `field:"required" json:"workerType" yaml:"workerType"`
}

