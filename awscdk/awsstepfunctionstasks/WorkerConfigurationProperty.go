package awsstepfunctionstasks


// Properties for the worker configuration.
//
// Example:
//   tasks.NewGlueStartJobRun(this, jsii.String("Task"), &GlueStartJobRunProps{
//   	GlueJobName: jsii.String("my-glue-job"),
//   	WorkerConfiguration: &WorkerConfigurationProperty{
//   		WorkerTypeV2: tasks.WorkerTypeV2_G_1X(),
//   		 // Worker type
//   		NumberOfWorkers: jsii.Number(2),
//   	},
//   })
//
type WorkerConfigurationProperty struct {
	// The number of workers of a defined `WorkerType` that are allocated when a job runs.
	NumberOfWorkers *float64 `field:"required" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The type of predefined worker that is allocated when a job runs.
	// Default: - must choose one of `workerType` or `workerTypeV2`.
	//
	// Deprecated: Use `workerTypeV2` for more flexibility in defining worker types.
	WorkerType WorkerType `field:"optional" json:"workerType" yaml:"workerType"`
	// The type of predefined worker that is allocated when a job runs.
	//
	// Can be one of the
	// predefined values or dynamic values using `WorkerTypeV2.of(...)`.
	// Default: - must choose one of `workerType` or `workerTypeV2`.
	//
	WorkerTypeV2 WorkerTypeV2 `field:"optional" json:"workerTypeV2" yaml:"workerTypeV2"`
}

