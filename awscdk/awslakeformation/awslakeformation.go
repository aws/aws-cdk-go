package awslakeformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::LakeFormation::DataLakeSettings`.
//
// The `AWS::LakeFormation::DataLakeSettings` resource is an AWS Lake Formation resource type that manages the data lake settings for your account. Note that the CloudFormation template only supports updating the `Admins` list. It does not support updating the [CreateDatabaseDefaultPermissions](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-settings.html#aws-lake-formation-api-aws-lake-formation-api-settings-DataLakeSettings) or [CreateTableDefaultPermissions](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-settings.html#aws-lake-formation-api-aws-lake-formation-api-settings-DataLakeSettings) . Those permissions can only be edited in the DataLakeSettings resource via the API.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   cfnDataLakeSettings := lakeformation.NewCfnDataLakeSettings(this, jsii.String("MyCfnDataLakeSettings"), &cfnDataLakeSettingsProps{
//   	admins: []interface{}{
//   		&dataLakePrincipalProperty{
//   			dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	trustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   })
//
type CfnDataLakeSettings interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// A list of AWS Lake Formation principals.
	Admins() interface{}
	SetAdmins(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::LakeFormation::DataLakeSettings.TrustedResourceOwners`.
	TrustedResourceOwners() *[]*string
	SetTrustedResourceOwners(val *[]*string)
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

// The jsii proxy struct for CfnDataLakeSettings
type jsiiProxy_CfnDataLakeSettings struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataLakeSettings) Admins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"admins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) TrustedResourceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::DataLakeSettings`.
func NewCfnDataLakeSettings(scope constructs.Construct, id *string, props *CfnDataLakeSettingsProps) CfnDataLakeSettings {
	_init_.Initialize()

	j := jsiiProxy_CfnDataLakeSettings{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::DataLakeSettings`.
func NewCfnDataLakeSettings_Override(c CfnDataLakeSettings, scope constructs.Construct, id *string, props *CfnDataLakeSettingsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings) SetAdmins(val interface{}) {
	_jsii_.Set(
		j,
		"admins",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings) SetTrustedResourceOwners(val *[]*string) {
	_jsii_.Set(
		j,
		"trustedResourceOwners",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataLakeSettings_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataLakeSettings_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
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
func CfnDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataLakeSettings_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The Lake Formation principal.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   dataLakePrincipalProperty := &dataLakePrincipalProperty{
//   	dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnDataLakeSettings_DataLakePrincipalProperty struct {
	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string `json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

// Properties for defining a `CfnDataLakeSettings`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   cfnDataLakeSettingsProps := &cfnDataLakeSettingsProps{
//   	admins: []interface{}{
//   		&dataLakePrincipalProperty{
//   			dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   		},
//   	},
//   	trustedResourceOwners: []*string{
//   		jsii.String("trustedResourceOwners"),
//   	},
//   }
//
type CfnDataLakeSettingsProps struct {
	// A list of AWS Lake Formation principals.
	Admins interface{} `json:"admins" yaml:"admins"`
	// `AWS::LakeFormation::DataLakeSettings.TrustedResourceOwners`.
	TrustedResourceOwners *[]*string `json:"trustedResourceOwners" yaml:"trustedResourceOwners"`
}

// A CloudFormation `AWS::LakeFormation::Permissions`.
//
// The `AWS::LakeFormation::Permissions` resource represents the permissions that a principal has on an AWS Glue Data Catalog resource (such as AWS Glue database or AWS Glue tables). When you upload a permissions stack, the permissions are granted to the principal and when you remove the stack, the permissions are revoked from the principal. If you remove a stack, and the principal does not have the permissions referenced in the stack then AWS Lake Formation will throw an error because you can’t call revoke on non-existing permissions. To successfully remove the stack, you’ll need to regrant those permissions and then remove the stack.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   cfnPermissions := lakeformation.NewCfnPermissions(this, jsii.String("MyCfnPermissions"), &cfnPermissionsProps{
//   	dataLakePrincipal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		databaseResource: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataLocationResource: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			s3Resource: jsii.String("s3Resource"),
//   		},
//   		tableResource: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableWildcard: &tableWildcardProperty{
//   			},
//   		},
//   		tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   })
//
type CfnPermissions interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The AWS Lake Formation principal.
	DataLakePrincipal() interface{}
	SetDataLakePrincipal(val interface{})
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
	// The permissions granted or revoked.
	Permissions() *[]*string
	SetPermissions(val *[]*string)
	// Indicates whether to grant the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption() *[]*string
	SetPermissionsWithGrantOption(val *[]*string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A structure for the resource.
	Resource() interface{}
	SetResource(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnPermissions
type jsiiProxy_CfnPermissions struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPermissions) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) DataLakePrincipal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLakePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Resource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissions) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::Permissions`.
func NewCfnPermissions(scope constructs.Construct, id *string, props *CfnPermissionsProps) CfnPermissions {
	_init_.Initialize()

	j := jsiiProxy_CfnPermissions{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnPermissions",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::Permissions`.
func NewCfnPermissions_Override(c CfnPermissions, scope constructs.Construct, id *string, props *CfnPermissionsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnPermissions",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPermissions) SetDataLakePrincipal(val interface{}) {
	_jsii_.Set(
		j,
		"dataLakePrincipal",
		val,
	)
}

func (j *jsiiProxy_CfnPermissions) SetPermissions(val *[]*string) {
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnPermissions) SetPermissionsWithGrantOption(val *[]*string) {
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_CfnPermissions) SetResource(val interface{}) {
	_jsii_.Set(
		j,
		"resource",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnPermissions_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnPermissions",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPermissions_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnPermissions",
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
func CfnPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermissions_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lakeformation.CfnPermissions",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissions) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnPermissions) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnPermissions) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnPermissions) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnPermissions) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnPermissions) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnPermissions) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnPermissions) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnPermissions) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPermissions) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissions) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A wildcard object, consisting of an optional list of excluded column names or indexes.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   columnWildcardProperty := &columnWildcardProperty{
//   	excludedColumnNames: []*string{
//   		jsii.String("excludedColumnNames"),
//   	},
//   }
//
type CfnPermissions_ColumnWildcardProperty struct {
	// Excludes column names.
	//
	// Any column with this name will be excluded.
	ExcludedColumnNames *[]*string `json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

// The Lake Formation principal.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   dataLakePrincipalProperty := &dataLakePrincipalProperty{
//   	dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnPermissions_DataLakePrincipalProperty struct {
	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string `json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

// A structure for a data location object where permissions are granted or revoked.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   dataLocationResourceProperty := &dataLocationResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	s3Resource: jsii.String("s3Resource"),
//   }
//
type CfnPermissions_DataLocationResourceProperty struct {
	// `CfnPermissions.DataLocationResourceProperty.CatalogId`.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// Currently not supported by AWS CloudFormation .
	S3Resource *string `json:"s3Resource" yaml:"s3Resource"`
}

// A structure for the database object.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   databaseResourceProperty := &databaseResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	name: jsii.String("name"),
//   }
//
type CfnPermissions_DatabaseResourceProperty struct {
	// `CfnPermissions.DatabaseResourceProperty.CatalogId`.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the database resource.
	//
	// Unique to the Data Catalog.
	Name *string `json:"name" yaml:"name"`
}

// A structure for the resource.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   resourceProperty := &resourceProperty{
//   	databaseResource: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	dataLocationResource: &dataLocationResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		s3Resource: jsii.String("s3Resource"),
//   	},
//   	tableResource: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   		tableWildcard: &tableWildcardProperty{
//   		},
//   	},
//   	tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		columnWildcard: &columnWildcardProperty{
//   			excludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnPermissions_ResourceProperty struct {
	// A structure for the database object.
	DatabaseResource interface{} `json:"databaseResource" yaml:"databaseResource"`
	// A structure for a data location object where permissions are granted or revoked.
	DataLocationResource interface{} `json:"dataLocationResource" yaml:"dataLocationResource"`
	// A structure for the table object.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	TableResource interface{} `json:"tableResource" yaml:"tableResource"`
	// Currently not supported by AWS CloudFormation .
	TableWithColumnsResource interface{} `json:"tableWithColumnsResource" yaml:"tableWithColumnsResource"`
}

// A structure for the table object.
//
// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableWildcard: &tableWildcardProperty{
//   	},
//   }
//
type CfnPermissions_TableResourceProperty struct {
	// `CfnPermissions.TableResourceProperty.CatalogId`.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	Name *string `json:"name" yaml:"name"`
	// An empty object representing all tables under a database.
	//
	// If this field is specified instead of the `Name` field, all tables under `DatabaseName` will have permission changes applied.
	TableWildcard interface{} `json:"tableWildcard" yaml:"tableWildcard"`
}

// A wildcard object representing every table under a database.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   tableWildcardProperty := &tableWildcardProperty{
//   }
//
type CfnPermissions_TableWildcardProperty struct {
}

// A structure for a table with columns object. This object is only used when granting a SELECT permission.
//
// This object must take a value for at least one of `ColumnsNames` , `ColumnsIndexes` , or `ColumnsWildcard` .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	columnWildcard: &columnWildcardProperty{
//   		excludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   }
//
type CfnPermissions_TableWithColumnsResourceProperty struct {
	// `CfnPermissions.TableWithColumnsResourceProperty.CatalogId`.
	CatalogId *string `json:"catalogId" yaml:"catalogId"`
	// The list of column names for the table.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnNames *[]*string `json:"columnNames" yaml:"columnNames"`
	// A wildcard specified by a `ColumnWildcard` object.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnWildcard interface{} `json:"columnWildcard" yaml:"columnWildcard"`
	// The name of the database for the table with columns resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The name of the table resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	Name *string `json:"name" yaml:"name"`
}

// Properties for defining a `CfnPermissions`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   cfnPermissionsProps := &cfnPermissionsProps{
//   	dataLakePrincipal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		databaseResource: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataLocationResource: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			s3Resource: jsii.String("s3Resource"),
//   		},
//   		tableResource: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableWildcard: &tableWildcardProperty{
//   			},
//   		},
//   		tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   }
//
type CfnPermissionsProps struct {
	// The AWS Lake Formation principal.
	DataLakePrincipal interface{} `json:"dataLakePrincipal" yaml:"dataLakePrincipal"`
	// A structure for the resource.
	Resource interface{} `json:"resource" yaml:"resource"`
	// The permissions granted or revoked.
	Permissions *[]*string `json:"permissions" yaml:"permissions"`
	// Indicates whether to grant the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption *[]*string `json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
}

// A CloudFormation `AWS::LakeFormation::Resource`.
//
// The `AWS::LakeFormation::Resource` represents the data (Amazon S3 buckets and folders) that is being registered with AWS Lake Formation . When a `Resource` type CloudFormation template is uploaded, an AWS Lake Formation [`RegisterResource`](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-credential-vending.html#aws-lake-formation-api-credential-vending-RegisterResource) API call is made to register the resource. When a `Resource` type CloudFormation template is removed, the AWS Lake Formation [`DeregisterResource`](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-credential-vending.html#aws-lake-formation-api-credential-vending-DeregisterResource) API is called.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   cfnResource := lakeformation.NewCfnResource(this, jsii.String("MyCfnResource"), &cfnResourceProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	useServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   })
//
type CfnResource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn() *string
	SetResourceArn(val *string)
	// The IAM role that registered a resource.
	RoleArn() *string
	SetRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.
	UseServiceLinkedRole() interface{}
	SetUseServiceLinkedRole(val interface{})
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

// The jsii proxy struct for CfnResource
type jsiiProxy_CfnResource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResource) UseServiceLinkedRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useServiceLinkedRole",
		&returns,
	)
	return returns
}


// Create a new `AWS::LakeFormation::Resource`.
func NewCfnResource(scope constructs.Construct, id *string, props *CfnResourceProps) CfnResource {
	_init_.Initialize()

	j := jsiiProxy_CfnResource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::LakeFormation::Resource`.
func NewCfnResource_Override(c CfnResource, scope constructs.Construct, id *string, props *CfnResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnResource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResource) SetResourceArn(val *string) {
	_jsii_.Set(
		j,
		"resourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnResource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnResource) SetUseServiceLinkedRole(val interface{}) {
	_jsii_.Set(
		j,
		"useServiceLinkedRole",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnResource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnResource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnResource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnResource",
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
func CfnResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lakeformation.CfnResource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResource) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnResource) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnResource) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnResource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnResource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnResource) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnResource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnResource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnResource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnResource`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import lakeformation "github.com/aws/aws-cdk-go/awscdk/aws_lakeformation"
//   cfnResourceProps := &cfnResourceProps{
//   	resourceArn: jsii.String("resourceArn"),
//   	useServiceLinkedRole: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnResourceProps struct {
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `json:"resourceArn" yaml:"resourceArn"`
	// Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.
	UseServiceLinkedRole interface{} `json:"useServiceLinkedRole" yaml:"useServiceLinkedRole"`
	// The IAM role that registered a resource.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
}

