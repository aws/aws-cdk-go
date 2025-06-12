//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func validateAwsCustomResourcePolicy_FromSdkCallsParameters(options *SdkCallsPolicyOptions) error {
	return nil
}

func validateAwsCustomResourcePolicy_FromStatementsParameters(statements *[]awsiam.PolicyStatement) error {
	return nil
}

