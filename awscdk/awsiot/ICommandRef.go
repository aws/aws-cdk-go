package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Command.
// Experimental.
type ICommandRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICommandRef
type jsiiProxy_ICommandRef struct {
	internal.Type__constructsIConstruct
}

