package awsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaEventSourceMappingUngracefulProperty := &LambdaEventSourceMappingUngracefulProperty{
//   	Behavior: jsii.String("behavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingungraceful.html
//
type CfnPlan_LambdaEventSourceMappingUngracefulProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-lambdaeventsourcemappingungraceful.html#cfn-arcregionswitch-plan-lambdaeventsourcemappingungraceful-behavior
	//
	// Default: - "skip".
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
}

