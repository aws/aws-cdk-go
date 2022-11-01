//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/constructs-go/constructs/v10"
)

func (t *jsiiProxy_TlsCertificate) validateBindParameters(_scope constructs.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateTlsCertificate_AcmParameters(certificate awscertificatemanager.ICertificate) error {
	if certificate == nil {
		return fmt.Errorf("parameter certificate is required, but nil was provided")
	}

	return nil
}

func validateTlsCertificate_FileParameters(certificateChainPath *string, privateKeyPath *string) error {
	if certificateChainPath == nil {
		return fmt.Errorf("parameter certificateChainPath is required, but nil was provided")
	}

	if privateKeyPath == nil {
		return fmt.Errorf("parameter privateKeyPath is required, but nil was provided")
	}

	return nil
}

func validateTlsCertificate_SdsParameters(secretName *string) error {
	if secretName == nil {
		return fmt.Errorf("parameter secretName is required, but nil was provided")
	}

	return nil
}

