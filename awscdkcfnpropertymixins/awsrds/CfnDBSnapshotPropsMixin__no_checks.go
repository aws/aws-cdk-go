//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDBSnapshotPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDBSnapshotPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDBSnapshotPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDBSnapshotPropsMixinParameters(props *CfnDBSnapshotMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

