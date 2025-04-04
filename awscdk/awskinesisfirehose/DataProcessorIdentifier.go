package awskinesisfirehose


// The key-value pair that identifies the underlying processor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataProcessorIdentifier := &DataProcessorIdentifier{
//   	ParameterName: jsii.String("parameterName"),
//   	ParameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processorparameter.html
//
type DataProcessorIdentifier struct {
	// The parameter name that corresponds to the processor resource's identifier.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessoryParameterProperty.ParameterName`.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The identifier of the underlying processor resource.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

