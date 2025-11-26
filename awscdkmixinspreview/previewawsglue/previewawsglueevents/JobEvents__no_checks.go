//go:build no_runtime_type_checking

package previewawsglueevents

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *JobEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (j *jsiiProxy_JobEvents) validateGlueJobRunStatusPatternParameters(options *JobEvents_GlueJobRunStatus_GlueJobRunStatusProps) error {
	return nil
}

func (j *jsiiProxy_JobEvents) validateGlueJobStateChangePatternParameters(options *JobEvents_GlueJobStateChange_GlueJobStateChangeProps) error {
	return nil
}

func validateJobEvents_FromJobParameters(jobRef interfacesawsglue.IJobRef) error {
	return nil
}

