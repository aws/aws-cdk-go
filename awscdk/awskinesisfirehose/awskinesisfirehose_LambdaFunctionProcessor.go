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

	if err := validateNewLambdaFunctionProcessorParameters(lambdaFunction, props); err != nil {
		panic(err)
	}
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
	if err := l.validateBindParameters(_scope, options); err != nil {
		panic(err)
	}
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

