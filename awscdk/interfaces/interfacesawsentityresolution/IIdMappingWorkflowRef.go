package interfacesawsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdMappingWorkflow.
// Experimental.
type IIdMappingWorkflowRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdMappingWorkflow resource.
	// Experimental.
	IdMappingWorkflowRef() *IdMappingWorkflowReference
}

// The jsii proxy for IIdMappingWorkflowRef
type jsiiProxy_IIdMappingWorkflowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIdMappingWorkflowRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIdMappingWorkflowRef) IdMappingWorkflowRef() *IdMappingWorkflowReference {
	var returns *IdMappingWorkflowReference
	_jsii_.Get(
		j,
		"idMappingWorkflowRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdMappingWorkflowRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdMappingWorkflowRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

