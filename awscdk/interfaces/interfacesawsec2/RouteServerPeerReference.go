package interfacesawsec2


// A reference to a RouteServerPeer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeServerPeerReference := &RouteServerPeerReference{
//   	RouteServerPeerArn: jsii.String("routeServerPeerArn"),
//   	RouteServerPeerId: jsii.String("routeServerPeerId"),
//   }
//
type RouteServerPeerReference struct {
	// The ARN of the RouteServerPeer resource.
	RouteServerPeerArn *string `field:"required" json:"routeServerPeerArn" yaml:"routeServerPeerArn"`
	// The Id of the RouteServerPeer resource.
	RouteServerPeerId *string `field:"required" json:"routeServerPeerId" yaml:"routeServerPeerId"`
}

