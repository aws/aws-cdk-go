//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVerifiedAccessGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVerifiedAccessGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVerifiedAccessGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVerifiedAccessGroupPropsMixinParameters(props *CfnVerifiedAccessGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

