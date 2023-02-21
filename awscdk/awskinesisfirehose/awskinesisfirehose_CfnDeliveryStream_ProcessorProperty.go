package awskinesisfirehose


// The `Processor` property specifies a data processor for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorProperty := &processorProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	parameters: []interface{}{
//   		&processorParameterProperty{
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
type CfnDeliveryStream_ProcessorProperty struct {
	// The type of processor.
	//
	// Valid values: `Lambda` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The processor parameters.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

