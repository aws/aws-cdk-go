//go:build no_runtime_type_checking

package awscertificatemanager

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Certificate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateMetricDaysToExpiryParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateCertificate_FromCertificateArnParameters(scope constructs.Construct, id *string, certificateArn *string) error {
	return nil
}

func validateCertificate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCertificate_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCertificate_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCertificateParameters(scope constructs.Construct, id *string, props *CertificateProps) error {
	return nil
}

