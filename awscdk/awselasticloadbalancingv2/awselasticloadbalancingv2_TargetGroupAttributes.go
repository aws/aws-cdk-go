package awselasticloadbalancingv2


// Properties to reference an existing target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupAttributes := &targetGroupAttributes{
//   	targetGroupArn: jsii.String("targetGroupArn"),
//
//   	// the properties below are optional
//   	loadBalancerArns: jsii.String("loadBalancerArns"),
//   }
//
type TargetGroupAttributes struct {
	// ARN of the target group.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
	// A Token representing the list of ARNs for the load balancer routing to this target group.
	LoadBalancerArns *string `field:"optional" json:"loadBalancerArns" yaml:"loadBalancerArns"`
}

