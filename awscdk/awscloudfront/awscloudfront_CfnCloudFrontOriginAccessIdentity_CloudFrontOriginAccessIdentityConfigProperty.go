package awscloudfront


// Origin access identity configuration.
//
// Send a `GET` request to the `/ *CloudFront API version* /CloudFront/identity ID/config` resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFrontOriginAccessIdentityConfigProperty := &cloudFrontOriginAccessIdentityConfigProperty{
//   	comment: jsii.String("comment"),
//   }
//
type CfnCloudFrontOriginAccessIdentity_CloudFrontOriginAccessIdentityConfigProperty struct {
	// A comment to describe the origin access identity.
	//
	// The comment cannot be longer than 128 characters.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
}

