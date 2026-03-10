//go:build !no_runtime_type_checking

package previewawsdevopsguruevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDevOpsGuruInsightSeverityUpgraded_DevOpsGuruInsightSeverityUpgradedPatternParameters(options *DevOpsGuruInsightSeverityUpgraded_DevOpsGuruInsightSeverityUpgradedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

