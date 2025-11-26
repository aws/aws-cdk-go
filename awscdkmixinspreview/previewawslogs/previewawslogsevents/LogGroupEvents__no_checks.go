//go:build no_runtime_type_checking

package previewawslogsevents

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogGroupEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *LogGroupEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func validateLogGroupEvents_FromLogGroupParameters(logGroupRef interfacesawslogs.ILogGroupRef) error {
	return nil
}

