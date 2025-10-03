package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdMappingWorkflow.
// Experimental.
type IIdMappingWorkflowRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdMappingWorkflowRef
type jsiiProxy_IIdMappingWorkflowRef struct {
	internal.Type__constructsIConstruct
}

