package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Prompt.
// Experimental.
type IPromptRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPromptRef
type jsiiProxy_IPromptRef struct {
	internal.Type__constructsIConstruct
}

