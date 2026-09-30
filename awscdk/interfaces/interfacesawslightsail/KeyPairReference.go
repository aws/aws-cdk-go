package interfacesawslightsail


// A reference to a KeyPair resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyPairReference := &KeyPairReference{
//   	KeyPairArn: jsii.String("keyPairArn"),
//   	KeyPairName: jsii.String("keyPairName"),
//   }
//
type KeyPairReference struct {
	// The ARN of the KeyPair resource.
	KeyPairArn *string `field:"required" json:"keyPairArn" yaml:"keyPairArn"`
	// The KeyPairName of the KeyPair resource.
	KeyPairName *string `field:"required" json:"keyPairName" yaml:"keyPairName"`
}

