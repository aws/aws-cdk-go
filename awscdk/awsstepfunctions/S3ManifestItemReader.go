package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Item Reader configuration for iterating over items in a S3 inventory manifest file stored in S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3ManifestItemReader := awscdk.Aws_stepfunctions.NewS3ManifestItemReader(&S3FileItemReaderProps{
//   	Bucket: bucket,
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	MaxItems: jsii.Number(123),
//   })
//
type S3ManifestItemReader interface {
	IItemReader
	// S3 Bucket containing a file with a list to iterate over.
	Bucket() awss3.IBucket
	InputType() *string
	// S3 key of a file with a list to iterate over.
	Key() *string
	// Limits the number of items passed to the Distributed Map state.
	// Default: - No maxItems.
	//
	MaxItems() *float64
	// ARN for the `getObject` method of the S3 API This API method is used to iterate all objects in the S3 bucket/prefix.
	Resource() *string
	// Compile policy statements to provide relevent permissions to the state machine.
	ProvidePolicyStatements() *[]awsiam.PolicyStatement
	// Renders the ItemReader configuration as JSON object.
	//
	// Returns: - JSON object.
	Render() interface{}
}

// The jsii proxy struct for S3ManifestItemReader
type jsiiProxy_S3ManifestItemReader struct {
	jsiiProxy_IItemReader
}

func (j *jsiiProxy_S3ManifestItemReader) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ManifestItemReader) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ManifestItemReader) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ManifestItemReader) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ManifestItemReader) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


func NewS3ManifestItemReader(props *S3FileItemReaderProps) S3ManifestItemReader {
	_init_.Initialize()

	if err := validateNewS3ManifestItemReaderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ManifestItemReader{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3ManifestItemReader",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3ManifestItemReader_Override(s S3ManifestItemReader, props *S3FileItemReaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3ManifestItemReader",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3ManifestItemReader) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		s,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ManifestItemReader) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

