package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.route53recoveryreadiness@Route53ApplicationRecoveryControllerReadinessCheckStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ApplicationRecoveryControllerReadinessCheckStatusChange := awscdkmixinspreview.Events.NewRoute53ApplicationRecoveryControllerReadinessCheckStatusChange()
//
// Experimental.
type Route53ApplicationRecoveryControllerReadinessCheckStatusChange interface {
}

// The jsii proxy struct for Route53ApplicationRecoveryControllerReadinessCheckStatusChange
type jsiiProxy_Route53ApplicationRecoveryControllerReadinessCheckStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewRoute53ApplicationRecoveryControllerReadinessCheckStatusChange() Route53ApplicationRecoveryControllerReadinessCheckStatusChange {
	_init_.Initialize()

	j := jsiiProxy_Route53ApplicationRecoveryControllerReadinessCheckStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerReadinessCheckStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRoute53ApplicationRecoveryControllerReadinessCheckStatusChange_Override(r Route53ApplicationRecoveryControllerReadinessCheckStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerReadinessCheckStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Route 53 Application Recovery Controller readiness check status change.
// Experimental.
func Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern(options *Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateRoute53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.Route53ApplicationRecoveryControllerReadinessCheckStatusChange",
		"route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

