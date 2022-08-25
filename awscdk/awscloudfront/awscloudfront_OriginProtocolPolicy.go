package awscloudfront


// Defines what protocols CloudFront will use to connect to an origin.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var loadBalancer applicationLoadBalancer
//
//   origin := origins.NewLoadBalancerV2Origin(loadBalancer, &loadBalancerV2OriginProps{
//   	connectionAttempts: jsii.Number(3),
//   	connectionTimeout: awscdk.Duration.seconds(jsii.Number(5)),
//   	readTimeout: awscdk.Duration.seconds(jsii.Number(45)),
//   	keepaliveTimeout: awscdk.Duration.seconds(jsii.Number(45)),
//   	protocolPolicy: cloudfront.originProtocolPolicy_MATCH_VIEWER,
//   })
//
type OriginProtocolPolicy string

const (
	// Connect on HTTP only.
	OriginProtocolPolicy_HTTP_ONLY OriginProtocolPolicy = "HTTP_ONLY"
	// Connect with the same protocol as the viewer.
	OriginProtocolPolicy_MATCH_VIEWER OriginProtocolPolicy = "MATCH_VIEWER"
	// Connect on HTTPS only.
	OriginProtocolPolicy_HTTPS_ONLY OriginProtocolPolicy = "HTTPS_ONLY"
)

