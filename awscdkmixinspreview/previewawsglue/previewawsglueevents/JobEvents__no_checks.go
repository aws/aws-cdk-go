//go:build no_runtime_type_checking

package previewawsglueevents

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (j *jsiiProxy_JobEvents) validateGlueJobRunStatusPatternParameters(options *GlueJobRunStatus_GlueJobRunStatusProps) error {
	return nil
}

func (j *jsiiProxy_JobEvents) validateGlueJobStateChangePatternParameters(options *GlueJobStateChange_GlueJobStateChangeProps) error {
	return nil
}

func validateJobEvents_FromJobParameters(jobRef interfacesawsglue.IJobRef) error {
	return nil
}

