//go:build no_runtime_type_checking

package mixinsawsappintegrations

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnEventIntegrationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnEventIntegrationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnEventIntegrationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnEventIntegrationPropsMixinParameters(props *CfnEventIntegrationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

