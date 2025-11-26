package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness"
)

// EventBridge event patterns for RecoveryGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var recoveryGroupRef IRecoveryGroupRef
//
//   recoveryGroupEvents := awscdkmixinspreview.Events.RecoveryGroupEvents_FromRecoveryGroup(recoveryGroupRef)
//
// Experimental.
type RecoveryGroupEvents interface {
	// EventBridge event pattern for RecoveryGroup Route 53 Application Recovery Controller recovery group readiness status change.
	// Experimental.
	Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePattern(options *RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for RecoveryGroupEvents
type jsiiProxy_RecoveryGroupEvents struct {
	_ byte // padding
}

// Create RecoveryGroupEvents from a RecoveryGroup reference.
// Experimental.
func RecoveryGroupEvents_FromRecoveryGroup(recoveryGroupRef interfacesawsroute53recoveryreadiness.IRecoveryGroupRef) RecoveryGroupEvents {
	_init_.Initialize()

	if err := validateRecoveryGroupEvents_FromRecoveryGroupParameters(recoveryGroupRef); err != nil {
		panic(err)
	}
	var returns RecoveryGroupEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents",
		"fromRecoveryGroup",
		[]interface{}{recoveryGroupRef},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RecoveryGroupEvents) Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePattern(options *RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangeProps) *awsevents.EventPattern {
	if err := r.validateRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

