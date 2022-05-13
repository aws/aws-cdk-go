// CDK Destinations Constructs for AWS Kinesis Firehose
package awscdkkinesisfirehosedestinationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha/v2"
	"github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for S3 record backup of a delivery stream.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//   app := cdk.NewApp()
//
//   stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))
//
//   bucket := s3.NewBucket(stack, jsii.String("Bucket"), &bucketProps{
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	autoDeleteObjects: jsii.Boolean(true),
//   })
//
//   backupBucket := s3.NewBucket(stack, jsii.String("BackupBucket"), &bucketProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   	autoDeleteObjects: jsii.Boolean(true),
//   })
//   logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &logGroupProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &nodejsFunctionProps{
//   	entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
//   	timeout: cdk.duration.minutes(jsii.Number(1)),
//   })
//
//   processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &dataProcessorProps{
//   	bufferInterval: cdk.*duration.seconds(jsii.Number(60)),
//   	bufferSize: cdk.size.mebibytes(jsii.Number(1)),
//   	retries: jsii.Number(1),
//   })
//
//   key := kms.NewKey(stack, jsii.String("Key"), &keyProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &keyProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   firehose.NewDeliveryStream(stack, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			logging: jsii.Boolean(true),
//   			logGroup: logGroup,
//   			processor: processor,
//   			compression: destinations.compression_GZIP(),
//   			dataOutputPrefix: jsii.String("regularPrefix"),
//   			errorOutputPrefix: jsii.String("errorPrefix"),
//   			bufferingInterval: cdk.*duration.seconds(jsii.Number(60)),
//   			bufferingSize: cdk.*size.mebibytes(jsii.Number(1)),
//   			encryptionKey: key,
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.backupMode_ALL,
//   				bucket: backupBucket,
//   				compression: destinations.*compression_ZIP(),
//   				dataOutputPrefix: jsii.String("backupPrefix"),
//   				errorOutputPrefix: jsii.String("backupErrorPrefix"),
//   				bufferingInterval: cdk.*duration.seconds(jsii.Number(60)),
//   				bufferingSize: cdk.*size.mebibytes(jsii.Number(1)),
//   				encryptionKey: backupKey,
//   			},
//   		}),
//   	},
//   })
//
//   app.synth()
//
// Experimental.
type BackupMode string

const (
	// All records are backed up.
	// Experimental.
	BackupMode_ALL BackupMode = "ALL"
	// Only records that failed to deliver or transform are backed up.
	// Experimental.
	BackupMode_FAILED BackupMode = "FAILED"
)

// Generic properties for defining a delivery stream destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import kinesisfirehose_destinations_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var compression compression
//   var dataProcessor iDataProcessor
//   var key key
//   var logGroup logGroup
//   var role role
//   var size size
//
//   commonDestinationProps := &commonDestinationProps{
//   	logging: jsii.Boolean(false),
//   	logGroup: logGroup,
//   	processor: dataProcessor,
//   	role: role,
//   	s3Backup: &destinationS3BackupProps{
//   		bucket: bucket,
//   		bufferingInterval: cdk.duration.minutes(jsii.Number(30)),
//   		bufferingSize: size,
//   		compression: compression,
//   		dataOutputPrefix: jsii.String("dataOutputPrefix"),
//   		encryptionKey: key,
//   		errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		logging: jsii.Boolean(false),
//   		logGroup: logGroup,
//   		mode: kinesisfirehose_destinations_alpha.backupMode_ALL,
//   	},
//   }
//
// Experimental.
type CommonDestinationProps struct {
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Experimental.
	Processor awscdkkinesisfirehosealpha.IDataProcessor `field:"optional" json:"processor" yaml:"processor"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The configuration for backing up source records to S3.
	// Experimental.
	S3Backup *DestinationS3BackupProps `field:"optional" json:"s3Backup" yaml:"s3Backup"`
}

// Common properties for defining a backup, intermediary, or final S3 destination for a Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_destinations_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var compression compression
//   var key key
//   var size size
//
//   commonDestinationS3Props := &commonDestinationS3Props{
//   	bufferingInterval: cdk.duration.minutes(jsii.Number(30)),
//   	bufferingSize: size,
//   	compression: compression,
//   	dataOutputPrefix: jsii.String("dataOutputPrefix"),
//   	encryptionKey: key,
//   	errorOutputPrefix: jsii.String("errorOutputPrefix"),
//   }
//
// Experimental.
type CommonDestinationS3Props struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
}

// Possible compression options Kinesis Data Firehose can use to compress data on delivery.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//   app := cdk.NewApp()
//
//   stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))
//
//   bucket := s3.NewBucket(stack, jsii.String("Bucket"), &bucketProps{
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	autoDeleteObjects: jsii.Boolean(true),
//   })
//
//   backupBucket := s3.NewBucket(stack, jsii.String("BackupBucket"), &bucketProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   	autoDeleteObjects: jsii.Boolean(true),
//   })
//   logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &logGroupProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &nodejsFunctionProps{
//   	entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
//   	timeout: cdk.duration.minutes(jsii.Number(1)),
//   })
//
//   processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &dataProcessorProps{
//   	bufferInterval: cdk.*duration.seconds(jsii.Number(60)),
//   	bufferSize: cdk.size.mebibytes(jsii.Number(1)),
//   	retries: jsii.Number(1),
//   })
//
//   key := kms.NewKey(stack, jsii.String("Key"), &keyProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &keyProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   firehose.NewDeliveryStream(stack, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			logging: jsii.Boolean(true),
//   			logGroup: logGroup,
//   			processor: processor,
//   			compression: destinations.compression_GZIP(),
//   			dataOutputPrefix: jsii.String("regularPrefix"),
//   			errorOutputPrefix: jsii.String("errorPrefix"),
//   			bufferingInterval: cdk.*duration.seconds(jsii.Number(60)),
//   			bufferingSize: cdk.*size.mebibytes(jsii.Number(1)),
//   			encryptionKey: key,
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.backupMode_ALL,
//   				bucket: backupBucket,
//   				compression: destinations.*compression_ZIP(),
//   				dataOutputPrefix: jsii.String("backupPrefix"),
//   				errorOutputPrefix: jsii.String("backupErrorPrefix"),
//   				bufferingInterval: cdk.*duration.seconds(jsii.Number(60)),
//   				bufferingSize: cdk.*size.mebibytes(jsii.Number(1)),
//   				encryptionKey: backupKey,
//   			},
//   		}),
//   	},
//   })
//
//   app.synth()
//
// Experimental.
type Compression interface {
	// the string value of the Compression.
	// Experimental.
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
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
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
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"GZIP",
		&returns,
	)
	return returns
}

func Compression_HADOOP_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"HADOOP_SNAPPY",
		&returns,
	)
	return returns
}

func Compression_SNAPPY() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"SNAPPY",
		&returns,
	)
	return returns
}

func Compression_ZIP() Compression {
	_init_.Initialize()
	var returns Compression
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.Compression",
		"ZIP",
		&returns,
	)
	return returns
}

// Properties for defining an S3 backup destination.
//
// S3 backup is available for all destinations, regardless of whether the final destination is S3 or not.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import firehose "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import destinations "github.com/aws/aws-cdk-go/awscdkkinesisfirehosedestinationsalpha"
//
//   app := cdk.NewApp()
//
//   stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))
//
//   bucket := s3.NewBucket(stack, jsii.String("Bucket"), &bucketProps{
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	autoDeleteObjects: jsii.Boolean(true),
//   })
//
//   backupBucket := s3.NewBucket(stack, jsii.String("BackupBucket"), &bucketProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   	autoDeleteObjects: jsii.Boolean(true),
//   })
//   logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &logGroupProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &nodejsFunctionProps{
//   	entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
//   	timeout: cdk.duration.minutes(jsii.Number(1)),
//   })
//
//   processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &dataProcessorProps{
//   	bufferInterval: cdk.*duration.seconds(jsii.Number(60)),
//   	bufferSize: cdk.size.mebibytes(jsii.Number(1)),
//   	retries: jsii.Number(1),
//   })
//
//   key := kms.NewKey(stack, jsii.String("Key"), &keyProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &keyProps{
//   	removalPolicy: cdk.*removalPolicy_DESTROY,
//   })
//
//   firehose.NewDeliveryStream(stack, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &s3BucketProps{
//   			logging: jsii.Boolean(true),
//   			logGroup: logGroup,
//   			processor: processor,
//   			compression: destinations.compression_GZIP(),
//   			dataOutputPrefix: jsii.String("regularPrefix"),
//   			errorOutputPrefix: jsii.String("errorPrefix"),
//   			bufferingInterval: cdk.*duration.seconds(jsii.Number(60)),
//   			bufferingSize: cdk.*size.mebibytes(jsii.Number(1)),
//   			encryptionKey: key,
//   			s3Backup: &destinationS3BackupProps{
//   				mode: destinations.backupMode_ALL,
//   				bucket: backupBucket,
//   				compression: destinations.*compression_ZIP(),
//   				dataOutputPrefix: jsii.String("backupPrefix"),
//   				errorOutputPrefix: jsii.String("backupErrorPrefix"),
//   				bufferingInterval: cdk.*duration.seconds(jsii.Number(60)),
//   				bufferingSize: cdk.*size.mebibytes(jsii.Number(1)),
//   				encryptionKey: backupKey,
//   			},
//   		}),
//   	},
//   })
//
//   app.synth()
//
// Experimental.
type DestinationS3BackupProps struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// The S3 bucket that will store data and failed records.
	// Experimental.
	Bucket awss3.IBucket `field:"optional" json:"bucket" yaml:"bucket"`
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Indicates the mode by which incoming records should be backed up to S3, if any.
	//
	// If `bucket` is provided, this will be implicitly set to `BackupMode.ALL`.
	// Experimental.
	Mode BackupMode `field:"optional" json:"mode" yaml:"mode"`
}

// An S3 bucket destination for data from a Kinesis Data Firehose delivery stream.
//
// Example:
//   var bucket bucket
//   // Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &dataProcessorProps{
//   	bufferInterval: awscdk.Duration.minutes(jsii.Number(5)),
//   	bufferSize: awscdk.Size.mebibytes(jsii.Number(5)),
//   	retries: jsii.Number(5),
//   })
//   s3Destination := destinations.NewS3Bucket(bucket, &s3BucketProps{
//   	processor: lambdaProcessor,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
// Experimental.
type S3Bucket interface {
	awscdkkinesisfirehosealpha.IDestination
	// Binds this destination to the Kinesis Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	// Experimental.
	Bind(scope constructs.Construct, _options *awscdkkinesisfirehosealpha.DestinationBindOptions) *awscdkkinesisfirehosealpha.DestinationConfig
}

// The jsii proxy struct for S3Bucket
type jsiiProxy_S3Bucket struct {
	internal.Type__awscdkkinesisfirehosealphaIDestination
}

// Experimental.
func NewS3Bucket(bucket awss3.IBucket, props *S3BucketProps) S3Bucket {
	_init_.Initialize()

	j := jsiiProxy_S3Bucket{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.S3Bucket",
		[]interface{}{bucket, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Bucket_Override(s S3Bucket, bucket awss3.IBucket, props *S3BucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisfirehose-destinations-alpha.S3Bucket",
		[]interface{}{bucket, props},
		s,
	)
}

func (s *jsiiProxy_S3Bucket) Bind(scope constructs.Construct, _options *awscdkkinesisfirehosealpha.DestinationBindOptions) *awscdkkinesisfirehosealpha.DestinationConfig {
	var returns *awscdkkinesisfirehosealpha.DestinationConfig

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
// Example:
//   var bucket bucket
//   // Provide a Lambda function that will transform records before delivery, with custom
//   // buffering and retry configuration
//   lambdaFunction := lambda.NewFunction(this, jsii.String("Processor"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_12_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("process-records"))),
//   })
//   lambdaProcessor := firehose.NewLambdaFunctionProcessor(lambdaFunction, &dataProcessorProps{
//   	bufferInterval: awscdk.Duration.minutes(jsii.Number(5)),
//   	bufferSize: awscdk.Size.mebibytes(jsii.Number(5)),
//   	retries: jsii.Number(5),
//   })
//   s3Destination := destinations.NewS3Bucket(bucket, &s3BucketProps{
//   	processor: lambdaProcessor,
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &deliveryStreamProps{
//   	destinations: []iDestination{
//   		s3Destination,
//   	},
//   })
//
// Experimental.
type S3BucketProps struct {
	// The length of time that Firehose buffers incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Duration.seconds(60)
	// Maximum: Duration.seconds(900)
	// Experimental.
	BufferingInterval awscdk.Duration `field:"optional" json:"bufferingInterval" yaml:"bufferingInterval"`
	// The size of the buffer that Kinesis Data Firehose uses for incoming data before delivering it to the S3 bucket.
	//
	// Minimum: Size.mebibytes(1)
	// Maximum: Size.mebibytes(128)
	// Experimental.
	BufferingSize awscdk.Size `field:"optional" json:"bufferingSize" yaml:"bufferingSize"`
	// The type of compression that Kinesis Data Firehose uses to compress the data that it delivers to the Amazon S3 bucket.
	//
	// The compression formats SNAPPY or ZIP cannot be specified for Amazon Redshift
	// destinations because they are not supported by the Amazon Redshift COPY operation
	// that reads from the S3 bucket.
	// Experimental.
	Compression Compression `field:"optional" json:"compression" yaml:"compression"`
	// A prefix that Kinesis Data Firehose evaluates and adds to records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	DataOutputPrefix *string `field:"optional" json:"dataOutputPrefix" yaml:"dataOutputPrefix"`
	// The AWS KMS key used to encrypt the data that it delivers to your Amazon S3 bucket.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3.
	//
	// This prefix appears immediately following the bucket name.
	// See: https://docs.aws.amazon.com/firehose/latest/dev/s3-prefixes.html
	//
	// Experimental.
	ErrorOutputPrefix *string `field:"optional" json:"errorOutputPrefix" yaml:"errorOutputPrefix"`
	// If true, log errors when data transformation or data delivery fails.
	//
	// If `logGroup` is provided, this will be implicitly set to `true`.
	// Experimental.
	Logging *bool `field:"optional" json:"logging" yaml:"logging"`
	// The CloudWatch log group where log streams will be created to hold error logs.
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The data transformation that should be performed on the data before writing to the destination.
	// Experimental.
	Processor awscdkkinesisfirehosealpha.IDataProcessor `field:"optional" json:"processor" yaml:"processor"`
	// The IAM role associated with this destination.
	//
	// Assumed by Kinesis Data Firehose to invoke processors and write to destinations.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The configuration for backing up source records to S3.
	// Experimental.
	S3Backup *DestinationS3BackupProps `field:"optional" json:"s3Backup" yaml:"s3Backup"`
}

