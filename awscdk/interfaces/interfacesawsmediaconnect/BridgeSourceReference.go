package interfacesawsmediaconnect


// A reference to a BridgeSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeSourceReference := &BridgeSourceReference{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	BridgeSourceName: jsii.String("bridgeSourceName"),
//   }
//
type BridgeSourceReference struct {
	// The BridgeArn of the BridgeSource resource.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The Name of the BridgeSource resource.
	BridgeSourceName *string `field:"required" json:"bridgeSourceName" yaml:"bridgeSourceName"`
}

