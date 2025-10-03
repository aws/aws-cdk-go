package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A NetworkAcl.
type INetworkAcl interface {
	INetworkAclRef
	awscdk.IResource
	// Add a new entry to the ACL.
	AddEntry(id *string, options *CommonNetworkAclEntryOptions) NetworkAclEntry
	// ID for the current Network ACL.
	NetworkAclId() *string
}

// The jsii proxy for INetworkAcl
type jsiiProxy_INetworkAcl struct {
	jsiiProxy_INetworkAclRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INetworkAcl) AddEntry(id *string, options *CommonNetworkAclEntryOptions) NetworkAclEntry {
	if err := i.validateAddEntryParameters(id, options); err != nil {
		panic(err)
	}
	var returns NetworkAclEntry

	_jsii_.Invoke(
		i,
		"addEntry",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkAcl) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_INetworkAcl) NetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAcl) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAcl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkAcl) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

