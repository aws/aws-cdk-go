package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateTemplate.
// Experimental.
type IStateTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStateTemplateRef
type jsiiProxy_IStateTemplateRef struct {
	internal.Type__constructsIConstruct
}

