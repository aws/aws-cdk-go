package awslicensemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::LicenseManager::Grant`.
type CfnGrant interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllowedOperations() *[]*string
	SetAllowedOperations(val *[]*string)
	AttrGrantArn() *string
	AttrVersion() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	GrantName() *string
	SetGrantName(val *string)
	HomeRegion() *string
	SetHomeRegion(val *string)
	LicenseArn() *string
	SetLicenseArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	Principals() *[]*string
	SetPrincipals(val *[]*string)
	Ref() *string
	Stack() awscdk.Stack
	Status() *string
	SetStatus(val *string)
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGrant
type jsiiProxy_CfnGrant struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGrant) AllowedOperations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) AttrGrantArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGrantArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) AttrVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) GrantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"grantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) HomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) LicenseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) Principals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGrant) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LicenseManager::Grant`.
func NewCfnGrant(scope constructs.Construct, id *string, props *CfnGrantProps) CfnGrant {
	_init_.Initialize()

	j := jsiiProxy_CfnGrant{}

	_jsii_.Create(
		"aws-cdk-lib.aws_licensemanager.CfnGrant",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LicenseManager::Grant`.
func NewCfnGrant_Override(c CfnGrant, scope constructs.Construct, id *string, props *CfnGrantProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_licensemanager.CfnGrant",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGrant) SetAllowedOperations(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedOperations",
		val,
	)
}

func (j *jsiiProxy_CfnGrant) SetGrantName(val *string) {
	_jsii_.Set(
		j,
		"grantName",
		val,
	)
}

func (j *jsiiProxy_CfnGrant) SetHomeRegion(val *string) {
	_jsii_.Set(
		j,
		"homeRegion",
		val,
	)
}

func (j *jsiiProxy_CfnGrant) SetLicenseArn(val *string) {
	_jsii_.Set(
		j,
		"licenseArn",
		val,
	)
}

func (j *jsiiProxy_CfnGrant) SetPrincipals(val *[]*string) {
	_jsii_.Set(
		j,
		"principals",
		val,
	)
}

func (j *jsiiProxy_CfnGrant) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnGrant_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_licensemanager.CfnGrant",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGrant_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_licensemanager.CfnGrant",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnGrant_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_licensemanager.CfnGrant",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGrant_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_licensemanager.CfnGrant",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGrant) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnGrant) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGrant) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnGrant) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGrant) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnGrant) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnGrant) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnGrant) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnGrant) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnGrant) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGrant) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnGrant) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnGrant) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnGrant) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnGrant) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::LicenseManager::Grant`.
type CfnGrantProps struct {
	// `AWS::LicenseManager::Grant.AllowedOperations`.
	AllowedOperations *[]*string `json:"allowedOperations"`
	// `AWS::LicenseManager::Grant.GrantName`.
	GrantName *string `json:"grantName"`
	// `AWS::LicenseManager::Grant.HomeRegion`.
	HomeRegion *string `json:"homeRegion"`
	// `AWS::LicenseManager::Grant.LicenseArn`.
	LicenseArn *string `json:"licenseArn"`
	// `AWS::LicenseManager::Grant.Principals`.
	Principals *[]*string `json:"principals"`
	// `AWS::LicenseManager::Grant.Status`.
	Status *string `json:"status"`
}

// A CloudFormation `AWS::LicenseManager::License`.
type CfnLicense interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrLicenseArn() *string
	AttrVersion() *string
	Beneficiary() *string
	SetBeneficiary(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConsumptionConfiguration() interface{}
	SetConsumptionConfiguration(val interface{})
	CreationStack() *[]*string
	Entitlements() interface{}
	SetEntitlements(val interface{})
	HomeRegion() *string
	SetHomeRegion(val *string)
	Issuer() interface{}
	SetIssuer(val interface{})
	LicenseMetadata() interface{}
	SetLicenseMetadata(val interface{})
	LicenseName() *string
	SetLicenseName(val *string)
	LogicalId() *string
	Node() constructs.Node
	ProductName() *string
	SetProductName(val *string)
	ProductSku() *string
	SetProductSku(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Status() *string
	SetStatus(val *string)
	UpdatedProperites() *map[string]interface{}
	Validity() interface{}
	SetValidity(val interface{})
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnLicense
type jsiiProxy_CfnLicense struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLicense) AttrLicenseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLicenseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) AttrVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Beneficiary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beneficiary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) ConsumptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Entitlements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitlements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) HomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Issuer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) LicenseMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) LicenseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) ProductName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) ProductSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicense) Validity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validity",
		&returns,
	)
	return returns
}


// Create a new `AWS::LicenseManager::License`.
func NewCfnLicense(scope constructs.Construct, id *string, props *CfnLicenseProps) CfnLicense {
	_init_.Initialize()

	j := jsiiProxy_CfnLicense{}

	_jsii_.Create(
		"aws-cdk-lib.aws_licensemanager.CfnLicense",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LicenseManager::License`.
func NewCfnLicense_Override(c CfnLicense, scope constructs.Construct, id *string, props *CfnLicenseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_licensemanager.CfnLicense",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLicense) SetBeneficiary(val *string) {
	_jsii_.Set(
		j,
		"beneficiary",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetConsumptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"consumptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetEntitlements(val interface{}) {
	_jsii_.Set(
		j,
		"entitlements",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetHomeRegion(val *string) {
	_jsii_.Set(
		j,
		"homeRegion",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetIssuer(val interface{}) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetLicenseMetadata(val interface{}) {
	_jsii_.Set(
		j,
		"licenseMetadata",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetLicenseName(val *string) {
	_jsii_.Set(
		j,
		"licenseName",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetProductName(val *string) {
	_jsii_.Set(
		j,
		"productName",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetProductSku(val *string) {
	_jsii_.Set(
		j,
		"productSku",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetStatus(val *string) {
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_CfnLicense) SetValidity(val interface{}) {
	_jsii_.Set(
		j,
		"validity",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnLicense_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_licensemanager.CfnLicense",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnLicense_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_licensemanager.CfnLicense",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnLicense_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_licensemanager.CfnLicense",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLicense_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_licensemanager.CfnLicense",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnLicense) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnLicense) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnLicense) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

// Adds an override to the synthesized CloudFormation resource.
//
// To add a
// property override, either use `addPropertyOverride` or prefix `path` with
// "Properties." (i.e. `Properties.TopicName`).
//
// If the override is nested, separate each nested level using a dot (.) in the path parameter.
// If there is an array as part of the nesting, specify the index in the path.
//
// To include a literal `.` in the property name, prefix with a `\`. In most
// programming languages you will need to write this as `"\\."` because the
// `\` itself will need to be escaped.
//
// For example,
// ```typescript
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
// ```
// would add the overrides
// ```json
// "Properties": {
//    "GlobalSecondaryIndexes": [
//      {
//        "Projection": {
//          "NonKeyAttributes": [ "myattribute" ]
//          ...
//        }
//        ...
//      },
//      {
//        "ProjectionType": "INCLUDE"
//        ...
//      },
//    ]
//    ...
// }
// ```
// Experimental.
func (c *jsiiProxy_CfnLicense) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnLicense) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnLicense) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnLicense) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnLicense) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnLicense) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnLicense) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnLicense) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLicense) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnLicense) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnLicense) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnLicense) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnLicense_BorrowConfigurationProperty struct {
	// `CfnLicense.BorrowConfigurationProperty.AllowEarlyCheckIn`.
	AllowEarlyCheckIn interface{} `json:"allowEarlyCheckIn"`
	// `CfnLicense.BorrowConfigurationProperty.MaxTimeToLiveInMinutes`.
	MaxTimeToLiveInMinutes *float64 `json:"maxTimeToLiveInMinutes"`
}

type CfnLicense_ConsumptionConfigurationProperty struct {
	// `CfnLicense.ConsumptionConfigurationProperty.BorrowConfiguration`.
	BorrowConfiguration interface{} `json:"borrowConfiguration"`
	// `CfnLicense.ConsumptionConfigurationProperty.ProvisionalConfiguration`.
	ProvisionalConfiguration interface{} `json:"provisionalConfiguration"`
	// `CfnLicense.ConsumptionConfigurationProperty.RenewType`.
	RenewType *string `json:"renewType"`
}

type CfnLicense_EntitlementProperty struct {
	// `CfnLicense.EntitlementProperty.Name`.
	Name *string `json:"name"`
	// `CfnLicense.EntitlementProperty.Unit`.
	Unit *string `json:"unit"`
	// `CfnLicense.EntitlementProperty.AllowCheckIn`.
	AllowCheckIn interface{} `json:"allowCheckIn"`
	// `CfnLicense.EntitlementProperty.MaxCount`.
	MaxCount *float64 `json:"maxCount"`
	// `CfnLicense.EntitlementProperty.Overage`.
	Overage interface{} `json:"overage"`
	// `CfnLicense.EntitlementProperty.Value`.
	Value *string `json:"value"`
}

type CfnLicense_IssuerDataProperty struct {
	// `CfnLicense.IssuerDataProperty.Name`.
	Name *string `json:"name"`
	// `CfnLicense.IssuerDataProperty.SignKey`.
	SignKey *string `json:"signKey"`
}

type CfnLicense_MetadataProperty struct {
	// `CfnLicense.MetadataProperty.Name`.
	Name *string `json:"name"`
	// `CfnLicense.MetadataProperty.Value`.
	Value *string `json:"value"`
}

type CfnLicense_ProvisionalConfigurationProperty struct {
	// `CfnLicense.ProvisionalConfigurationProperty.MaxTimeToLiveInMinutes`.
	MaxTimeToLiveInMinutes *float64 `json:"maxTimeToLiveInMinutes"`
}

type CfnLicense_ValidityDateFormatProperty struct {
	// `CfnLicense.ValidityDateFormatProperty.Begin`.
	Begin *string `json:"begin"`
	// `CfnLicense.ValidityDateFormatProperty.End`.
	End *string `json:"end"`
}

// Properties for defining a `AWS::LicenseManager::License`.
type CfnLicenseProps struct {
	// `AWS::LicenseManager::License.ConsumptionConfiguration`.
	ConsumptionConfiguration interface{} `json:"consumptionConfiguration"`
	// `AWS::LicenseManager::License.Entitlements`.
	Entitlements interface{} `json:"entitlements"`
	// `AWS::LicenseManager::License.HomeRegion`.
	HomeRegion *string `json:"homeRegion"`
	// `AWS::LicenseManager::License.Issuer`.
	Issuer interface{} `json:"issuer"`
	// `AWS::LicenseManager::License.LicenseName`.
	LicenseName *string `json:"licenseName"`
	// `AWS::LicenseManager::License.ProductName`.
	ProductName *string `json:"productName"`
	// `AWS::LicenseManager::License.Validity`.
	Validity interface{} `json:"validity"`
	// `AWS::LicenseManager::License.Beneficiary`.
	Beneficiary *string `json:"beneficiary"`
	// `AWS::LicenseManager::License.LicenseMetadata`.
	LicenseMetadata interface{} `json:"licenseMetadata"`
	// `AWS::LicenseManager::License.ProductSKU`.
	ProductSku *string `json:"productSku"`
	// `AWS::LicenseManager::License.Status`.
	Status *string `json:"status"`
}

