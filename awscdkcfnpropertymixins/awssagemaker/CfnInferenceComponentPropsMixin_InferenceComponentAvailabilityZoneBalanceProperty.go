package awssagemaker


// Configuration for balancing inference component copies across Availability Zones.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceComponentAvailabilityZoneBalanceProperty := &InferenceComponentAvailabilityZoneBalanceProperty{
//   	EnforcementMode: jsii.String("enforcementMode"),
//   	MaxImbalance: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentavailabilityzonebalance.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentAvailabilityZoneBalanceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentavailabilityzonebalance.html#cfn-sagemaker-inferencecomponent-inferencecomponentavailabilityzonebalance-enforcementmode
	//
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
	// The maximum allowed difference in the number of inference component copies between any two Availability Zones.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentavailabilityzonebalance.html#cfn-sagemaker-inferencecomponent-inferencecomponentavailabilityzonebalance-maximbalance
	//
	MaxImbalance *float64 `field:"optional" json:"maxImbalance" yaml:"maxImbalance"`
}

