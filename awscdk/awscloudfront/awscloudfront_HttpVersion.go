package awscloudfront


// Maximum HTTP version to support.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Configure a distribution to use HTTP/2 and HTTP/3
//   // Configure a distribution to use HTTP/2 and HTTP/3
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewHttpOrigin(jsii.String("www.example.com")),
//   	},
//   	httpVersion: cloudfront.httpVersion_HTTP2_AND_3,
//   })
//
type HttpVersion string

const (
	// HTTP 1.1.
	HttpVersion_HTTP1_1 HttpVersion = "HTTP1_1"
	// HTTP 2.
	HttpVersion_HTTP2 HttpVersion = "HTTP2"
	// HTTP 2 and HTTP 3.
	HttpVersion_HTTP2_AND_3 HttpVersion = "HTTP2_AND_3"
	// HTTP 3.
	HttpVersion_HTTP3 HttpVersion = "HTTP3"
)

