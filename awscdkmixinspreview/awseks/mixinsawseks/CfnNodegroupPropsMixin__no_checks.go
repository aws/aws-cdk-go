//go:build no_runtime_type_checking

package mixinsawseks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNodegroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNodegroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNodegroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNodegroupPropsMixinParameters(props *CfnNodegroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

