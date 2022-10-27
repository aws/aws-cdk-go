//go:build no_runtime_type_checking

package awsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAction) validateBindParameters(scope constructs.Construct, options *ActionBindOptions) error {
	return nil
}

