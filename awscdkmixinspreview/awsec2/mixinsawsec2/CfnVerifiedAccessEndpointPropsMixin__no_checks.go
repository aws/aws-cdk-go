//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVerifiedAccessEndpointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVerifiedAccessEndpointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVerifiedAccessEndpointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVerifiedAccessEndpointPropsMixinParameters(props *CfnVerifiedAccessEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

