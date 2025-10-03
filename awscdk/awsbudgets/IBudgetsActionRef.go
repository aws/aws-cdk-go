package awsbudgets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BudgetsAction.
// Experimental.
type IBudgetsActionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBudgetsActionRef
type jsiiProxy_IBudgetsActionRef struct {
	internal.Type__constructsIConstruct
}

