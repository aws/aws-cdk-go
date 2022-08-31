//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReportGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ReportGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ReportGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_ReportGroup) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_ReportGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_ReportGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateReportGroup_FromReportGroupNameParameters(scope constructs.Construct, id *string, reportGroupName *string) error {
	return nil
}

func validateReportGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateReportGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewReportGroupParameters(scope constructs.Construct, id *string, props *ReportGroupProps) error {
	return nil
}

