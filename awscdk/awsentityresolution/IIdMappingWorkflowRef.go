package awsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdMappingWorkflow.
// Experimental.
type IIdMappingWorkflowRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a IdMappingWorkflow resource.
	// Experimental.
	IdMappingWorkflowRef() *IdMappingWorkflowReference
}

// The jsii proxy for IIdMappingWorkflowRef
type jsiiProxy_IIdMappingWorkflowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IIdMappingWorkflowRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

