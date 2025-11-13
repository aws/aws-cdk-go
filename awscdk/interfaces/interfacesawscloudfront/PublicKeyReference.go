package interfacesawscloudfront


// A reference to a PublicKey resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicKeyReference := &PublicKeyReference{
//   	PublicKeyId: jsii.String("publicKeyId"),
//   }
//
type PublicKeyReference struct {
	// The Id of the PublicKey resource.
	PublicKeyId *string `field:"required" json:"publicKeyId" yaml:"publicKeyId"`
}

