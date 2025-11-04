package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The data processor to extract message after decompression of CloudWatch Logs.
//
// This processor must used with `DecompressionProcessor`.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	Processors: []IDataProcessor{
//   		firehose.NewDecompressionProcessor(),
//   		firehose.NewCloudWatchLogProcessor(&CloudWatchLogProcessorOptions{
//   			DataMessageExtraction: jsii.Boolean(true),
//   		}),
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/Message_extraction.html
//
type CloudWatchLogProcessor interface {
	IDataProcessor
	// The constructor props of the DataProcessor.
	Props() *DataProcessorProps
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for CloudWatchLogProcessor
type jsiiProxy_CloudWatchLogProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_CloudWatchLogProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewCloudWatchLogProcessor(options *CloudWatchLogProcessorOptions) CloudWatchLogProcessor {
	_init_.Initialize()

	if err := validateNewCloudWatchLogProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchLogProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.CloudWatchLogProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewCloudWatchLogProcessor_Override(c CloudWatchLogProcessor, options *CloudWatchLogProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.CloudWatchLogProcessor",
		[]interface{}{options},
		c,
	)
}

func (c *jsiiProxy_CloudWatchLogProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	if err := c.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

