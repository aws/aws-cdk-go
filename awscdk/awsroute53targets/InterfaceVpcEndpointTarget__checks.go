//go:build !no_runtime_type_checking

package awsroute53targets

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

func (i *jsiiProxy_InterfaceVpcEndpointTarget) validateBindParameters(record awsroute53.IRecordSet) error {
	if record == nil {
		return fmt.Errorf("parameter record is required, but nil was provided")
	}

	return nil
}

func validateNewInterfaceVpcEndpointTargetParameters(vpcEndpoint awsec2.InterfaceVpcEndpoint) error {
	if vpcEndpoint == nil {
		return fmt.Errorf("parameter vpcEndpoint is required, but nil was provided")
	}

	return nil
}

