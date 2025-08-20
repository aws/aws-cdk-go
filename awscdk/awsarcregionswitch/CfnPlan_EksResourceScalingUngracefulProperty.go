package awsarcregionswitch


// The ungraceful settings for AWS EKS resource scaling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksResourceScalingUngracefulProperty := &EksResourceScalingUngracefulProperty{
//   	MinimumSuccessPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingungraceful.html
//
type CfnPlan_EksResourceScalingUngracefulProperty struct {
	// The minimum success percentage for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-eksresourcescalingungraceful.html#cfn-arcregionswitch-plan-eksresourcescalingungraceful-minimumsuccesspercentage
	//
	MinimumSuccessPercentage *float64 `field:"required" json:"minimumSuccessPercentage" yaml:"minimumSuccessPercentage"`
}

