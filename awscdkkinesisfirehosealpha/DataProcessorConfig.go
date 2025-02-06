package awscdkkinesisfirehosealpha


// The full configuration of a data processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//
//   dataProcessorConfig := &DataProcessorConfig{
//   	ProcessorIdentifier: &DataProcessorIdentifier{
//   		ParameterName: jsii.String("parameterName"),
//   		ParameterValue: jsii.String("parameterValue"),
//   	},
//   	ProcessorType: jsii.String("processorType"),
//   }
//
// Deprecated.
type DataProcessorConfig struct {
	// The key-value pair that identifies the underlying processor resource.
	// Deprecated.
	ProcessorIdentifier *DataProcessorIdentifier `field:"required" json:"processorIdentifier" yaml:"processorIdentifier"`
	// The type of the underlying processor resource.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessorProperty.Type`.
	//
	// Example:
	//   "Lambda"
	//
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processor.html#cfn-kinesisfirehose-deliverystream-processor-type
	//
	// Deprecated.
	ProcessorType *string `field:"required" json:"processorType" yaml:"processorType"`
}

