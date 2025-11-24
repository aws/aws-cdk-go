//go:build no_runtime_type_checking

package mixinsawsodb

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCloudVmClusterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCloudVmClusterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCloudVmClusterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCloudVmClusterPropsMixinParameters(props *CfnCloudVmClusterMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

