package awscleanrooms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscleanrooms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Membership.
// Experimental.
type IMembershipRef interface {
	constructs.IConstruct
	// A reference to a Membership resource.
	// Experimental.
	MembershipRef() *MembershipReference
}

// The jsii proxy for IMembershipRef
type jsiiProxy_IMembershipRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMembershipRef) MembershipRef() *MembershipReference {
	var returns *MembershipReference
	_jsii_.Get(
		j,
		"membershipRef",
		&returns,
	)
	return returns
}

