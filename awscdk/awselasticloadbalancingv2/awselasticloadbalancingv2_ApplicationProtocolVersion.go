package awselasticloadbalancingv2


// Load balancing protocol version for application load balancers.
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
// Experimental.
type ApplicationProtocolVersion string

const (
	// GRPC.
	// Experimental.
	ApplicationProtocolVersion_GRPC ApplicationProtocolVersion = "GRPC"
	// HTTP1.
	// Experimental.
	ApplicationProtocolVersion_HTTP1 ApplicationProtocolVersion = "HTTP1"
	// HTTP2.
	// Experimental.
	ApplicationProtocolVersion_HTTP2 ApplicationProtocolVersion = "HTTP2"
)

