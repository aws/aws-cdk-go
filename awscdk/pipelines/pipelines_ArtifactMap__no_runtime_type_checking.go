//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ArtifactMap) validateToCodePipelineParameters(x FileSet) error {
	return nil
}

