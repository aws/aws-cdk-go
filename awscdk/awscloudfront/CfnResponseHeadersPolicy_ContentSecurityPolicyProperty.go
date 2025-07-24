package awscloudfront


// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
//
// For more information about the `Content-Security-Policy` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentSecurityPolicyProperty := &ContentSecurityPolicyProperty{
//   	ContentSecurityPolicy: jsii.String("contentSecurityPolicy"),
//   	Override: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-contentsecuritypolicy.html
//
type CfnResponseHeadersPolicy_ContentSecurityPolicyProperty struct {
	// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-contentsecuritypolicy.html#cfn-cloudfront-responseheaderspolicy-contentsecuritypolicy-contentsecuritypolicy
	//
	ContentSecurityPolicy *string `field:"required" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// A Boolean that determines whether CloudFront overrides the `Content-Security-Policy` HTTP response header received from the origin with the one specified in this response headers policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-contentsecuritypolicy.html#cfn-cloudfront-responseheaderspolicy-contentsecuritypolicy-override
	//
	Override interface{} `field:"required" json:"override" yaml:"override"`
}

