//go:build no_runtime_type_checking

package awselasticloadbalancing

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILoadBalancer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

