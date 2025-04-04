package awsecs


// The details of the Auto Scaling group for the capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingGroupProviderProperty := &AutoScalingGroupProviderProperty{
//   	AutoScalingGroupArn: jsii.String("autoScalingGroupArn"),
//
//   	// the properties below are optional
//   	ManagedDraining: jsii.String("managedDraining"),
//   	ManagedScaling: &ManagedScalingProperty{
//   		InstanceWarmupPeriod: jsii.Number(123),
//   		MaximumScalingStepSize: jsii.Number(123),
//   		MinimumScalingStepSize: jsii.Number(123),
//   		Status: jsii.String("status"),
//   		TargetCapacity: jsii.Number(123),
//   	},
//   	ManagedTerminationProtection: jsii.String("managedTerminationProtection"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html
//
type CfnCapacityProvider_AutoScalingGroupProviderProperty struct {
	// The Amazon Resource Name (ARN) that identifies the Auto Scaling group, or the Auto Scaling group name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-autoscalinggrouparn
	//
	AutoScalingGroupArn *string `field:"required" json:"autoScalingGroupArn" yaml:"autoScalingGroupArn"`
	// The managed draining option for the Auto Scaling group capacity provider.
	//
	// When you enable this, Amazon ECS manages and gracefully drains the EC2 container instances that are in the Auto Scaling group capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-manageddraining
	//
	ManagedDraining *string `field:"optional" json:"managedDraining" yaml:"managedDraining"`
	// The managed scaling settings for the Auto Scaling group capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-managedscaling
	//
	ManagedScaling interface{} `field:"optional" json:"managedScaling" yaml:"managedScaling"`
	// The managed termination protection setting to use for the Auto Scaling group capacity provider.
	//
	// This determines whether the Auto Scaling group has managed termination protection. The default is off.
	//
	// > When using managed termination protection, managed scaling must also be used otherwise managed termination protection doesn't work.
	//
	// When managed termination protection is on, Amazon ECS prevents the Amazon EC2 instances in an Auto Scaling group that contain tasks from being terminated during a scale-in action. The Auto Scaling group and each instance in the Auto Scaling group must have instance protection from scale-in actions on as well. For more information, see [Instance Protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html#instance-protection) in the *AWS Auto Scaling User Guide* .
	//
	// When managed termination protection is off, your Amazon EC2 instances aren't protected from termination when the Auto Scaling group scales in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autoscalinggroupprovider.html#cfn-ecs-capacityprovider-autoscalinggroupprovider-managedterminationprotection
	//
	ManagedTerminationProtection *string `field:"optional" json:"managedTerminationProtection" yaml:"managedTerminationProtection"`
}

