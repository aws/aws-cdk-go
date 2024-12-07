package awselasticloadbalancingv2


// How to interpret the load balancing target identifiers.
//
// Example:
//   var vpc vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &ApplicationTargetGroupProps{
//   	TargetType: elbv2.TargetType_IP,
//   	Port: jsii.Number(50051),
//   	Protocol: elbv2.ApplicationProtocol_HTTP,
//   	ProtocolVersion: elbv2.ApplicationProtocolVersion_GRPC,
//   	HealthCheck: &HealthCheck{
//   		Enabled: jsii.Boolean(true),
//   		HealthyGrpcCodes: jsii.String("0-99"),
//   	},
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

