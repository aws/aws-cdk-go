package awskinesisfirehosedestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible compression options Kinesis Data Firehose can use to compress data on delivery.
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

	if err := validateCompression_OfParameters(value); err != nil {
		panic(err)
	}
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

