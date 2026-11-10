//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
)

func (w *jsiiProxy_WebFunctionEndpointGrants) validateInvokeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateWebFunctionEndpointGrants_FromWebFunctionEndpointParameters(endpoint interfacesawslambda.IWebFunctionEndpointRef) error {
	if endpoint == nil {
		return fmt.Errorf("parameter endpoint is required, but nil was provided")
	}

	return nil
}

