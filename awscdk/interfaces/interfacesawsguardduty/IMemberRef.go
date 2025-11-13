package interfacesawsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Member.
// Experimental.
type IMemberRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Member resource.
	// Experimental.
	MemberRef() *MemberReference
}

// The jsii proxy for IMemberRef
type jsiiProxy_IMemberRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMemberRef) MemberRef() *MemberReference {
	var returns *MemberReference
	_jsii_.Get(
		j,
		"memberRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemberRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemberRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

