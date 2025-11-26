//go:build !no_runtime_type_checking

package previewawsguarddutyevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty"
)

func (d *jsiiProxy_DetectorEvents) validateGuardDutyFindingPatternParameters(options *DetectorEvents_GuardDutyFinding_GuardDutyFindingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateDetectorEvents_FromDetectorParameters(detectorRef interfacesawsguardduty.IDetectorRef) error {
	if detectorRef == nil {
		return fmt.Errorf("parameter detectorRef is required, but nil was provided")
	}

	return nil
}

