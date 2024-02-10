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
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a bucket
//   bucket := s3.NewBucket(this, jsii.String("Bucket"))
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	ItemReader: sfn.NewS3JsonItemReader(&S3FileItemReaderProps{
//   		Bucket: bucket,
//   		Key: jsii.String("my-key.json"),
//   	}),
//   	ResultWriter: sfn.NewResultWriter(&ResultWriterProps{
//   		Bucket: bucket,
//   		Prefix: jsii.String("my-prefix"),
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
type ResultWriter interface {
	// S3 Bucket in which to save Map Run results.
	Bucket() awss3.IBucket
	// S3 prefix in which to save Map Run results.
	// Default: - No prefix.
	//
	Prefix() *string
	// Compile policy statements to provide relevent permissions to the state machine.
	ProvidePolicyStatements() *[]awsiam.PolicyStatement
	// Render ResultWriter in ASL JSON format.
	Render() interface{}
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

func (r *jsiiProxy_ResultWriter) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

