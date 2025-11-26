//go:build !no_runtime_type_checking

package previewawsathenaevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsathena"
)

func (w *jsiiProxy_WorkGroupEvents) validateAthenaQueryStateChangePatternParameters(options *WorkGroupEvents_AthenaQueryStateChange_AthenaQueryStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWorkGroupEvents_FromWorkGroupParameters(workGroupRef interfacesawsathena.IWorkGroupRef) error {
	if workGroupRef == nil {
		return fmt.Errorf("parameter workGroupRef is required, but nil was provided")
	}

	return nil
}

