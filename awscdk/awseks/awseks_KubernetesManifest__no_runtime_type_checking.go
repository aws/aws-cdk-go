//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func validateKubernetesManifest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKubernetesManifestParameters(scope constructs.Construct, id *string, props *KubernetesManifestProps) error {
	return nil
}

