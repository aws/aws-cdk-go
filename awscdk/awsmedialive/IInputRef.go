package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Input.
// Experimental.
type IInputRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInputRef
type jsiiProxy_IInputRef struct {
	internal.Type__constructsIConstruct
}

