//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnPrefixListPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnPrefixListPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnPrefixListPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnPrefixListPropsMixinParameters(props *CfnPrefixListMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

