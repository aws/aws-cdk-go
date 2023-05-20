package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The executable properties related to the Glue job's GlueVersion, JobType and code.
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
// Experimental.
type JobExecutable interface {
	// Called during Job initialization to get JobExecutableConfig.
	// Experimental.
	Bind() *JobExecutableConfig
}

// The jsii proxy struct for JobExecutable
type jsiiProxy_JobExecutable struct {
	_ byte // padding
}

// Create a custom JobExecutable.
// Experimental.
func JobExecutable_Of(config *JobExecutableConfig) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_OfParameters(config); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"of",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Create Python executable props for Apache Spark ETL job.
// Experimental.
func JobExecutable_PythonEtl(props *PythonSparkJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonEtlParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonEtl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for Ray jobs.
// Experimental.
func JobExecutable_PythonRay(props *PythonRayExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonRayParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonRay",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for python shell jobs.
// Experimental.
func JobExecutable_PythonShell(props *PythonShellExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonShellParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonShell",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Python executable props for Apache Spark Streaming job.
// Experimental.
func JobExecutable_PythonStreaming(props *PythonSparkJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_PythonStreamingParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"pythonStreaming",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Scala executable props for Apache Spark ETL job.
// Experimental.
func JobExecutable_ScalaEtl(props *ScalaJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_ScalaEtlParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"scalaEtl",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create Scala executable props for Apache Spark Streaming job.
// Experimental.
func JobExecutable_ScalaStreaming(props *ScalaJobExecutableProps) JobExecutable {
	_init_.Initialize()

	if err := validateJobExecutable_ScalaStreamingParameters(props); err != nil {
		panic(err)
	}
	var returns JobExecutable

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobExecutable",
		"scalaStreaming",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobExecutable) Bind() *JobExecutableConfig {
	var returns *JobExecutableConfig

	_jsii_.Invoke(
		j,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

