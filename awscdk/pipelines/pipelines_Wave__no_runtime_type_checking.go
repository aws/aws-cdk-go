//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Wave) validateAddStageParameters(stage awscdk.Stage, options *AddStageOpts) error {
	return nil
}

func validateNewWaveParameters(id *string, props *WaveProps) error {
	return nil
}

