//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIntegrationResourcePropertyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIntegrationResourcePropertyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIntegrationResourcePropertyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIntegrationResourcePropertyPropsMixinParameters(props *CfnIntegrationResourcePropertyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

