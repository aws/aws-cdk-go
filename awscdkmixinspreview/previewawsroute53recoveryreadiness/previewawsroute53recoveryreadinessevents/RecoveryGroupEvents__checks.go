//go:build !no_runtime_type_checking

package previewawsroute53recoveryreadinessevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness"
)

func (r *jsiiProxy_RecoveryGroupEvents) validateRoute53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangePatternParameters(options *RecoveryGroupEvents_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChange_Route53ApplicationRecoveryControllerRecoveryGroupReadinessStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateRecoveryGroupEvents_FromRecoveryGroupParameters(recoveryGroupRef interfacesawsroute53recoveryreadiness.IRecoveryGroupRef) error {
	if recoveryGroupRef == nil {
		return fmt.Errorf("parameter recoveryGroupRef is required, but nil was provided")
	}

	return nil
}

