//go:build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPortfolio) validateAddProductParameters(product IProduct) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateConstrainCloudFormationParametersParameters(product IProduct, options *CloudFormationRuleConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateConstrainTagUpdatesParameters(product IProduct, options *TagUpdateConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateDeployWithStackSetsParameters(product IProduct, options *StackSetsConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateGiveAccessToGroupParameters(group awsiam.IGroup) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateGiveAccessToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateGiveAccessToUserParameters(user awsiam.IUser) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateNotifyOnStackEventsParameters(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateSetLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateSetLocalLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateSetLocalLaunchRoleNameParameters(product IProduct, launchRoleName *string, options *CommonConstraintOptions) error {
	return nil
}

func (i *jsiiProxy_IPortfolio) validateShareWithAccountParameters(accountId *string, options *PortfolioShareOptions) error {
	return nil
}

