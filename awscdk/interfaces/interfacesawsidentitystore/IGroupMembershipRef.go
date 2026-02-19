package interfacesawsidentitystore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupMembership.
// Experimental.
type IGroupMembershipRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GroupMembership resource.
	// Experimental.
	GroupMembershipRef() *GroupMembershipReference
}

// The jsii proxy for IGroupMembershipRef
type jsiiProxy_IGroupMembershipRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGroupMembershipRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGroupMembershipRef) GroupMembershipRef() *GroupMembershipReference {
	var returns *GroupMembershipReference
	_jsii_.Get(
		j,
		"groupMembershipRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupMembershipRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGroupMembershipRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

