package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Item Reader configuration for iterating over items in a JSON array stored in a S3 file.
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
type S3JsonItemReader interface {
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

// The jsii proxy struct for S3JsonItemReader
type jsiiProxy_S3JsonItemReader struct {
	jsiiProxy_IItemReader
}

func (j *jsiiProxy_S3JsonItemReader) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonItemReader) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonItemReader) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonItemReader) MaxItems() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxItems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3JsonItemReader) Resource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}


func NewS3JsonItemReader(props *S3FileItemReaderProps) S3JsonItemReader {
	_init_.Initialize()

	if err := validateNewS3JsonItemReaderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3JsonItemReader{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3JsonItemReader",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewS3JsonItemReader_Override(s S3JsonItemReader, props *S3FileItemReaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.S3JsonItemReader",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_S3JsonItemReader) ProvidePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		s,
		"providePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3JsonItemReader) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

