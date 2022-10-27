//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
)

func (m *jsiiProxy_MutualTlsCertificate) validateBindParameters(_scope awscdk.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateMutualTlsCertificate_AcmParameters(certificate awscertificatemanager.ICertificate) error {
	if certificate == nil {
		return fmt.Errorf("parameter certificate is required, but nil was provided")
	}

	return nil
}

func validateMutualTlsCertificate_FileParameters(certificateChainPath *string, privateKeyPath *string) error {
	if certificateChainPath == nil {
		return fmt.Errorf("parameter certificateChainPath is required, but nil was provided")
	}

	if privateKeyPath == nil {
		return fmt.Errorf("parameter privateKeyPath is required, but nil was provided")
	}

	return nil
}

func validateMutualTlsCertificate_SdsParameters(secretName *string) error {
	if secretName == nil {
		return fmt.Errorf("parameter secretName is required, but nil was provided")
	}

	return nil
}

