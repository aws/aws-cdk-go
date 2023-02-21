//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISource) validateBindParameters(scope constructs.Construct, project IProject) error {
	return nil
}

