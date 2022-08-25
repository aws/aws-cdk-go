package awselasticloadbalancingv2


// How to interpret the load balancing target identifiers.
//
// Example:
//   var vpc vpc
//
//
//   tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &applicationTargetGroupProps{
//   	targetType: elbv2.targetType_IP,
//   	port: jsii.Number(50051),
//   	protocol: elbv2.applicationProtocol_HTTP,
//   	protocolVersion: elbv2.applicationProtocolVersion_GRPC,
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(true),
//   		healthyGrpcCodes: jsii.String("0-99"),
//   	},
//   	vpc: vpc,
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

