//go:build no_runtime_type_checking

package mixinsawss3tables

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNamespacePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNamespacePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNamespacePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNamespacePropsMixinParameters(props *CfnNamespaceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

