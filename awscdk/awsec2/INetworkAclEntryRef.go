package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkAclEntry.
// Experimental.
type INetworkAclEntryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NetworkAclEntry resource.
	// Experimental.
	NetworkAclEntryRef() *NetworkAclEntryReference
}

// The jsii proxy for INetworkAclEntryRef
type jsiiProxy_INetworkAclEntryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_INetworkAclEntryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

