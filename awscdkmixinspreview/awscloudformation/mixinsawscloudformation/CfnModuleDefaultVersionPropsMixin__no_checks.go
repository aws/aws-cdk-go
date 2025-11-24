//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnModuleDefaultVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnModuleDefaultVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnModuleDefaultVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnModuleDefaultVersionPropsMixinParameters(props *CfnModuleDefaultVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

