package awscloudfront


// Determines whether CloudFront includes the `Referrer-Policy` HTTP response header and the header's value.
//
// For more information about the `Referrer-Policy` HTTP response header, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referrerPolicyProperty := &ReferrerPolicyProperty{
//   	Override: jsii.Boolean(false),
//   	ReferrerPolicy: jsii.String("referrerPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.html
//
type CfnResponseHeadersPolicy_ReferrerPolicyProperty struct {
	// A Boolean that determines whether CloudFront overrides the `Referrer-Policy` HTTP response header received from the origin with the one specified in this response headers policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.html#cfn-cloudfront-responseheaderspolicy-referrerpolicy-override
	//
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// The value of the `Referrer-Policy` HTTP response header. Valid values are:.
	//
	// - `no-referrer`
	// - `no-referrer-when-downgrade`
	// - `origin`
	// - `origin-when-cross-origin`
	// - `same-origin`
	// - `strict-origin`
	// - `strict-origin-when-cross-origin`
	// - `unsafe-url`
	//
	// For more information about these values, see [Referrer-Policy](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) in the MDN Web Docs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-referrerpolicy.html#cfn-cloudfront-responseheaderspolicy-referrerpolicy-referrerpolicy
	//
	ReferrerPolicy *string `field:"required" json:"referrerPolicy" yaml:"referrerPolicy"`
}

