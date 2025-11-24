//go:build no_runtime_type_checking

package mixinsawsglue

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConnectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConnectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConnectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConnectionPropsMixinParameters(props *CfnConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

