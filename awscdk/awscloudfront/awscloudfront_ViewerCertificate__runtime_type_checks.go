//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
)

func validateViewerCertificate_FromAcmCertificateParameters(certificate awscertificatemanager.ICertificate, options *ViewerCertificateOptions) error {
	if certificate == nil {
		return fmt.Errorf("parameter certificate is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateViewerCertificate_FromIamCertificateParameters(iamCertificateId *string, options *ViewerCertificateOptions) error {
	if iamCertificateId == nil {
		return fmt.Errorf("parameter iamCertificateId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

