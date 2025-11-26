package previewawscloudfrontmixins


// Determines whether CloudFront includes the `X-XSS-Protection` HTTP response header and the header's value.
//
// For more information about the `X-XSS-Protection` HTTP response header, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   xSSProtectionProperty := &XSSProtectionProperty{
//   	ModeBlock: jsii.Boolean(false),
//   	Override: jsii.Boolean(false),
//   	Protection: jsii.Boolean(false),
//   	ReportUri: jsii.String("reportUri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html
//
type CfnResponseHeadersPolicyPropsMixin_XSSProtectionProperty struct {
	// A Boolean that determines whether CloudFront includes the `mode=block` directive in the `X-XSS-Protection` header.
	//
	// For more information about this directive, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-modeblock
	//
	ModeBlock interface{} `field:"optional" json:"modeBlock" yaml:"modeBlock"`
	// A Boolean that determines whether CloudFront overrides the `X-XSS-Protection` HTTP response header received from the origin with the one specified in this response headers policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-override
	//
	Override interface{} `field:"optional" json:"override" yaml:"override"`
	// A Boolean that determines the value of the `X-XSS-Protection` HTTP response header.
	//
	// When this setting is `true` , the value of the `X-XSS-Protection` header is `1` . When this setting is `false` , the value of the `X-XSS-Protection` header is `0` .
	//
	// For more information about these settings, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-protection
	//
	Protection interface{} `field:"optional" json:"protection" yaml:"protection"`
	// A reporting URI, which CloudFront uses as the value of the `report` directive in the `X-XSS-Protection` header.
	//
	// You cannot specify a `ReportUri` when `ModeBlock` is `true` .
	//
	// For more information about using a reporting URL, see [X-XSS-Protection](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-xssprotection.html#cfn-cloudfront-responseheaderspolicy-xssprotection-reporturi
	//
	ReportUri *string `field:"optional" json:"reportUri" yaml:"reportUri"`
}

