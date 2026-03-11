//go:build no_runtime_type_checking

package awssqs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnQueuePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnQueuePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnQueuePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnQueuePropsMixinParameters(props *CfnQueueMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

