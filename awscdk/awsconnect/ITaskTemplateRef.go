package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskTemplate.
// Experimental.
type ITaskTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITaskTemplateRef
type jsiiProxy_ITaskTemplateRef struct {
	internal.Type__constructsIConstruct
}

