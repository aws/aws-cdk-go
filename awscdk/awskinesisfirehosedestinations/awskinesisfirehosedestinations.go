package awskinesisfirehosedestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehosedestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// Generic properties for defining a delivery stream destination.
// Experimental.
type CommonDestinationProps struct {
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `json:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// An S3 bucket destination for data from a Kinesis Data Firehose delivery stream.
// Experimental.
type S3Bucket interface {
	awskinesisfirehose.IDestination
	Bind(scope constructs.Construct, _options *awskinesisfirehose.DestinationBindOptions) *awskinesisfirehose.DestinationConfig
}

// The jsii proxy struct for S3Bucket
type jsiiProxy_S3Bucket struct {
	internal.Type__awskinesisfirehoseIDestination
}

// Experimental.
func NewS3Bucket(bucket awss3.IBucket, props *S3BucketProps) S3Bucket {
	_init_.Initialize()

	j := jsiiProxy_S3Bucket{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose_destinations.S3Bucket",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Bucket_Override(s S3Bucket, bucket awss3.IBucket, props *S3BucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose_destinations.S3Bucket",
		[]interface{}{bucket, props},
		s,
	)
}

// Binds this destination to the Kinesis Data Firehose delivery stream.
//
// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
// Experimental.
func (s *jsiiProxy_S3Bucket) Bind(scope constructs.Construct, _options *awskinesisfirehose.DestinationBindOptions) *awskinesisfirehose.DestinationConfig {
	var returns *awskinesisfirehose.DestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, _options},
		&returns,
	)

	return returns
}

// Props for defining an S3 destination of a Kinesis Data Firehose delivery stream.
// Experimental.
type S3BucketProps struct {
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `json:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

