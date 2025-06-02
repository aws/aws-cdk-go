//go:build no_runtime_type_checking

package awsdocdb

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDatabaseCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

