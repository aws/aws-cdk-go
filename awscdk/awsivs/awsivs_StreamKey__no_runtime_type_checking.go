//go:build no_runtime_type_checking

package awsivs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamKey) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StreamKey) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StreamKey) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_StreamKey) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_StreamKey) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateStreamKey_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStreamKey_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewStreamKeyParameters(scope constructs.Construct, id *string, props *StreamKeyProps) error {
	return nil
}

