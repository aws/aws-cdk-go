//go:build no_runtime_type_checking

package awscertificatemanager

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateCertificate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PrivateCertificate) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PrivateCertificate) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PrivateCertificate) validateMetricDaysToExpiryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validatePrivateCertificate_FromCertificateArnParameters(scope constructs.Construct, id *string, certificateArn *string) error {
	return nil
}

func validatePrivateCertificate_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrivateCertificate_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrivateCertificate_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPrivateCertificateParameters(scope constructs.Construct, id *string, props *PrivateCertificateProps) error {
	return nil
}

