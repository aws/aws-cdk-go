package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProjectMembership.
// Experimental.
type IProjectMembershipRef interface {
	constructs.IConstruct
	// A reference to a ProjectMembership resource.
	// Experimental.
	ProjectMembershipRef() *ProjectMembershipReference
}

// The jsii proxy for IProjectMembershipRef
type jsiiProxy_IProjectMembershipRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IProjectMembershipRef) ProjectMembershipRef() *ProjectMembershipReference {
	var returns *ProjectMembershipReference
	_jsii_.Get(
		j,
		"projectMembershipRef",
		&returns,
	)
	return returns
}

