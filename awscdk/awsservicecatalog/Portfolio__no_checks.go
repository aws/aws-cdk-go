//go:build no_runtime_type_checking

package awsservicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Portfolio) validateAddProductParameters(product IProduct) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateAssociateTagOptionsParameters(tagOptions TagOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateConstrainCloudFormationParametersParameters(product IProduct, options *CloudFormationRuleConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateConstrainTagUpdatesParameters(product IProduct, options *TagUpdateConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateDeployWithStackSetsParameters(product IProduct, options *StackSetsConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateGenerateUniqueHashParameters(value *string) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateGiveAccessToGroupParameters(group awsiam.IGroup) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateGiveAccessToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateGiveAccessToUserParameters(user awsiam.IUser) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateNotifyOnStackEventsParameters(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateSetLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateSetLocalLaunchRoleParameters(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateSetLocalLaunchRoleNameParameters(product IProduct, launchRoleName *string, options *CommonConstraintOptions) error {
	return nil
}

func (p *jsiiProxy_Portfolio) validateShareWithAccountParameters(accountId *string, options *PortfolioShareOptions) error {
	return nil
}

func validatePortfolio_FromPortfolioArnParameters(scope constructs.Construct, id *string, portfolioArn *string) error {
	return nil
}

func validatePortfolio_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePortfolio_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePortfolio_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPortfolioParameters(scope constructs.Construct, id *string, props *PortfolioProps) error {
	return nil
}

