//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func validateStageDeployment_FromStageParameters(stage awscdk.Stage, props *StageDeploymentProps) error {
	return nil
}

