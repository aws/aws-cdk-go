//go:build no_runtime_type_checking

package awscdkivsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RecordingConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RecordingConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RecordingConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRecordingConfiguration_FromArnParameters(scope constructs.Construct, id *string, recordingConfigurationArn *string) error {
	return nil
}

func validateRecordingConfiguration_FromRecordingConfigurationIdParameters(scope constructs.Construct, id *string, recordingConfigurationId *string) error {
	return nil
}

func validateRecordingConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRecordingConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRecordingConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRecordingConfigurationParameters(scope constructs.Construct, id *string, props *RecordingConfigurationProps) error {
	return nil
}

