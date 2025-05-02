package awscodebuild


// Defines when a new instance is auto-scaled into the compute fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetTrackingScalingConfigurationProperty := &TargetTrackingScalingConfigurationProperty{
//   	MetricType: jsii.String("metricType"),
//   	TargetValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-targettrackingscalingconfiguration.html
//
type CfnFleet_TargetTrackingScalingConfigurationProperty struct {
	// The metric type to determine auto-scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-targettrackingscalingconfiguration.html#cfn-codebuild-fleet-targettrackingscalingconfiguration-metrictype
	//
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// The value of `metricType` when to start scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-targettrackingscalingconfiguration.html#cfn-codebuild-fleet-targettrackingscalingconfiguration-targetvalue
	//
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

