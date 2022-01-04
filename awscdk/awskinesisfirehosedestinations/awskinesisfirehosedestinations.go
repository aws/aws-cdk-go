package awskinesisfirehosedestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/awskinesisfirehosedestinations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v3"
)

// Options for S3 record backup of a delivery stream.
//
// TODO: EXAMPLE
//
// Experimental.
type BackupMode string

const (
	BackupMode_ALL BackupMode = "ALL"
	BackupMode_FAILED BackupMode = "FAILED"
)

// Generic properties for defining a delivery stream destination.
//
// TODO: EXAMPLE
//
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
	// The data transformation that should be performed on the data before writing to the destination.
	// Experimental.
	Processor awskinesisfirehose.IDataProcessor `json:"processor"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The configuration for backing up source records to S3.
	// Experimental.
	S3Backup *DestinationS3BackupProps `json:"s3Backup"`
}

// Common properties for defining a backup, intermediary, or final S3 destination for a Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
// Experimental.
type CommonDestinationS3Props struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `json:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `json:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `json:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `json:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `json:"errorOutputPrefix"`
}

// Possible compression options Kinesis Data Firehose can use to compress data on delivery.
//
// TODO: EXAMPLE
//
// Experimental.
type Compression interface {
	Value() *string
}

// The jsii proxy struct for Compression
type jsiiProxy_Compression struct {
	_ byte // padding
}

func (j *jsiiProxy_Compression) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates a new Compression instance with a custom value.
// Experimental.
func Compression_Of(value *string) Compression {
	_init_.Initialize()

	var returns Compression

	_jsii_.StaticInvoke(
		"monocdk.aws_kinesisfirehose_destinations.Compression",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Compression_GZIP() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"monocdk.aws_kinesisfirehose_destinations.Compression",
		"GZIP",
		&returns,
	)
	return returns
}

func Compression_HADOOP_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"monocdk.aws_kinesisfirehose_destinations.Compression",
		"HADOOP_SNAPPY",
		&returns,
	)
	return returns
}

func Compression_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"monocdk.aws_kinesisfirehose_destinations.Compression",
		"SNAPPY",
		&returns,
	)
	return returns
}

func Compression_ZIP() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"monocdk.aws_kinesisfirehose_destinations.Compression",
		"ZIP",
		&returns,
	)
	return returns
}

// Properties for defining an S3 backup destination.
//
// S3 backup is available for all destinations, regardless of whether the final destination is S3 or not.
//
// TODO: EXAMPLE
//
// Experimental.
type DestinationS3BackupProps struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `json:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `json:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `json:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `json:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `json:"errorOutputPrefix"`
	// The S3 bucket that will store data and failed records.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `json:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// Indicates the mode by which incoming records should be backed up to S3, if any.
	//
	// If `bucket` is provided, this will be implicitly set to `BackupMode.ALL`.
	// Experimental.
	Mode BackupMode `json:"mode"`
}

// An S3 bucket destination for data from a Kinesis Data Firehose delivery stream.
//
// TODO: EXAMPLE
//
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
//
// TODO: EXAMPLE
//
// Experimental.
type S3BucketProps struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `json:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `json:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `json:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `json:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `json:"errorOutputPrefix"`
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `json:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Experimental.
	Processor awskinesisfirehose.IDataProcessor `json:"processor"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The configuration for backing up source records to S3.
	// Experimental.
	S3Backup *DestinationS3BackupProps `json:"s3Backup"`
}

