package awslicensemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::LicenseManager::Grant`.
//
// Specifies a grant.
//
// A grant shares the use of license entitlements with specific AWS accounts . For more information, see [Granted licenses](https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html) in the *AWS License Manager User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   cfnGrant := licensemanager.NewCfnGrant(this, jsii.String("MyCfnGrant"), &cfnGrantProps{
//   	allowedOperations: []*string{
//   		jsii.String("allowedOperations"),
//   	},
//   	grantName: jsii.String("grantName"),
//   	homeRegion: jsii.String("homeRegion"),
//   	licenseArn: jsii.String("licenseArn"),
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	status: jsii.String("status"),
//   })
//
type CfnGrant interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Allowed operations for the grant.
	AllowedOperations() *[]*string
	SetAllowedOperations(val *[]*string)
	// The Amazon Resource Name (ARN) of the grant.
	AttrGrantArn() *string
	// The grant version.
	AttrVersion() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Grant name.
	GrantName() *string
	SetGrantName(val *string)
	// Home Region of the grant.
	HomeRegion() *string
	SetHomeRegion(val *string)
	// License ARN.
	LicenseArn() *string
	SetLicenseArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// The grant principals.
	Principals() *[]*string
	SetPrincipals(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Granted license status.
	Status() *string
	SetStatus(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnGrant) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnGrant) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnGrant) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnGrant) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnGrant) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnGrant) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnGrant) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnGrant) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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

func (c *jsiiProxy_CfnGrant) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnGrant`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   cfnGrantProps := &cfnGrantProps{
//   	allowedOperations: []*string{
//   		jsii.String("allowedOperations"),
//   	},
//   	grantName: jsii.String("grantName"),
//   	homeRegion: jsii.String("homeRegion"),
//   	licenseArn: jsii.String("licenseArn"),
//   	principals: []*string{
//   		jsii.String("principals"),
//   	},
//   	status: jsii.String("status"),
//   }
//
type CfnGrantProps struct {
	// Allowed operations for the grant.
	AllowedOperations *[]*string `json:"allowedOperations" yaml:"allowedOperations"`
	// Grant name.
	GrantName *string `json:"grantName" yaml:"grantName"`
	// Home Region of the grant.
	HomeRegion *string `json:"homeRegion" yaml:"homeRegion"`
	// License ARN.
	LicenseArn *string `json:"licenseArn" yaml:"licenseArn"`
	// The grant principals.
	Principals *[]*string `json:"principals" yaml:"principals"`
	// Granted license status.
	Status *string `json:"status" yaml:"status"`
}

// A CloudFormation `AWS::LicenseManager::License`.
//
// Specifies a granted license.
//
// Granted licenses are licenses for products that your organization purchased from AWS Marketplace or directly from a seller who integrated their software with managed entitlements. For more information, see [Granted licenses](https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html) in the *AWS License Manager User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   cfnLicense := licensemanager.NewCfnLicense(this, jsii.String("MyCfnLicense"), &cfnLicenseProps{
//   	consumptionConfiguration: &consumptionConfigurationProperty{
//   		borrowConfiguration: &borrowConfigurationProperty{
//   			allowEarlyCheckIn: jsii.Boolean(false),
//   			maxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		provisionalConfiguration: &provisionalConfigurationProperty{
//   			maxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		renewType: jsii.String("renewType"),
//   	},
//   	entitlements: []interface{}{
//   		&entitlementProperty{
//   			name: jsii.String("name"),
//   			unit: jsii.String("unit"),
//
//   			// the properties below are optional
//   			allowCheckIn: jsii.Boolean(false),
//   			maxCount: jsii.Number(123),
//   			overage: jsii.Boolean(false),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	homeRegion: jsii.String("homeRegion"),
//   	issuer: &issuerDataProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		signKey: jsii.String("signKey"),
//   	},
//   	licenseName: jsii.String("licenseName"),
//   	productName: jsii.String("productName"),
//   	validity: &validityDateFormatProperty{
//   		begin: jsii.String("begin"),
//   		end: jsii.String("end"),
//   	},
//
//   	// the properties below are optional
//   	beneficiary: jsii.String("beneficiary"),
//   	licenseMetadata: []interface{}{
//   		&metadataProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	productSku: jsii.String("productSku"),
//   	status: jsii.String("status"),
//   })
//
type CfnLicense interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the license.
	AttrLicenseArn() *string
	// The license version.
	AttrVersion() *string
	// License beneficiary.
	Beneficiary() *string
	SetBeneficiary(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Configuration for consumption of the license.
	ConsumptionConfiguration() interface{}
	SetConsumptionConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// License entitlements.
	Entitlements() interface{}
	SetEntitlements(val interface{})
	// Home Region of the license.
	HomeRegion() *string
	SetHomeRegion(val *string)
	// License issuer.
	Issuer() interface{}
	SetIssuer(val interface{})
	// License metadata.
	LicenseMetadata() interface{}
	SetLicenseMetadata(val interface{})
	// License name.
	LicenseName() *string
	SetLicenseName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Product name.
	ProductName() *string
	SetProductName(val *string)
	// Product SKU.
	ProductSku() *string
	SetProductSku(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// License status.
	Status() *string
	SetStatus(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Date and time range during which the license is valid, in ISO8601-UTC format.
	Validity() interface{}
	SetValidity(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
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
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (c *jsiiProxy_CfnLicense) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLicense) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLicense) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLicense) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLicense) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLicense) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLicense) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

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

func (c *jsiiProxy_CfnLicense) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

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

func (c *jsiiProxy_CfnLicense) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Details about a borrow configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   borrowConfigurationProperty := &borrowConfigurationProperty{
//   	allowEarlyCheckIn: jsii.Boolean(false),
//   	maxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
type CfnLicense_BorrowConfigurationProperty struct {
	// Indicates whether early check-ins are allowed.
	AllowEarlyCheckIn interface{} `json:"allowEarlyCheckIn" yaml:"allowEarlyCheckIn"`
	// Maximum time for the borrow configuration, in minutes.
	MaxTimeToLiveInMinutes *float64 `json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

// Details about a consumption configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   consumptionConfigurationProperty := &consumptionConfigurationProperty{
//   	borrowConfiguration: &borrowConfigurationProperty{
//   		allowEarlyCheckIn: jsii.Boolean(false),
//   		maxTimeToLiveInMinutes: jsii.Number(123),
//   	},
//   	provisionalConfiguration: &provisionalConfigurationProperty{
//   		maxTimeToLiveInMinutes: jsii.Number(123),
//   	},
//   	renewType: jsii.String("renewType"),
//   }
//
type CfnLicense_ConsumptionConfigurationProperty struct {
	// Details about a borrow configuration.
	BorrowConfiguration interface{} `json:"borrowConfiguration" yaml:"borrowConfiguration"`
	// Details about a provisional configuration.
	ProvisionalConfiguration interface{} `json:"provisionalConfiguration" yaml:"provisionalConfiguration"`
	// Renewal frequency.
	RenewType *string `json:"renewType" yaml:"renewType"`
}

// Describes a resource entitled for use with a license.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   entitlementProperty := &entitlementProperty{
//   	name: jsii.String("name"),
//   	unit: jsii.String("unit"),
//
//   	// the properties below are optional
//   	allowCheckIn: jsii.Boolean(false),
//   	maxCount: jsii.Number(123),
//   	overage: jsii.Boolean(false),
//   	value: jsii.String("value"),
//   }
//
type CfnLicense_EntitlementProperty struct {
	// Entitlement name.
	Name *string `json:"name" yaml:"name"`
	// Entitlement unit.
	Unit *string `json:"unit" yaml:"unit"`
	// Indicates whether check-ins are allowed.
	AllowCheckIn interface{} `json:"allowCheckIn" yaml:"allowCheckIn"`
	// Maximum entitlement count.
	//
	// Use if the unit is not None.
	MaxCount *float64 `json:"maxCount" yaml:"maxCount"`
	// Indicates whether overages are allowed.
	Overage interface{} `json:"overage" yaml:"overage"`
	// Entitlement resource.
	//
	// Use only if the unit is None.
	Value *string `json:"value" yaml:"value"`
}

// Details associated with the issuer of a license.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   issuerDataProperty := &issuerDataProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	signKey: jsii.String("signKey"),
//   }
//
type CfnLicense_IssuerDataProperty struct {
	// Issuer name.
	Name *string `json:"name" yaml:"name"`
	// Asymmetric KMS key from AWS Key Management Service .
	//
	// The KMS key must have a key usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.
	SignKey *string `json:"signKey" yaml:"signKey"`
}

// Describes key/value pairs.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   metadataProperty := &metadataProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnLicense_MetadataProperty struct {
	// The key name.
	Name *string `json:"name" yaml:"name"`
	// The value.
	Value *string `json:"value" yaml:"value"`
}

// Details about a provisional configuration.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   provisionalConfigurationProperty := &provisionalConfigurationProperty{
//   	maxTimeToLiveInMinutes: jsii.Number(123),
//   }
//
type CfnLicense_ProvisionalConfigurationProperty struct {
	// Maximum time for the provisional configuration, in minutes.
	MaxTimeToLiveInMinutes *float64 `json:"maxTimeToLiveInMinutes" yaml:"maxTimeToLiveInMinutes"`
}

// Date and time range during which the license is valid, in ISO8601-UTC format.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   validityDateFormatProperty := &validityDateFormatProperty{
//   	begin: jsii.String("begin"),
//   	end: jsii.String("end"),
//   }
//
type CfnLicense_ValidityDateFormatProperty struct {
	// Start of the time range.
	Begin *string `json:"begin" yaml:"begin"`
	// End of the time range.
	End *string `json:"end" yaml:"end"`
}

// Properties for defining a `CfnLicense`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import licensemanager "github.com/aws/aws-cdk-go/awscdk/aws_licensemanager"
//   cfnLicenseProps := &cfnLicenseProps{
//   	consumptionConfiguration: &consumptionConfigurationProperty{
//   		borrowConfiguration: &borrowConfigurationProperty{
//   			allowEarlyCheckIn: jsii.Boolean(false),
//   			maxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		provisionalConfiguration: &provisionalConfigurationProperty{
//   			maxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		renewType: jsii.String("renewType"),
//   	},
//   	entitlements: []interface{}{
//   		&entitlementProperty{
//   			name: jsii.String("name"),
//   			unit: jsii.String("unit"),
//
//   			// the properties below are optional
//   			allowCheckIn: jsii.Boolean(false),
//   			maxCount: jsii.Number(123),
//   			overage: jsii.Boolean(false),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	homeRegion: jsii.String("homeRegion"),
//   	issuer: &issuerDataProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		signKey: jsii.String("signKey"),
//   	},
//   	licenseName: jsii.String("licenseName"),
//   	productName: jsii.String("productName"),
//   	validity: &validityDateFormatProperty{
//   		begin: jsii.String("begin"),
//   		end: jsii.String("end"),
//   	},
//
//   	// the properties below are optional
//   	beneficiary: jsii.String("beneficiary"),
//   	licenseMetadata: []interface{}{
//   		&metadataProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	productSku: jsii.String("productSku"),
//   	status: jsii.String("status"),
//   }
//
type CfnLicenseProps struct {
	// Configuration for consumption of the license.
	ConsumptionConfiguration interface{} `json:"consumptionConfiguration" yaml:"consumptionConfiguration"`
	// License entitlements.
	Entitlements interface{} `json:"entitlements" yaml:"entitlements"`
	// Home Region of the license.
	HomeRegion *string `json:"homeRegion" yaml:"homeRegion"`
	// License issuer.
	Issuer interface{} `json:"issuer" yaml:"issuer"`
	// License name.
	LicenseName *string `json:"licenseName" yaml:"licenseName"`
	// Product name.
	ProductName *string `json:"productName" yaml:"productName"`
	// Date and time range during which the license is valid, in ISO8601-UTC format.
	Validity interface{} `json:"validity" yaml:"validity"`
	// License beneficiary.
	Beneficiary *string `json:"beneficiary" yaml:"beneficiary"`
	// License metadata.
	LicenseMetadata interface{} `json:"licenseMetadata" yaml:"licenseMetadata"`
	// Product SKU.
	ProductSku *string `json:"productSku" yaml:"productSku"`
	// License status.
	Status *string `json:"status" yaml:"status"`
}

