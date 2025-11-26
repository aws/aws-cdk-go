//go:build no_runtime_type_checking

package previewawsmskmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServerlessClusterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServerlessClusterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServerlessClusterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServerlessClusterPropsMixinParameters(props *CfnServerlessClusterMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

