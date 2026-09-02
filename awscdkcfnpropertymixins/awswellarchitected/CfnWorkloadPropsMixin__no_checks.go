//go:build no_runtime_type_checking

package awswellarchitected

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkloadPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkloadPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkloadPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkloadPropsMixinParameters(props *CfnWorkloadMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

