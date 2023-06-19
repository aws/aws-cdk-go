package awscloudfront


// Defines what protocols CloudFront will use to connect to an origin.
//
// Example:
//   import elbv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var loadBalancer applicationLoadBalancer
//
//   origin := origins.NewLoadBalancerV2Origin(loadBalancer, &LoadBalancerV2OriginProps{
//   	ConnectionAttempts: jsii.Number(3),
//   	ConnectionTimeout: awscdk.Duration_Seconds(jsii.Number(5)),
//   	ReadTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
//   	KeepaliveTimeout: awscdk.Duration_*Seconds(jsii.Number(45)),
//   	ProtocolPolicy: cloudfront.OriginProtocolPolicy_MATCH_VIEWER,
//   })
//
// Experimental.
type OriginProtocolPolicy string

const (
	// Connect on HTTP only.
	// Experimental.
	OriginProtocolPolicy_HTTP_ONLY OriginProtocolPolicy = "HTTP_ONLY"
	// Connect with the same protocol as the viewer.
	// Experimental.
	OriginProtocolPolicy_MATCH_VIEWER OriginProtocolPolicy = "MATCH_VIEWER"
	// Connect on HTTPS only.
	// Experimental.
	OriginProtocolPolicy_HTTPS_ONLY OriginProtocolPolicy = "HTTPS_ONLY"
)

