package awslambda


// A target tracking scaling policy for the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingScalingPolicyProperty := &TargetTrackingScalingPolicyProperty{
//   	PredefinedMetricType: jsii.String("predefinedMetricType"),
//   	TargetValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.html
//
type CfnCapacityProvider_TargetTrackingScalingPolicyProperty struct {
	// The predefined metric for target tracking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.html#cfn-lambda-capacityprovider-targettrackingscalingpolicy-predefinedmetrictype
	//
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// The target value for the metric as a percentage (for example, 70.0 for 70%).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.html#cfn-lambda-capacityprovider-targettrackingscalingpolicy-targetvalue
	//
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}

