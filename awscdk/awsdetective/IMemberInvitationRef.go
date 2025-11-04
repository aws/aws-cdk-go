package awsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MemberInvitation.
// Experimental.
type IMemberInvitationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MemberInvitation resource.
	// Experimental.
	MemberInvitationRef() *MemberInvitationReference
}

// The jsii proxy for IMemberInvitationRef
type jsiiProxy_IMemberInvitationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IMemberInvitationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

