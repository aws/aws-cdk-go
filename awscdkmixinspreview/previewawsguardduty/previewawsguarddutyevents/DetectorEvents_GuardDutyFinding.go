package previewawsguarddutyevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.guardduty@GuardDutyFinding event types for Detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   guardDutyFinding := #error#.NewGuardDutyFinding()
//
// Experimental.
type DetectorEvents_GuardDutyFinding interface {
}

// The jsii proxy struct for DetectorEvents_GuardDutyFinding
type jsiiProxy_DetectorEvents_GuardDutyFinding struct {
	_ byte // padding
}

// Experimental.
func NewDetectorEvents_GuardDutyFinding() DetectorEvents_GuardDutyFinding {
	_init_.Initialize()

	j := jsiiProxy_DetectorEvents_GuardDutyFinding{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.events.DetectorEvents.GuardDutyFinding",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDetectorEvents_GuardDutyFinding_Override(d DetectorEvents_GuardDutyFinding) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_guardduty.events.DetectorEvents.GuardDutyFinding",
		nil, // no parameters
		d,
	)
}

