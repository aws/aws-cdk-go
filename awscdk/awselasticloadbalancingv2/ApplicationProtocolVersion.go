package awselasticloadbalancingv2


// Load balancing protocol version for application load balancers.
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
type ApplicationProtocolVersion string

const (
	// GRPC.
	ApplicationProtocolVersion_GRPC ApplicationProtocolVersion = "GRPC"
	// HTTP1.
	ApplicationProtocolVersion_HTTP1 ApplicationProtocolVersion = "HTTP1"
	// HTTP2.
	ApplicationProtocolVersion_HTTP2 ApplicationProtocolVersion = "HTTP2"
)

