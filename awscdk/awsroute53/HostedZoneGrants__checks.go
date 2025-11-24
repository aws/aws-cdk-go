//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (h *jsiiProxy_HostedZoneGrants) validateDelegationParameters(grantee awsiam.IGrantable, delegationOptions *GrantDelegationOptions) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(delegationOptions, func() string { return "parameter delegationOptions" }); err != nil {
		return err
	}

	return nil
}

func validateHostedZoneGrants_FromHostedZoneParameters(hostedZone INamedHostedZoneRef) error {
	if hostedZone == nil {
		return fmt.Errorf("parameter hostedZone is required, but nil was provided")
	}

	return nil
}

