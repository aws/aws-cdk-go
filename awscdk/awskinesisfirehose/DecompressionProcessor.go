package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The data processor to decompress CloudWatch Logs.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	Processors: []IDataProcessor{
//   		firehose.NewDecompressionProcessor(),
//   		firehose.NewAppendDelimiterToRecordProcessor(),
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/writing-with-cloudwatch-logs-decompression.html
//
type DecompressionProcessor interface {
	IDataProcessor
	// The constructor props of the DataProcessor.
	Props() *DataProcessorProps
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for DecompressionProcessor
type jsiiProxy_DecompressionProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_DecompressionProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewDecompressionProcessor(options *DecompressionProcessorOptions) DecompressionProcessor {
	_init_.Initialize()

	if err := validateNewDecompressionProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_DecompressionProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.DecompressionProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewDecompressionProcessor_Override(d DecompressionProcessor, options *DecompressionProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.DecompressionProcessor",
		[]interface{}{options},
		d,
	)
}

func (d *jsiiProxy_DecompressionProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	if err := d.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

