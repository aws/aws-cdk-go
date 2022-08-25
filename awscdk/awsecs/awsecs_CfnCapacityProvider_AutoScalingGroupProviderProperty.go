package awsecs


// The `AutoScalingGroupProvider` property specifies the details of the Auto Scaling group for the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingGroupProviderProperty := &autoScalingGroupProviderProperty{
//   	autoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   	// the properties below are optional
//   	managedScaling: &managedScalingProperty{
//   		instanceWarmupPeriod: jsii.Number(123),
//   		maximumScalingStepSize: jsii.Number(123),
//   		minimumScalingStepSize: jsii.Number(123),
//   		status: jsii.String("status"),
//   		targetCapacity: jsii.Number(123),
//   	},
//   	managedTerminationProtection: jsii.String("managedTerminationProtection"),
//   }
//
type CfnCapacityProvider_AutoScalingGroupProviderProperty struct {
	// The Amazon Resource Name (ARN) or short name that identifies the Auto Scaling group.
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// The managed scaling settings for the Auto Scaling group capacity provider.
	ManagedScaling interface{} `field:"optional" json:"managedScaling" yaml:"managedScaling"`
	// The managed termination protection setting to use for the Auto Scaling group capacity provider.
	//
	// This determines whether the Auto Scaling group has managed termination protection. The default is disabled.
	//
	// > When using managed termination protection, managed scaling must also be used otherwise managed termination protection doesn't work.
	//
	// When managed termination protection is enabled, Amazon ECS prevents the Amazon EC2 instances in an Auto Scaling group that contain tasks from being terminated during a scale-in action. The Auto Scaling group and each instance in the Auto Scaling group must have instance protection from scale-in actions enabled as well. For more information, see [Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the *AWS Auto Scaling User Guide* .
	//
	// When managed termination protection is disabled, your Amazon EC2 instances aren't protected from termination when the Auto Scaling group scales in.
	ManagedTerminationProtection *string `field:"optional" json:"managedTerminationProtection" yaml:"managedTerminationProtection"`
}

