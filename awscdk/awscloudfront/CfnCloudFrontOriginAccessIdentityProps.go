package awscloudfront


// Properties for defining a `CfnCloudFrontOriginAccessIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudFrontOriginAccessIdentityProps := &CfnCloudFrontOriginAccessIdentityProps{
//   	CloudFrontOriginAccessIdentityConfig: &CloudFrontOriginAccessIdentityConfigProperty{
//   		Comment: jsii.String("comment"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html
//
type CfnCloudFrontOriginAccessIdentityProps struct {
	// The current configuration information for the identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html#cfn-cloudfront-cloudfrontoriginaccessidentity-cloudfrontoriginaccessidentityconfig
	//
	CloudFrontOriginAccessIdentityConfig interface{} `field:"required" json:"cloudFrontOriginAccessIdentityConfig" yaml:"cloudFrontOriginAccessIdentityConfig"`
}

