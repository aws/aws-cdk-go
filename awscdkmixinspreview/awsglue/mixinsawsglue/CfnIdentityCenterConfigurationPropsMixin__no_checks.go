//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnIdentityCenterConfigurationPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnIdentityCenterConfigurationPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnIdentityCenterConfigurationPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnIdentityCenterConfigurationPropsMixinParameters(props *CfnIdentityCenterConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

