//go:build !no_runtime_type_checking

package previewawssesevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAdvisorRecommendationStatusOpen_AdvisorRecommendationStatusOpenPatternParameters(options *AdvisorRecommendationStatusOpen_AdvisorRecommendationStatusOpenProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

