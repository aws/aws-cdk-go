package awsbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BudgetsAction.
// Experimental.
type IBudgetsActionRef interface {
	constructs.IConstruct
	// A reference to a BudgetsAction resource.
	// Experimental.
	BudgetsActionRef() *BudgetsActionReference
}

// The jsii proxy for IBudgetsActionRef
type jsiiProxy_IBudgetsActionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBudgetsActionRef) BudgetsActionRef() *BudgetsActionReference {
	var returns *BudgetsActionReference
	_jsii_.Get(
		j,
		"budgetsActionRef",
		&returns,
	)
	return returns
}

