package awsquicksight

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a dataset.
//
// This operation doesn't support datasets that include uploaded files as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagRuleConfigurations interface{}
//
//   cfnDataSet := awscdk.Aws_quicksight.NewCfnDataSet(this, jsii.String("MyCfnDataSet"), &CfnDataSetProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	ColumnGroups: []interface{}{
//   		&ColumnGroupProperty{
//   			GeoSpatialColumnGroup: &GeoSpatialColumnGroupProperty{
//   				Columns: []*string{
//   					jsii.String("columns"),
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CountryCode: jsii.String("countryCode"),
//   			},
//   		},
//   	},
//   	ColumnLevelPermissionRules: []interface{}{
//   		&ColumnLevelPermissionRuleProperty{
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			Principals: []*string{
//   				jsii.String("principals"),
//   			},
//   		},
//   	},
//   	DataSetId: jsii.String("dataSetId"),
//   	DatasetParameters: []interface{}{
//   		&DatasetParameterProperty{
//   			DateTimeDatasetParameter: &DateTimeDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &DateTimeDatasetParameterDefaultValuesProperty{
//   					StaticValues: []*string{
//   						jsii.String("staticValues"),
//   					},
//   				},
//   				TimeGranularity: jsii.String("timeGranularity"),
//   			},
//   			DecimalDatasetParameter: &DecimalDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &DecimalDatasetParameterDefaultValuesProperty{
//   					StaticValues: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   			IntegerDatasetParameter: &IntegerDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &IntegerDatasetParameterDefaultValuesProperty{
//   					StaticValues: []interface{}{
//   						jsii.Number(123),
//   					},
//   				},
//   			},
//   			StringDatasetParameter: &StringDatasetParameterProperty{
//   				Id: jsii.String("id"),
//   				Name: jsii.String("name"),
//   				ValueType: jsii.String("valueType"),
//
//   				// the properties below are optional
//   				DefaultValues: &StringDatasetParameterDefaultValuesProperty{
//   					StaticValues: []*string{
//   						jsii.String("staticValues"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DataSetRefreshProperties: &DataSetRefreshPropertiesProperty{
//   		RefreshConfiguration: &RefreshConfigurationProperty{
//   			IncrementalRefresh: &IncrementalRefreshProperty{
//   				LookbackWindow: &LookbackWindowProperty{
//   					ColumnName: jsii.String("columnName"),
//   					Size: jsii.Number(123),
//   					SizeUnit: jsii.String("sizeUnit"),
//   				},
//   			},
//   		},
//   	},
//   	DataSetUsageConfiguration: &DataSetUsageConfigurationProperty{
//   		DisableUseAsDirectQuerySource: jsii.Boolean(false),
//   		DisableUseAsImportedSource: jsii.Boolean(false),
//   	},
//   	FieldFolders: map[string]interface{}{
//   		"fieldFoldersKey": &FieldFolderProperty{
//   			"columns": []*string{
//   				jsii.String("columns"),
//   			},
//   			"description": jsii.String("description"),
//   		},
//   	},
//   	ImportMode: jsii.String("importMode"),
//   	IngestionWaitPolicy: &IngestionWaitPolicyProperty{
//   		IngestionWaitTimeInHours: jsii.Number(123),
//   		WaitForSpiceIngestion: jsii.Boolean(false),
//   	},
//   	LogicalTableMap: map[string]interface{}{
//   		"logicalTableMapKey": &LogicalTableProperty{
//   			"alias": jsii.String("alias"),
//   			"source": &LogicalTableSourceProperty{
//   				"dataSetArn": jsii.String("dataSetArn"),
//   				"joinInstruction": &JoinInstructionProperty{
//   					"leftOperand": jsii.String("leftOperand"),
//   					"onClause": jsii.String("onClause"),
//   					"rightOperand": jsii.String("rightOperand"),
//   					"type": jsii.String("type"),
//
//   					// the properties below are optional
//   					"leftJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   					"rightJoinKeyProperties": &JoinKeyPropertiesProperty{
//   						"uniqueKey": jsii.Boolean(false),
//   					},
//   				},
//   				"physicalTableId": jsii.String("physicalTableId"),
//   			},
//
//   			// the properties below are optional
//   			"dataTransforms": []interface{}{
//   				&TransformOperationProperty{
//   					"castColumnTypeOperation": &CastColumnTypeOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnType": jsii.String("newColumnType"),
//
//   						// the properties below are optional
//   						"format": jsii.String("format"),
//   						"subType": jsii.String("subType"),
//   					},
//   					"createColumnsOperation": &CreateColumnsOperationProperty{
//   						"columns": []interface{}{
//   							&CalculatedColumnProperty{
//   								"columnId": jsii.String("columnId"),
//   								"columnName": jsii.String("columnName"),
//   								"expression": jsii.String("expression"),
//   							},
//   						},
//   					},
//   					"filterOperation": &FilterOperationProperty{
//   						"conditionExpression": jsii.String("conditionExpression"),
//   					},
//   					"overrideDatasetParameterOperation": &OverrideDatasetParameterOperationProperty{
//   						"parameterName": jsii.String("parameterName"),
//
//   						// the properties below are optional
//   						"newDefaultValues": &NewDefaultValuesProperty{
//   							"dateTimeStaticValues": []*string{
//   								jsii.String("dateTimeStaticValues"),
//   							},
//   							"decimalStaticValues": []interface{}{
//   								jsii.Number(123),
//   							},
//   							"integerStaticValues": []interface{}{
//   								jsii.Number(123),
//   							},
//   							"stringStaticValues": []*string{
//   								jsii.String("stringStaticValues"),
//   							},
//   						},
//   						"newParameterName": jsii.String("newParameterName"),
//   					},
//   					"projectOperation": &ProjectOperationProperty{
//   						"projectedColumns": []*string{
//   							jsii.String("projectedColumns"),
//   						},
//   					},
//   					"renameColumnOperation": &RenameColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"newColumnName": jsii.String("newColumnName"),
//   					},
//   					"tagColumnOperation": &TagColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"tags": []ColumnTagProperty{
//   							&ColumnTagProperty{
//   								"columnDescription": &ColumnDescriptionProperty{
//   									"text": jsii.String("text"),
//   								},
//   								"columnGeographicRole": jsii.String("columnGeographicRole"),
//   							},
//   						},
//   					},
//   					"untagColumnOperation": &UntagColumnOperationProperty{
//   						"columnName": jsii.String("columnName"),
//   						"tagNames": []*string{
//   							jsii.String("tagNames"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   		},
//   	},
//   	PhysicalTableMap: map[string]interface{}{
//   		"physicalTableMapKey": &PhysicalTableProperty{
//   			"customSql": &CustomSqlProperty{
//   				"columns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"subType": jsii.String("subType"),
//   					},
//   				},
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"name": jsii.String("name"),
//   				"sqlQuery": jsii.String("sqlQuery"),
//   			},
//   			"relationalTable": &RelationalTableProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"subType": jsii.String("subType"),
//   					},
//   				},
//   				"name": jsii.String("name"),
//
//   				// the properties below are optional
//   				"catalog": jsii.String("catalog"),
//   				"schema": jsii.String("schema"),
//   			},
//   			"s3Source": &S3SourceProperty{
//   				"dataSourceArn": jsii.String("dataSourceArn"),
//   				"inputColumns": []interface{}{
//   					&InputColumnProperty{
//   						"name": jsii.String("name"),
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"subType": jsii.String("subType"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				"uploadSettings": &UploadSettingsProperty{
//   					"containsHeader": jsii.Boolean(false),
//   					"delimiter": jsii.String("delimiter"),
//   					"format": jsii.String("format"),
//   					"startFromRow": jsii.Number(123),
//   					"textQualifier": jsii.String("textQualifier"),
//   				},
//   			},
//   		},
//   	},
//   	RowLevelPermissionDataSet: &RowLevelPermissionDataSetProperty{
//   		Arn: jsii.String("arn"),
//   		PermissionPolicy: jsii.String("permissionPolicy"),
//
//   		// the properties below are optional
//   		FormatVersion: jsii.String("formatVersion"),
//   		Namespace: jsii.String("namespace"),
//   		Status: jsii.String("status"),
//   	},
//   	RowLevelPermissionTagConfiguration: &RowLevelPermissionTagConfigurationProperty{
//   		TagRules: []interface{}{
//   			&RowLevelPermissionTagRuleProperty{
//   				ColumnName: jsii.String("columnName"),
//   				TagKey: jsii.String("tagKey"),
//
//   				// the properties below are optional
//   				MatchAllValue: jsii.String("matchAllValue"),
//   				TagMultiValueDelimiter: jsii.String("tagMultiValueDelimiter"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Status: jsii.String("status"),
//   		TagRuleConfigurations: tagRuleConfigurations,
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dataset.html
//
type CfnDataSet interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The Amazon Resource Name (ARN) of the dataset.
	AttrArn() *string
	// <p>The amount of SPICE capacity used by this dataset.
	//
	// This is 0 if the dataset isn't
	//             imported into SPICE.</p>
	AttrConsumedSpiceCapacityInBytes() awscdk.IResolvable
	// The time this dataset version was created.
	AttrCreatedTime() *string
	// The time this dataset version was last updated.
	AttrLastUpdatedTime() *string
	// <p>The list of columns after all transforms.
	//
	// These columns are available in templates,
	//             analyses, and dashboards.</p>
	AttrOutputColumns() awscdk.IResolvable
	// The AWS account ID.
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Groupings of columns that work together in certain Amazon QuickSight features.
	ColumnGroups() interface{}
	SetColumnGroups(val interface{})
	// A set of one or more definitions of a `ColumnLevelPermissionRule` .
	ColumnLevelPermissionRules() interface{}
	SetColumnLevelPermissionRules(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An ID for the dataset that you want to create.
	DataSetId() *string
	SetDataSetId(val *string)
	// The parameters that are declared in a dataset.
	DatasetParameters() interface{}
	SetDatasetParameters(val interface{})
	// The refresh properties of a dataset.
	DataSetRefreshProperties() interface{}
	SetDataSetRefreshProperties(val interface{})
	// The usage configuration to apply to child datasets that reference this dataset as a source.
	DataSetUsageConfiguration() interface{}
	SetDataSetUsageConfiguration(val interface{})
	// The folder that contains fields and nested subfolders for your dataset.
	FieldFolders() interface{}
	SetFieldFolders(val interface{})
	// Indicates whether you want to import the data into SPICE.
	ImportMode() *string
	SetImportMode(val *string)
	// The wait policy to use when creating or updating a Dataset.
	IngestionWaitPolicy() interface{}
	SetIngestionWaitPolicy(val interface{})
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
	// Configures the combination and transformation of the data from the physical tables.
	LogicalTableMap() interface{}
	SetLogicalTableMap(val interface{})
	// The display name for the dataset.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// A list of resource permissions on the dataset.
	Permissions() interface{}
	SetPermissions(val interface{})
	// Declares the physical tables that are available in the underlying data sources.
	PhysicalTableMap() interface{}
	SetPhysicalTableMap(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The row-level security configuration for the data that you want to create.
	RowLevelPermissionDataSet() interface{}
	SetRowLevelPermissionDataSet(val interface{})
	// The element you can use to define tags for row-level security.
	RowLevelPermissionTagConfiguration() interface{}
	SetRowLevelPermissionTagConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.
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
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html#resources-section-structure-logicalid
	//
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

// The jsii proxy struct for CfnDataSet
type jsiiProxy_CfnDataSet struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnDataSet) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrConsumedSpiceCapacityInBytes() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrConsumedSpiceCapacityInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrCreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrLastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AttrOutputColumns() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrOutputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) ColumnGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) ColumnLevelPermissionRules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnLevelPermissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) DataSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) DatasetParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) DataSetRefreshProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSetRefreshProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) DataSetUsageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSetUsageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) FieldFolders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) ImportMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) IngestionWaitPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingestionWaitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) LogicalTableMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Permissions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) PhysicalTableMap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) RowLevelPermissionDataSet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowLevelPermissionDataSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) RowLevelPermissionTagConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowLevelPermissionTagConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSet) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnDataSet(scope constructs.Construct, id *string, props *CfnDataSetProps) CfnDataSet {
	_init_.Initialize()

	if err := validateNewCfnDataSetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataSet{}

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnDataSet_Override(c CfnDataSet, scope constructs.Construct, id *string, props *CfnDataSetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSet)SetAwsAccountId(val *string) {
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetColumnGroups(val interface{}) {
	if err := j.validateSetColumnGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnGroups",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetColumnLevelPermissionRules(val interface{}) {
	if err := j.validateSetColumnLevelPermissionRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnLevelPermissionRules",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetDataSetId(val *string) {
	_jsii_.Set(
		j,
		"dataSetId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetDatasetParameters(val interface{}) {
	if err := j.validateSetDatasetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetDataSetRefreshProperties(val interface{}) {
	if err := j.validateSetDataSetRefreshPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSetRefreshProperties",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetDataSetUsageConfiguration(val interface{}) {
	if err := j.validateSetDataSetUsageConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSetUsageConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetFieldFolders(val interface{}) {
	if err := j.validateSetFieldFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldFolders",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetImportMode(val *string) {
	_jsii_.Set(
		j,
		"importMode",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetIngestionWaitPolicy(val interface{}) {
	if err := j.validateSetIngestionWaitPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestionWaitPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetLogicalTableMap(val interface{}) {
	if err := j.validateSetLogicalTableMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logicalTableMap",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetPermissions(val interface{}) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetPhysicalTableMap(val interface{}) {
	if err := j.validateSetPhysicalTableMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalTableMap",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetRowLevelPermissionDataSet(val interface{}) {
	if err := j.validateSetRowLevelPermissionDataSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowLevelPermissionDataSet",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetRowLevelPermissionTagConfiguration(val interface{}) {
	if err := j.validateSetRowLevelPermissionTagConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowLevelPermissionTagConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSet)SetTagsRaw(val *[]*awscdk.CfnTag) {
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
func CfnDataSet_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSet_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDataSet_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSet_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
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
func CfnDataSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSet_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_quicksight.CfnDataSet",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataSet) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataSet) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSet) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSet) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataSet) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataSet) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataSet) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataSet) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataSet) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDataSet) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDataSet) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataSet) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSet) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataSet) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDataSet) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDataSet) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataSet) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

