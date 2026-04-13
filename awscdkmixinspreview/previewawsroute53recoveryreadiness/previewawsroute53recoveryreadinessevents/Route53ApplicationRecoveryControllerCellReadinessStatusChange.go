package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerCellReadinessStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerCellReadinessStatusChange := awscdkmixinspreview.Events.NewRoute53ApplicationRecoveryControllerCellReadinessStatusChange()
//
// Experimental.
type Route53ApplicationRecoveryControllerCellReadinessStatusChange interface {
}

// The jsii proxy struct for Route53ApplicationRecoveryControllerCellReadinessStatusChange
type jsiiProxy_Route53ApplicationRecoveryControllerCellReadinessStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewRoute53ApplicationRecoveryControllerCellReadinessStatusChange() Route53ApplicationRecoveryControllerCellReadinessStatusChange {
	_init_.Initialize()

	j := jsiiProxy_Route53ApplicationRecoveryControllerCellReadinessStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerCellReadinessStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRoute53ApplicationRecoveryControllerCellReadinessStatusChange_Override(r Route53ApplicationRecoveryControllerCellReadinessStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerCellReadinessStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Route 53 Application Recovery Controller cell readiness status change.
// Experimental.
func Route53ApplicationRecoveryControllerCellReadinessStatusChange_EventPattern(options *Route53ApplicationRecoveryControllerCellReadinessStatusChange_Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRoute53ApplicationRecoveryControllerCellReadinessStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerCellReadinessStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

