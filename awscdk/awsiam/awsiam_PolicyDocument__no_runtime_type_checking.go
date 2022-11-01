//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PolicyDocument) validateResolveParameters(context awscdk.IResolveContext) error {
	return nil
}

func validatePolicyDocument_FromJsonParameters(obj interface{}) error {
	return nil
}

func validateNewPolicyDocumentParameters(props *PolicyDocumentProps) error {
	return nil
}

