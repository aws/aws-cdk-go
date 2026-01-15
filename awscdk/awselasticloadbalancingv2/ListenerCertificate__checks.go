//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscertificatemanager"
)

func validateListenerCertificate_FromArnParameters(certificateArn *string) error {
	if certificateArn == nil {
		return fmt.Errorf("parameter certificateArn is required, but nil was provided")
	}

	return nil
}

func validateListenerCertificate_FromCertificateManagerParameters(acmCertificate interfacesawscertificatemanager.ICertificateRef) error {
	if acmCertificate == nil {
		return fmt.Errorf("parameter acmCertificate is required, but nil was provided")
	}

	return nil
}

func validateNewListenerCertificateParameters(certificateArn *string) error {
	if certificateArn == nil {
		return fmt.Errorf("parameter certificateArn is required, but nil was provided")
	}

	return nil
}

