//go:build !no_runtime_type_checking

package previewawsomicsevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics"
)

func (s *jsiiProxy_SequenceStoreEvents) validateReadSetActivationJobStatusChangePatternParameters(options *SequenceStoreEvents_ReadSetActivationJobStatusChange_ReadSetActivationJobStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SequenceStoreEvents) validateReadSetExportJobStatusChangePatternParameters(options *SequenceStoreEvents_ReadSetExportJobStatusChange_ReadSetExportJobStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SequenceStoreEvents) validateReadSetImportJobStatusChangePatternParameters(options *SequenceStoreEvents_ReadSetImportJobStatusChange_ReadSetImportJobStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_SequenceStoreEvents) validateReadSetStatusChangePatternParameters(options *SequenceStoreEvents_ReadSetStatusChange_ReadSetStatusChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateSequenceStoreEvents_FromSequenceStoreParameters(sequenceStoreRef interfacesawsomics.ISequenceStoreRef) error {
	if sequenceStoreRef == nil {
		return fmt.Errorf("parameter sequenceStoreRef is required, but nil was provided")
	}

	return nil
}

