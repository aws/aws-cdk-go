//go:build no_runtime_type_checking

package awsgamelift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFleetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFleetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFleetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFleetPropsMixinParameters(props *CfnFleetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

