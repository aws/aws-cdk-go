package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProjectMembership.
// Experimental.
type IProjectMembershipRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ProjectMembership resource.
	// Experimental.
	ProjectMembershipRef() *ProjectMembershipReference
}

// The jsii proxy for IProjectMembershipRef
type jsiiProxy_IProjectMembershipRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IProjectMembershipRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProjectMembershipRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

