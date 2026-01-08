//go:build no_runtime_type_checking

package awsinspector

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAssessmentTemplate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

