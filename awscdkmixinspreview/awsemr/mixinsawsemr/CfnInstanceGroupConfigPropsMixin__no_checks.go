//go:build no_runtime_type_checking

package mixinsawsemr

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceGroupConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceGroupConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstanceGroupConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstanceGroupConfigPropsMixinParameters(props *CfnInstanceGroupConfigMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

