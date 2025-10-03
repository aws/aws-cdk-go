package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProjectMembership.
// Experimental.
type IProjectMembershipRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProjectMembershipRef
type jsiiProxy_IProjectMembershipRef struct {
	internal.Type__constructsIConstruct
}

