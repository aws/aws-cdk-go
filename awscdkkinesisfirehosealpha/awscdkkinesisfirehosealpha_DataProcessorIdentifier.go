// The CDK Construct Library for AWS::KinesisFirehose
package awscdkkinesisfirehosealpha


// The key-value pair that identifies the underlying processor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import kinesisfirehose_alpha "github.com/aws/aws-cdk-go/awscdkkinesisfirehosealpha"
//
//   dataProcessorIdentifier := &dataProcessorIdentifier{
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processorparameter.html
//
// Experimental.
type DataProcessorIdentifier struct {
	// The parameter name that corresponds to the processor resource's identifier.
	//
	// Must be an accepted value in `CfnDeliveryStream.ProcessoryParameterProperty.ParameterName`.
	// Experimental.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The identifier of the underlying processor resource.
	// Experimental.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

