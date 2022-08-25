// Version 2 of the AWS Cloud Development Kit library
package awscdk


// With scheduled actions, the group size properties of an Auto Scaling group can change at any time.
//
// When you update a
// stack with an Auto Scaling group and scheduled action, AWS CloudFormation always sets the group size property values of
// your Auto Scaling group to the values that are defined in the AWS::AutoScaling::AutoScalingGroup resource of your template,
// even if a scheduled action is in effect.
//
// If you do not want AWS CloudFormation to change any of the group size property values when you have a scheduled action in
// effect, use the AutoScalingScheduledAction update policy to prevent AWS CloudFormation from changing the MinSize, MaxSize,
// or DesiredCapacity properties unless you have modified these values in your template.\
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingScheduledAction := &cfnAutoScalingScheduledAction{
//   	ignoreUnmodifiedGroupSizeProperties: jsii.Boolean(false),
//   }
//
type CfnAutoScalingScheduledAction struct {
	IgnoreUnmodifiedGroupSizeProperties *bool `field:"optional" json:"ignoreUnmodifiedGroupSizeProperties" yaml:"ignoreUnmodifiedGroupSizeProperties"`
}

