//go:build no_runtime_type_checking

package alexaask

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSkillPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSkillPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSkillPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSkillPropsMixinParameters(props *CfnSkillMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

