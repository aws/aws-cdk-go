//go:build no_runtime_type_checking

package mixinsawsiam

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceProfilePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceProfilePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstanceProfilePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstanceProfilePropsMixinParameters(props *CfnInstanceProfileMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

