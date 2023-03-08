package awskinesisfirehose


// The `ProcessingConfiguration` property configures data processing for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processingConfigurationProperty := &processingConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   	processors: []interface{}{
//   		&processorProperty{
//   			type: jsii.String("type"),
//
//   			// the properties below are optional
//   			parameters: []interface{}{
//   				&processorParameterProperty{
//   					parameterName: jsii.String("parameterName"),
//   					parameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeliveryStream_ProcessingConfigurationProperty struct {
	// Indicates whether data processing is enabled (true) or disabled (false).
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The data processors.
	Processors interface{} `field:"optional" json:"processors" yaml:"processors"`
}

