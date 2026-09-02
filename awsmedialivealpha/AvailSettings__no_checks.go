//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateAvailSettings_EsamParameters(props *EsamSettings) error {
	return nil
}

func validateAvailSettings_SpliceInsertParameters(props *Scte35SpliceInsertSettings) error {
	return nil
}

func validateAvailSettings_TimeSignalAposParameters(props *Scte35TimeSignalAposSettings) error {
	return nil
}

