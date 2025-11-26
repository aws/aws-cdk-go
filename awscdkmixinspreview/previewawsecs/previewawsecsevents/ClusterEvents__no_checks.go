//go:build no_runtime_type_checking

package previewawsecsevents

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *ClusterEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	return nil
}

func (c *jsiiProxy_ClusterEvents) validateECSContainerInstanceStateChangePatternParameters(options *ClusterEvents_ECSContainerInstanceStateChange_ECSContainerInstanceStateChangeProps) error {
	return nil
}

func (c *jsiiProxy_ClusterEvents) validateECSServiceActionPatternParameters(options *ClusterEvents_ECSServiceAction_ECSServiceActionProps) error {
	return nil
}

func (c *jsiiProxy_ClusterEvents) validateECSTaskStateChangePatternParameters(options *ClusterEvents_ECSTaskStateChange_ECSTaskStateChangeProps) error {
	return nil
}

func validateClusterEvents_FromClusterParameters(clusterRef interfacesawsecs.IClusterRef) error {
	return nil
}

