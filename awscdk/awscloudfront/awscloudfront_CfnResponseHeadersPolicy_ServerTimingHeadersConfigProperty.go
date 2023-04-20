package awscloudfront


// A configuration for enabling the `Server-Timing` header in HTTP responses sent from CloudFront.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverTimingHeadersConfigProperty := &ServerTimingHeadersConfigProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	SamplingRate: jsii.Number(123),
//   }
//
type CfnResponseHeadersPolicy_ServerTimingHeadersConfigProperty struct {
	// A Boolean that determines whether CloudFront adds the `Server-Timing` header to HTTP responses that it sends in response to requests that match a cache behavior that's associated with this response headers policy.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A number 0–100 (inclusive) that specifies the percentage of responses that you want CloudFront to add the `Server-Timing` header to.
	//
	// When you set the sampling rate to 100, CloudFront adds the `Server-Timing` header to the HTTP response for every request that matches the cache behavior that this response headers policy is attached to. When you set it to 50, CloudFront adds the header to 50% of the responses for requests that match the cache behavior. You can set the sampling rate to any number 0–100 with up to four decimal places.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

