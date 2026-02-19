package interfacesawsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MemberInvitation.
// Experimental.
type IMemberInvitationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MemberInvitation resource.
	// Experimental.
	MemberInvitationRef() *MemberInvitationReference
}

// The jsii proxy for IMemberInvitationRef
type jsiiProxy_IMemberInvitationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMemberInvitationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMemberInvitationRef) MemberInvitationRef() *MemberInvitationReference {
	var returns *MemberInvitationReference
	_jsii_.Get(
		j,
		"memberInvitationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemberInvitationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMemberInvitationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

