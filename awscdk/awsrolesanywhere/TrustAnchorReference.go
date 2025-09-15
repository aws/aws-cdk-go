package awsrolesanywhere


// A reference to a TrustAnchor resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trustAnchorReference := &TrustAnchorReference{
//   	TrustAnchorArn: jsii.String("trustAnchorArn"),
//   	TrustAnchorId: jsii.String("trustAnchorId"),
//   }
//
type TrustAnchorReference struct {
	// The ARN of the TrustAnchor resource.
	TrustAnchorArn *string `field:"required" json:"trustAnchorArn" yaml:"trustAnchorArn"`
	// The TrustAnchorId of the TrustAnchor resource.
	TrustAnchorId *string `field:"required" json:"trustAnchorId" yaml:"trustAnchorId"`
}

