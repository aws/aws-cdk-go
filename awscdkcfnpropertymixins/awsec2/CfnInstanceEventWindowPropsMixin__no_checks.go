//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInstanceEventWindowPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInstanceEventWindowPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInstanceEventWindowPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInstanceEventWindowPropsMixinParameters(props *CfnInstanceEventWindowMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

