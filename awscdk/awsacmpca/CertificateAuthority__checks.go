//go:build !no_runtime_type_checking

package awsacmpca

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateCertificateAuthority_FromCertificateAuthorityArnParameters(scope constructs.Construct, id *string, certificateAuthorityArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if certificateAuthorityArn == nil {
		return fmt.Errorf("parameter certificateAuthorityArn is required, but nil was provided")
	}

	return nil
}

