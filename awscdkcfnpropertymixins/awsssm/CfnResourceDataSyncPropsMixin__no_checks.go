//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnResourceDataSyncPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnResourceDataSyncPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnResourceDataSyncPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnResourceDataSyncPropsMixinParameters(props *CfnResourceDataSyncMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

