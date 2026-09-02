//go:build no_runtime_type_checking

package awstextract

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAdapterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAdapterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAdapterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAdapterPropsMixinParameters(props *CfnAdapterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

