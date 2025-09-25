package awsivs


// A reference to a PublicKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicKeyReference := &PublicKeyReference{
//   	PublicKeyArn: jsii.String("publicKeyArn"),
//   }
//
type PublicKeyReference struct {
	// The Arn of the PublicKey resource.
	PublicKeyArn *string `field:"required" json:"publicKeyArn" yaml:"publicKeyArn"`
}

