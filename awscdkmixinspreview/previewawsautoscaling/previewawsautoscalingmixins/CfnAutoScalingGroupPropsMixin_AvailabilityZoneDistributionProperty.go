package previewawsautoscalingmixins


// `AvailabilityZoneDistribution` is a property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   availabilityZoneDistributionProperty := &AvailabilityZoneDistributionProperty{
//   	CapacityDistributionStrategy: jsii.String("capacityDistributionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzonedistribution.html
//
type CfnAutoScalingGroupPropsMixin_AvailabilityZoneDistributionProperty struct {
	// If launches fail in an Availability Zone, the following strategies are available. The default is `balanced-best-effort` .
	//
	// - `balanced-only` - If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.
	// - `balanced-best-effort` - If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzonedistribution.html#cfn-autoscaling-autoscalinggroup-availabilityzonedistribution-capacitydistributionstrategy
	//
	CapacityDistributionStrategy *string `field:"optional" json:"capacityDistributionStrategy" yaml:"capacityDistributionStrategy"`
}

