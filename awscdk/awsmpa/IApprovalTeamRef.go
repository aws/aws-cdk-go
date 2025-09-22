package awsmpa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmpa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApprovalTeam.
// Experimental.
type IApprovalTeamRef interface {
	constructs.IConstruct
	// A reference to a ApprovalTeam resource.
	// Experimental.
	ApprovalTeamRef() *ApprovalTeamReference
}

// The jsii proxy for IApprovalTeamRef
type jsiiProxy_IApprovalTeamRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApprovalTeamRef) ApprovalTeamRef() *ApprovalTeamReference {
	var returns *ApprovalTeamReference
	_jsii_.Get(
		j,
		"approvalTeamRef",
		&returns,
	)
	return returns
}

