package awslex

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Bot.
// Experimental.
type IBotRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBotRef
type jsiiProxy_IBotRef struct {
	internal.Type__constructsIConstruct
}

