//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapprunner

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVpcConnector) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

