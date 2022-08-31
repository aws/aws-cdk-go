package awskinesisfirehosedestinations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Possible compression options Kinesis Data Firehose can use to compress data on delivery.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//   import lambdanodejs "github.com/aws/aws-cdk-go/awscdk"
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
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

