package interfacesawsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchingWorkflow.
// Experimental.
type IMatchingWorkflowRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MatchingWorkflow resource.
	// Experimental.
	MatchingWorkflowRef() *MatchingWorkflowReference
}

// The jsii proxy for IMatchingWorkflowRef
type jsiiProxy_IMatchingWorkflowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMatchingWorkflowRef) MatchingWorkflowRef() *MatchingWorkflowReference {
	var returns *MatchingWorkflowReference
	_jsii_.Get(
		j,
		"matchingWorkflowRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMatchingWorkflowRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMatchingWorkflowRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

