package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectPeer.
// Experimental.
type IConnectPeerRef interface {
	constructs.IConstruct
	// A reference to a ConnectPeer resource.
	// Experimental.
	ConnectPeerRef() *ConnectPeerReference
}

// The jsii proxy for IConnectPeerRef
type jsiiProxy_IConnectPeerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IConnectPeerRef) ConnectPeerRef() *ConnectPeerReference {
	var returns *ConnectPeerReference
	_jsii_.Get(
		j,
		"connectPeerRef",
		&returns,
	)
	return returns
}

