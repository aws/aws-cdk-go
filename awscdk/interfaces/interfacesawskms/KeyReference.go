package interfacesawskms


// A reference to a Key resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyReference := &KeyReference{
//   	KeyArn: jsii.String("keyArn"),
//   	KeyId: jsii.String("keyId"),
//   }
//
type KeyReference struct {
	// The ARN of the Key resource.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
	// The KeyId of the Key resource.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

