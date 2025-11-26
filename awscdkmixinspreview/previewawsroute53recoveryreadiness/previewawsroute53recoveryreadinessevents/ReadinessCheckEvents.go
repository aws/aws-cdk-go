package previewawsroute53recoveryreadinessevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness"
)

// EventBridge event patterns for ReadinessCheck.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var readinessCheckRef IReadinessCheckRef
//
//   readinessCheckEvents := awscdkmixinspreview.Events.ReadinessCheckEvents_FromReadinessCheck(readinessCheckRef)
//
// Experimental.
type ReadinessCheckEvents interface {
	// EventBridge event pattern for ReadinessCheck Route 53 Application Recovery Controller readiness check status change.
	// Experimental.
	Route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern(options *ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ReadinessCheckEvents
type jsiiProxy_ReadinessCheckEvents struct {
	_ byte // padding
}

// Create ReadinessCheckEvents from a ReadinessCheck reference.
// Experimental.
func ReadinessCheckEvents_FromReadinessCheck(readinessCheckRef interfacesawsroute53recoveryreadiness.IReadinessCheckRef) ReadinessCheckEvents {
	_init_.Initialize()

	if err := validateReadinessCheckEvents_FromReadinessCheckParameters(readinessCheckRef); err != nil {
		panic(err)
	}
	var returns ReadinessCheckEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53recoveryreadiness.events.ReadinessCheckEvents",
		"fromReadinessCheck",
		[]interface{}{readinessCheckRef},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReadinessCheckEvents) Route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern(options *ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps) *awsevents.EventPattern {
	if err := r.validateRoute53ApplicationRecoveryControllerReadinessCheckStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		r,
		"route53ApplicationRecoveryControllerReadinessCheckStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

