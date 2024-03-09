package awsiotsitewise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an asset model from specified property and hierarchy definitions.
//
// You create assets from asset models. With asset models, you can easily create assets of the same type that have standardized definitions. Each asset created from a model inherits the asset model's property and hierarchy definitions. For more information, see [Defining asset models](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetModel := awscdk.Aws_iotsitewise.NewCfnAssetModel(this, jsii.String("MyCfnAssetModel"), &CfnAssetModelProps{
//   	AssetModelName: jsii.String("assetModelName"),
//
//   	// the properties below are optional
//   	AssetModelCompositeModels: []interface{}{
//   		&AssetModelCompositeModelProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			ComposedAssetModelId: jsii.String("composedAssetModelId"),
//   			CompositeModelProperties: []interface{}{
//   				&AssetModelPropertyProperty{
//   					DataType: jsii.String("dataType"),
//   					Name: jsii.String("name"),
//   					Type: &PropertyTypeProperty{
//   						TypeName: jsii.String("typeName"),
//
//   						// the properties below are optional
//   						Attribute: &AttributeProperty{
//   							DefaultValue: jsii.String("defaultValue"),
//   						},
//   						Metric: &MetricProperty{
//   							Expression: jsii.String("expression"),
//   							Variables: []interface{}{
//   								&ExpressionVariableProperty{
//   									Name: jsii.String("name"),
//   									Value: &VariableValueProperty{
//   										HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   										HierarchyId: jsii.String("hierarchyId"),
//   										HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   										PropertyExternalId: jsii.String("propertyExternalId"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyLogicalId: jsii.String("propertyLogicalId"),
//   										PropertyPath: []interface{}{
//   											&PropertyPathDefinitionProperty{
//   												Name: jsii.String("name"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							Window: &MetricWindowProperty{
//   								Tumbling: &TumblingWindowProperty{
//   									Interval: jsii.String("interval"),
//
//   									// the properties below are optional
//   									Offset: jsii.String("offset"),
//   								},
//   							},
//   						},
//   						Transform: &TransformProperty{
//   							Expression: jsii.String("expression"),
//   							Variables: []interface{}{
//   								&ExpressionVariableProperty{
//   									Name: jsii.String("name"),
//   									Value: &VariableValueProperty{
//   										HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   										HierarchyId: jsii.String("hierarchyId"),
//   										HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   										PropertyExternalId: jsii.String("propertyExternalId"),
//   										PropertyId: jsii.String("propertyId"),
//   										PropertyLogicalId: jsii.String("propertyLogicalId"),
//   										PropertyPath: []interface{}{
//   											&PropertyPathDefinitionProperty{
//   												Name: jsii.String("name"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					DataTypeSpec: jsii.String("dataTypeSpec"),
//   					ExternalId: jsii.String("externalId"),
//   					Id: jsii.String("id"),
//   					LogicalId: jsii.String("logicalId"),
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			ParentAssetModelCompositeModelExternalId: jsii.String("parentAssetModelCompositeModelExternalId"),
//   			Path: []*string{
//   				jsii.String("path"),
//   			},
//   		},
//   	},
//   	AssetModelDescription: jsii.String("assetModelDescription"),
//   	AssetModelExternalId: jsii.String("assetModelExternalId"),
//   	AssetModelHierarchies: []interface{}{
//   		&AssetModelHierarchyProperty{
//   			ChildAssetModelId: jsii.String("childAssetModelId"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			LogicalId: jsii.String("logicalId"),
//   		},
//   	},
//   	AssetModelProperties: []interface{}{
//   		&AssetModelPropertyProperty{
//   			DataType: jsii.String("dataType"),
//   			Name: jsii.String("name"),
//   			Type: &PropertyTypeProperty{
//   				TypeName: jsii.String("typeName"),
//
//   				// the properties below are optional
//   				Attribute: &AttributeProperty{
//   					DefaultValue: jsii.String("defaultValue"),
//   				},
//   				Metric: &MetricProperty{
//   					Expression: jsii.String("expression"),
//   					Variables: []interface{}{
//   						&ExpressionVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: &VariableValueProperty{
//   								HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   								HierarchyId: jsii.String("hierarchyId"),
//   								HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   								PropertyExternalId: jsii.String("propertyExternalId"),
//   								PropertyId: jsii.String("propertyId"),
//   								PropertyLogicalId: jsii.String("propertyLogicalId"),
//   								PropertyPath: []interface{}{
//   									&PropertyPathDefinitionProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					Window: &MetricWindowProperty{
//   						Tumbling: &TumblingWindowProperty{
//   							Interval: jsii.String("interval"),
//
//   							// the properties below are optional
//   							Offset: jsii.String("offset"),
//   						},
//   					},
//   				},
//   				Transform: &TransformProperty{
//   					Expression: jsii.String("expression"),
//   					Variables: []interface{}{
//   						&ExpressionVariableProperty{
//   							Name: jsii.String("name"),
//   							Value: &VariableValueProperty{
//   								HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   								HierarchyId: jsii.String("hierarchyId"),
//   								HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   								PropertyExternalId: jsii.String("propertyExternalId"),
//   								PropertyId: jsii.String("propertyId"),
//   								PropertyLogicalId: jsii.String("propertyLogicalId"),
//   								PropertyPath: []interface{}{
//   									&PropertyPathDefinitionProperty{
//   										Name: jsii.String("name"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			DataTypeSpec: jsii.String("dataTypeSpec"),
//   			ExternalId: jsii.String("externalId"),
//   			Id: jsii.String("id"),
//   			LogicalId: jsii.String("logicalId"),
//   			Unit: jsii.String("unit"),
//   		},
//   	},
//   	AssetModelType: jsii.String("assetModelType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-assetmodel.html
//
type CfnAssetModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The composite asset models that are part of this asset model.
	AssetModelCompositeModels() interface{}
	SetAssetModelCompositeModels(val interface{})
	// A description for the asset model.
	AssetModelDescription() *string
	SetAssetModelDescription(val *string)
	// The external ID of the asset model.
	AssetModelExternalId() *string
	SetAssetModelExternalId(val *string)
	// The hierarchy definitions of the asset model.
	AssetModelHierarchies() interface{}
	SetAssetModelHierarchies(val interface{})
	// A unique, friendly name for the asset model.
	AssetModelName() *string
	SetAssetModelName(val *string)
	// The property definitions of the asset model.
	AssetModelProperties() interface{}
	SetAssetModelProperties(val interface{})
	// The type of the asset model (ASSET_MODEL OR COMPONENT_MODEL).
	AssetModelType() *string
	SetAssetModelType(val *string)
	// The ARN of the asset model, which has the following format.
	AttrAssetModelArn() *string
	// The ID of the asset model.
	AttrAssetModelId() *string
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
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// A list of key-value pairs that contain metadata for the asset.
	TagsRaw() *[]*awscdk.CfnTag
	SetTagsRaw(val *[]*awscdk.CfnTag)
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

// The jsii proxy struct for CfnAssetModel
type jsiiProxy_CfnAssetModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnAssetModel) AssetModelCompositeModels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelCompositeModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelHierarchies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelHierarchies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AssetModelType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AttrAssetModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssetModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) AttrAssetModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAssetModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetModel) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnAssetModel(scope constructs.Construct, id *string, props *CfnAssetModelProps) CfnAssetModel {
	_init_.Initialize()

	if err := validateNewCfnAssetModelParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssetModel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iotsitewise.CfnAssetModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnAssetModel_Override(c CfnAssetModel, scope constructs.Construct, id *string, props *CfnAssetModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iotsitewise.CfnAssetModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelCompositeModels(val interface{}) {
	if err := j.validateSetAssetModelCompositeModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetModelCompositeModels",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelDescription(val *string) {
	_jsii_.Set(
		j,
		"assetModelDescription",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelExternalId(val *string) {
	_jsii_.Set(
		j,
		"assetModelExternalId",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelHierarchies(val interface{}) {
	if err := j.validateSetAssetModelHierarchiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetModelHierarchies",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelName(val *string) {
	if err := j.validateSetAssetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetModelName",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelProperties(val interface{}) {
	if err := j.validateSetAssetModelPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetModelProperties",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetAssetModelType(val *string) {
	_jsii_.Set(
		j,
		"assetModelType",
		val,
	)
}

func (j *jsiiProxy_CfnAssetModel)SetTagsRaw(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAssetModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetModel_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotsitewise.CfnAssetModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnAssetModel_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetModel_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotsitewise.CfnAssetModel",
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
func CfnAssetModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iotsitewise.CfnAssetModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssetModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_iotsitewise.CfnAssetModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssetModel) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAssetModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAssetModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAssetModel) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnAssetModel) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAssetModel) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAssetModel) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAssetModel) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAssetModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnAssetModel) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnAssetModel) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssetModel) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

