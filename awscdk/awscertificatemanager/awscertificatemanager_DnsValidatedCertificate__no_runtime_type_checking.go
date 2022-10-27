//go:build no_runtime_type_checking

package awscertificatemanager

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsValidatedCertificate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DnsValidatedCertificate) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DnsValidatedCertificate) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DnsValidatedCertificate) validateMetricDaysToExpiryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DnsValidatedCertificate) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DnsValidatedCertificate) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDnsValidatedCertificate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDnsValidatedCertificate_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDnsValidatedCertificateParameters(scope constructs.Construct, id *string, props *DnsValidatedCertificateProps) error {
	return nil
}

