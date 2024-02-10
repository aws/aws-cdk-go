package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Item Reader configuration for iterating over objects in an S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3ObjectsItemReader := awscdk.Aws_stepfunctions.NewS3ObjectsItemReader(&S3ObjectsItemReaderProps{
//   	Bucket: bucket,
//
//   	// the properties below are optional
//   	MaxItems: jsii.Number(123),
//   	Prefix: jsii.String("prefix"),
//   })
//
type S3ObjectsItemReader interface {
	IItemReader
	// S3 Bucket containing objects to iterate over.
	Bucket() awss3.IBucket
	// Limits the number of items passed to the Distributed Map state.
	// Default: - Distributed Map state will iterate over all items provided by the ItemReader.
	//
	MaxItems() *float64
	// S3 prefix used to limit objects to iterate over.
	// Default: - No prefix.
	//
	Prefix() *string
	// ARN for the `listObjectsV2` method of the S3 API This API method is used to iterate all objects in the S3 bucket/prefix.
	Resource() *string
	// Compile policy statements to provide relevent permissions to the state machine.
	ProvidePolicyStatements() *[]awsiam.PolicyStatement
	// Renders the ItemReader configuration as JSON object.
	//
	// Returns: - JSON object.
	Render() interface{}
}

// The jsii proxy struct for S3ObjectsItemReader
type jsiiProxy_S3ObjectsItemReader struct {
	jsiiProxy_IItemReader
}

func (j *jsiiProxy_S3ObjectsItemReader) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectsItemReader) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectsItemReader) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectsItemReader) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


func NewS3ObjectsItemReader(props *S3ObjectsItemReaderProps) S3ObjectsItemReader {
	_init_.Initialize()

	if err := validateNewS3ObjectsItemReaderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ObjectsItemReader{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3ObjectsItemReader",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3ObjectsItemReader_Override(s S3ObjectsItemReader, props *S3ObjectsItemReaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3ObjectsItemReader",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3ObjectsItemReader) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		s,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectsItemReader) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

