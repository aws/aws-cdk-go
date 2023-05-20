package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AWS Glue version determines the versions of Apache Spark and Python that are available to the job.
//
// Example:
//   glue.NewJob(this, jsii.String("PythonSparkStreamingJob"), &JobProps{
//   	Executable: glue.JobExecutable_PythonStreaming(&PythonSparkJobExecutableProps{
//   		GlueVersion: glue.GlueVersion_V4_0(),
//   		PythonVersion: glue.PythonVersion_THREE,
//   		Script: glue.Code_FromAsset(path.join(__dirname, jsii.String("job-script/hello_world.py"))),
//   	}),
//   	Description: jsii.String("an example Python Streaming job"),
//   })
//
// See: https://docs.aws.amazon.com/glue/latest/dg/add-job.html.
//
// If you need to use a GlueVersion that doesn't exist as a static member, you
// can instantiate a `GlueVersion` object, e.g: `GlueVersion.of('1.5')`.
//
// Experimental.
type GlueVersion interface {
	// The name of this GlueVersion, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for GlueVersion
type jsiiProxy_GlueVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_GlueVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom Glue version.
// Experimental.
func GlueVersion_Of(version *string) GlueVersion {
	_init_.Initialize()

	if err := validateGlueVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns GlueVersion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func GlueVersion_V0_9() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V0_9",
		&returns,
	)
	return returns
}

func GlueVersion_V1_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V1_0",
		&returns,
	)
	return returns
}

func GlueVersion_V2_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V2_0",
		&returns,
	)
	return returns
}

func GlueVersion_V3_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V3_0",
		&returns,
	)
	return returns
}

func GlueVersion_V4_0() GlueVersion {
	_init_.Initialize()
	var returns GlueVersion
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.GlueVersion",
		"V4_0",
		&returns,
	)
	return returns
}

