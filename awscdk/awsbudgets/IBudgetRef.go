package awsbudgets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbudgets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Budget.
// Experimental.
type IBudgetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBudgetRef
type jsiiProxy_IBudgetRef struct {
	internal.Type__constructsIConstruct
}

