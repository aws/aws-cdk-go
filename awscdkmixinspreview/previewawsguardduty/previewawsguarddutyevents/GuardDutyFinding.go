package previewawsguarddutyevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.guardduty@GuardDutyFinding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardDutyFinding := awscdkmixinspreview.Events.NewGuardDutyFinding()
//
// Experimental.
type GuardDutyFinding interface {
}

// The jsii proxy struct for GuardDutyFinding
type jsiiProxy_GuardDutyFinding struct {
	_ byte // padding
}

// Experimental.
func NewGuardDutyFinding() GuardDutyFinding {
	_init_.Initialize()

	j := jsiiProxy_GuardDutyFinding{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.events.GuardDutyFinding",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewGuardDutyFinding_Override(g GuardDutyFinding) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.events.GuardDutyFinding",
		nil, // no parameters
		g,
	)
}

// EventBridge event pattern for GuardDuty Finding.
// Experimental.
func GuardDutyFinding_GuardDutyFindingPattern(options *GuardDutyFinding_GuardDutyFindingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateGuardDutyFinding_GuardDutyFindingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_guardduty.events.GuardDutyFinding",
		"guardDutyFindingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

