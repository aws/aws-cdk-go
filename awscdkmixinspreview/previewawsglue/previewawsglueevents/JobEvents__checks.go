//go:build !no_runtime_type_checking

package previewawsglueevents

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
)

func (j *jsiiProxy_JobEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *JobEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JobEvents) validateGlueJobRunStatusPatternParameters(options *JobEvents_GlueJobRunStatus_GlueJobRunStatusProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JobEvents) validateGlueJobStateChangePatternParameters(options *JobEvents_GlueJobStateChange_GlueJobStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateJobEvents_FromJobParameters(jobRef interfacesawsglue.IJobRef) error {
	if jobRef == nil {
		return fmt.Errorf("parameter jobRef is required, but nil was provided")
	}

	return nil
}

