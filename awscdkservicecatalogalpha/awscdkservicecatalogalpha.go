// The CDK Construct Library for AWS::ServiceCatalog
package awscdkservicecatalogalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog Cloudformation Product.
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
}

// Properties of product version (also known as a provisioning artifact).
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
// Experimental.
type CloudFormationTemplateConfig struct {
	// The http url of the template in S3.
	// Experimental.
	HttpUrl *string `json:"httpUrl"`
}

// Properties for governance mechanisms and constraints.
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
	// Configure deployment options using AWS Cloudformaiton StackSets.
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
	// Experimental.
	SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
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
// Experimental.
type MessageLanguage string

const (
	MessageLanguage_EN MessageLanguage = "EN"
	MessageLanguage_JP MessageLanguage = "JP"
	MessageLanguage_ZH MessageLanguage = "ZH"
)

// A Service Catalog portfolio.
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

// Configure deployment options using AWS Cloudformaiton StackSets.
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
// Experimental.
func (p *jsiiProxy_Portfolio) SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	_jsii_.InvokeVoid(
		p,
		"setLaunchRole",
		[]interface{}{product, launchRole, options},
	)
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
	// TagOptions associated directly on portfolio.
	// Experimental.
	TagOptions TagOptions `json:"tagOptions"`
}

// Options for portfolio share.
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

// Properties for deploying with Stackset, which creates a StackSet constraint.
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
// Experimental.
type TemplateRuleAssertion struct {
	// The assertion condition.
	// Experimental.
	Assert awscdk.ICfnRuleConditionExpression `json:"assert"`
	// The description for the asssertion.
	// Experimental.
	Description *string `json:"description"`
}

