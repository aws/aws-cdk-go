//go:build no_runtime_type_checking

package mixinsawsentityresolution

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIdNamespacePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIdNamespacePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIdNamespacePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIdNamespacePropsMixinParameters(props *CfnIdNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

