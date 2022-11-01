//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca"
)

func (t *jsiiProxy_TlsValidationTrust) validateBindParameters(scope awscdk.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func validateTlsValidationTrust_AcmParameters(certificateAuthorities *[]awsacmpca.ICertificateAuthority) error {
	if certificateAuthorities == nil {
		return fmt.Errorf("parameter certificateAuthorities is required, but nil was provided")
	}

	return nil
}

func validateTlsValidationTrust_FileParameters(certificateChain *string) error {
	if certificateChain == nil {
		return fmt.Errorf("parameter certificateChain is required, but nil was provided")
	}

	return nil
}

func validateTlsValidationTrust_SdsParameters(secretName *string) error {
	if secretName == nil {
		return fmt.Errorf("parameter secretName is required, but nil was provided")
	}

	return nil
}

