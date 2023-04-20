package awskinesisfirehosedestinations

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Properties for defining an S3 backup destination.
//
// S3 backup is available for all destinations, regardless of whether the final destination is S3 or not.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//
//   stack := cdk.NewStack(app, jsii.String("aws-cdk-firehose-delivery-stream-s3-all-properties"))
//
//   bucket := s3.NewBucket(stack, jsii.String("Bucket"), &BucketProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//
//   backupBucket := s3.NewBucket(stack, jsii.String("BackupBucket"), &BucketProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	AutoDeleteObjects: jsii.Boolean(true),
//   })
//   logGroup := logs.NewLogGroup(stack, jsii.String("LogGroup"), &LogGroupProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   dataProcessorFunction := lambdanodejs.NewNodejsFunction(stack, jsii.String("DataProcessorFunction"), &NodejsFunctionProps{
//   	Entry: path.join(__dirname, jsii.String("lambda-data-processor.js")),
//   	Timeout: cdk.Duration_Minutes(jsii.Number(1)),
//   })
//
//   processor := firehose.NewLambdaFunctionProcessor(dataProcessorFunction, &DataProcessorProps{
//   	BufferInterval: cdk.Duration_Seconds(jsii.Number(60)),
//   	BufferSize: cdk.Size_Mebibytes(jsii.Number(1)),
//   	Retries: jsii.Number(1),
//   })
//
//   key := kms.NewKey(stack, jsii.String("Key"), &KeyProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   backupKey := kms.NewKey(stack, jsii.String("BackupKey"), &KeyProps{
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   })
//
//   firehose.NewDeliveryStream(stack, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destinations: []iDestination{
//   		destinations.NewS3Bucket(bucket, &S3BucketProps{
//   			Logging: jsii.Boolean(true),
//   			LogGroup: logGroup,
//   			Processor: processor,
//   			Compression: destinations.Compression_GZIP(),
//   			DataOutputPrefix: jsii.String("regularPrefix"),
//   			ErrorOutputPrefix: jsii.String("errorPrefix"),
//   			BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
//   			BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
//   			EncryptionKey: key,
//   			S3Backup: &DestinationS3BackupProps{
//   				Mode: destinations.BackupMode_ALL,
//   				Bucket: backupBucket,
//   				Compression: destinations.Compression_ZIP(),
//   				DataOutputPrefix: jsii.String("backupPrefix"),
//   				ErrorOutputPrefix: jsii.String("backupErrorPrefix"),
//   				BufferingInterval: cdk.Duration_*Seconds(jsii.Number(60)),
//   				BufferingSize: cdk.Size_*Mebibytes(jsii.Number(1)),
//   				EncryptionKey: backupKey,
//   			},
//   		}),
//   	},
//   })
//
//   app.Synth()
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

