package awselasticloadbalancingv2


// Properties to reference an existing target group.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var stack stack
//
//
//   targetGroup := elbv2.applicationTargetGroup.fromTargetGroupAttributes(stack, jsii.String("MyTargetGroup"), &targetGroupAttributes{
//   	targetGroupArn: fn_ImportValue(jsii.String("TargetGroupArn")),
//   	loadBalancerArns: *fn_*ImportValue(jsii.String("LoadBalancerArn")),
//   })
//
//   targetGroupMetrics := targetGroup.metrics
//
type TargetGroupAttributes struct {
	// ARN of the target group.
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
	// A Token representing the list of ARNs for the load balancer routing to this target group.
	LoadBalancerArns *string `field:"optional" json:"loadBalancerArns" yaml:"loadBalancerArns"`
}

