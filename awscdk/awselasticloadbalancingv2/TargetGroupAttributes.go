package awselasticloadbalancingv2


// Properties to reference an existing target group.
//
// Example:
//   var stack stack
//
//
//   targetGroup := elbv2.ApplicationTargetGroup_FromTargetGroupAttributes(this, jsii.String("MyTargetGroup"), &TargetGroupAttributes{
//   	TargetGroupArn: awscdk.Fn_ImportValue(jsii.String("TargetGroupArn")),
//   	LoadBalancerArns: awscdk.Fn_*ImportValue(jsii.String("LoadBalancerArn")),
//   })
//
//   targetGroupMetrics := targetGroup.Metrics
//
type TargetGroupAttributes struct {
	// ARN of the target group.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
	// A Token representing the list of ARNs for the load balancer routing to this target group.
	LoadBalancerArns *string `field:"optional" json:"loadBalancerArns" yaml:"loadBalancerArns"`
}

