//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILayerVersion) validateAddPermissionParameters(id *string, permission *LayerVersionPermission) error {
	return nil
}

