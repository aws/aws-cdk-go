//go:build no_runtime_type_checking

package mixinsawsomics

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVariantStorePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVariantStorePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVariantStorePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVariantStorePropsMixinParameters(props *CfnVariantStoreMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

