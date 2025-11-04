package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The data processor to append new line delimiter to each record.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	Processors: []IDataProcessor{
//   		firehose.NewAppendDelimiterToRecordProcessor(),
//   	},
//   })
//   firehose.NewDeliveryStream(this, jsii.String("Delivery Stream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/dynamic-partitioning-s3bucketprefix.html#dynamic-partitioning-new-line-delimiter
//
type AppendDelimiterToRecordProcessor interface {
	IDataProcessor
	// The constructor props of the DataProcessor.
	Props() *DataProcessorProps
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for AppendDelimiterToRecordProcessor
type jsiiProxy_AppendDelimiterToRecordProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_AppendDelimiterToRecordProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewAppendDelimiterToRecordProcessor() AppendDelimiterToRecordProcessor {
	_init_.Initialize()

	j := jsiiProxy_AppendDelimiterToRecordProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.AppendDelimiterToRecordProcessor",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAppendDelimiterToRecordProcessor_Override(a AppendDelimiterToRecordProcessor) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.AppendDelimiterToRecordProcessor",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AppendDelimiterToRecordProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	if err := a.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

