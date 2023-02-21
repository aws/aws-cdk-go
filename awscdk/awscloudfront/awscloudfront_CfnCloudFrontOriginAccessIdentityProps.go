package awscloudfront


// Properties for defining a `CfnCloudFrontOriginAccessIdentity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudFrontOriginAccessIdentityProps := &cfnCloudFrontOriginAccessIdentityProps{
//   	cloudFrontOriginAccessIdentityConfig: &cloudFrontOriginAccessIdentityConfigProperty{
//   		comment: jsii.String("comment"),
//   	},
//   }
//
type CfnCloudFrontOriginAccessIdentityProps struct {
	// The current configuration information for the identity.
	CloudFrontOriginAccessIdentityConfig interface{} `field:"required" json:"cloudFrontOriginAccessIdentityConfig" yaml:"cloudFrontOriginAccessIdentityConfig"`
}

