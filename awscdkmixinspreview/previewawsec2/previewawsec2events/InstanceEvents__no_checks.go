//go:build no_runtime_type_checking

package previewawsec2events

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *InstanceEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (i *jsiiProxy_InstanceEvents) validateEC2InstanceStateChangeNotificationPatternParameters(options *InstanceEvents_EC2InstanceStateChangeNotification_EC2InstanceStateChangeNotificationProps) error {
	return nil
}

func (i *jsiiProxy_InstanceEvents) validateEC2SpotInstanceInterruptionWarningPatternParameters(options *InstanceEvents_EC2SpotInstanceInterruptionWarning_EC2SpotInstanceInterruptionWarningProps) error {
	return nil
}

func validateInstanceEvents_FromInstanceParameters(instanceRef interfacesawsec2.IInstanceRef) error {
	return nil
}

