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
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   /**
//    * Tree view of bucket:
//    *  my-bucket
//    *  |
//    *  +--item1
//    *  |
//    *  +--otherItem
//    *  |
//    *  +--item2
//    *  |
//    *  ...
//    */
//   bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
//   	BucketName: jsii.String("my-bucket"),
//   })
//
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("DistributedMap"), &DistributedMapProps{
//   	ItemReader: sfn.NewS3ObjectsItemReader(&S3ObjectsItemReaderProps{
//   		Bucket: *Bucket,
//   		Prefix: jsii.String("item"),
//   	}),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass")))
//
type S3ObjectsItemReader interface {
	IItemReader
	// S3 Bucket containing objects to iterate over.
	Bucket() awss3.IBucket
	// S3 bucket name containing objects to iterate over or a file with a list to iterate over, as JsonPath.
	BucketNamePath() *string
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
	Render(queryLanguage QueryLanguage) interface{}
	// Validate that ItemReader contains exactly either.
	// See: bucketNamePath.
	//
	ValidateItemReader() *[]*string
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

func (j *jsiiProxy_S3ObjectsItemReader) BucketNamePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNamePath",
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

func (s *jsiiProxy_S3ObjectsItemReader) Render(queryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"render",
		[]interface{}{queryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectsItemReader) ValidateItemReader() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validateItemReader",
		nil, // no parameters
		&returns,
	)

	return returns
}

