//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVirtualService) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

