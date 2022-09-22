//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamResource) validateResourceArnsParameters(api GraphqlApi) error {
	return nil
}

func validateIamResource_OfTypeParameters(type_ *string) error {
	return nil
}

