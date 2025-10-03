package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Ruleset.
// Experimental.
type IRulesetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRulesetRef
type jsiiProxy_IRulesetRef struct {
	internal.Type__constructsIConstruct
}

