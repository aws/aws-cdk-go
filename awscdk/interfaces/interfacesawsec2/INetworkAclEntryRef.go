package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkAclEntry.
// Experimental.
type INetworkAclEntryRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkAclEntry resource.
	// Experimental.
	NetworkAclEntryRef() *NetworkAclEntryReference
}

// The jsii proxy for INetworkAclEntryRef
type jsiiProxy_INetworkAclEntryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_INetworkAclEntryRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetworkAclEntryRef) NetworkAclEntryRef() *NetworkAclEntryReference {
	var returns *NetworkAclEntryReference
	_jsii_.Get(
		j,
		"networkAclEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAclEntryRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAclEntryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

