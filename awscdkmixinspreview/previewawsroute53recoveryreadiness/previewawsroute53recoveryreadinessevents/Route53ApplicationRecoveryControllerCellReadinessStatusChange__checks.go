//go:build !no_runtime_type_checking

package previewawsroute53recoveryreadinessevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateRoute53ApplicationRecoveryControllerCellReadinessStatusChange_EventPatternParameters(options *Route53ApplicationRecoveryControllerCellReadinessStatusChange_Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

