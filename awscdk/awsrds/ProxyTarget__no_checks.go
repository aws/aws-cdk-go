//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProxyTarget) validateBindParameters(proxy DatabaseProxy) error {
	return nil
}

func validateProxyTarget_FromClusterParameters(cluster IDatabaseCluster) error {
	return nil
}

func validateProxyTarget_FromInstanceParameters(instance IDatabaseInstance) error {
	return nil
}

