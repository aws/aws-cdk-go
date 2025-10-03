package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Cell.
// Experimental.
type ICellRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICellRef
type jsiiProxy_ICellRef struct {
	internal.Type__constructsIConstruct
}

