package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange := awscdkmixinspreview.Events.NewRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange()
//
// Experimental.
type Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange interface {
}

// The jsii proxy struct for Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange
type jsiiProxy_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange() Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange {
	_init_.Initialize()

	j := jsiiProxy_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Override(r Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Route 53 Application Recovery Controller recovery group readiness status change.
// Experimental.
func Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_EventPattern(options *Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

