package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingArnsProperty := &AutoScalingArnsProperty{
//   	ApplicationAutoScalingPolicies: []*string{
//   		jsii.String("applicationAutoScalingPolicies"),
//   	},
//   	ScalableTarget: jsii.String("scalableTarget"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-autoscalingarns.html
//
type CfnExpressGatewayService_AutoScalingArnsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-autoscalingarns.html#cfn-ecs-expressgatewayservice-autoscalingarns-applicationautoscalingpolicies
	//
	ApplicationAutoScalingPolicies *[]*string `field:"optional" json:"applicationAutoScalingPolicies" yaml:"applicationAutoScalingPolicies"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-autoscalingarns.html#cfn-ecs-expressgatewayservice-autoscalingarns-scalabletarget
	//
	ScalableTarget *string `field:"optional" json:"scalableTarget" yaml:"scalableTarget"`
}

