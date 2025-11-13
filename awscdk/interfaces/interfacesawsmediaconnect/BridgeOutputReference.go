package interfacesawsmediaconnect


// A reference to a BridgeOutput resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeOutputReference := &BridgeOutputReference{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	BridgeOutputName: jsii.String("bridgeOutputName"),
//   }
//
type BridgeOutputReference struct {
	// The BridgeArn of the BridgeOutput resource.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The Name of the BridgeOutput resource.
	BridgeOutputName *string `field:"required" json:"bridgeOutputName" yaml:"bridgeOutputName"`
}

