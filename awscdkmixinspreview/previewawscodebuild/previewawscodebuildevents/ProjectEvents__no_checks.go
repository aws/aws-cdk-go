//go:build no_runtime_type_checking

package previewawscodebuildevents

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_ProjectEvents) validateCodeBuildBuildPhaseChangePatternParameters(options *CodeBuildBuildPhaseChange_CodeBuildBuildPhaseChangeProps) error {
	return nil
}

func (p *jsiiProxy_ProjectEvents) validateCodeBuildBuildStateChangePatternParameters(options *CodeBuildBuildStateChange_CodeBuildBuildStateChangeProps) error {
	return nil
}

func validateProjectEvents_FromProjectParameters(projectRef interfacesawscodebuild.IProjectRef) error {
	return nil
}

