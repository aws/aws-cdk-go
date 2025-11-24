//go:build no_runtime_type_checking

package mixinsawsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBackupVaultPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBackupVaultPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBackupVaultPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBackupVaultPropsMixinParameters(props *CfnBackupVaultMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

