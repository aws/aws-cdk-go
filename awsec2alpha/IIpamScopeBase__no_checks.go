//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IIpamScopeBase) validateAddPoolParameters(id *string, options *PoolOptions) error {
	return nil
}

