//go:build !no_runtime_type_checking

package previewawslogs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (x *jsiiProxy_XRayLogsDelivery) validateBindParameters(scope constructs.IConstruct, logType *string, sourceResourceArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if logType == nil {
		return fmt.Errorf("parameter logType is required, but nil was provided")
	}

	if sourceResourceArn == nil {
		return fmt.Errorf("parameter sourceResourceArn is required, but nil was provided")
	}

	return nil
}

func validateNewXRayLogsDeliveryParameters(props *RecordFieldDeliveryProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

