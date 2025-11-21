//go:build !no_runtime_type_checking

package awsimagebuilderalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DistributionConfiguration) validateAddAmiDistributionsParameters(amiDistributions *[]*AmiDistribution) error {
	for idx_7a3048, v := range *amiDistributions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter amiDistributions[%#v]", idx_7a3048) }); err != nil {
			return err
		}
	}

	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateAddContainerDistributionsParameters(containerDistributions *[]*ContainerDistribution) error {
	for idx_3252cc, v := range *containerDistributions {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter containerDistributions[%#v]", idx_3252cc) }); err != nil {
			return err
		}
	}

	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (d *jsiiProxy_DistributionConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGrantParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DistributionConfiguration) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateDistributionConfiguration_FromDistributionConfigurationArnParameters(scope constructs.Construct, id *string, distributionConfigurationArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if distributionConfigurationArn == nil {
		return fmt.Errorf("parameter distributionConfigurationArn is required, but nil was provided")
	}

	return nil
}

func validateDistributionConfiguration_FromDistributionConfigurationNameParameters(scope constructs.Construct, id *string, distributionConfigurationName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if distributionConfigurationName == nil {
		return fmt.Errorf("parameter distributionConfigurationName is required, but nil was provided")
	}

	return nil
}

func validateDistributionConfiguration_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateDistributionConfiguration_IsDistributionConfigurationParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateDistributionConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateDistributionConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewDistributionConfigurationParameters(scope constructs.Construct, id *string, props *DistributionConfigurationProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

