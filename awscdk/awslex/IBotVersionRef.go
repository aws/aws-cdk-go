package awslex

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BotVersion.
// Experimental.
type IBotVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBotVersionRef
type jsiiProxy_IBotVersionRef struct {
	internal.Type__constructsIConstruct
}

