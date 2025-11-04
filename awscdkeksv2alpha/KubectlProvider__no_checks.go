//go:build no_runtime_type_checking

package awscdkeksv2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateKubectlProvider_FromKubectlProviderAttributesParameters(scope constructs.Construct, id *string, attrs *KubectlProviderAttributes) error {
	return nil
}

func validateKubectlProvider_GetKubectlProviderParameters(scope constructs.Construct, cluster ICluster) error {
	return nil
}

func validateKubectlProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKubectlProviderParameters(scope constructs.Construct, id *string, props *KubectlProviderProps) error {
	return nil
}

