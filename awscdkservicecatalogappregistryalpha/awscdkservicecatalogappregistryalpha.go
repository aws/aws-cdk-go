// The CDK Construct Library for AWS::ServiceCatalogAppRegistry
package awscdkservicecatalogappregistryalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkservicecatalogappregistryalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog AppRegistry Application.
//
// TODO: EXAMPLE
//
// Experimental.
type Application interface {
	awscdk.Resource
	IApplication
	ApplicationArn() *string
	ApplicationId() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateAttributeGroup(attributeGroup IAttributeGroup)
	AssociateStack(stack awscdk.Stack)
	GeneratePhysicalName() *string
	GenerateUniqueHash(resourceAddress *string) *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Application
type jsiiProxy_Application struct {
	internal.Type__awscdkResource
	jsiiProxy_IApplication
}

func (j *jsiiProxy_Application) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplication(scope constructs.Construct, id *string, props *ApplicationProps) Application {
	_init_.Initialize()

	j := jsiiProxy_Application{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplication_Override(a Application, scope constructs.Construct, id *string, props *ApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports an Application construct that represents an external application.
// Experimental.
func Application_FromApplicationArn(scope constructs.Construct, id *string, applicationArn *string) IApplication {
	_init_.Initialize()

	var returns IApplication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		"fromApplicationArn",
		[]interface{}{scope, id, applicationArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Application_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Application_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.Application",
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
func (a *jsiiProxy_Application) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Associate an attribute group with application If the attribute group is already associated, it will ignore duplicate request.
// Experimental.
func (a *jsiiProxy_Application) AssociateAttributeGroup(attributeGroup IAttributeGroup) {
	_jsii_.InvokeVoid(
		a,
		"associateAttributeGroup",
		[]interface{}{attributeGroup},
	)
}

// Associate a stack with the application If the resource is already associated, it will ignore duplicate request.
//
// A stack can only be associated with one application.
// Experimental.
func (a *jsiiProxy_Application) AssociateStack(stack awscdk.Stack) {
	_jsii_.InvokeVoid(
		a,
		"associateStack",
		[]interface{}{stack},
	)
}

// Experimental.
func (a *jsiiProxy_Application) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create a unique id.
// Experimental.
func (a *jsiiProxy_Application) GenerateUniqueHash(resourceAddress *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generateUniqueHash",
		[]interface{}{resourceAddress},
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
func (a *jsiiProxy_Application) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_Application) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_Application) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Service Catalog AppRegistry Application.
//
// TODO: EXAMPLE
//
// Experimental.
type ApplicationProps struct {
	// Enforces a particular physical application name.
	// Experimental.
	ApplicationName *string `json:"applicationName"`
	// Description for application.
	// Experimental.
	Description *string `json:"description"`
}

// A Service Catalog AppRegistry Attribute Group.
//
// TODO: EXAMPLE
//
// Experimental.
type AttributeGroup interface {
	awscdk.Resource
	IAttributeGroup
	AttributeGroupArn() *string
	AttributeGroupId() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for AttributeGroup
type jsiiProxy_AttributeGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IAttributeGroup
}

func (j *jsiiProxy_AttributeGroup) AttributeGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) AttributeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttributeGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAttributeGroup(scope constructs.Construct, id *string, props *AttributeGroupProps) AttributeGroup {
	_init_.Initialize()

	j := jsiiProxy_AttributeGroup{}

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeGroup_Override(a AttributeGroup, scope constructs.Construct, id *string, props *AttributeGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		[]interface{}{scope, id, props},
		a,
	)
}

// Imports an attribute group construct that represents an external attribute group.
// Experimental.
func AttributeGroup_FromAttributeGroupArn(scope constructs.Construct, id *string, attributeGroupArn *string) IAttributeGroup {
	_init_.Initialize()

	var returns IAttributeGroup

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		"fromAttributeGroupArn",
		[]interface{}{scope, id, attributeGroupArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AttributeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AttributeGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-servicecatalogappregistry-alpha.AttributeGroup",
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
func (a *jsiiProxy_AttributeGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_AttributeGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_AttributeGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_AttributeGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AttributeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a Service Catalog AppRegistry Attribute Group.
//
// TODO: EXAMPLE
//
// Experimental.
type AttributeGroupProps struct {
	// Enforces a particular physical attribute group name.
	// Experimental.
	AttributeGroupName *string `json:"attributeGroupName"`
	// A JSON of nested key-value pairs that represent the attributes in the group.
	//
	// Attributes maybe an empty JSON '{}', but must be explicitly stated.
	// Experimental.
	Attributes *map[string]interface{} `json:"attributes"`
	// Description for attribute group.
	// Experimental.
	Description *string `json:"description"`
}

// A Service Catalog AppRegistry Application.
// Experimental.
type IApplication interface {
	awscdk.IResource
	// Associate thisapplication with an attribute group.
	// Experimental.
	AssociateAttributeGroup(attributeGroup IAttributeGroup)
	// Associate this application with a CloudFormation stack.
	// Experimental.
	AssociateStack(stack awscdk.Stack)
	// The ARN of the application.
	// Experimental.
	ApplicationArn() *string
	// The ID of the application.
	// Experimental.
	ApplicationId() *string
}

// The jsii proxy for IApplication
type jsiiProxy_IApplication struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplication) AssociateAttributeGroup(attributeGroup IAttributeGroup) {
	_jsii_.InvokeVoid(
		i,
		"associateAttributeGroup",
		[]interface{}{attributeGroup},
	)
}

func (i *jsiiProxy_IApplication) AssociateStack(stack awscdk.Stack) {
	_jsii_.InvokeVoid(
		i,
		"associateStack",
		[]interface{}{stack},
	)
}

func (j *jsiiProxy_IApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

// A Service Catalog AppRegistry Attribute Group.
// Experimental.
type IAttributeGroup interface {
	awscdk.IResource
	// The ARN of the attribute group.
	// Experimental.
	AttributeGroupArn() *string
	// The ID of the attribute group.
	// Experimental.
	AttributeGroupId() *string
}

// The jsii proxy for IAttributeGroup
type jsiiProxy_IAttributeGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAttributeGroup) AttributeGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAttributeGroup) AttributeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeGroupId",
		&returns,
	)
	return returns
}

