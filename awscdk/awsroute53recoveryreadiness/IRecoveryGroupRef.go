package awsroute53recoveryreadiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecoveryGroup.
// Experimental.
type IRecoveryGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRecoveryGroupRef
type jsiiProxy_IRecoveryGroupRef struct {
	internal.Type__constructsIConstruct
}

