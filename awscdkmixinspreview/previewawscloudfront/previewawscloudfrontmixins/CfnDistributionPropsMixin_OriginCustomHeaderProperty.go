package previewawscloudfrontmixins


// A complex type that contains `HeaderName` and `HeaderValue` elements, if any, for this distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   originCustomHeaderProperty := &OriginCustomHeaderProperty{
//   	HeaderName: jsii.String("headerName"),
//   	HeaderValue: jsii.String("headerValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origincustomheader.html
//
type CfnDistributionPropsMixin_OriginCustomHeaderProperty struct {
	// The name of a header that you want CloudFront to send to your origin.
	//
	// For more information, see [Adding Custom Headers to Origin Requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/forward-custom-headers.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origincustomheader.html#cfn-cloudfront-distribution-origincustomheader-headername
	//
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The value for the header that you specified in the `HeaderName` field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origincustomheader.html#cfn-cloudfront-distribution-origincustomheader-headervalue
	//
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
}

