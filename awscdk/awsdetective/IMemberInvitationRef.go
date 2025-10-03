package awsdetective

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdetective/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MemberInvitation.
// Experimental.
type IMemberInvitationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMemberInvitationRef
type jsiiProxy_IMemberInvitationRef struct {
	internal.Type__constructsIConstruct
}

