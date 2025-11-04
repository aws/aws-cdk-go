package awsbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Budget.
// Experimental.
type IBudgetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Budget resource.
	// Experimental.
	BudgetRef() *BudgetReference
}

// The jsii proxy for IBudgetRef
type jsiiProxy_IBudgetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IBudgetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBudgetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

