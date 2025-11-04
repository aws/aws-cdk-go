package awssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assignment.
// Experimental.
type IAssignmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Assignment resource.
	// Experimental.
	AssignmentRef() *AssignmentReference
}

// The jsii proxy for IAssignmentRef
type jsiiProxy_IAssignmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAssignmentRef) AssignmentRef() *AssignmentReference {
	var returns *AssignmentReference
	_jsii_.Get(
		j,
		"assignmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssignmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssignmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

