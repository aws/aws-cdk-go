//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTypeActivationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTypeActivationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTypeActivationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTypeActivationPropsMixinParameters(props *CfnTypeActivationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

