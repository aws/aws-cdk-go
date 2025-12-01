package awslambda


// The scaling configuration for the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderScalingConfigProperty := &CapacityProviderScalingConfigProperty{
//   	MaxVCpuCount: jsii.Number(123),
//   	ScalingMode: jsii.String("scalingMode"),
//   	ScalingPolicies: []interface{}{
//   		&TargetTrackingScalingPolicyProperty{
//   			PredefinedMetricType: jsii.String("predefinedMetricType"),
//   			TargetValue: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderscalingconfig.html
//
type CfnCapacityProvider_CapacityProviderScalingConfigProperty struct {
	// The maximum number of EC2 instances that the capacity provider can scale up to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderscalingconfig.html#cfn-lambda-capacityprovider-capacityproviderscalingconfig-maxvcpucount
	//
	MaxVCpuCount *float64 `field:"optional" json:"maxVCpuCount" yaml:"maxVCpuCount"`
	// The scaling mode for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderscalingconfig.html#cfn-lambda-capacityprovider-capacityproviderscalingconfig-scalingmode
	//
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
	// A list of target tracking scaling policies for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-capacityproviderscalingconfig.html#cfn-lambda-capacityprovider-capacityproviderscalingconfig-scalingpolicies
	//
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
}

