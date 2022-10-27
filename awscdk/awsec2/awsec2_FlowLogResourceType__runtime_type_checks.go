//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateFlowLogResourceType_FromNetworkInterfaceIdParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateFlowLogResourceType_FromSubnetParameters(subnet ISubnet) error {
	if subnet == nil {
		return fmt.Errorf("parameter subnet is required, but nil was provided")
	}

	return nil
}

func validateFlowLogResourceType_FromVpcParameters(vpc IVpc) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FlowLogResourceType) validateSetResourceIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_FlowLogResourceType) validateSetResourceTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

