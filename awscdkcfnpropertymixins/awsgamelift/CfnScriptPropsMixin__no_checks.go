//go:build no_runtime_type_checking

package awsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnScriptPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnScriptPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnScriptPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnScriptPropsMixinParameters(props *CfnScriptMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

