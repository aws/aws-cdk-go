package interfacesawsnetworkmanager


// A reference to a ConnectPeer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectPeerReference := &ConnectPeerReference{
//   	ConnectPeerId: jsii.String("connectPeerId"),
//   }
//
type ConnectPeerReference struct {
	// The ConnectPeerId of the ConnectPeer resource.
	ConnectPeerId *string `field:"required" json:"connectPeerId" yaml:"connectPeerId"`
}

