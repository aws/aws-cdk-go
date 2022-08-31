//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stage) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Stage) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Stage) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Stage) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateStage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStage_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewStageParameters(scope constructs.Construct, id *string, props *StageProps) error {
	return nil
}

