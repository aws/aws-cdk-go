//go:build no_runtime_type_checking

package mixinsawsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCustomDBEngineVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCustomDBEngineVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCustomDBEngineVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCustomDBEngineVersionPropsMixinParameters(props *CfnCustomDBEngineVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

