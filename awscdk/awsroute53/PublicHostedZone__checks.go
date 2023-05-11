//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PublicHostedZone) validateAddDelegationParameters(delegate IPublicHostedZone, opts *ZoneDelegationOptions) error {
	if delegate == nil {
		return fmt.Errorf("parameter delegate is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(opts, func() string { return "parameter opts" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateAddVpcParameters(_vpc awsec2.IVpc) error {
	if _vpc == nil {
		return fmt.Errorf("parameter _vpc is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PublicHostedZone) validateGrantDelegationParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validatePublicHostedZone_FromHostedZoneAttributesParameters(scope constructs.Construct, id *string, attrs *HostedZoneAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validatePublicHostedZone_FromHostedZoneIdParameters(scope constructs.Construct, id *string, hostedZoneId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if hostedZoneId == nil {
		return fmt.Errorf("parameter hostedZoneId is required, but nil was provided")
	}

	return nil
}

func validatePublicHostedZone_FromLookupParameters(scope constructs.Construct, id *string, query *HostedZoneProviderProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(query, func() string { return "parameter query" }); err != nil {
		return err
	}

	return nil
}

func validatePublicHostedZone_FromPublicHostedZoneAttributesParameters(scope constructs.Construct, id *string, attrs *PublicHostedZoneAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validatePublicHostedZone_FromPublicHostedZoneIdParameters(scope constructs.Construct, id *string, publicHostedZoneId *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if publicHostedZoneId == nil {
		return fmt.Errorf("parameter publicHostedZoneId is required, but nil was provided")
	}

	return nil
}

func validatePublicHostedZone_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePublicHostedZone_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validatePublicHostedZone_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewPublicHostedZoneParameters(scope constructs.Construct, id *string, props *PublicHostedZoneProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

