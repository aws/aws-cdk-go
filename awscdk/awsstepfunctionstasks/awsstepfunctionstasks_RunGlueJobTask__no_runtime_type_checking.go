//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RunGlueJobTask) validateBindParameters(task awsstepfunctions.Task) error {
	return nil
}

func validateNewRunGlueJobTaskParameters(glueJobName *string, props *RunGlueJobTaskProps) error {
	return nil
}

