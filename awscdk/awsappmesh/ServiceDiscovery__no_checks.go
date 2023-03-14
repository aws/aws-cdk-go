//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceDiscovery) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateServiceDiscovery_CloudMapParameters(service awsservicediscovery.IService) error {
	return nil
}

func validateServiceDiscovery_DnsParameters(hostname *string) error {
	return nil
}

