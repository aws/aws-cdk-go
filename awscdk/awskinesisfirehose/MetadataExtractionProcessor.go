package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The data processor for dynamic partitioning with inline parsing.
//
// Example:
//   var bucket Bucket
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	DynamicPartitioning: &DynamicPartitioningProps{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	Processors: []IDataProcessor{
//   		firehose.MetadataExtractionProcessor_Jq16(map[string]*string{
//   			"customer_id": jsii.String(".customer_id"),
//   			"device": jsii.String(".type.device"),
//   			"year": jsii.String(".event_timestamp|strftime(\"%Y\")"),
//   		}),
//   	},
//   	DataOutputPrefix: jsii.String("!{partitionKeyFromQuery:year}/!{partitionKeyFromQuery:device}/!{partitionKeyFromQuery:customer_id}/"),
//   })
//   firehose.NewDeliveryStream(this, jsii.String("DeliveryStream"), &DeliveryStreamProps{
//   	Destination: s3Destination,
//   })
//
// See: https://docs.aws.amazon.com/firehose/latest/dev/dynamic-partitioning-partitioning-keys.html
//
type MetadataExtractionProcessor interface {
	IDataProcessor
	// The constructor props of the DataProcessor.
	Props() *DataProcessorProps
	// Binds this processor to a destination of a delivery stream.
	//
	// Implementers should use this method to grant processor invocation permissions to the provided stream and return the
	// necessary configuration to register as a processor.
	Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig
}

// The jsii proxy struct for MetadataExtractionProcessor
type jsiiProxy_MetadataExtractionProcessor struct {
	jsiiProxy_IDataProcessor
}

func (j *jsiiProxy_MetadataExtractionProcessor) Props() *DataProcessorProps {
	var returns *DataProcessorProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewMetadataExtractionProcessor(options *MetadataExtractionProcessorOptions, keys *[]*string) MetadataExtractionProcessor {
	_init_.Initialize()

	if err := validateNewMetadataExtractionProcessorParameters(options, keys); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetadataExtractionProcessor{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.MetadataExtractionProcessor",
		[]interface{}{options, keys},
		&j,
	)

	return &j
}

func NewMetadataExtractionProcessor_Override(m MetadataExtractionProcessor, options *MetadataExtractionProcessorOptions, keys *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.MetadataExtractionProcessor",
		[]interface{}{options, keys},
		m,
	)
}

// Creates the inline parsing configuration with JQ 1.6 engine.
func MetadataExtractionProcessor_Jq16(query *map[string]*string) MetadataExtractionProcessor {
	_init_.Initialize()

	if err := validateMetadataExtractionProcessor_Jq16Parameters(query); err != nil {
		panic(err)
	}
	var returns MetadataExtractionProcessor

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.MetadataExtractionProcessor",
		"jq16",
		[]interface{}{query},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetadataExtractionProcessor) Bind(scope constructs.Construct, options *DataProcessorBindOptions) *DataProcessorConfig {
	if err := m.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DataProcessorConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

