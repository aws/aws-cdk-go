//go:build no_runtime_type_checking

package mixinsawsconnect

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRoutingProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRoutingProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRoutingProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRoutingProfilePropsMixinParameters(props *CfnRoutingProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

