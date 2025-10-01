package awsbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Budget.
// Experimental.
type IBudgetRef interface {
	constructs.IConstruct
	// A reference to a Budget resource.
	// Experimental.
	BudgetRef() *BudgetReference
}

// The jsii proxy for IBudgetRef
type jsiiProxy_IBudgetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IBudgetRef) BudgetRef() *BudgetReference {
	var returns *BudgetReference
	_jsii_.Get(
		j,
		"budgetRef",
		&returns,
	)
	return returns
}

