package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerCellReadinessStatusChange event types for Cell.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerCellReadinessStatusChange := #error#.NewRoute53ApplicationRecoveryControllerCellReadinessStatusChange()
//
// Experimental.
type CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange interface {
}

// The jsii proxy struct for CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange
type jsiiProxy_CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewCellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange() CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange {
	_init_.Initialize()

	j := jsiiProxy_CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents.Route53ApplicationRecoveryControllerCellReadinessStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange_Override(c CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.CellEvents.Route53ApplicationRecoveryControllerCellReadinessStatusChange",
		nil, // no parameters
		c,
	)
}

