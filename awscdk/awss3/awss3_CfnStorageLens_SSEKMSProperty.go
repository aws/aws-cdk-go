package awss3


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sSEKMSProperty := &sSEKMSProperty{
//   	keyId: jsii.String("keyId"),
//   }
//
type CfnStorageLens_SSEKMSProperty struct {
	// `CfnStorageLens.SSEKMSProperty.KeyId`.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

