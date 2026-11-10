//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBackupAccessPointPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBackupAccessPointPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBackupAccessPointPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBackupAccessPointPropsMixinParameters(props *CfnBackupAccessPointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

