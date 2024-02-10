package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AWS Glue runtime determines the runtime engine of the job.
//
// Example:
//   glue.NewJob(this, jsii.String("RayJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonRay(&PythonRayExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE_NINE,
//   		Runtime: glue.Runtime_RAY_TWO_FOUR(),
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script"), jsii.String("hello_world.py"))),
//   	}),
//   	WorkerType: glue.WorkerType_Z_2X(),
//   	WorkerCount: jsii.Number(2),
//   	Description: jsii.String("an example Ray job"),
//   })
//
// Experimental.
type Runtime interface {
	// The name of this Runtime.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom runtime.
// Experimental.
func Runtime_Of(runtime *string) Runtime {
	_init_.Initialize()

	if err := validateRuntime_OfParameters(runtime); err != nil {
		panic(err)
	}
	var returns Runtime

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.Runtime",
		"of",
		[]interface{}{runtime},
		&returns,
	)

	return returns
}

func Runtime_RAY_TWO_FOUR() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.Runtime",
		"RAY_TWO_FOUR",
		&returns,
	)
	return returns
}

