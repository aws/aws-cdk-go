package awscloudfront


// Properties of CloudFront OriginAccessIdentity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originAccessIdentityProps := &originAccessIdentityProps{
//   	comment: jsii.String("comment"),
//   }
//
type OriginAccessIdentityProps struct {
	// Any comments you want to include about the origin access identity.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

