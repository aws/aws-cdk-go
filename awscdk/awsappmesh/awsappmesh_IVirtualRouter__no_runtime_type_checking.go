//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVirtualRouter) validateAddRouteParameters(id *string, props *RouteBaseProps) error {
	return nil
}

