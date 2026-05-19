//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEvaluator) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEvaluator) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

