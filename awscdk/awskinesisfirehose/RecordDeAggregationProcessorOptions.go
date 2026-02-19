package awskinesisfirehose


// Props for RecordDeAggregationProcessor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recordDeAggregationProcessorOptions := &RecordDeAggregationProcessorOptions{
//   	SubRecordType: awscdk.Aws_kinesisfirehose.SubRecordType_JSON,
//
//   	// the properties below are optional
//   	Delimiter: jsii.String("delimiter"),
//   }
//
type RecordDeAggregationProcessorOptions struct {
	// The sub-record type to deaggregate input records.
	SubRecordType SubRecordType `field:"required" json:"subRecordType" yaml:"subRecordType"`
	// The custom delimiter when subRecordType is DELIMITED.
	//
	// Must be specified in the base64-encoded format.
	// Default: - No delimiter.
	//
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

