package awselasticloadbalancingv2


// Properties to reference an existing target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupImportProps := &targetGroupImportProps{
//   	targetGroupArn: jsii.String("targetGroupArn"),
//
//   	// the properties below are optional
//   	defaultPort: jsii.String("defaultPort"),
//   	loadBalancerArns: jsii.String("loadBalancerArns"),
//   }
//
// Deprecated: Use TargetGroupAttributes instead.
type TargetGroupImportProps struct {
	// ARN of the target group.
	// Deprecated: Use TargetGroupAttributes instead.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
	// Port target group is listening on.
	// Deprecated: - This property is unused and the wrong type. No need to use it.
	DefaultPort *string `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// A Token representing the list of ARNs for the load balancer routing to this target group.
	// Deprecated: Use TargetGroupAttributes instead.
	LoadBalancerArns *string `field:"optional" json:"loadBalancerArns" yaml:"loadBalancerArns"`
}

