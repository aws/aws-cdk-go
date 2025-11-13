package interfacesawsmediaconnect


// A reference to a Bridge resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeReference := &BridgeReference{
//   	BridgeArn: jsii.String("bridgeArn"),
//   }
//
type BridgeReference struct {
	// The BridgeArn of the Bridge resource.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
}

