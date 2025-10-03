package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Membership.
// Experimental.
type IMembershipRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMembershipRef
type jsiiProxy_IMembershipRef struct {
	internal.Type__constructsIConstruct
}

