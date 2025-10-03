package awslex

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslex/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BotAlias.
// Experimental.
type IBotAliasRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBotAliasRef
type jsiiProxy_IBotAliasRef struct {
	internal.Type__constructsIConstruct
}

