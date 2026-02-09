package cxapi


// Properties of a discovered key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyContextResponse := &KeyContextResponse{
//   	KeyId: jsii.String("keyId"),
//   }
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type KeyContextResponse struct {
	// Id of the key.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

