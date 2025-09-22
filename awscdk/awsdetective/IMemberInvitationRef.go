package awsdetective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MemberInvitation.
// Experimental.
type IMemberInvitationRef interface {
	constructs.IConstruct
	// A reference to a MemberInvitation resource.
	// Experimental.
	MemberInvitationRef() *MemberInvitationReference
}

// The jsii proxy for IMemberInvitationRef
type jsiiProxy_IMemberInvitationRef struct {
	internal.Type__constructsIConstruct
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

