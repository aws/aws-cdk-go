package awscloudfront


// A configuration for a set of security-related HTTP response headers.
//
// CloudFront adds these headers to HTTP responses that it sends for requests that match a cache behavior associated with this response headers policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityHeadersConfigProperty := &SecurityHeadersConfigProperty{
//   	ContentSecurityPolicy: &ContentSecurityPolicyProperty{
//   		ContentSecurityPolicy: jsii.String("contentSecurityPolicy"),
//   		Override: jsii.Boolean(false),
//   	},
//   	ContentTypeOptions: &ContentTypeOptionsProperty{
//   		Override: jsii.Boolean(false),
//   	},
//   	FrameOptions: &FrameOptionsProperty{
//   		FrameOption: jsii.String("frameOption"),
//   		Override: jsii.Boolean(false),
//   	},
//   	ReferrerPolicy: &ReferrerPolicyProperty{
//   		Override: jsii.Boolean(false),
//   		ReferrerPolicy: jsii.String("referrerPolicy"),
//   	},
//   	StrictTransportSecurity: &StrictTransportSecurityProperty{
//   		AccessControlMaxAgeSec: jsii.Number(123),
//   		Override: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		IncludeSubdomains: jsii.Boolean(false),
//   		Preload: jsii.Boolean(false),
//   	},
//   	XssProtection: &XSSProtectionProperty{
//   		Override: jsii.Boolean(false),
//   		Protection: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		ModeBlock: jsii.Boolean(false),
//   		ReportUri: jsii.String("reportUri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html
//
type CfnResponseHeadersPolicy_SecurityHeadersConfigProperty struct {
	// The policy directives and their values that CloudFront includes as values for the `Content-Security-Policy` HTTP response header.
	//
	// For more information about the `Content-Security-Policy` HTTP response header, see [Content-Security-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html#cfn-cloudfront-responseheaderspolicy-securityheadersconfig-contentsecuritypolicy
	//
	ContentSecurityPolicy interface{} `field:"optional" json:"contentSecurityPolicy" yaml:"contentSecurityPolicy"`
	// Determines whether CloudFront includes the `X-Content-Type-Options` HTTP response header with its value set to `nosniff` .
	//
	// For more information about the `X-Content-Type-Options` HTTP response header, see [X-Content-Type-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html#cfn-cloudfront-responseheaderspolicy-securityheadersconfig-contenttypeoptions
	//
	ContentTypeOptions interface{} `field:"optional" json:"contentTypeOptions" yaml:"contentTypeOptions"`
	// Determines whether CloudFront includes the `X-Frame-Options` HTTP response header and the header's value.
	//
	// For more information about the `X-Frame-Options` HTTP response header, see [X-Frame-Options](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html#cfn-cloudfront-responseheaderspolicy-securityheadersconfig-frameoptions
	//
	FrameOptions interface{} `field:"optional" json:"frameOptions" yaml:"frameOptions"`
	// Determines whether CloudFront includes the `Referrer-Policy` HTTP response header and the header's value.
	//
	// For more information about the `Referrer-Policy` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html#cfn-cloudfront-responseheaderspolicy-securityheadersconfig-referrerpolicy
	//
	ReferrerPolicy interface{} `field:"optional" json:"referrerPolicy" yaml:"referrerPolicy"`
	// Determines whether CloudFront includes the `Strict-Transport-Security` HTTP response header and the header's value.
	//
	// For more information about the `Strict-Transport-Security` HTTP response header, see [Strict-Transport-Security](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html#cfn-cloudfront-responseheaderspolicy-securityheadersconfig-stricttransportsecurity
	//
	StrictTransportSecurity interface{} `field:"optional" json:"strictTransportSecurity" yaml:"strictTransportSecurity"`
	// Determines whether CloudFront includes the `X-XSS-Protection` HTTP response header and the header's value.
	//
	// For more information about the `X-XSS-Protection` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-securityheadersconfig.html#cfn-cloudfront-responseheaderspolicy-securityheadersconfig-xssprotection
	//
	XssProtection interface{} `field:"optional" json:"xssProtection" yaml:"xssProtection"`
}

