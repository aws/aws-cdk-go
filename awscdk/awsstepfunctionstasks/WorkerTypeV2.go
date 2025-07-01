package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerTypeV2 that doesn't exist as a static member, you
// can instantiate a `WorkerTypeV2` object, e.g: `WorkerTypeV2.of('other type')`.
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
type WorkerTypeV2 interface {
	// The name of this WorkerType, as expected by Job resource.
	Name() *string
}

// The jsii proxy struct for WorkerTypeV2
type jsiiProxy_WorkerTypeV2 struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkerTypeV2) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom worker type.
func WorkerTypeV2_Of(workerType *string) WorkerTypeV2 {
	_init_.Initialize()

	if err := validateWorkerTypeV2_OfParameters(workerType); err != nil {
		panic(err)
	}
	var returns WorkerTypeV2

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"of",
		[]interface{}{workerType},
		&returns,
	)

	return returns
}

func WorkerTypeV2_G_025X() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"G_025X",
		&returns,
	)
	return returns
}

func WorkerTypeV2_G_1X() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"G_1X",
		&returns,
	)
	return returns
}

func WorkerTypeV2_G_2X() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"G_2X",
		&returns,
	)
	return returns
}

func WorkerTypeV2_G_4X() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"G_4X",
		&returns,
	)
	return returns
}

func WorkerTypeV2_G_8X() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"G_8X",
		&returns,
	)
	return returns
}

func WorkerTypeV2_STANDARD() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"STANDARD",
		&returns,
	)
	return returns
}

func WorkerTypeV2_Z_2X() WorkerTypeV2 {
	_init_.Initialize()
	var returns WorkerTypeV2
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.WorkerTypeV2",
		"Z_2X",
		&returns,
	)
	return returns
}

