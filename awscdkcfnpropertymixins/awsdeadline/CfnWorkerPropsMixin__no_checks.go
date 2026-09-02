//go:build no_runtime_type_checking

package awsdeadline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkerPropsMixinParameters(props *CfnWorkerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

