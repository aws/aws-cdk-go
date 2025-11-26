//go:build no_runtime_type_checking

package previewawsbackupmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBackupPlanPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBackupPlanPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBackupPlanPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBackupPlanPropsMixinParameters(props *CfnBackupPlanMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

