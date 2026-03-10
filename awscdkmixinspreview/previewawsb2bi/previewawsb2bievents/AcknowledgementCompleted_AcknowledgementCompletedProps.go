package previewawsb2bievents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.b2bi@AcknowledgementCompleted event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acknowledgementCompletedProps := &AcknowledgementCompletedProps{
//   	AckErrorCodeDetected: []*string{
//   		jsii.String("ackErrorCodeDetected"),
//   	},
//   	AckFileS3Attributes: &AckFileS3Attributes{
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
//   	AckX12Type: []*string{
//   		jsii.String("ackX12Type"),
//   	},
//   	AckX12Version: []*string{
//   		jsii.String("ackX12Version"),
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
//   	InputX12TransactionSet: []*string{
//   		jsii.String("inputX12TransactionSet"),
//   	},
//   	InputX12Version: []*string{
//   		jsii.String("inputX12Version"),
//   	},
//   	StartTimestamp: []*string{
//   		jsii.String("startTimestamp"),
//   	},
//   	TradingPartnerId: []*string{
//   		jsii.String("tradingPartnerId"),
//   	},
//   	TransformerJobId: []*string{
//   		jsii.String("transformerJobId"),
//   	},
//   }
//
// Experimental.
type AcknowledgementCompleted_AcknowledgementCompletedProps struct {
	// ack-error-code-detected property.
	//
	// Specify an array of string values to match this event if the actual value of ack-error-code-detected is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AckErrorCodeDetected *[]*string `field:"optional" json:"ackErrorCodeDetected" yaml:"ackErrorCodeDetected"`
	// ack-file-s3-attributes property.
	//
	// Specify an array of string values to match this event if the actual value of ack-file-s3-attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AckFileS3Attributes *AcknowledgementCompleted_AckFileS3Attributes `field:"optional" json:"ackFileS3Attributes" yaml:"ackFileS3Attributes"`
	// ack-x12-type property.
	//
	// Specify an array of string values to match this event if the actual value of ack-x12-type is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AckX12Type *[]*string `field:"optional" json:"ackX12Type" yaml:"ackX12Type"`
	// ack-x12-version property.
	//
	// Specify an array of string values to match this event if the actual value of ack-x12-version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AckX12Version *[]*string `field:"optional" json:"ackX12Version" yaml:"ackX12Version"`
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
	// input-file-s3-attributes property.
	//
	// Specify an array of string values to match this event if the actual value of input-file-s3-attributes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputFileS3Attributes *AcknowledgementCompleted_InputFileS3Attributes `field:"optional" json:"inputFileS3Attributes" yaml:"inputFileS3Attributes"`
	// input-x12-transaction-set property.
	//
	// Specify an array of string values to match this event if the actual value of input-x12-transaction-set is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputX12TransactionSet *[]*string `field:"optional" json:"inputX12TransactionSet" yaml:"inputX12TransactionSet"`
	// input-x12-version property.
	//
	// Specify an array of string values to match this event if the actual value of input-x12-version is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InputX12Version *[]*string `field:"optional" json:"inputX12Version" yaml:"inputX12Version"`
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
	// transformer-job-id property.
	//
	// Specify an array of string values to match this event if the actual value of transformer-job-id is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TransformerJobId *[]*string `field:"optional" json:"transformerJobId" yaml:"transformerJobId"`
}

