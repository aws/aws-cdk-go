package awscloudfront


// The HTTP version(s) to enable on the distribution.
//
// Note: Setting `HTTP3` enables HTTP 3 only (not HTTP 2). To support both HTTP 2 and HTTP 3, use `HTTP2_AND_3`.
//
// Example:
//   // Configure a distribution to use HTTP/2 and HTTP/3
//   // Configure a distribution to use HTTP/2 and HTTP/3
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   	},
//   	HttpVersion: cloudfront.HttpVersion_HTTP2_AND_3,
//   })
//
// See: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html#cloudfront-UpdateDistribution-request-HttpVersion
//
type HttpVersion string

const (
	// HTTP 1.1.
	HttpVersion_HTTP1_1 HttpVersion = "HTTP1_1"
	// HTTP 2.
	HttpVersion_HTTP2 HttpVersion = "HTTP2"
	// HTTP 2 and HTTP 3.
	HttpVersion_HTTP2_AND_3 HttpVersion = "HTTP2_AND_3"
	// HTTP 3 only (does not include HTTP 2).
	HttpVersion_HTTP3 HttpVersion = "HTTP3"
)

