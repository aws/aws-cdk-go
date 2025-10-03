package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Script.
// Experimental.
type IScriptRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScriptRef
type jsiiProxy_IScriptRef struct {
	internal.Type__constructsIConstruct
}

