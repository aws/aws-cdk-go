package previewawslambdamixins


// A scaling policy for the capacity provider that automatically adjusts capacity to maintain a target value for a specific metric.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetTrackingScalingPolicyProperty := &TargetTrackingScalingPolicyProperty{
//   	PredefinedMetricType: jsii.String("predefinedMetricType"),
//   	TargetValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.html
//
type CfnCapacityProviderPropsMixin_TargetTrackingScalingPolicyProperty struct {
	// The predefined metric type to track for scaling decisions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.html#cfn-lambda-capacityprovider-targettrackingscalingpolicy-predefinedmetrictype
	//
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// The target value for the metric that the scaling policy attempts to maintain through scaling actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-capacityprovider-targettrackingscalingpolicy.html#cfn-lambda-capacityprovider-targettrackingscalingpolicy-targetvalue
	//
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

