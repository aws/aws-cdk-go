package awsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdMappingWorkflow.
// Experimental.
type IIdMappingWorkflowRef interface {
	constructs.IConstruct
	// A reference to a IdMappingWorkflow resource.
	// Experimental.
	IdMappingWorkflowRef() *IdMappingWorkflowReference
}

// The jsii proxy for IIdMappingWorkflowRef
type jsiiProxy_IIdMappingWorkflowRef struct {
	internal.Type__constructsIConstruct
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

