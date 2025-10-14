package awsroute53recoveryreadiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53recoveryreadiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecoveryGroup.
// Experimental.
type IRecoveryGroupRef interface {
	constructs.IConstruct
	// A reference to a RecoveryGroup resource.
	// Experimental.
	RecoveryGroupRef() *RecoveryGroupReference
}

// The jsii proxy for IRecoveryGroupRef
type jsiiProxy_IRecoveryGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRecoveryGroupRef) RecoveryGroupRef() *RecoveryGroupReference {
	var returns *RecoveryGroupReference
	_jsii_.Get(
		j,
		"recoveryGroupRef",
		&returns,
	)
	return returns
}

