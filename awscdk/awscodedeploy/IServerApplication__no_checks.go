//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IServerApplication) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

