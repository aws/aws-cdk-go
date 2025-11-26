package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerReadinessCheckStatusChange event types for ReadinessCheck.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerReadinessCheckStatusChange := #error#.NewRoute53ApplicationRecoveryControllerReadinessCheckStatusChange()
//
// Experimental.
type ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange interface {
}

// The jsii proxy struct for ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange
type jsiiProxy_ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange() ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents.Route53ApplicationRecoveryControllerReadinessCheckStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Override(r ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents.Route53ApplicationRecoveryControllerReadinessCheckStatusChange",
		nil, // no parameters
		r,
	)
}

