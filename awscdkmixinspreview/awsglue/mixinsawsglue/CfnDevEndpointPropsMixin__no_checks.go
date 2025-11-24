//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDevEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDevEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDevEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDevEndpointPropsMixinParameters(props *CfnDevEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

