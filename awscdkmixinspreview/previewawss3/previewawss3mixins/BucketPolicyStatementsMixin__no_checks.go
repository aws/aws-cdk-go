//go:build no_runtime_type_checking

package previewawss3mixins

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketPolicyStatementsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (b *jsiiProxy_BucketPolicyStatementsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucketPolicyStatementsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewBucketPolicyStatementsMixinParameters(statements *[]awsiam.PolicyStatement) error {
	return nil
}

