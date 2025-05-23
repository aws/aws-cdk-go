package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Configuration for writing Distributed Map state results to S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   resultWriter := awscdk.Aws_stepfunctions.NewResultWriter(&ResultWriterProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   })
//
// Deprecated: use {@link ResultWriterV2 } instead.
type ResultWriter interface {
	// S3 Bucket in which to save Map Run results.
	// Deprecated: use {@link ResultWriterV2 } instead.
	Bucket() awss3.IBucket
	// S3 prefix in which to save Map Run results.
	// Default: - No prefix.
	//
	// Deprecated: use {@link ResultWriterV2 } instead.
	Prefix() *string
	// Compile policy statements to provide relevent permissions to the state machine.
	// Deprecated: use {@link ResultWriterV2 } instead.
	ProvidePolicyStatements() *[]awsiam.PolicyStatement
	// Render ResultWriter in ASL JSON format.
	// Deprecated: use {@link ResultWriterV2 } instead.
	Render(queryLanguage QueryLanguage) interface{}
}

// The jsii proxy struct for ResultWriter
type jsiiProxy_ResultWriter struct {
	_ byte // padding
}

func (j *jsiiProxy_ResultWriter) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResultWriter) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}


// Deprecated: use {@link ResultWriterV2 } instead.
func NewResultWriter(props *ResultWriterProps) ResultWriter {
	_init_.Initialize()

	if err := validateNewResultWriterParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResultWriter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ResultWriter",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: use {@link ResultWriterV2 } instead.
func NewResultWriter_Override(r ResultWriter, props *ResultWriterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.ResultWriter",
		[]interface{}{props},
		r,
	)
}

func (r *jsiiProxy_ResultWriter) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		r,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResultWriter) Render(queryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"render",
		[]interface{}{queryLanguage},
		&returns,
	)

	return returns
}

