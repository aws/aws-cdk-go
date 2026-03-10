//go:build !no_runtime_type_checking

package previewawsqldbevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateQLDBLedgerStateChange_QldbLedgerStateChangePatternParameters(options *QLDBLedgerStateChange_QLDBLedgerStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

