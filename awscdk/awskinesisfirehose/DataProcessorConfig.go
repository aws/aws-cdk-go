package awskinesisfirehose


// The full configuration of a data processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProcessorConfig := &DataProcessorConfig{
//   	ProcessorIdentifier: &DataProcessorIdentifier{
//   		ParameterName: jsii.String("parameterName"),
//   		ParameterValue: jsii.String("parameterValue"),
//   	},
//   	ProcessorType: jsii.String("processorType"),
//
//   	// the properties below are optional
//   	Parameters: []ProcessorParameterProperty{
//   		&ProcessorParameterProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisfirehose-deliverystream-processor.html
//
type DataProcessorConfig struct {
	// The key-value pair that identifies the underlying processor resource.
	//
	// Ignored when the `parameters` is specified.
	ProcessorIdentifier *DataProcessorIdentifier `field:"required" json:"processorIdentifier" yaml:"processorIdentifier"`
	// The type of processor.
	ProcessorType *string `field:"required" json:"processorType" yaml:"processorType"`
	// The processor parameters.
	// Default: - No processor parameters.
	//
	Parameters *[]*CfnDeliveryStream_ProcessorParameterProperty `field:"optional" json:"parameters" yaml:"parameters"`
}

