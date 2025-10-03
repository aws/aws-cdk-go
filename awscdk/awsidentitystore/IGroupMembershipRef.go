package awsidentitystore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupMembership.
// Experimental.
type IGroupMembershipRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGroupMembershipRef
type jsiiProxy_IGroupMembershipRef struct {
	internal.Type__constructsIConstruct
}

