//go:build no_runtime_type_checking

package previewawsecsevents

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (c *jsiiProxy_ClusterEvents) validateEcsContainerInstanceStateChangePatternParameters(options *ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) error {
	return nil
}

func (c *jsiiProxy_ClusterEvents) validateEcsServiceActionPatternParameters(options *ECSServiceAction_ECSServiceActionProps) error {
	return nil
}

func (c *jsiiProxy_ClusterEvents) validateEcsTaskStateChangePatternParameters(options *ECSTaskStateChange_ECSTaskStateChangeProps) error {
	return nil
}

func validateClusterEvents_FromClusterParameters(clusterRef interfacesawsecs.IClusterRef) error {
	return nil
}

