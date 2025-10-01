package awsec2


// A reference to a KeyPair resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyPairReference := &KeyPairReference{
//   	KeyName: jsii.String("keyName"),
//   }
//
type KeyPairReference struct {
	// The KeyName of the KeyPair resource.
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
}

