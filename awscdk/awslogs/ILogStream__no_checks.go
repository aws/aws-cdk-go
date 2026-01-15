//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILogStream) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

