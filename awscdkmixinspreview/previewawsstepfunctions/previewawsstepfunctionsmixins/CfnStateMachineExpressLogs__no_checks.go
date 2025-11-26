//go:build no_runtime_type_checking

package previewawsstepfunctionsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStateMachineExpressLogs) validateToLogGroupParameters(logGroup interfacesawslogs.ILogGroupRef) error {
	return nil
}

