// The CDK Construct Library for AWS::Glue
package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of predefined worker that is allocated when a job runs.
//
// If you need to use a WorkerType that doesn't exist as a static member, you
// can instantiate a `WorkerType` object, e.g: `WorkerType.of('other type')`.
//
// Example:
//   glue.NewJob(this, jsii.String("RayJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonRay(&PythonRayExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE_NINE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	WorkerType: glue.WorkerType_Z_2X(),
//   	WorkerCount: jsii.Number(2),
//   	Description: jsii.String("an example Ray job"),
//   })
//
// Experimental.
type WorkerType interface {
	// The name of this WorkerType, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for WorkerType
type jsiiProxy_WorkerType struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkerType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom worker type.
// Experimental.
func WorkerType_Of(workerType *string) WorkerType {
	_init_.Initialize()

	if err := validateWorkerType_OfParameters(workerType); err != nil {
		panic(err)
	}
	var returns WorkerType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"of",
		[]interface{}{workerType},
		&returns,
	)

	return returns
}

func WorkerType_G_025X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"G_025X",
		&returns,
	)
	return returns
}

func WorkerType_G_1X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"G_1X",
		&returns,
	)
	return returns
}

func WorkerType_G_2X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"G_2X",
		&returns,
	)
	return returns
}

func WorkerType_STANDARD() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"STANDARD",
		&returns,
	)
	return returns
}

func WorkerType_Z_2X() WorkerType {
	_init_.Initialize()
	var returns WorkerType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.WorkerType",
		"Z_2X",
		&returns,
	)
	return returns
}

