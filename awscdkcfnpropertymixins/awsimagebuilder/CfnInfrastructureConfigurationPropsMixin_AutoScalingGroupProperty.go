package awsimagebuilder


// Identifies the Amazon EC2 Auto Scaling group that Image Builder takes the build instance from, and how long Image Builder waits when the group has no instance available.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   autoScalingGroupProperty := &AutoScalingGroupProperty{
//   	CapacityWaitTimeoutSeconds: jsii.Number(123),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-autoscalinggroup.html
//
type CfnInfrastructureConfigurationPropsMixin_AutoScalingGroupProperty struct {
	// The number of seconds that Image Builder waits for an instance to become available in the group before the build fails.
	//
	// If you set this value to 0 or omit it, the build fails as soon as the group has no instance available.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-autoscalinggroup.html#cfn-imagebuilder-infrastructureconfiguration-autoscalinggroup-capacitywaittimeoutseconds
	//
	CapacityWaitTimeoutSeconds *float64 `field:"optional" json:"capacityWaitTimeoutSeconds" yaml:"capacityWaitTimeoutSeconds"`
	// The name of the Auto Scaling group, in the same account and AWS Region as the infrastructure configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-infrastructureconfiguration-autoscalinggroup.html#cfn-imagebuilder-infrastructureconfiguration-autoscalinggroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

