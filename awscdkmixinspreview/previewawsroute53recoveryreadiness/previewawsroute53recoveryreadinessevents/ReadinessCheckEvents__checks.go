//go:build !no_runtime_type_checking

package previewawsroute53recoveryreadinessevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoveryreadiness"
)

func (r *jsiiProxy_ReadinessCheckEvents) validateRoute53ApplicationRecoveryControllerReadinessCheckStatusChangePatternParameters(options *ReadinessCheckEvents_Route53ApplicationRecoveryControllerReadinessCheckStatusChange_Route53ApplicationRecoveryControllerReadinessCheckStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateReadinessCheckEvents_FromReadinessCheckParameters(readinessCheckRef interfacesawsroute53recoveryreadiness.IReadinessCheckRef) error {
	if readinessCheckRef == nil {
		return fmt.Errorf("parameter readinessCheckRef is required, but nil was provided")
	}

	return nil
}

