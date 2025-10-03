package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionAlias.
// Experimental.
type IConnectionAliasRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectionAliasRef
type jsiiProxy_IConnectionAliasRef struct {
	internal.Type__constructsIConstruct
}

