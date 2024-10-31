package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availabilityZoneDistributionProperty := &AvailabilityZoneDistributionProperty{
//   	CapacityDistributionStrategy: jsii.String("capacityDistributionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzonedistribution.html
//
type CfnAutoScalingGroup_AvailabilityZoneDistributionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-availabilityzonedistribution.html#cfn-autoscaling-autoscalinggroup-availabilityzonedistribution-capacitydistributionstrategy
	//
	CapacityDistributionStrategy *string `field:"optional" json:"capacityDistributionStrategy" yaml:"capacityDistributionStrategy"`
}

