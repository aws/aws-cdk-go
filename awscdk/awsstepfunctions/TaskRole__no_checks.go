//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func validateTaskRole_FromRoleParameters(role awsiam.IRole) error {
	return nil
}

func validateTaskRole_FromRoleArnJsonPathParameters(expression *string) error {
	return nil
}

