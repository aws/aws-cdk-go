//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDataQualityRulesetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDataQualityRulesetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDataQualityRulesetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDataQualityRulesetPropsMixinParameters(props *CfnDataQualityRulesetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

