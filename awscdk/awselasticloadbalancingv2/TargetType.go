package awselasticloadbalancingv2


// How to interpret the load balancing target identifiers.
//
// Example:
//   var vpc vpc
//
//
//   // Target group with slow start mode enabled
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &ApplicationTargetGroupProps{
//   	TargetType: elbv2.TargetType_INSTANCE,
//   	SlowStart: awscdk.Duration_Seconds(jsii.Number(60)),
//   	Port: jsii.Number(80),
//   	Vpc: Vpc,
//   })
//
type TargetType string

const (
	// Targets identified by instance ID.
	TargetType_INSTANCE TargetType = "INSTANCE"
	// Targets identified by IP address.
	TargetType_IP TargetType = "IP"
	// Target is a single Lambda Function.
	TargetType_LAMBDA TargetType = "LAMBDA"
	// Target is a single Application Load Balancer.
	TargetType_ALB TargetType = "ALB"
)

