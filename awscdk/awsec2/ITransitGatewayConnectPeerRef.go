package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayConnectPeer.
// Experimental.
type ITransitGatewayConnectPeerRef interface {
	constructs.IConstruct
	// A reference to a TransitGatewayConnectPeer resource.
	// Experimental.
	TransitGatewayConnectPeerRef() *TransitGatewayConnectPeerReference
}

// The jsii proxy for ITransitGatewayConnectPeerRef
type jsiiProxy_ITransitGatewayConnectPeerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITransitGatewayConnectPeerRef) TransitGatewayConnectPeerRef() *TransitGatewayConnectPeerReference {
	var returns *TransitGatewayConnectPeerReference
	_jsii_.Get(
		j,
		"transitGatewayConnectPeerRef",
		&returns,
	)
	return returns
}

