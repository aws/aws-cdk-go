//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRouteTablePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRouteTablePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRouteTablePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRouteTablePropsMixinParameters(props *CfnRouteTableMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

