//go:build no_runtime_type_checking

package mixinsawsiottwinmaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnScenePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnScenePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnScenePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnScenePropsMixinParameters(props *CfnSceneMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

