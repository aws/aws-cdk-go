//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceVersionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourceVersionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceVersionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceVersionPropsMixinParameters(props *CfnResourceVersionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

