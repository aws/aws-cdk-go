package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectPeer.
// Experimental.
type IConnectPeerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectPeer resource.
	// Experimental.
	ConnectPeerRef() *ConnectPeerReference
}

// The jsii proxy for IConnectPeerRef
type jsiiProxy_IConnectPeerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IConnectPeerRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IConnectPeerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectPeerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

