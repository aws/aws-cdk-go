package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PromptVersion.
// Experimental.
type IPromptVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPromptVersionRef
type jsiiProxy_IPromptVersionRef struct {
	internal.Type__constructsIConstruct
}

