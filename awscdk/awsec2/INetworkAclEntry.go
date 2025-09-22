package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A NetworkAclEntry.
type INetworkAclEntry interface {
	INetworkAclEntryRef
	awscdk.IResource
	// The network ACL.
	NetworkAcl() INetworkAcl
}

// The jsii proxy for INetworkAclEntry
type jsiiProxy_INetworkAclEntry struct {
	jsiiProxy_INetworkAclEntryRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INetworkAclEntry) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_INetworkAclEntry) NetworkAcl() INetworkAcl {
	var returns INetworkAcl
	_jsii_.Get(
		j,
		"networkAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAclEntry) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAclEntry) NetworkAclEntryRef() *NetworkAclEntryReference {
	var returns *NetworkAclEntryReference
	_jsii_.Get(
		j,
		"networkAclEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAclEntry) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAclEntry) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

