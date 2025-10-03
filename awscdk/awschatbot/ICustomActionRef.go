package awschatbot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomAction.
// Experimental.
type ICustomActionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomActionRef
type jsiiProxy_ICustomActionRef struct {
	internal.Type__constructsIConstruct
}

