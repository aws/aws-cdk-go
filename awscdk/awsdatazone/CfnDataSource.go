package awsdatazone

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DataZone::DataSource` resource specifies an Amazon DataZone data source that is used to import technical metadata of assets (data) from the source databases or data warehouses into Amazon DataZone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataSource := awscdk.Aws_datazone.NewCfnDataSource(this, jsii.String("MyCfnDataSource"), &CfnDataSourceProps{
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Name: jsii.String("name"),
//   	ProjectIdentifier: jsii.String("projectIdentifier"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AssetFormsInput: []interface{}{
//   		&FormInputProperty{
//   			FormName: jsii.String("formName"),
//
//   			// the properties below are optional
//   			Content: jsii.String("content"),
//   			TypeIdentifier: jsii.String("typeIdentifier"),
//   			TypeRevision: jsii.String("typeRevision"),
//   		},
//   	},
//   	Configuration: &DataSourceConfigurationInputProperty{
//   		GlueRunConfiguration: &GlueRunConfigurationInputProperty{
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//
//   					// the properties below are optional
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			AutoImportDataQualityResult: jsii.Boolean(false),
//   			CatalogName: jsii.String("catalogName"),
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   		},
//   		RedshiftRunConfiguration: &RedshiftRunConfigurationInputProperty{
//   			RelationalFilterConfigurations: []interface{}{
//   				&RelationalFilterConfigurationProperty{
//   					DatabaseName: jsii.String("databaseName"),
//
//   					// the properties below are optional
//   					FilterExpressions: []interface{}{
//   						&FilterExpressionProperty{
//   							Expression: jsii.String("expression"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					SchemaName: jsii.String("schemaName"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			DataAccessRole: jsii.String("dataAccessRole"),
//   			RedshiftCredentialConfiguration: &RedshiftCredentialConfigurationProperty{
//   				SecretManagerArn: jsii.String("secretManagerArn"),
//   			},
//   			RedshiftStorage: &RedshiftStorageProperty{
//   				RedshiftClusterSource: &RedshiftClusterStorageProperty{
//   					ClusterName: jsii.String("clusterName"),
//   				},
//   				RedshiftServerlessSource: &RedshiftServerlessStorageProperty{
//   					WorkgroupName: jsii.String("workgroupName"),
//   				},
//   			},
//   		},
//   		SageMakerRunConfiguration: &SageMakerRunConfigurationInputProperty{
//   			TrackingAssets: map[string][]*string{
//   				"trackingAssetsKey": []*string{
//   					jsii.String("trackingAssets"),
//   				},
//   			},
//   		},
//   	},
//   	ConnectionIdentifier: jsii.String("connectionIdentifier"),
//   	Description: jsii.String("description"),
//   	EnableSetting: jsii.String("enableSetting"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	PublishOnImport: jsii.Boolean(false),
//   	Recommendation: &RecommendationConfigurationProperty{
//   		EnableBusinessNameGeneration: jsii.Boolean(false),
//   	},
//   	Schedule: &ScheduleConfigurationProperty{
//   		Schedule: jsii.String("schedule"),
//   		Timezone: jsii.String("timezone"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-datasource.html
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The metadata forms attached to the assets that the data source works with.
	AssetFormsInput() interface{}
	SetAssetFormsInput(val interface{})
	// The connection ID that's part of the data source summary.
	AttrConnectionId() *string
	// The timestamp of when the data source was created.
	AttrCreatedAt() *string
	// The ID of the Amazon DataZone domain in which the data source exists.
	AttrDomainId() *string
	// The ID of the environment in which the data source exists.
	AttrEnvironmentId() *string
	// The identifier of the data source run.
	AttrId() *string
	// The count of the assets created during the last data source run.
	AttrLastRunAssetCount() awscdk.IResolvable
	// The timestamp of when the data source run was last performed.
	AttrLastRunAt() *string
	// The status of the last data source run.
	AttrLastRunStatus() *string
	// The project ID included in the data source run activity.
	AttrProjectId() *string
	// The status of the data source.
	AttrStatus() *string
	// The timestamp of when the data source was updated.
	AttrUpdatedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The configuration of the data source.
	Configuration() interface{}
	SetConfiguration(val interface{})
	// The unique identifier of a connection used to fetch relevant parameters from connection during Datasource run.
	ConnectionIdentifier() *string
	SetConnectionIdentifier(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the data source.
	Description() *string
	SetDescription(val *string)
	// The ID of the Amazon DataZone domain where the data source is created.
	DomainIdentifier() *string
	SetDomainIdentifier(val *string)
	// Specifies whether the data source is enabled.
	EnableSetting() *string
	SetEnableSetting(val *string)
	// The unique identifier of the Amazon DataZone environment to which the data source publishes assets.
	EnvironmentIdentifier() *string
	SetEnvironmentIdentifier(val *string)
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
	// The name of the data source.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The identifier of the Amazon DataZone project in which you want to add this data source.
	ProjectIdentifier() *string
	SetProjectIdentifier(val *string)
	// Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.
	PublishOnImport() interface{}
	SetPublishOnImport(val interface{})
	// Specifies whether the business name generation is to be enabled for this data source.
	Recommendation() interface{}
	SetRecommendation(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The schedule of the data source runs.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The type of the data source.
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) AssetFormsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetFormsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrDomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrLastRunAssetCount() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrLastRunAssetCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrLastRunAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastRunAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrLastRunStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastRunStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Configuration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ConnectionIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DomainIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) EnableSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enableSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) EnvironmentIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ProjectIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) PublishOnImport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishOnImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Recommendation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recommendation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	if err := validateNewCfnDataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_datazone.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_datazone.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource)SetAssetFormsInput(val interface{}) {
	if err := j.validateSetAssetFormsInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetFormsInput",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetConfiguration(val interface{}) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetConnectionIdentifier(val *string) {
	_jsii_.Set(
		j,
		"connectionIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetDomainIdentifier(val *string) {
	if err := j.validateSetDomainIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetEnableSetting(val *string) {
	_jsii_.Set(
		j,
		"enableSetting",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetEnvironmentIdentifier(val *string) {
	_jsii_.Set(
		j,
		"environmentIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetProjectIdentifier(val *string) {
	if err := j.validateSetProjectIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetPublishOnImport(val interface{}) {
	if err := j.validateSetPublishOnImportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishOnImport",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetRecommendation(val interface{}) {
	if err := j.validateSetRecommendationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendation",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetSchedule(val interface{}) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDataSource_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnDataSource",
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
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_datazone.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_datazone.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSource) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDataSource) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

