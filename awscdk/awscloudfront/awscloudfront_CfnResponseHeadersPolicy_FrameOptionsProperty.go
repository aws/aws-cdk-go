package awscloudfront


// Determines whether CloudFront includes the `X-Frame-Options` HTTP response header and the header’s value.
//
// For more information about the `X-Frame-Options` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameOptionsProperty := &frameOptionsProperty{
//   	frameOption: jsii.String("frameOption"),
//   	override: jsii.Boolean(false),
//   }
//
type CfnResponseHeadersPolicy_FrameOptionsProperty struct {
	// The value of the `X-Frame-Options` HTTP response header. Valid values are `DENY` and `SAMEORIGIN` .
	//
	// For more information about these values, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
	FrameOption *string `field:"required" json:"frameOption" yaml:"frameOption"`
	// A Boolean that determines whether CloudFront overrides the `X-Frame-Options` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

