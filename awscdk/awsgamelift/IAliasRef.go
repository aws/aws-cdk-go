package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alias.
// Experimental.
type IAliasRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAliasRef
type jsiiProxy_IAliasRef struct {
	internal.Type__constructsIConstruct
}

