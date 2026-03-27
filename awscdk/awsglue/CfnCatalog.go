package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a catalog in the Glue Data Catalog.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCatalog := awscdk.Aws_glue.NewCfnCatalog(this, jsii.String("MyCfnCatalog"), &CfnCatalogProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   	CatalogProperties: &CatalogPropertiesProperty{
//   		CustomProperties: map[string]*string{
//   			"customPropertiesKey": jsii.String("customProperties"),
//   		},
//   		DataLakeAccessProperties: &DataLakeAccessPropertiesProperty{
//   			AllowFullTableExternalDataAccess: jsii.String("allowFullTableExternalDataAccess"),
//   			CatalogType: jsii.String("catalogType"),
//   			DataLakeAccess: jsii.Boolean(false),
//   			DataTransferRole: jsii.String("dataTransferRole"),
//   			KmsKey: jsii.String("kmsKey"),
//   			ManagedWorkgroupName: jsii.String("managedWorkgroupName"),
//   			ManagedWorkgroupStatus: jsii.String("managedWorkgroupStatus"),
//   			RedshiftDatabaseName: jsii.String("redshiftDatabaseName"),
//   		},
//   	},
//   	CreateDatabaseDefaultPermissions: []interface{}{
//   		&PrincipalPermissionsProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	CreateTableDefaultPermissions: []interface{}{
//   		&PrincipalPermissionsProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FederatedCatalog: &FederatedCatalogProperty{
//   		ConnectionName: jsii.String("connectionName"),
//   		Identifier: jsii.String("identifier"),
//   	},
//   	OverwriteChildResourcePermissionsWithDefault: jsii.String("overwriteChildResourcePermissionsWithDefault"),
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetRedshiftCatalog: &TargetRedshiftCatalogProperty{
//   		CatalogArn: jsii.String("catalogArn"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-catalog.html
//
type CfnCatalog interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsglue.ICatalogRef
	awscdk.ITaggableV2
	// Allows third-party engines to access data in Amazon S3 locations that are registered with Lake Formation.
	AllowFullTableExternalDataAccess() *string
	SetAllowFullTableExternalDataAccess(val *string)
	// The ID of the catalog.
	AttrCatalogId() *string
	// Additional key-value properties for the catalog.
	AttrCatalogPropertiesCustomProperties() awscdk.IResolvable
	// The name of the managed workgroup associated with the catalog.
	AttrCatalogPropertiesDataLakeAccessPropertiesManagedWorkgroupName() *string
	// The status of the managed workgroup.
	AttrCatalogPropertiesDataLakeAccessPropertiesManagedWorkgroupStatus() *string
	// The name of the Redshift database.
	AttrCatalogPropertiesDataLakeAccessPropertiesRedshiftDatabaseName() *string
	// The time at which the catalog was created.
	AttrCreateTime() *float64
	// The Amazon Resource Name (ARN) of the catalog.
	AttrResourceArn() *string
	// The time at which the catalog was last updated.
	AttrUpdateTime() *float64
	// A structure that specifies data lake access properties and other custom properties.
	CatalogProperties() interface{}
	SetCatalogProperties(val interface{})
	// A reference to a Catalog resource.
	CatalogRef() *interfacesawsglue.CatalogReference
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// An array of PrincipalPermissions objects for default database permissions.
	CreateDatabaseDefaultPermissions() interface{}
	SetCreateDatabaseDefaultPermissions(val interface{})
	// An array of PrincipalPermissions objects for default table permissions.
	CreateTableDefaultPermissions() interface{}
	SetCreateTableDefaultPermissions(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description of the catalog.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// A FederatedCatalog structure that references an entity outside the Glue Data Catalog.
	FederatedCatalog() interface{}
	SetFederatedCatalog(val interface{})
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
	// The name of the catalog to create.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Specifies whether to overwrite child resource permissions with the default permissions.
	OverwriteChildResourcePermissionsWithDefault() *string
	SetOverwriteChildResourcePermissionsWithDefault(val *string)
	// A map of key-value pairs that define parameters and properties of the catalog.
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
	// An array of key-value pairs to apply to this resource.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// A structure that describes a target catalog for resource linking.
	TargetRedshiftCatalog() interface{}
	SetTargetRedshiftCatalog(val interface{})
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnCatalog
type jsiiProxy_CfnCatalog struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsglueICatalogRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnCatalog) AllowFullTableExternalDataAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrCatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCatalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrCatalogPropertiesCustomProperties() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCatalogPropertiesCustomProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrCatalogPropertiesDataLakeAccessPropertiesManagedWorkgroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCatalogPropertiesDataLakeAccessPropertiesManagedWorkgroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrCatalogPropertiesDataLakeAccessPropertiesManagedWorkgroupStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCatalogPropertiesDataLakeAccessPropertiesManagedWorkgroupStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrCatalogPropertiesDataLakeAccessPropertiesRedshiftDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCatalogPropertiesDataLakeAccessPropertiesRedshiftDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrCreateTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) AttrUpdateTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrUpdateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CatalogProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CatalogRef() *interfacesawsglue.CatalogReference {
	var returns *interfacesawsglue.CatalogReference
	_jsii_.Get(
		j,
		"catalogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CreateDatabaseDefaultPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CreateTableDefaultPermissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createTableDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) FederatedCatalog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"federatedCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) OverwriteChildResourcePermissionsWithDefault() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteChildResourcePermissionsWithDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) TargetRedshiftCatalog() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetRedshiftCatalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCatalog) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Catalog`.
func NewCfnCatalog(scope constructs.Construct, id *string, props *CfnCatalogProps) CfnCatalog {
	_init_.Initialize()

	if err := validateNewCfnCatalogParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCatalog{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnCatalog",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Catalog`.
func NewCfnCatalog_Override(c CfnCatalog, scope constructs.Construct, id *string, props *CfnCatalogProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnCatalog",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCatalog)SetAllowFullTableExternalDataAccess(val *string) {
	_jsii_.Set(
		j,
		"allowFullTableExternalDataAccess",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetCatalogProperties(val interface{}) {
	if err := j.validateSetCatalogPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogProperties",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetCreateDatabaseDefaultPermissions(val interface{}) {
	if err := j.validateSetCreateDatabaseDefaultPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabaseDefaultPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetCreateTableDefaultPermissions(val interface{}) {
	if err := j.validateSetCreateTableDefaultPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTableDefaultPermissions",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetFederatedCatalog(val interface{}) {
	if err := j.validateSetFederatedCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federatedCatalog",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetOverwriteChildResourcePermissionsWithDefault(val *string) {
	_jsii_.Set(
		j,
		"overwriteChildResourcePermissionsWithDefault",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetParameters(val interface{}) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnCatalog)SetTargetRedshiftCatalog(val interface{}) {
	if err := j.validateSetTargetRedshiftCatalogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRedshiftCatalog",
		val,
	)
}

// Checks whether the given object is a CfnCatalog.
func CfnCatalog_IsCfnCatalog(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCatalog_IsCfnCatalogParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCatalog",
		"isCfnCatalog",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCatalog_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCatalog_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCatalog",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCatalog_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCatalog_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCatalog",
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
func CfnCatalog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCatalog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnCatalog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCatalog_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnCatalog",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCatalog) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCatalog) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCatalog) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCatalog) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCatalog) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCatalog) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCatalog) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCatalog) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCatalog) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCatalog) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCatalog) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCatalog) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCatalog) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCatalog) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCatalog) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCatalog) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCatalog) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCatalog) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCatalog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCatalog) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnCatalog) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

