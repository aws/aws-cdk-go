//go:build no_runtime_type_checking

package previewawsecsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterSettings) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_ClusterSettings) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateClusterSettings_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewClusterSettingsParameters(settings *[]*awsecs.CfnCluster_ClusterSettingsProperty) error {
	return nil
}

