//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InvokeActivity) validateBindParameters(_task awsstepfunctions.Task) error {
	return nil
}

func validateNewInvokeActivityParameters(activity awsstepfunctions.IActivity, props *InvokeActivityProps) error {
	return nil
}

