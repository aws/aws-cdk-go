//go:build !no_runtime_type_checking

package awsservicecatalog

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

func (i *jsiiProxy_IPortfolio) validateAddProductParameters(product IProduct) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IPortfolio) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	if tagOptions == nil {
		return fmt.Errorf("parameter tagOptions is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IPortfolio) validateConstrainCloudFormationParametersParameters(product IProduct, options *CloudFormationRuleConstraintOptions) error {
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

func (i *jsiiProxy_IPortfolio) validateConstrainTagUpdatesParameters(product IProduct, options *TagUpdateConstraintOptions) error {
	if product == nil {
		return fmt.Errorf("parameter product is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IPortfolio) validateDeployWithStackSetsParameters(product IProduct, options *StackSetsConstraintOptions) error {
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

func (i *jsiiProxy_IPortfolio) validateGiveAccessToGroupParameters(group awsiam.IGroup) error {
	if group == nil {
		return fmt.Errorf("parameter group is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IPortfolio) validateGiveAccessToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IPortfolio) validateGiveAccessToUserParameters(user awsiam.IUser) error {
	if user == nil {
		return fmt.Errorf("parameter user is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IPortfolio) validateNotifyOnStackEventsParameters(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) error {
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

func (i *jsiiProxy_IPortfolio) validateSetLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
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

func (i *jsiiProxy_IPortfolio) validateSetLocalLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
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

func (i *jsiiProxy_IPortfolio) validateSetLocalLaunchRoleNameParameters(product IProduct, launchRoleName *string, options *CommonConstraintOptions) error {
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

func (i *jsiiProxy_IPortfolio) validateShareWithAccountParameters(accountId *string, options *PortfolioShareOptions) error {
	if accountId == nil {
		return fmt.Errorf("parameter accountId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

