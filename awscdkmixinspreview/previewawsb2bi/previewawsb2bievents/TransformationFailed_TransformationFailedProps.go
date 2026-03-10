package previewawsb2bievents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.b2bi@TransformationFailed event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformationFailedProps := &TransformationFailedProps{
//   	AckErrorCodeDetected: []*string{
//   		jsii.String("ackErrorCodeDetected"),
//   	},
//   	AckGenerationStatus: []*string{
//   		jsii.String("ackGenerationStatus"),
//   	},
//   	EndTimestamp: []*string{
//   		jsii.String("endTimestamp"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	FailureCode: []*string{
//   		jsii.String("failureCode"),
//   	},
//   	FailureMessage: []*string{
//   		jsii.String("failureMessage"),
//   	},
//   	InputFileS3Attributes: &InputFileS3Attributes{
//   		Bucket: []*string{
//   			jsii.String("bucket"),
//   		},
//   		ObjectKey: []*string{
//   			jsii.String("objectKey"),
//   		},
//   		ObjectSizeBytes: []*string{
//   			jsii.String("objectSizeBytes"),
//   		},
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	TradingPartnerId: []*string{
//   		jsii.String("tradingPartnerId"),
//   	},
//   	X12TransactionSet: []*string{
//   		jsii.String("x12TransactionSet"),
//   	},
//   	X12Version: []*string{
//   		jsii.String("x12Version"),
//   	},
//   }
//
// Experimental.
type TransformationFailed_TransformationFailedProps struct {
	// ack-error-code-detected property.
	//
	// Specify an array of string values to match this event if the actual value of ack-error-code-detected is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AckErrorCodeDetected *[]*string `field:"optional" json:"ackErrorCodeDetected" yaml:"ackErrorCodeDetected"`
	// ack-generation-status property.
	//
	// Specify an array of string values to match this event if the actual value of ack-generation-status is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AckGenerationStatus *[]*string `field:"optional" json:"ackGenerationStatus" yaml:"ackGenerationStatus"`
	// end-timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of end-timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EndTimestamp *[]*string `field:"optional" json:"endTimestamp" yaml:"endTimestamp"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// failure-code property.
	//
	// Specify an array of string values to match this event if the actual value of failure-code is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureCode *[]*string `field:"optional" json:"failureCode" yaml:"failureCode"`
	// failure-message property.
	//
	// Specify an array of string values to match this event if the actual value of failure-message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailureMessage *[]*string `field:"optional" json:"failureMessage" yaml:"failureMessage"`
	// input-file-s3-attributes property.
	//
	// Specify an array of string values to match this event if the actual value of input-file-s3-attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputFileS3Attributes *TransformationFailed_InputFileS3Attributes `field:"optional" json:"inputFileS3Attributes" yaml:"inputFileS3Attributes"`
	// start-timestamp property.
	//
	// Specify an array of string values to match this event if the actual value of start-timestamp is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTimestamp *[]*string `field:"optional" json:"startTimestamp" yaml:"startTimestamp"`
	// trading-partner-id property.
	//
	// Specify an array of string values to match this event if the actual value of trading-partner-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TradingPartnerId *[]*string `field:"optional" json:"tradingPartnerId" yaml:"tradingPartnerId"`
	// x12-transaction-set property.
	//
	// Specify an array of string values to match this event if the actual value of x12-transaction-set is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	X12TransactionSet *[]*string `field:"optional" json:"x12TransactionSet" yaml:"x12TransactionSet"`
	// x12-version property.
	//
	// Specify an array of string values to match this event if the actual value of x12-version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	X12Version *[]*string `field:"optional" json:"x12Version" yaml:"x12Version"`
}

