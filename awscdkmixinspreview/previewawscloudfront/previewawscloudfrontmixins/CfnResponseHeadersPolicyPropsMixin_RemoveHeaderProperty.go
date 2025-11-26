package previewawscloudfrontmixins


// The name of an HTTP header that CloudFront removes from HTTP responses to requests that match the cache behavior that this response headers policy is attached to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   removeHeaderProperty := &RemoveHeaderProperty{
//   	Header: jsii.String("header"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-removeheader.html
//
type CfnResponseHeadersPolicyPropsMixin_RemoveHeaderProperty struct {
	// The HTTP header name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-removeheader.html#cfn-cloudfront-responseheaderspolicy-removeheader-header
	//
	Header *string `field:"optional" json:"header" yaml:"header"`
}

