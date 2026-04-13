//go:build !no_runtime_type_checking

package previewawsconnectevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateContactLensPostCallRulesMatched_EventPatternParameters(options *ContactLensPostCallRulesMatched_ContactLensPostCallRulesMatchedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

