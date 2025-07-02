//go:build !no_runtime_type_checking

package awsservicecatalog

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_Portfolio) validateAddProductParameters(product IProduct) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	if tagOptions == nil {
		return fmt.Errorf("parameter tagOptions is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateConstrainCloudFormationParametersParameters(product IProduct, options *CloudFormationRuleConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateConstrainTagUpdatesParameters(product IProduct, options *TagUpdateConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateDeployWithStackSetsParameters(product IProduct, options *StackSetsConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateGenerateUniqueHashParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (p *jsiiProxy_Portfolio) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateGiveAccessToGroupParameters(group awsiam.IGroup) error {
	if group == nil {
		return fmt.Errorf("parameter group is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateGiveAccessToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateGiveAccessToUserParameters(user awsiam.IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateNotifyOnStackEventsParameters(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateSetLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if launchRole == nil {
		return fmt.Errorf("parameter launchRole is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateSetLocalLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if launchRole == nil {
		return fmt.Errorf("parameter launchRole is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateSetLocalLaunchRoleNameParameters(product IProduct, launchRoleName *string, options *CommonConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if launchRoleName == nil {
		return fmt.Errorf("parameter launchRoleName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_Portfolio) validateShareWithAccountParameters(accountId *string, options *PortfolioShareOptions) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validatePortfolio_FromPortfolioArnParameters(scope constructs.Construct, id *string, portfolioArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if portfolioArn == nil {
		return fmt.Errorf("parameter portfolioArn is required, but nil was provided")
	}

	return nil
}

func validatePortfolio_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validatePortfolio_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validatePortfolio_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewPortfolioParameters(scope constructs.Construct, id *string, props *PortfolioProps) error {
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

