//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLogDestination) validateBindParameters(scope constructs.Construct, flowLog FlowLog) error {
	return nil
}

func validateFlowLogDestination_ToS3Parameters(options *S3DestinationOptions) error {
	return nil
}

