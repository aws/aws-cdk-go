//go:build no_runtime_type_checking

package mixinsawsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBackupSelectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBackupSelectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBackupSelectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBackupSelectionPropsMixinParameters(props *CfnBackupSelectionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

