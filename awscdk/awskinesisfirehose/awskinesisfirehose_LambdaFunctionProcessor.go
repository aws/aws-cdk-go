package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// Use an AWS Lambda function to transform records.
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
type LambdaFunctionProcessor interface {
	IDataProcessor
	// The constructor props of the LambdaFunctionProcessor.
	// Experimental.
	Props() *DataProcessorProps
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	// Experimental.
	Bind(_scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for LambdaFunctionProcessor
type jsiiProxy_LambdaFunctionProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_LambdaFunctionProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunctionProcessor(lambdaFunction awslambda.IFunction, props *DataProcessorProps) LambdaFunctionProcessor {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionProcessor{}

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.LambdaFunctionProcessor",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionProcessor_Override(l LambdaFunctionProcessor, lambdaFunction awslambda.IFunction, props *DataProcessorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_kinesisfirehose.LambdaFunctionProcessor",
		[]interface{}{lambdaFunction, props},
		l,
	)
}

func (l *jsiiProxy_LambdaFunctionProcessor) Bind(_scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

