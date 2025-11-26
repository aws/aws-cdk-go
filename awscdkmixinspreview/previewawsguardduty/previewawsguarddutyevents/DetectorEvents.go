package previewawsguarddutyevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty"
)

// EventBridge event patterns for Detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var detectorRef IDetectorRef
//
//   detectorEvents := awscdkmixinspreview.Events.DetectorEvents_FromDetector(detectorRef)
//
// Experimental.
type DetectorEvents interface {
	// EventBridge event pattern for Detector GuardDuty Finding.
	// Experimental.
	GuardDutyFindingPattern(options *DetectorEvents_GuardDutyFinding_GuardDutyFindingProps) *awsevents.EventPattern
}

// The jsii proxy struct for DetectorEvents
type jsiiProxy_DetectorEvents struct {
	_ byte // padding
}

// Create DetectorEvents from a Detector reference.
// Experimental.
func DetectorEvents_FromDetector(detectorRef interfacesawsguardduty.IDetectorRef) DetectorEvents {
	_init_.Initialize()

	if err := validateDetectorEvents_FromDetectorParameters(detectorRef); err != nil {
		panic(err)
	}
	var returns DetectorEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_guardduty.events.DetectorEvents",
		"fromDetector",
		[]interface{}{detectorRef},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DetectorEvents) GuardDutyFindingPattern(options *DetectorEvents_GuardDutyFinding_GuardDutyFindingProps) *awsevents.EventPattern {
	if err := d.validateGuardDutyFindingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		d,
		"guardDutyFindingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

