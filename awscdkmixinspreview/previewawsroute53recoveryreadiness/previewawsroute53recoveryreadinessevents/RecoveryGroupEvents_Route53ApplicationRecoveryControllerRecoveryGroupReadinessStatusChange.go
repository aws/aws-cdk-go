package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange event types for RecoveryGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange := #error#.NewRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange()
//
// Experimental.
type RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange interface {
}

// The jsii proxy struct for RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange
type jsiiProxy_RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewRecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange() RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange {
	_init_.Initialize()

	j := jsiiProxy_RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Override(r RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.RecoveryGroupEvents.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange",
		nil, // no parameters
		r,
	)
}

