//go:build !no_runtime_type_checking

package awscdkgluealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

func validateConnectionNetwork_SubnetParameters(subnet awsec2.ISubnet) error {
	if subnet == nil {
		return fmt.Errorf("parameter subnet is required, but nil was provided")
	}

	return nil
}

func validateConnectionNetwork_VpcParameters(vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(vpcSubnets, func() string { return "parameter vpcSubnets" }); err != nil {
		return err
	}

	return nil
}

