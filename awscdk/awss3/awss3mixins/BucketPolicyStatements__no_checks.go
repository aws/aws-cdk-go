//go:build no_runtime_type_checking

package awss3mixins

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketPolicyStatements) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (b *jsiiProxy_BucketPolicyStatements) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucketPolicyStatements_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewBucketPolicyStatementsParameters(statements *[]awsiam.PolicyStatement) error {
	return nil
}

