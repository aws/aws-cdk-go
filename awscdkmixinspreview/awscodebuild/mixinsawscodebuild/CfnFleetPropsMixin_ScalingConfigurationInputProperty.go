package mixinsawscodebuild


// The scaling configuration input of a compute fleet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingConfigurationInputProperty := &ScalingConfigurationInputProperty{
//   	MaxCapacity: jsii.Number(123),
//   	ScalingType: jsii.String("scalingType"),
//   	TargetTrackingScalingConfigs: []interface{}{
//   		&TargetTrackingScalingConfigurationProperty{
//   			MetricType: jsii.String("metricType"),
//   			TargetValue: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-scalingconfigurationinput.html
//
type CfnFleetPropsMixin_ScalingConfigurationInputProperty struct {
	// The maximum number of instances in the ﬂeet when auto-scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-scalingconfigurationinput.html#cfn-codebuild-fleet-scalingconfigurationinput-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The scaling type for a compute fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-scalingconfigurationinput.html#cfn-codebuild-fleet-scalingconfigurationinput-scalingtype
	//
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
	// A list of `TargetTrackingScalingConfiguration` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-fleet-scalingconfigurationinput.html#cfn-codebuild-fleet-scalingconfigurationinput-targettrackingscalingconfigs
	//
	TargetTrackingScalingConfigs interface{} `field:"optional" json:"targetTrackingScalingConfigs" yaml:"targetTrackingScalingConfigs"`
}

