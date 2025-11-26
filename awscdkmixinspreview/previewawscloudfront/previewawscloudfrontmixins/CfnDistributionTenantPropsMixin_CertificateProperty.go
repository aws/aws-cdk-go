package previewawscloudfrontmixins


// The Certificate Manager (ACM) certificate associated with your distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   certificateProperty := &CertificateProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-certificate.html
//
type CfnDistributionTenantPropsMixin_CertificateProperty struct {
	// The Amazon Resource Name (ARN) of the ACM certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributiontenant-certificate.html#cfn-cloudfront-distributiontenant-certificate-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

