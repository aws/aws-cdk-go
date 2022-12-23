package awskinesisfirehose


// The `ProcessorParameter` property specifies a processor parameter in a data processor for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorParameterProperty := &processorParameterProperty{
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnDeliveryStream_ProcessorParameterProperty struct {
	// The name of the parameter.
	//
	// Currently the following default values are supported: 3 for `NumberOfRetries` and 60 for the `BufferIntervalInSeconds` . The `BufferSizeInMBs` ranges between 0.2 MB and up to 3MB. The default buffering hint is 1MB for all destinations, except Splunk. For Splunk, the default buffering hint is 256 KB.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The parameter value.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

