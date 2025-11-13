package interfacesawscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Membership.
// Experimental.
type IMembershipRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Membership resource.
	// Experimental.
	MembershipRef() *MembershipReference
}

// The jsii proxy for IMembershipRef
type jsiiProxy_IMembershipRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMembershipRef) MembershipRef() *MembershipReference {
	var returns *MembershipReference
	_jsii_.Get(
		j,
		"membershipRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMembershipRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMembershipRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

