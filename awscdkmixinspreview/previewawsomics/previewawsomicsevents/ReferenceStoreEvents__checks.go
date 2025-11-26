//go:build !no_runtime_type_checking

package previewawsomicsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics"
)

func (r *jsiiProxy_ReferenceStoreEvents) validateReferenceImportJobStatusChangePatternParameters(options *ReferenceStoreEvents_ReferenceImportJobStatusChange_ReferenceImportJobStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_ReferenceStoreEvents) validateReferenceStatusChangePatternParameters(options *ReferenceStoreEvents_ReferenceStatusChange_ReferenceStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateReferenceStoreEvents_FromReferenceStoreParameters(referenceStoreRef interfacesawsomics.IReferenceStoreRef) error {
	if referenceStoreRef == nil {
		return fmt.Errorf("parameter referenceStoreRef is required, but nil was provided")
	}

	return nil
}

