//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PromptRouter) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validatePromptRouter_FromDefaultIdParameters(defaultRouter DefaultPromptRouterIdentifier, region *string) error {
	return nil
}

func validateNewPromptRouterParameters(props *PromptRouterProps, region *string) error {
	return nil
}

