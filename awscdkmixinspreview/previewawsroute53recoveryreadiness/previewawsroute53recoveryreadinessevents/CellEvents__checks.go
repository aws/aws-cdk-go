//go:build !no_runtime_type_checking

package previewawsroute53recoveryreadinessevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness"
)

func (c *jsiiProxy_CellEvents) validateRoute53ApplicationRecoveryControllerCellReadinessStatusChangePatternParameters(options *CellEvents_Route53ApplicationRecoveryControllerCellReadinessStatusChange_Route53ApplicationRecoveryControllerCellReadinessStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateCellEvents_FromCellParameters(cellRef interfacesawsroute53recoveryreadiness.ICellRef) error {
	if cellRef == nil {
		return fmt.Errorf("parameter cellRef is required, but nil was provided")
	}

	return nil
}

