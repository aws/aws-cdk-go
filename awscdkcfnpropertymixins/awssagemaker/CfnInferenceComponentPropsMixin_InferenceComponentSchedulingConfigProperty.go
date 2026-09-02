package awssagemaker


// The scheduling configuration that determines how inference component copies are placed across available instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceComponentSchedulingConfigProperty := &InferenceComponentSchedulingConfigProperty{
//   	AvailabilityZoneBalance: &InferenceComponentAvailabilityZoneBalanceProperty{
//   		EnforcementMode: jsii.String("enforcementMode"),
//   		MaxImbalance: jsii.Number(123),
//   	},
//   	PlacementStrategy: jsii.String("placementStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentschedulingconfig.html
//
type CfnInferenceComponentPropsMixin_InferenceComponentSchedulingConfigProperty struct {
	// Configuration for balancing inference component copies across Availability Zones.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentschedulingconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentschedulingconfig-availabilityzonebalance
	//
	AvailabilityZoneBalance interface{} `field:"optional" json:"availabilityZoneBalance" yaml:"availabilityZoneBalance"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-inferencecomponentschedulingconfig.html#cfn-sagemaker-inferencecomponent-inferencecomponentschedulingconfig-placementstrategy
	//
	PlacementStrategy *string `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
}

