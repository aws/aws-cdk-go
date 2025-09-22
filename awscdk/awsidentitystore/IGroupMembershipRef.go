package awsidentitystore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupMembership.
// Experimental.
type IGroupMembershipRef interface {
	constructs.IConstruct
	// A reference to a GroupMembership resource.
	// Experimental.
	GroupMembershipRef() *GroupMembershipReference
}

// The jsii proxy for IGroupMembershipRef
type jsiiProxy_IGroupMembershipRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGroupMembershipRef) GroupMembershipRef() *GroupMembershipReference {
	var returns *GroupMembershipReference
	_jsii_.Get(
		j,
		"groupMembershipRef",
		&returns,
	)
	return returns
}

