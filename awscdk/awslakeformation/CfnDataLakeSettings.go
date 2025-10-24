package awslakeformation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::LakeFormation::DataLakeSettings` resource is an AWS Lake Formation resource type that manages the data lake settings for your account.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdkgluealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack Stack
//   var accountId string
//
//
//   tagKey := "aws"
//   tagValues := []*string{
//   	"dev",
//   }
//
//   database := awscdkgluealpha.NewDatabase(this, jsii.String("Database"))
//
//   table := awscdkgluealpha.NewS3Table(this, jsii.String("Table"), &S3TableProps{
//   	Database: Database,
//   	Columns: []Column{
//   		&Column{
//   			Name: jsii.String("col1"),
//   			Type: awscdkgluealpha.Schema_STRING(),
//   		},
//   		&Column{
//   			Name: jsii.String("col2"),
//   			Type: awscdkgluealpha.Schema_STRING(),
//   		},
//   	},
//   	DataFormat: awscdkgluealpha.DataFormat_CSV(),
//   })
//
//   synthesizer := stack.Synthesizer.(DefaultStackSynthesizer)
//   awscdk.NewCfnDataLakeSettings(this, jsii.String("DataLakeSettings"), &CfnDataLakeSettingsProps{
//   	Admins: []interface{}{
//   		&DataLakePrincipalProperty{
//   			DataLakePrincipalIdentifier: stack.FormatArn(&ArnComponents{
//   				Service: jsii.String("iam"),
//   				Resource: jsii.String("role"),
//   				Region: jsii.String(""),
//   				Account: accountId,
//   				ResourceName: jsii.String("Admin"),
//   			}),
//   		},
//   		&DataLakePrincipalProperty{
//   			// The CDK cloudformation execution role.
//   			DataLakePrincipalIdentifier: synthesizer.cloudFormationExecutionRoleArn.replace(jsii.String("${AWS::Partition}"), jsii.String("aws")),
//   		},
//   	},
//   })
//
//   tag := awscdk.NewCfnTag(this, jsii.String("Tag"), &CfnTagProps{
//   	CatalogId: accountId,
//   	TagKey: jsii.String(TagKey),
//   	TagValues: TagValues,
//   })
//
//   lfTagPairProperty := &LFTagPairProperty{
//   	CatalogId: accountId,
//   	TagKey: jsii.String(TagKey),
//   	TagValues: TagValues,
//   }
//
//   tagAssociation := awscdk.NewCfnTagAssociation(this, jsii.String("TagAssociation"), &CfnTagAssociationProps{
//   	LfTags: []interface{}{
//   		lfTagPairProperty,
//   	},
//   	Resource: &ResourceProperty{
//   		TableWithColumns: &TableWithColumnsResourceProperty{
//   			DatabaseName: database.DatabaseName,
//   			ColumnNames: []*string{
//   				jsii.String("col1"),
//   				jsii.String("col2"),
//   			},
//   			CatalogId: accountId,
//   			Name: table.TableName,
//   		},
//   	},
//   })
//
//   tagAssociation.Node.AddDependency(tag)
//   tagAssociation.Node.AddDependency(table)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-datalakesettings.html
//
type CfnDataLakeSettings interface {
	awscdk.CfnResource
	IDataLakeSettingsRef
	awscdk.IInspectable
	// A list of AWS Lake Formation principals.
	Admins() interface{}
	SetAdmins(val interface{})
	// Whether to allow Amazon EMR clusters or other third-party query engines to access data managed by Lake Formation .
	AllowExternalDataFiltering() interface{}
	SetAllowExternalDataFiltering(val interface{})
	// Specifies whether query engines and applications can get credentials without IAM session tags if the user has full table access.
	AllowFullTableExternalDataAccess() interface{}
	SetAllowFullTableExternalDataAccess(val interface{})
	AttrId() *string
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	AuthorizedSessionTagValueList() *[]*string
	SetAuthorizedSessionTagValueList(val *[]*string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Specifies whether access control on a newly created database is managed by Lake Formation permissions or exclusively by IAM permissions.
	CreateDatabaseDefaultPermissions() interface{}
	SetCreateDatabaseDefaultPermissions(val interface{})
	// Specifies whether access control on a newly created table is managed by Lake Formation permissions or exclusively by IAM permissions.
	CreateTableDefaultPermissions() interface{}
	SetCreateTableDefaultPermissions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A reference to a DataLakeSettings resource.
	DataLakeSettingsRef() *DataLakeSettingsReference
	// A list of the account IDs of AWS accounts with Amazon EMR clusters or third-party engines that are allwed to perform data filtering.
	ExternalDataFilteringAllowList() interface{}
	SetExternalDataFilteringAllowList(val interface{})
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
	// Specifies whether the data lake settings are updated by adding new values to the current settings ( `APPEND` ) or by replacing the current settings with new settings ( `REPLACE` ).
	MutationType() *string
	SetMutationType(val *string)
	// The tree node.
	Node() constructs.Node
	// A key-value map that provides an additional configuration on your data lake.
	Parameters() interface{}
	SetParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of UTF-8 strings.
	TrustedResourceOwners() *[]*string
	SetTrustedResourceOwners(val *[]*string)
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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
	jsiiProxy_IDataLakeSettingsRef
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

func (j *jsiiProxy_CfnDataLakeSettings) AllowExternalDataFiltering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalDataFiltering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) AllowFullTableExternalDataAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) AuthorizedSessionTagValueList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedSessionTagValueList",
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

func (j *jsiiProxy_CfnDataLakeSettings) CreateDatabaseDefaultPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) CreateTableDefaultPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createTableDefaultPermissions",
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

func (j *jsiiProxy_CfnDataLakeSettings) DataLakeSettingsRef() *DataLakeSettingsReference {
	var returns *DataLakeSettingsReference
	_jsii_.Get(
		j,
		"dataLakeSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataLakeSettings) ExternalDataFilteringAllowList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalDataFilteringAllowList",
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

func (j *jsiiProxy_CfnDataLakeSettings) MutationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutationType",
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

func (j *jsiiProxy_CfnDataLakeSettings) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
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

func (j *jsiiProxy_CfnDataLakeSettings) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnDataLakeSettings(scope constructs.Construct, id *string, props *CfnDataLakeSettingsProps) CfnDataLakeSettings {
	_init_.Initialize()

	if err := validateNewCfnDataLakeSettingsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataLakeSettings{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDataLakeSettings_Override(c CfnDataLakeSettings, scope constructs.Construct, id *string, props *CfnDataLakeSettingsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetAdmins(val interface{}) {
	if err := j.validateSetAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admins",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetAllowExternalDataFiltering(val interface{}) {
	if err := j.validateSetAllowExternalDataFilteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExternalDataFiltering",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetAllowFullTableExternalDataAccess(val interface{}) {
	if err := j.validateSetAllowFullTableExternalDataAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFullTableExternalDataAccess",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetAuthorizedSessionTagValueList(val *[]*string) {
	_jsii_.Set(
		j,
		"authorizedSessionTagValueList",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetCreateDatabaseDefaultPermissions(val interface{}) {
	if err := j.validateSetCreateDatabaseDefaultPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabaseDefaultPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetCreateTableDefaultPermissions(val interface{}) {
	if err := j.validateSetCreateTableDefaultPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTableDefaultPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetExternalDataFilteringAllowList(val interface{}) {
	if err := j.validateSetExternalDataFilteringAllowListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalDataFilteringAllowList",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetMutationType(val *string) {
	_jsii_.Set(
		j,
		"mutationType",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataLakeSettings)SetTrustedResourceOwners(val *[]*string) {
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

	if err := validateCfnDataLakeSettings_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDataLakeSettings_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataLakeSettings_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lakeformation.CfnDataLakeSettings",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataLakeSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataLakeSettings) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataLakeSettings) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

