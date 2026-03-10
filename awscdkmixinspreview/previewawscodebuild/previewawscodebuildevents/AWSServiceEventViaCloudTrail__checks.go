//go:build !no_runtime_type_checking

package previewawscodebuildevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAWSServiceEventViaCloudTrail_AwsServiceEventViaCloudTrailPatternParameters(options *AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

