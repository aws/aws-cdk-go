//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsekslegacy

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KubernetesResource) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (k *jsiiProxy_KubernetesResource) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateKubernetesResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKubernetesResourceParameters(scope awscdk.Construct, id *string, props *KubernetesResourceProps) error {
	return nil
}

