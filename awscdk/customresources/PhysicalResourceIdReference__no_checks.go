//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PhysicalResourceIdReference) validateResolveParameters(_arg awscdk.IResolveContext) error {
	return nil
}

