package awscloudfront


// Determines whether CloudFront includes the `X-Content-Type-Options` HTTP response header with its value set to `nosniff` .
//
// For more information about the `X-Content-Type-Options` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentTypeOptionsProperty := &contentTypeOptionsProperty{
//   	override: jsii.Boolean(false),
//   }
//
type CfnResponseHeadersPolicy_ContentTypeOptionsProperty struct {
	// A Boolean that determines whether CloudFront overrides the `X-Content-Type-Options` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

