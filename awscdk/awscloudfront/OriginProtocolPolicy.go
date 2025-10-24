package awscloudfront


// Defines what protocols CloudFront will use to connect to an origin.
//
// Example:
//   var loadBalancer ApplicationLoadBalancer
//
//   origin := origins.NewLoadBalancerV2Origin(loadBalancer, &LoadBalancerV2OriginProps{
//   	ConnectionAttempts: jsii.Number(3),
//   	ConnectionTimeout: awscdk.Duration_Seconds(jsii.Number(5)),
//   	ReadTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
//   	ResponseCompletionTimeout: awscdk.Duration_*Seconds(jsii.Number(120)),
//   	KeepaliveTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
//   	ProtocolPolicy: cloudfront.OriginProtocolPolicy_MATCH_VIEWER,
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

