package awsautoscaling


// A structure that specifies a tag for the `Tags` property of [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html) resource.
//
// For more information, see [Tag Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide* . You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#aws-properties-as-group--examples) section of the `AWS::AutoScaling::AutoScalingGroup` resource.
//
// CloudFormation adds the following tags to all Auto Scaling groups and associated instances:
//
// - aws:cloudformation:stack-name
// - aws:cloudformation:stack-id
// - aws:cloudformation:logical-id.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagPropertyProperty := &tagPropertyProperty{
//   	key: jsii.String("key"),
//   	propagateAtLaunch: jsii.Boolean(false),
//   	value: jsii.String("value"),
//   }
//
type CfnAutoScalingGroup_TagPropertyProperty struct {
	// The tag key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Set to `true` if you want CloudFormation to copy the tag to EC2 instances that are launched as part of the Auto Scaling group.
	//
	// Set to `false` if you want the tag attached only to the Auto Scaling group and not copied to any instances launched as part of the Auto Scaling group.
	PropagateAtLaunch interface{} `field:"required" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// The tag value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

