package awscdkgluealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkgluealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The job type.
//
// If you need to use a JobType that doesn't exist as a static member, you
// can instantiate a `JobType` object, e.g: `JobType.of('other name')`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import glue_alpha "github.com/aws/aws-cdk-go/awscdkgluealpha"
//
//   jobType := glue_alpha.JobType_ETL()
//
// Experimental.
type JobType interface {
	// The name of this JobType, as expected by Job resource.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for JobType
type jsiiProxy_JobType struct {
	_ byte // padding
}

func (j *jsiiProxy_JobType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Custom type name.
// Experimental.
func JobType_Of(name *string) JobType {
	_init_.Initialize()

	if err := validateJobType_OfParameters(name); err != nil {
		panic(err)
	}
	var returns JobType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-glue-alpha.JobType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func JobType_ETL() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"ETL",
		&returns,
	)
	return returns
}

func JobType_PYTHON_SHELL() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"PYTHON_SHELL",
		&returns,
	)
	return returns
}

func JobType_RAY() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"RAY",
		&returns,
	)
	return returns
}

func JobType_STREAMING() JobType {
	_init_.Initialize()
	var returns JobType
	_jsii_.StaticGet(
		"@aws-cdk/aws-glue-alpha.JobType",
		"STREAMING",
		&returns,
	)
	return returns
}

