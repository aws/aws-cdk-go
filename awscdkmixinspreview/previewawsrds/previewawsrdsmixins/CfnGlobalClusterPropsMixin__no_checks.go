//go:build no_runtime_type_checking

package previewawsrdsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGlobalClusterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGlobalClusterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGlobalClusterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGlobalClusterPropsMixinParameters(props *CfnGlobalClusterMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

