package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Build.
// Experimental.
type IBuildRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBuildRef
type jsiiProxy_IBuildRef struct {
	internal.Type__constructsIConstruct
}

