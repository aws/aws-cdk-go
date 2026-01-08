package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicecatalog"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog portfolio.
type IPortfolio interface {
	interfacesawsservicecatalog.IPortfolioRef
	awscdk.IResource
	// Associate portfolio with the given product.
	AddProduct(product IProduct)
	// Associate Tag Options.
	//
	// A TagOption is a key-value pair managed in AWS Service Catalog.
	// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
	AssociateTagOptions(tagOptions TagOptions)
	// Set provisioning rules for the product.
	ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions)
	// Add a Resource Update Constraint.
	ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions)
	// Configure deployment options using AWS Cloudformation StackSets.
	DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions)
	// Associate portfolio with an IAM Group.
	GiveAccessToGroup(group awsiam.IGroup)
	// Associate portfolio with an IAM Role.
	GiveAccessToRole(role awsiam.IRole)
	// Associate portfolio with an IAM User.
	GiveAccessToUser(user awsiam.IUser)
	// Add notifications for supplied topics on the provisioned product.
	NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// This sets the launch role using the role arn which is tied to the account this role exists in.
	// This is useful if you will be provisioning products from the account where this role exists.
	// If you intend to share the portfolio across accounts, use a local launch role.
	SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// The role name will be referenced by in the local account and must be set explicitly.
	// This is useful when sharing the portfolio with multiple accounts.
	SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// The role will be referenced by name in the local account instead of a static role arn.
	// A role with this name will automatically be created and assumable by Service Catalog in this account.
	// This is useful when sharing the portfolio with multiple accounts.
	SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole
	// Initiate a portfolio share with another account.
	ShareWithAccount(accountId *string, options *PortfolioShareOptions)
	// The ARN of the portfolio.
	PortfolioArn() *string
	// The ID of the portfolio.
	PortfolioId() *string
}

// The jsii proxy for IPortfolio
type jsiiProxy_IPortfolio struct {
	internal.Type__interfacesawsservicecatalogIPortfolioRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPortfolio) AddProduct(product IProduct) {
	if err := i.validateAddProductParameters(product); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addProduct",
		[]interface{}{product},
	)
}

func (i *jsiiProxy_IPortfolio) AssociateTagOptions(tagOptions TagOptions) {
	if err := i.validateAssociateTagOptionsParameters(tagOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

func (i *jsiiProxy_IPortfolio) ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions) {
	if err := i.validateConstrainCloudFormationParametersParameters(product, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"constrainCloudFormationParameters",
		[]interface{}{product, options},
	)
}

func (i *jsiiProxy_IPortfolio) ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions) {
	if err := i.validateConstrainTagUpdatesParameters(product, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"constrainTagUpdates",
		[]interface{}{product, options},
	)
}

func (i *jsiiProxy_IPortfolio) DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions) {
	if err := i.validateDeployWithStackSetsParameters(product, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"deployWithStackSets",
		[]interface{}{product, options},
	)
}

func (i *jsiiProxy_IPortfolio) GiveAccessToGroup(group awsiam.IGroup) {
	if err := i.validateGiveAccessToGroupParameters(group); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"giveAccessToGroup",
		[]interface{}{group},
	)
}

func (i *jsiiProxy_IPortfolio) GiveAccessToRole(role awsiam.IRole) {
	if err := i.validateGiveAccessToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"giveAccessToRole",
		[]interface{}{role},
	)
}

func (i *jsiiProxy_IPortfolio) GiveAccessToUser(user awsiam.IUser) {
	if err := i.validateGiveAccessToUserParameters(user); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"giveAccessToUser",
		[]interface{}{user},
	)
}

func (i *jsiiProxy_IPortfolio) NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) {
	if err := i.validateNotifyOnStackEventsParameters(product, topic, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"notifyOnStackEvents",
		[]interface{}{product, topic, options},
	)
}

func (i *jsiiProxy_IPortfolio) SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	if err := i.validateSetLaunchRoleParameters(product, launchRole, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"setLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

func (i *jsiiProxy_IPortfolio) SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	if err := i.validateSetLocalLaunchRoleParameters(product, launchRole, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"setLocalLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

func (i *jsiiProxy_IPortfolio) SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole {
	if err := i.validateSetLocalLaunchRoleNameParameters(product, launchRoleName, options); err != nil {
		panic(err)
	}
	var returns awsiam.IRole

	_jsii_.Invoke(
		i,
		"setLocalLaunchRoleName",
		[]interface{}{product, launchRoleName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IPortfolio) ShareWithAccount(accountId *string, options *PortfolioShareOptions) {
	if err := i.validateShareWithAccountParameters(accountId, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"shareWithAccount",
		[]interface{}{accountId, options},
	)
}

func (i *jsiiProxy_IPortfolio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IPortfolio) PortfolioArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolio) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolio) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolio) PortfolioRef() *interfacesawsservicecatalog.PortfolioReference {
	var returns *interfacesawsservicecatalog.PortfolioReference
	_jsii_.Get(
		j,
		"portfolioRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPortfolio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

