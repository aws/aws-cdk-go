//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProxyConfiguration) validateBindParameters(_scope constructs.Construct, _taskDefinition TaskDefinition) error {
	return nil
}

