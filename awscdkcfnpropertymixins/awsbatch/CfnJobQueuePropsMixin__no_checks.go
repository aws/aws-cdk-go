//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnJobQueuePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnJobQueuePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnJobQueuePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnJobQueuePropsMixinParameters(props *CfnJobQueueMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

