//go:build no_runtime_type_checking

package mixinsawsquicksight

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCConnectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCConnectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCConnectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCConnectionPropsMixinParameters(props *CfnVPCConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

