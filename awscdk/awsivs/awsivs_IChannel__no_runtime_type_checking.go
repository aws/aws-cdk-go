//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsivs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IChannel) validateAddStreamKeyParameters(id *string) error {
	return nil
}

