package awsiotthingsgraph

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotthingsgraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowTemplate.
// Experimental.
type IFlowTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFlowTemplateRef
type jsiiProxy_IFlowTemplateRef struct {
	internal.Type__constructsIConstruct
}

