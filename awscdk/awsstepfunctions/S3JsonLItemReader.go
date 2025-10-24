package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Item Reader configuration for iterating over the rows of the JSONL file stored in S3.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket Bucket
//
//   s3JsonLItemReader := awscdk.Aws_stepfunctions.NewS3JsonLItemReader(&S3FileItemReaderProps{
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	Bucket: bucket,
//   	BucketNamePath: jsii.String("bucketNamePath"),
//   	MaxItems: jsii.Number(123),
//   })
//
type S3JsonLItemReader interface {
	IItemReader
	// S3 Bucket containing a file with a list to iterate over.
	Bucket() awss3.IBucket
	// S3 bucket name containing objects to iterate over or a file with a list to iterate over, as JsonPath.
	BucketNamePath() *string
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
	Render(queryLanguage QueryLanguage) interface{}
	// Validate that ItemReader contains exactly either.
	// See: bucketNamePath.
	//
	ValidateItemReader() *[]*string
}

// The jsii proxy struct for S3JsonLItemReader
type jsiiProxy_S3JsonLItemReader struct {
	jsiiProxy_IItemReader
}

func (j *jsiiProxy_S3JsonLItemReader) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonLItemReader) BucketNamePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNamePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonLItemReader) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonLItemReader) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonLItemReader) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonLItemReader) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


func NewS3JsonLItemReader(props *S3FileItemReaderProps) S3JsonLItemReader {
	_init_.Initialize()

	if err := validateNewS3JsonLItemReaderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3JsonLItemReader{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3JsonLItemReader",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3JsonLItemReader_Override(s S3JsonLItemReader, props *S3FileItemReaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3JsonLItemReader",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3JsonLItemReader) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		s,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3JsonLItemReader) Render(queryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"render",
		[]interface{}{queryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3JsonLItemReader) ValidateItemReader() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validateItemReader",
		nil, // no parameters
		&returns,
	)

	return returns
}

