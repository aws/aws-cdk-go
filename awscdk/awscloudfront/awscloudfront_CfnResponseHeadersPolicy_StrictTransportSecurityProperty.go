package awscloudfront


// Determines whether CloudFront includes the `Strict-Transport-Security` HTTP response header and the header’s value.
//
// For more information about the `Strict-Transport-Security` HTTP response header, see [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   strictTransportSecurityProperty := &strictTransportSecurityProperty{
//   	accessControlMaxAgeSec: jsii.Number(123),
//   	override: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	includeSubdomains: jsii.Boolean(false),
//   	preload: jsii.Boolean(false),
//   }
//
type CfnResponseHeadersPolicy_StrictTransportSecurityProperty struct {
	// A number that CloudFront uses as the value for the `max-age` directive in the `Strict-Transport-Security` HTTP response header.
	AccessControlMaxAgeSec *float64 `field:"required" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
	// A Boolean that determines whether CloudFront overrides the `Strict-Transport-Security` HTTP response header received from the origin with the one specified in this response headers policy.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// A Boolean that determines whether CloudFront includes the `includeSubDomains` directive in the `Strict-Transport-Security` HTTP response header.
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// A Boolean that determines whether CloudFront includes the `preload` directive in the `Strict-Transport-Security` HTTP response header.
	Preload interface{} `field:"optional" json:"preload" yaml:"preload"`
}

