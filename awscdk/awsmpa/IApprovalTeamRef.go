package awsmpa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmpa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApprovalTeam.
// Experimental.
type IApprovalTeamRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApprovalTeamRef
type jsiiProxy_IApprovalTeamRef struct {
	internal.Type__constructsIConstruct
}

