// The CDK Construct Library for AWS::ServiceCatalog
package awscdkservicecatalogalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog Cloudformation Product.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationProduct interface {
	Product
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	ProductArn() *string
	ProductId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateTagOptions(tagOptions TagOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for CloudFormationProduct
type jsiiProxy_CloudFormationProduct struct {
	jsiiProxy_Product
}

func (j *jsiiProxy_CloudFormationProduct) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationProduct) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationProduct) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationProduct) ProductArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationProduct) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudFormationProduct(scope constructs.Construct, id *string, props *CloudFormationProductProps) CloudFormationProduct {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationProduct{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProduct",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFormationProduct_Override(c CloudFormationProduct, scope constructs.Construct, id *string, props *CloudFormationProductProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProduct",
		[]interface{}{scope, id, props},
		c,
	)
}

// Creates a Product construct that represents an external product.
// Experimental.
func CloudFormationProduct_FromProductArn(scope constructs.Construct, id *string, productArn *string) IProduct {
	_init_.Initialize()

	var returns IProduct

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProduct",
		"fromProductArn",
		[]interface{}{scope, id, productArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CloudFormationProduct_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProduct",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func CloudFormationProduct_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationProduct",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CloudFormationProduct) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associate Tag Options.
//
// A TagOption is a key-value pair managed in AWS Service Catalog.
// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
// Experimental.
func (c *jsiiProxy_CloudFormationProduct) AssociateTagOptions(tagOptions TagOptions) {
	_jsii_.InvokeVoid(
		c,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

// Experimental.
func (c *jsiiProxy_CloudFormationProduct) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (c *jsiiProxy_CloudFormationProduct) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (c *jsiiProxy_CloudFormationProduct) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CloudFormationProduct) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Cloudformation Product.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationProductProps struct {
	// The owner of the product.
	// Experimental.
	Owner *string `json:"owner"`
	// The name of the product.
	// Experimental.
	ProductName *string `json:"productName"`
	// The configuration of the product version.
	// Experimental.
	ProductVersions *[]*CloudFormationProductVersion `json:"productVersions"`
	// The description of the product.
	// Experimental.
	Description *string `json:"description"`
	// The distributor of the product.
	// Experimental.
	Distributor *string `json:"distributor"`
	// The language code.
	//
	// Controls language for logging and errors.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
	// Whether to give provisioning artifacts a new unique identifier when the product attributes or provisioning artifacts is updated.
	// Experimental.
	ReplaceProductVersionIds *bool `json:"replaceProductVersionIds"`
	// The support information about the product.
	// Experimental.
	SupportDescription *string `json:"supportDescription"`
	// The contact email for product support.
	// Experimental.
	SupportEmail *string `json:"supportEmail"`
	// The contact URL for product support.
	// Experimental.
	SupportUrl *string `json:"supportUrl"`
	// TagOptions associated directly to a product.
	// Experimental.
	TagOptions TagOptions `json:"tagOptions"`
}

// Properties of product version (also known as a provisioning artifact).
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationProductVersion struct {
	// The S3 template that points to the provisioning version template.
	// Experimental.
	CloudFormationTemplate CloudFormationTemplate `json:"cloudFormationTemplate"`
	// The description of the product version.
	// Experimental.
	Description *string `json:"description"`
	// The name of the product version.
	// Experimental.
	ProductVersionName *string `json:"productVersionName"`
	// Whether the specified product template will be validated by CloudFormation.
	//
	// If turned off, an invalid template configuration can be stored.
	// Experimental.
	ValidateTemplate *bool `json:"validateTemplate"`
}

// Properties for provisoning rule constraint.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationRuleConstraintOptions struct {
	// The description of the constraint.
	// Experimental.
	Description *string `json:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
	// The rule with condition and assertions to apply to template.
	// Experimental.
	Rule *TemplateRule `json:"rule"`
}

// Represents the Product Provisioning Artifact Template.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationTemplate interface {
	Bind(scope constructs.Construct) *CloudFormationTemplateConfig
}

// The jsii proxy struct for CloudFormationTemplate
type jsiiProxy_CloudFormationTemplate struct {
	_ byte // padding
}

// Experimental.
func NewCloudFormationTemplate_Override(c CloudFormationTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationTemplate",
		nil, // no parameters
		c,
	)
}

// Loads the provisioning artifacts template from a local disk path.
// Experimental.
func CloudFormationTemplate_FromAsset(path *string, options *awss3assets.AssetOptions) CloudFormationTemplate {
	_init_.Initialize()

	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationTemplate",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Creates a product with the resources defined in the given product stack.
// Experimental.
func CloudFormationTemplate_FromProductStack(productStack ProductStack) CloudFormationTemplate {
	_init_.Initialize()

	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationTemplate",
		"fromProductStack",
		[]interface{}{productStack},
		&returns,
	)

	return returns
}

// Template from URL.
// Experimental.
func CloudFormationTemplate_FromUrl(url *string) CloudFormationTemplate {
	_init_.Initialize()

	var returns CloudFormationTemplate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.CloudFormationTemplate",
		"fromUrl",
		[]interface{}{url},
		&returns,
	)

	return returns
}

// Called when the product is initialized to allow this object to bind to the stack, add resources and have fun.
// Experimental.
func (c *jsiiProxy_CloudFormationTemplate) Bind(scope constructs.Construct) *CloudFormationTemplateConfig {
	var returns *CloudFormationTemplateConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Result of binding `Template` into a `Product`.
//
// TODO: EXAMPLE
//
// Experimental.
type CloudFormationTemplateConfig struct {
	// The http url of the template in S3.
	// Experimental.
	HttpUrl *string `json:"httpUrl"`
}

// Properties for governance mechanisms and constraints.
//
// TODO: EXAMPLE
//
// Experimental.
type CommonConstraintOptions struct {
	// The description of the constraint.
	// Experimental.
	Description *string `json:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
}

// A Service Catalog portfolio.
// Experimental.
type IPortfolio interface {
	awscdk.IResource
	// Associate portfolio with the given product.
	// Experimental.
	AddProduct(product IProduct)
	// Associate Tag Options.
	//
	// A TagOption is a key-value pair managed in AWS Service Catalog.
	// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
	// Experimental.
	AssociateTagOptions(tagOptions TagOptions)
	// Set provisioning rules for the product.
	// Experimental.
	ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions)
	// Add a Resource Update Constraint.
	// Experimental.
	ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions)
	// Configure deployment options using AWS Cloudformation StackSets.
	// Experimental.
	DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions)
	// Associate portfolio with an IAM Group.
	// Experimental.
	GiveAccessToGroup(group awsiam.IGroup)
	// Associate portfolio with an IAM Role.
	// Experimental.
	GiveAccessToRole(role awsiam.IRole)
	// Associate portfolio with an IAM User.
	// Experimental.
	GiveAccessToUser(user awsiam.IUser)
	// Add notifications for supplied topics on the provisioned product.
	// Experimental.
	NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// This sets the launch role using the role arn which is tied to the account this role exists in.
	// This is useful if you will be provisioning products from the account where this role exists.
	// If you intend to share the portfolio across accounts, use a local launch role.
	// Experimental.
	SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// The role name will be referenced by in the local account and must be set explicitly.
	// This is useful when sharing the portfolio with multiple accounts.
	// Experimental.
	SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// The role will be referenced by name in the local account instead of a static role arn.
	// A role with this name will automatically be created and assumable by Service Catalog in this account.
	// This is useful when sharing the portfolio with multiple accounts.
	// Experimental.
	SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole
	// Initiate a portfolio share with another account.
	// Experimental.
	ShareWithAccount(accountId *string, options *PortfolioShareOptions)
	// The ARN of the portfolio.
	// Experimental.
	PortfolioArn() *string
	// The ID of the portfolio.
	// Experimental.
	PortfolioId() *string
}

// The jsii proxy for IPortfolio
type jsiiProxy_IPortfolio struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPortfolio) AddProduct(product IProduct) {
	_jsii_.InvokeVoid(
		i,
		"addProduct",
		[]interface{}{product},
	)
}

func (i *jsiiProxy_IPortfolio) AssociateTagOptions(tagOptions TagOptions) {
	_jsii_.InvokeVoid(
		i,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

func (i *jsiiProxy_IPortfolio) ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions) {
	_jsii_.InvokeVoid(
		i,
		"constrainCloudFormationParameters",
		[]interface{}{product, options},
	)
}

func (i *jsiiProxy_IPortfolio) ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions) {
	_jsii_.InvokeVoid(
		i,
		"constrainTagUpdates",
		[]interface{}{product, options},
	)
}

func (i *jsiiProxy_IPortfolio) DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions) {
	_jsii_.InvokeVoid(
		i,
		"deployWithStackSets",
		[]interface{}{product, options},
	)
}

func (i *jsiiProxy_IPortfolio) GiveAccessToGroup(group awsiam.IGroup) {
	_jsii_.InvokeVoid(
		i,
		"giveAccessToGroup",
		[]interface{}{group},
	)
}

func (i *jsiiProxy_IPortfolio) GiveAccessToRole(role awsiam.IRole) {
	_jsii_.InvokeVoid(
		i,
		"giveAccessToRole",
		[]interface{}{role},
	)
}

func (i *jsiiProxy_IPortfolio) GiveAccessToUser(user awsiam.IUser) {
	_jsii_.InvokeVoid(
		i,
		"giveAccessToUser",
		[]interface{}{user},
	)
}

func (i *jsiiProxy_IPortfolio) NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		i,
		"notifyOnStackEvents",
		[]interface{}{product, topic, options},
	)
}

func (i *jsiiProxy_IPortfolio) SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		i,
		"setLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

func (i *jsiiProxy_IPortfolio) SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		i,
		"setLocalLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

func (i *jsiiProxy_IPortfolio) SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole {
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
	_jsii_.InvokeVoid(
		i,
		"shareWithAccount",
		[]interface{}{accountId, options},
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

// A Service Catalog product, currently only supports type CloudFormationProduct.
// Experimental.
type IProduct interface {
	awscdk.IResource
	// Associate Tag Options.
	//
	// A TagOption is a key-value pair managed in AWS Service Catalog.
	// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
	// Experimental.
	AssociateTagOptions(tagOptions TagOptions)
	// The ARN of the product.
	// Experimental.
	ProductArn() *string
	// The id of the product.
	// Experimental.
	ProductId() *string
}

// The jsii proxy for IProduct
type jsiiProxy_IProduct struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProduct) AssociateTagOptions(tagOptions TagOptions) {
	_jsii_.InvokeVoid(
		i,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

func (j *jsiiProxy_IProduct) ProductArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProduct) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

// The language code.
//
// Used for error and logging messages for end users.
// The default behavior if not specified is English.
//
// TODO: EXAMPLE
//
// Experimental.
type MessageLanguage string

const (
	MessageLanguage_EN MessageLanguage = "EN"
	MessageLanguage_JP MessageLanguage = "JP"
	MessageLanguage_ZH MessageLanguage = "ZH"
)

// A Service Catalog portfolio.
//
// TODO: EXAMPLE
//
// Experimental.
type Portfolio interface {
	awscdk.Resource
	IPortfolio
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	PortfolioArn() *string
	PortfolioId() *string
	Stack() awscdk.Stack
	AddProduct(product IProduct)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateTagOptions(tagOptions TagOptions)
	ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions)
	ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions)
	DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions)
	GeneratePhysicalName() *string
	GenerateUniqueHash(value *string) *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GiveAccessToGroup(group awsiam.IGroup)
	GiveAccessToRole(role awsiam.IRole)
	GiveAccessToUser(user awsiam.IUser)
	NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions)
	SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole
	ShareWithAccount(accountId *string, options *PortfolioShareOptions)
	ToString() *string
}

// The jsii proxy struct for Portfolio
type jsiiProxy_Portfolio struct {
	internal.Type__awscdkResource
	jsiiProxy_IPortfolio
}

func (j *jsiiProxy_Portfolio) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) PortfolioArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewPortfolio(scope constructs.Construct, id *string, props *PortfolioProps) Portfolio {
	_init_.Initialize()

	j := jsiiProxy_Portfolio{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.Portfolio",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPortfolio_Override(p Portfolio, scope constructs.Construct, id *string, props *PortfolioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.Portfolio",
		[]interface{}{scope, id, props},
		p,
	)
}

// Creates a Portfolio construct that represents an external portfolio.
// Experimental.
func Portfolio_FromPortfolioArn(scope constructs.Construct, id *string, portfolioArn *string) IPortfolio {
	_init_.Initialize()

	var returns IPortfolio

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.Portfolio",
		"fromPortfolioArn",
		[]interface{}{scope, id, portfolioArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Portfolio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.Portfolio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Portfolio_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.Portfolio",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Associate portfolio with the given product.
// Experimental.
func (p *jsiiProxy_Portfolio) AddProduct(product IProduct) {
	_jsii_.InvokeVoid(
		p,
		"addProduct",
		[]interface{}{product},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (p *jsiiProxy_Portfolio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associate Tag Options.
//
// A TagOption is a key-value pair managed in AWS Service Catalog.
// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
// Experimental.
func (p *jsiiProxy_Portfolio) AssociateTagOptions(tagOptions TagOptions) {
	_jsii_.InvokeVoid(
		p,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

// Set provisioning rules for the product.
// Experimental.
func (p *jsiiProxy_Portfolio) ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"constrainCloudFormationParameters",
		[]interface{}{product, options},
	)
}

// Add a Resource Update Constraint.
// Experimental.
func (p *jsiiProxy_Portfolio) ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"constrainTagUpdates",
		[]interface{}{product, options},
	)
}

// Configure deployment options using AWS Cloudformation StackSets.
// Experimental.
func (p *jsiiProxy_Portfolio) DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"deployWithStackSets",
		[]interface{}{product, options},
	)
}

// Experimental.
func (p *jsiiProxy_Portfolio) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create a unique id based off the L1 CfnPortfolio or the arn of an imported portfolio.
// Experimental.
func (p *jsiiProxy_Portfolio) GenerateUniqueHash(value *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generateUniqueHash",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (p *jsiiProxy_Portfolio) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (p *jsiiProxy_Portfolio) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Associate portfolio with an IAM Group.
// Experimental.
func (p *jsiiProxy_Portfolio) GiveAccessToGroup(group awsiam.IGroup) {
	_jsii_.InvokeVoid(
		p,
		"giveAccessToGroup",
		[]interface{}{group},
	)
}

// Associate portfolio with an IAM Role.
// Experimental.
func (p *jsiiProxy_Portfolio) GiveAccessToRole(role awsiam.IRole) {
	_jsii_.InvokeVoid(
		p,
		"giveAccessToRole",
		[]interface{}{role},
	)
}

// Associate portfolio with an IAM User.
// Experimental.
func (p *jsiiProxy_Portfolio) GiveAccessToUser(user awsiam.IUser) {
	_jsii_.InvokeVoid(
		p,
		"giveAccessToUser",
		[]interface{}{user},
	)
}

// Add notifications for supplied topics on the provisioned product.
// Experimental.
func (p *jsiiProxy_Portfolio) NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"notifyOnStackEvents",
		[]interface{}{product, topic, options},
	)
}

// Force users to assume a certain role when launching a product.
//
// This sets the launch role using the role arn which is tied to the account this role exists in.
// This is useful if you will be provisioning products from the account where this role exists.
// If you intend to share the portfolio across accounts, use a local launch role.
// Experimental.
func (p *jsiiProxy_Portfolio) SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"setLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

// Force users to assume a certain role when launching a product.
//
// The role name will be referenced by in the local account and must be set explicitly.
// This is useful when sharing the portfolio with multiple accounts.
// Experimental.
func (p *jsiiProxy_Portfolio) SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"setLocalLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

// Force users to assume a certain role when launching a product.
//
// The role will be referenced by name in the local account instead of a static role arn.
// A role with this name will automatically be created and assumable by Service Catalog in this account.
// This is useful when sharing the portfolio with multiple accounts.
// Experimental.
func (p *jsiiProxy_Portfolio) SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole {
	var returns awsiam.IRole

	_jsii_.Invoke(
		p,
		"setLocalLaunchRoleName",
		[]interface{}{product, launchRoleName, options},
		&returns,
	)

	return returns
}

// Initiate a portfolio share with another account.
// Experimental.
func (p *jsiiProxy_Portfolio) ShareWithAccount(accountId *string, options *PortfolioShareOptions) {
	_jsii_.InvokeVoid(
		p,
		"shareWithAccount",
		[]interface{}{accountId, options},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_Portfolio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Portfolio.
//
// TODO: EXAMPLE
//
// Experimental.
type PortfolioProps struct {
	// The name of the portfolio.
	// Experimental.
	DisplayName *string `json:"displayName"`
	// The provider name.
	// Experimental.
	ProviderName *string `json:"providerName"`
	// Description for portfolio.
	// Experimental.
	Description *string `json:"description"`
	// The message language.
	//
	// Controls language for
	// status logging and errors.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
	// TagOptions associated directly to a portfolio.
	// Experimental.
	TagOptions TagOptions `json:"tagOptions"`
}

// Options for portfolio share.
//
// TODO: EXAMPLE
//
// Experimental.
type PortfolioShareOptions struct {
	// The message language of the share.
	//
	// Controls status and error message language for share.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
	// Whether to share tagOptions as a part of the portfolio share.
	// Experimental.
	ShareTagOptions *bool `json:"shareTagOptions"`
}

// Abstract class for Service Catalog Product.
//
// TODO: EXAMPLE
//
// Experimental.
type Product interface {
	awscdk.Resource
	IProduct
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	ProductArn() *string
	ProductId() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateTagOptions(tagOptions TagOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Product
type jsiiProxy_Product struct {
	internal.Type__awscdkResource
	jsiiProxy_IProduct
}

func (j *jsiiProxy_Product) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Product) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Product) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Product) ProductArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Product) ProductId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Product) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewProduct_Override(p Product, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.Product",
		[]interface{}{scope, id, props},
		p,
	)
}

// Creates a Product construct that represents an external product.
// Experimental.
func Product_FromProductArn(scope constructs.Construct, id *string, productArn *string) IProduct {
	_init_.Initialize()

	var returns IProduct

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.Product",
		"fromProductArn",
		[]interface{}{scope, id, productArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Product_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.Product",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Product_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.Product",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (p *jsiiProxy_Product) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associate Tag Options.
//
// A TagOption is a key-value pair managed in AWS Service Catalog.
// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
// Experimental.
func (p *jsiiProxy_Product) AssociateTagOptions(tagOptions TagOptions) {
	_jsii_.InvokeVoid(
		p,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

// Experimental.
func (p *jsiiProxy_Product) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (p *jsiiProxy_Product) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (p *jsiiProxy_Product) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_Product) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A Service Catalog product stack, which is similar in form to a Cloudformation nested stack.
//
// You can add the resources to this stack that you want to define for your service catalog product.
//
// This stack will not be treated as an independent deployment
// artifact (won't be listed in "cdk list" or deployable through "cdk deploy"),
// but rather only synthesized as a template and uploaded as an asset to S3.
//
// TODO: EXAMPLE
//
// Experimental.
type ProductStack interface {
	awscdk.Stack
	Account() *string
	ArtifactId() *string
	AvailabilityZones() *[]*string
	Dependencies() *[]awscdk.Stack
	Environment() *string
	Nested() *bool
	NestedStackParent() awscdk.Stack
	NestedStackResource() awscdk.CfnResource
	Node() constructs.Node
	NotificationArns() *[]*string
	Partition() *string
	Region() *string
	StackId() *string
	StackName() *string
	Synthesizer() awscdk.IStackSynthesizer
	Tags() awscdk.TagManager
	TemplateFile() *string
	TemplateOptions() awscdk.ITemplateOptions
	TerminationProtection() *bool
	UrlSuffix() *string
	AddDependency(target awscdk.Stack, reason *string)
	AddTransform(transform *string)
	AllocateLogicalId(cfnElement awscdk.CfnElement) *string
	ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string
	FormatArn(components *awscdk.ArnComponents) *string
	GetLogicalId(element awscdk.CfnElement) *string
	RenameLogicalId(oldId *string, newId *string)
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	Resolve(obj interface{}) interface{}
	SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents
	ToJsonString(obj interface{}, space *float64) *string
	ToString() *string
}

// The jsii proxy struct for ProductStack
type jsiiProxy_ProductStack struct {
	internal.Type__awscdkStack
}

func (j *jsiiProxy_ProductStack) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Dependencies() *[]awscdk.Stack {
	var returns *[]awscdk.Stack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Nested() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) NestedStackParent() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"nestedStackParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) NestedStackResource() awscdk.CfnResource {
	var returns awscdk.CfnResource
	_jsii_.Get(
		j,
		"nestedStackResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Synthesizer() awscdk.IStackSynthesizer {
	var returns awscdk.IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) TemplateOptions() awscdk.ITemplateOptions {
	var returns awscdk.ITemplateOptions
	_jsii_.Get(
		j,
		"templateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProductStack) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


// Experimental.
func NewProductStack(scope constructs.Construct, id *string) ProductStack {
	_init_.Initialize()

	j := jsiiProxy_ProductStack{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.ProductStack",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewProductStack_Override(p ProductStack, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.ProductStack",
		[]interface{}{scope, id},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ProductStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.ProductStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Stack.
//
// We do attribute detection since we can't reliably use 'instanceof'.
// Experimental.
func ProductStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.ProductStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Looks up the first stack scope in which `construct` is defined.
//
// Fails if there is no stack up the tree.
// Experimental.
func ProductStack_Of(construct constructs.IConstruct) awscdk.Stack {
	_init_.Initialize()

	var returns awscdk.Stack

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalog-alpha.ProductStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a dependency between this stack and another stack.
//
// This can be used to define dependencies between any two stacks within an
// app, and also supports nested stacks.
// Experimental.
func (p *jsiiProxy_ProductStack) AddDependency(target awscdk.Stack, reason *string) {
	_jsii_.InvokeVoid(
		p,
		"addDependency",
		[]interface{}{target, reason},
	)
}

// Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.
//
// Duplicate values are removed when stack is synthesized.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
//
// Experimental.
func (p *jsiiProxy_ProductStack) AddTransform(transform *string) {
	_jsii_.InvokeVoid(
		p,
		"addTransform",
		[]interface{}{transform},
	)
}

// Returns the naming scheme used to allocate logical IDs.
//
// By default, uses
// the `HashedAddressingScheme` but this method can be overridden to customize
// this behavior.
//
// In order to make sure logical IDs are unique and stable, we hash the resource
// construct tree path (i.e. toplevel/secondlevel/.../myresource) and add it as
// a suffix to the path components joined without a separator (CloudFormation
// IDs only allow alphanumeric characters).
//
// The result will be:
//
//    <path.join('')><md5(path.join('/')>
//      "human"      "hash"
//
// If the "human" part of the ID exceeds 240 characters, we simply trim it so
// the total ID doesn't exceed CloudFormation's 255 character limit.
//
// We only take 8 characters from the md5 hash (0.000005 chance of collision).
//
// Special cases:
//
// - If the path only contains a single component (i.e. it's a top-level
//    resource), we won't add the hash to it. The hash is not needed for
//    disamiguation and also, it allows for a more straightforward migration an
//    existing CloudFormation template to a CDK stack without logical ID changes
//    (or renames).
// - For aesthetic reasons, if the last components of the path are the same
//    (i.e. `L1/L2/Pipeline/Pipeline`), they will be de-duplicated to make the
//    resulting human portion of the ID more pleasing: `L1L2Pipeline<HASH>`
//    instead of `L1L2PipelinePipeline<HASH>`
// - If a component is named "Default" it will be omitted from the path. This
//    allows refactoring higher level abstractions around constructs without affecting
//    the IDs of already deployed resources.
// - If a component is named "Resource" it will be omitted from the user-visible
//    path, but included in the hash. This reduces visual noise in the human readable
//    part of the identifier.
// Experimental.
func (p *jsiiProxy_ProductStack) AllocateLogicalId(cfnElement awscdk.CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"allocateLogicalId",
		[]interface{}{cfnElement},
		&returns,
	)

	return returns
}

// Create a CloudFormation Export for a value.
//
// Returns a string representing the corresponding `Fn.importValue()`
// expression for this Export. You can control the name for the export by
// passing the `name` option.
//
// If you don't supply a value for `name`, the value you're exporting must be
// a Resource attribute (for example: `bucket.bucketName`) and it will be
// given the same name as the automatic cross-stack reference that would be created
// if you used the attribute in another Stack.
//
// One of the uses for this method is to *remove* the relationship between
// two Stacks established by automatic cross-stack references. It will
// temporarily ensure that the CloudFormation Export still exists while you
// remove the reference from the consuming stack. After that, you can remove
// the resource and the manual export.
//
// ## Example
//
// Here is how the process works. Let's say there are two stacks,
// `producerStack` and `consumerStack`, and `producerStack` has a bucket
// called `bucket`, which is referenced by `consumerStack` (perhaps because
// an AWS Lambda Function writes into it, or something like that).
//
// It is not safe to remove `producerStack.bucket` because as the bucket is being
// deleted, `consumerStack` might still be using it.
//
// Instead, the process takes two deployments:
//
// ### Deployment 1: break the relationship
//
// - Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
//    stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
//    remove the Lambda Function altogether).
// - In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
//    will make sure the CloudFormation Export continues to exist while the relationship
//    between the two stacks is being broken.
// - Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).
//
// ### Deployment 2: remove the bucket resource
//
// - You are now free to remove the `bucket` resource from `producerStack`.
// - Don't forget to remove the `exportValue()` call as well.
// - Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).
// Experimental.
func (p *jsiiProxy_ProductStack) ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"exportValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

// Creates an ARN from components.
//
// If `partition`, `region` or `account` are not specified, the stack's
// partition, region and account will be used.
//
// If any component is the empty string, an empty string will be inserted
// into the generated ARN at the location that component corresponds to.
//
// The ARN will be formatted as follows:
//
//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}}{resource-name}
//
// The required ARN pieces that are omitted will be taken from the stack that
// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
// can be 'undefined'.
// Experimental.
func (p *jsiiProxy_ProductStack) FormatArn(components *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"formatArn",
		[]interface{}{components},
		&returns,
	)

	return returns
}

// Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.
//
// This method is called when a `CfnElement` is created and used to render the
// initial logical identity of resources. Logical ID renames are applied at
// this stage.
//
// This method uses the protected method `allocateLogicalId` to render the
// logical ID for an element. To modify the naming scheme, extend the `Stack`
// class and override this method.
// Experimental.
func (p *jsiiProxy_ProductStack) GetLogicalId(element awscdk.CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getLogicalId",
		[]interface{}{element},
		&returns,
	)

	return returns
}

// Rename a generated logical identities.
//
// To modify the naming scheme strategy, extend the `Stack` class and
// override the `allocateLogicalId` method.
// Experimental.
func (p *jsiiProxy_ProductStack) RenameLogicalId(oldId *string, newId *string) {
	_jsii_.InvokeVoid(
		p,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

// Indicate that a context key was expected.
//
// Contains instructions which will be emitted into the cloud assembly on how
// the key should be supplied.
// Experimental.
func (p *jsiiProxy_ProductStack) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		p,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

// Resolve a tokenized value in the context of the current stack.
// Experimental.
func (p *jsiiProxy_ProductStack) Resolve(obj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Splits the provided ARN into its components.
//
// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
// and a Token representing a dynamic CloudFormation expression
// (in which case the returned components will also be dynamic CloudFormation expressions,
// encoded as Tokens).
// Experimental.
func (p *jsiiProxy_ProductStack) SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents {
	var returns *awscdk.ArnComponents

	_jsii_.Invoke(
		p,
		"splitArn",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

// Convert an object, potentially containing tokens, to a JSON string.
// Experimental.
func (p *jsiiProxy_ProductStack) ToJsonString(obj interface{}, space *float64) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toJsonString",
		[]interface{}{obj, space},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_ProductStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for deploying with Stackset, which creates a StackSet constraint.
//
// TODO: EXAMPLE
//
// Experimental.
type StackSetsConstraintOptions struct {
	// The description of the constraint.
	// Experimental.
	Description *string `json:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
	// List of accounts to deploy stacks to.
	// Experimental.
	Accounts *[]*string `json:"accounts"`
	// IAM role used to administer the StackSets configuration.
	// Experimental.
	AdminRole awsiam.IRole `json:"adminRole"`
	// IAM role used to provision the products in the Stacks.
	// Experimental.
	ExecutionRoleName *string `json:"executionRoleName"`
	// List of regions to deploy stacks to.
	// Experimental.
	Regions *[]*string `json:"regions"`
	// Wether to allow end users to create, update, and delete stacks.
	// Experimental.
	AllowStackSetInstanceOperations *bool `json:"allowStackSetInstanceOperations"`
}

// Defines a Tag Option, which are similar to tags but have multiple values per key.
//
// TODO: EXAMPLE
//
// Experimental.
type TagOptions interface {
	TagOptionsMap() *map[string]*[]*string
}

// The jsii proxy struct for TagOptions
type jsiiProxy_TagOptions struct {
	_ byte // padding
}

func (j *jsiiProxy_TagOptions) TagOptionsMap() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"tagOptionsMap",
		&returns,
	)
	return returns
}


// Experimental.
func NewTagOptions(tagOptionsMap *map[string]*[]*string) TagOptions {
	_init_.Initialize()

	j := jsiiProxy_TagOptions{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.TagOptions",
		[]interface{}{tagOptionsMap},
		&j,
	)

	return &j
}

// Experimental.
func NewTagOptions_Override(t TagOptions, tagOptionsMap *map[string]*[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalog-alpha.TagOptions",
		[]interface{}{tagOptionsMap},
		t,
	)
}

// Properties for ResourceUpdateConstraint.
//
// TODO: EXAMPLE
//
// Experimental.
type TagUpdateConstraintOptions struct {
	// The description of the constraint.
	// Experimental.
	Description *string `json:"description"`
	// The language code.
	//
	// Configures the language for error messages from service catalog.
	// Experimental.
	MessageLanguage MessageLanguage `json:"messageLanguage"`
	// Toggle for if users should be allowed to change/update tags on provisioned products.
	// Experimental.
	Allow *bool `json:"allow"`
}

// Defines the provisioning template constraints.
//
// TODO: EXAMPLE
//
// Experimental.
type TemplateRule struct {
	// A list of assertions that make up the rule.
	// Experimental.
	Assertions *[]*TemplateRuleAssertion `json:"assertions"`
	// Name of the rule.
	// Experimental.
	RuleName *string `json:"ruleName"`
	// Specify when to apply rule with a rule-specific intrinsic function.
	// Experimental.
	Condition awscdk.ICfnRuleConditionExpression `json:"condition"`
}

// An assertion within a template rule, defined by intrinsic functions.
//
// TODO: EXAMPLE
//
// Experimental.
type TemplateRuleAssertion struct {
	// The assertion condition.
	// Experimental.
	Assert awscdk.ICfnRuleConditionExpression `json:"assert"`
	// The description for the asssertion.
	// Experimental.
	Description *string `json:"description"`
}

