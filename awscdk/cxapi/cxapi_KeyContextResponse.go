package cxapi


// Properties of a discovered key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyContextResponse := &keyContextResponse{
//   	keyId: jsii.String("keyId"),
//   }
//
// Experimental.
type KeyContextResponse struct {
	// Id of the key.
	// Experimental.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

