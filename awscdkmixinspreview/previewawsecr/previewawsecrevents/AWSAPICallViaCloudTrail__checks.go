//go:build !no_runtime_type_checking

package previewawsecrevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAWSAPICallViaCloudTrail_EventPatternParameters(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

