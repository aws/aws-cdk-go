package interfacesawsidentitystore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AllGroupMemberships.
// Experimental.
type IAllGroupMembershipsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AllGroupMemberships resource.
	// Experimental.
	AllGroupMembershipsRef() *AllGroupMembershipsReference
}

// The jsii proxy for IAllGroupMembershipsRef
type jsiiProxy_IAllGroupMembershipsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAllGroupMembershipsRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAllGroupMembershipsRef) AllGroupMembershipsRef() *AllGroupMembershipsReference {
	var returns *AllGroupMembershipsReference
	_jsii_.Get(
		j,
		"allGroupMembershipsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAllGroupMembershipsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAllGroupMembershipsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

