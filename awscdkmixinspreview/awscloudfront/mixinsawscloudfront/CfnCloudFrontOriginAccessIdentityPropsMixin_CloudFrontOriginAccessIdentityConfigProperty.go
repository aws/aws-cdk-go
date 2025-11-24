package mixinsawscloudfront


// Origin access identity configuration.
//
// Send a `GET` request to the `/ *CloudFront API version* /CloudFront/identity ID/config` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudFrontOriginAccessIdentityConfigProperty := &CloudFrontOriginAccessIdentityConfigProperty{
//   	Comment: jsii.String("comment"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig.html
//
type CfnCloudFrontOriginAccessIdentityPropsMixin_CloudFrontOriginAccessIdentityConfigProperty struct {
	// A comment to describe the origin access identity.
	//
	// The comment cannot be longer than 128 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig-comment
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

