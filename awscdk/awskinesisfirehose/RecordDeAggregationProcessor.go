package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The data processor for multi record deaggrecation.
//
// Record deaggregation by JSON or by delimiter is capped at 500 per record.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	DynamicPartitioning: &DynamicPartitioningProps{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	Processors: []IDataProcessor{
//   		firehose.RecordDeAggregationProcessor_Json(),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/dynamic-partitioning-multirecord-deaggergation.html
//
type RecordDeAggregationProcessor interface {
	IDataProcessor
	// The constructor props of the DataProcessor.
	Props() *DataProcessorProps
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for RecordDeAggregationProcessor
type jsiiProxy_RecordDeAggregationProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_RecordDeAggregationProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewRecordDeAggregationProcessor(options *RecordDeAggregationProcessorOptions) RecordDeAggregationProcessor {
	_init_.Initialize()

	if err := validateNewRecordDeAggregationProcessorParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_RecordDeAggregationProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.RecordDeAggregationProcessor",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewRecordDeAggregationProcessor_Override(r RecordDeAggregationProcessor, options *RecordDeAggregationProcessorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.RecordDeAggregationProcessor",
		[]interface{}{options},
		r,
	)
}

// Perform deaggregation based on a specified custom delimiter.
func RecordDeAggregationProcessor_Delimited(delimiter *string) RecordDeAggregationProcessor {
	_init_.Initialize()

	if err := validateRecordDeAggregationProcessor_DelimitedParameters(delimiter); err != nil {
		panic(err)
	}
	var returns RecordDeAggregationProcessor

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.RecordDeAggregationProcessor",
		"delimited",
		[]interface{}{delimiter},
		&returns,
	)

	return returns
}

// Perform deaggregation from JSON objects on a single line with no delimiter or newline-delimited (JSONL).
func RecordDeAggregationProcessor_Json() RecordDeAggregationProcessor {
	_init_.Initialize()

	var returns RecordDeAggregationProcessor

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.RecordDeAggregationProcessor",
		"json",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecordDeAggregationProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	if err := r.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

