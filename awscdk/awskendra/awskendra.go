package awskendra

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Kendra::DataSource`.
//
// Specifies a data source that you use to with an Amazon Kendra index.
//
// You specify a name, connector type and description for your data source.
//
// TODO: EXAMPLE
//
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataSourceConfiguration() interface{}
	SetDataSourceConfiguration(val interface{})
	Description() *string
	SetDescription(val *string)
	IndexId() *string
	SetIndexId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	Schedule() *string
	SetSchedule(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
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

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DataSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceConfiguration",
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

func (j *jsiiProxy_CfnDataSource) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
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

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Schedule() *string {
	var returns *string
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

func (j *jsiiProxy_CfnDataSource) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
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


// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource(scope constructs.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope constructs.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDataSourceConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"dataSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetIndexId(val *string) {
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
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
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnDataSource",
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
		"aws-cdk-lib.aws_kendra.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies access control list files for the documents in a data source.
//
// TODO: EXAMPLE
//
type CfnDataSource_AccessControlListConfigurationProperty struct {
	// Path to the AWS S3 bucket that contains the access control list files.
	KeyPath *string `json:"keyPath" yaml:"keyPath"`
}

// Provides information about the column that should be used for filtering the query response by groups.
//
// TODO: EXAMPLE
//
type CfnDataSource_AclConfigurationProperty struct {
	// A list of groups, separated by semi-colons, that filters a query response based on user context.
	//
	// The document is only returned to users that are in one of the groups specified in the `UserContext` field of the [Query](https://docs.aws.amazon.com/kendra/latest/dg/API_Query.html) operation.
	AllowedGroupsColumnName *string `json:"allowedGroupsColumnName" yaml:"allowedGroupsColumnName"`
}

// Provides information about how Amazon Kendra should use the columns of a database in an index.
//
// TODO: EXAMPLE
//
type CfnDataSource_ColumnConfigurationProperty struct {
	// One to five columns that indicate when a document in the database has changed.
	ChangeDetectingColumns *[]*string `json:"changeDetectingColumns" yaml:"changeDetectingColumns"`
	// The column that contains the contents of the document.
	DocumentDataColumnName *string `json:"documentDataColumnName" yaml:"documentDataColumnName"`
	// The column that provides the document's unique identifier.
	DocumentIdColumnName *string `json:"documentIdColumnName" yaml:"documentIdColumnName"`
	// The column that contains the title of the document.
	DocumentTitleColumnName *string `json:"documentTitleColumnName" yaml:"documentTitleColumnName"`
	// An array of objects that map database column names to the corresponding fields in an index.
	//
	// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
}

// Specifies the attachment settings for the Confluence data source.
//
// Attachment settings are optional, if you don't specify settings attachments, Amazon Kendra won't index them.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceAttachmentConfigurationProperty struct {
	// Defines how attachment metadata fields should be mapped to index fields.
	//
	// Before you can map a field, you must first create an index field with a matching type using the console or the `UpdateIndex` operation.
	//
	// If you specify the `AttachentFieldMappings` parameter, you must specify at least one field mapping.
	AttachmentFieldMappings interface{} `json:"attachmentFieldMappings" yaml:"attachmentFieldMappings"`
	// Indicates whether Amazon Kendra indexes attachments to the pages and blogs in the Confluence data source.
	CrawlAttachments interface{} `json:"crawlAttachments" yaml:"crawlAttachments"`
}

// Defines the mapping between a field in the Confluence data source to a Amazon Kendra index field.
//
// You must first create the index field using the `UpdateIndex` operation.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceAttachmentToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	//
	// You must first create the index field using the `UpdateIndex` operation.
	DataSourceFieldName *string `json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Specifies the blog settings for the Confluence data source.
//
// Blogs are always indexed unless filtered from the index by the `ExclusionPatterns` or `InclusionPatterns` fields in the `ConfluenceConfiguration` type.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceBlogConfigurationProperty struct {
	// Defines how blog metadata fields should be mapped to index fields.
	//
	// Before you can map a field, you must first create an index field with a matching type using the console or the `UpdateIndex` operation.
	//
	// If you specify the `BlogFieldMappings` parameter, you must specify at least one field mapping.
	BlogFieldMappings interface{} `json:"blogFieldMappings" yaml:"blogFieldMappings"`
}

// Defines the mapping between a blog field in the Confluence data source to a Amazon Kendra index field.
//
// You must first create the index field using the `UpdateIndex` operation.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceBlogToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	DataSourceFieldName *string `json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Provides configuration information for data sources that connect to Confluence.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key/value pairs required to connect to your Confluence server.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - username - The user name or email address of a user with administrative privileges for the Confluence server.
	// - password - The password associated with the user logging in to the Confluence server.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The URL of your Confluence instance.
	//
	// Use the full URL of the server. For example, `https://server.example.com:port/` . You can also use an IP address, for example, `https://192.168.1.113/` .
	ServerUrl *string `json:"serverUrl" yaml:"serverUrl"`
	// Specifies the version of the Confluence installation that you are connecting to.
	Version *string `json:"version" yaml:"version"`
	// Specifies configuration information for indexing attachments to Confluence blogs and pages.
	AttachmentConfiguration interface{} `json:"attachmentConfiguration" yaml:"attachmentConfiguration"`
	// Specifies configuration information for indexing Confluence blogs.
	BlogConfiguration interface{} `json:"blogConfiguration" yaml:"blogConfiguration"`
	// A list of regular expression patterns that apply to a URL on the Confluence server.
	//
	// An exclusion pattern can apply to a blog post, a page, a space, or an attachment. Items that match the pattern are excluded from the index. Items that don't match the pattern are included in the index. If a item matches both an exclusion pattern and an inclusion pattern, the item isn't included in the index.
	ExclusionPatterns *[]*string `json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of regular expression patterns that apply to a URL on the Confluence server.
	//
	// An inclusion pattern can apply to a blog post, a page, a space, or an attachment. Items that match the patterns are included in the index. Items that don't match the pattern are excluded from the index. If an item matches both an inclusion pattern and an exclusion pattern, the item isn't included in the index.
	InclusionPatterns *[]*string `json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Specifies configuration information for indexing Confluence pages.
	PageConfiguration interface{} `json:"pageConfiguration" yaml:"pageConfiguration"`
	// Specifies configuration information for indexing Confluence spaces.
	SpaceConfiguration interface{} `json:"spaceConfiguration" yaml:"spaceConfiguration"`
	// Specifies the information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Specifies the page settings for the Confluence data source.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluencePageConfigurationProperty struct {
	// Defines how page metadata fields should be mapped to index fields.
	//
	// Before you can map a field, you must first create an index field with a matching type using the console or the `UpdateIndex` operation.
	//
	// If you specify the `PageFieldMappings` parameter, you must specify at least one field mapping.
	PageFieldMappings interface{} `json:"pageFieldMappings" yaml:"pageFieldMappings"`
}

// Defines the mapping between a field in the Confluence data source to a Amazon Kendra index field.
//
// You must first create the index field using the `UpdateIndex` operation.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluencePageToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	DataSourceFieldName *string `json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Specifies the configuration for indexing Confluence spaces.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceSpaceConfigurationProperty struct {
	// Specifies whether Amazon Kendra should index archived spaces.
	CrawlArchivedSpaces interface{} `json:"crawlArchivedSpaces" yaml:"crawlArchivedSpaces"`
	// Specifies whether Amazon Kendra should index personal spaces.
	//
	// Users can add restrictions to items in personal spaces. If personal spaces are indexed, queries without user context information may return restricted items from a personal space in their results. For more information, see [Filtering on user context](https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html) .
	CrawlPersonalSpaces interface{} `json:"crawlPersonalSpaces" yaml:"crawlPersonalSpaces"`
	// A list of space keys of Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are not indexed. If a space is in both the `ExcludeSpaces` and the `IncludeSpaces` list, the space is excluded.
	ExcludeSpaces *[]*string `json:"excludeSpaces" yaml:"excludeSpaces"`
	// A list of space keys for Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are indexed. Spaces that aren't in the list aren't indexed. A space in the list must exist. Otherwise, Amazon Kendra logs an error when the data source is synchronized. If a space is in both the `IncludeSpaces` and the `ExcludeSpaces` list, the space is excluded.
	IncludeSpaces *[]*string `json:"includeSpaces" yaml:"includeSpaces"`
	// Defines how space metadata fields should be mapped to index fields.
	//
	// Before you can map a field, you must first create an index field with a matching type using the console or the `UpdateIndex` operation.
	//
	// If you specify the `SpaceFieldMappings` parameter, you must specify at least one field mapping.
	SpaceFieldMappings interface{} `json:"spaceFieldMappings" yaml:"spaceFieldMappings"`
}

// Defines the mapping between a field in the Confluence data source to a Amazon Kendra index field.
//
// You must first create the index field using the `UpdateIndex` operation.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConfluenceSpaceToIndexFieldMappingProperty struct {
	// The name of the field in the data source.
	DataSourceFieldName *string `json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the index field to map to the Confluence data source field.
	//
	// The index field type must match the Confluence field type.
	IndexFieldName *string `json:"indexFieldName" yaml:"indexFieldName"`
	// The format for date fields in the data source.
	//
	// If the field specified in `DataSourceFieldName` is a date field you must specify the date format. If the field is not a date field, an exception is thrown.
	DateFieldFormat *string `json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Provides the information necessary to connect to a database.
//
// TODO: EXAMPLE
//
type CfnDataSource_ConnectionConfigurationProperty struct {
	// The name of the host for the database.
	//
	// Can be either a string (host.subdomain.domain.tld) or an IPv4 or IPv6 address.
	DatabaseHost *string `json:"databaseHost" yaml:"databaseHost"`
	// The name of the database containing the document data.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The port that the database uses for connections.
	DatabasePort *float64 `json:"databasePort" yaml:"databasePort"`
	// The Amazon Resource Name (ARN) of credentials stored in AWS Secrets Manager .
	//
	// The credentials should be a user/password pair. For more information, see [Using a Database Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-database.html) . For more information about AWS Secrets Manager , see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The name of the table that contains the document data.
	TableName *string `json:"tableName" yaml:"tableName"`
}

// Configuration information for an Amazon Kendra data source.
//
// TODO: EXAMPLE
//
type CfnDataSource_DataSourceConfigurationProperty struct {
	// Provides configuration information for connecting to a Confluence data source.
	ConfluenceConfiguration interface{} `json:"confluenceConfiguration" yaml:"confluenceConfiguration"`
	// Provides information necessary to create a data source connector for a database.
	DatabaseConfiguration interface{} `json:"databaseConfiguration" yaml:"databaseConfiguration"`
	// Provides configuration for data sources that connect to Google Drive.
	GoogleDriveConfiguration interface{} `json:"googleDriveConfiguration" yaml:"googleDriveConfiguration"`
	// Provides configuration for data sources that connect to Microsoft OneDrive.
	OneDriveConfiguration interface{} `json:"oneDriveConfiguration" yaml:"oneDriveConfiguration"`
	// Provides information to create a data source connector for a document repository in an Amazon S3 bucket.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
	// Provides configuration information for data sources that connect to a Salesforce site.
	SalesforceConfiguration interface{} `json:"salesforceConfiguration" yaml:"salesforceConfiguration"`
	// Provides configuration for data sources that connect to ServiceNow instances.
	ServiceNowConfiguration interface{} `json:"serviceNowConfiguration" yaml:"serviceNowConfiguration"`
	// Provides information necessary to create a data source connector for a Microsoft SharePoint site.
	SharePointConfiguration interface{} `json:"sharePointConfiguration" yaml:"sharePointConfiguration"`
	// Provides the configuration information required for Amazon Kendra Web Crawler.
	WebCrawlerConfiguration interface{} `json:"webCrawlerConfiguration" yaml:"webCrawlerConfiguration"`
	// Provides the configuration information to connect to WorkDocs as your data source.
	WorkDocsConfiguration interface{} `json:"workDocsConfiguration" yaml:"workDocsConfiguration"`
}

// Maps a column or attribute in the data source to an index field.
//
// You must first create the fields in the index using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation.
//
// TODO: EXAMPLE
//
type CfnDataSource_DataSourceToIndexFieldMappingProperty struct {
	// The name of the column or attribute in the data source.
	DataSourceFieldName *string `json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// The name of the field in the index.
	IndexFieldName *string `json:"indexFieldName" yaml:"indexFieldName"`
	// The type of data stored in the column or attribute.
	DateFieldFormat *string `json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

// Provides information for connecting to an Amazon VPC.
//
// TODO: EXAMPLE
//
type CfnDataSource_DataSourceVpcConfigurationProperty struct {
	// A list of identifiers of security groups within your Amazon VPC.
	//
	// The security groups should enable Amazon Kendra to connect to the data source.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of identifiers for subnets within your Amazon VPC.
	//
	// The subnets should be able to connect to each other in the VPC, and they should have outgoing access to the Internet through a NAT device.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

// Provides the information necessary to connect a database to an index.
//
// TODO: EXAMPLE
//
type CfnDataSource_DatabaseConfigurationProperty struct {
	// Information about where the index should get the document information from the database.
	ColumnConfiguration interface{} `json:"columnConfiguration" yaml:"columnConfiguration"`
	// The information necessary to connect to a database.
	ConnectionConfiguration interface{} `json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// The type of database engine that runs the database.
	DatabaseEngineType *string `json:"databaseEngineType" yaml:"databaseEngineType"`
	// Information about the database column that provides information for user context filtering.
	AclConfiguration interface{} `json:"aclConfiguration" yaml:"aclConfiguration"`
	// Provides information about how Amazon Kendra uses quote marks around SQL identifiers when querying a database data source.
	SqlConfiguration interface{} `json:"sqlConfiguration" yaml:"sqlConfiguration"`
	// Provides information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes.
//
// Each metadata file contains metadata about a single document.
//
// TODO: EXAMPLE
//
type CfnDataSource_DocumentsMetadataConfigurationProperty struct {
	// A prefix used to filter metadata configuration files in the AWS S3 bucket.
	//
	// The S3 bucket might contain multiple metadata files. Use `S3Prefix` to include only the desired metadata files.
	S3Prefix *string `json:"s3Prefix" yaml:"s3Prefix"`
}

// Provides configuration information for data sources that connect to Google Drive.
//
// TODO: EXAMPLE
//
type CfnDataSource_GoogleDriveConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of a AWS Secrets Manager secret that contains the credentials required to connect to Google Drive.
	//
	// For more information, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// A list of MIME types to exclude from the index. All documents matching the specified MIME type are excluded.
	//
	// For a list of MIME types, see [Using a Google Workspace Drive data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-google-drive.html) .
	ExcludeMimeTypes *[]*string `json:"excludeMimeTypes" yaml:"excludeMimeTypes"`
	// A list of identifiers or shared drives to exclude from the index.
	//
	// All files and folders stored on the shared drive are excluded.
	ExcludeSharedDrives *[]*string `json:"excludeSharedDrives" yaml:"excludeSharedDrives"`
	// A list of email addresses of the users.
	//
	// Documents owned by these users are excluded from the index. Documents shared with excluded users are indexed unless they are excluded in another way.
	ExcludeUserAccounts *[]*string `json:"excludeUserAccounts" yaml:"excludeUserAccounts"`
	// A list of regular expression patterns that apply to the path on Google Drive.
	//
	// Items that match the pattern are excluded from the index from both shared drives and users' My Drives. Items that don't match the pattern are included in the index. If an item matches both an exclusion pattern and an inclusion pattern, it is excluded from the index.
	ExclusionPatterns *[]*string `json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// Defines mapping between a field in the Google Drive and a Amazon Kendra index field.
	//
	// If you are using the console, you can define index fields when creating the mapping. If you are using the API, you must first create the field using the `UpdateIndex` operation.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns that apply to path on Google Drive.
	//
	// Items that match the pattern are included in the index from both shared drives and users' My Drives. Items that don't match the pattern are excluded from the index. If an item matches both an inclusion pattern and an exclusion pattern, it is excluded from the index.
	InclusionPatterns *[]*string `json:"inclusionPatterns" yaml:"inclusionPatterns"`
}

// Provides configuration information for data sources that connect to OneDrive.
//
// TODO: EXAMPLE
//
type CfnDataSource_OneDriveConfigurationProperty struct {
	// A list of user accounts whose documents should be indexed.
	OneDriveUsers interface{} `json:"oneDriveUsers" yaml:"oneDriveUsers"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the user name and password to connect to OneDrive.
	//
	// The user named should be the application ID for the OneDrive application, and the password is the application key for the OneDrive application.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The Azure Active Directory domain of the organization.
	TenantDomain *string `json:"tenantDomain" yaml:"tenantDomain"`
	// A Boolean value that specifies whether local groups are disabled ( `True` ) or enabled ( `False` ).
	DisableLocalGroups interface{} `json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// List of regular expressions applied to documents.
	//
	// Items that match the exclusion pattern are not indexed. If you provide both an inclusion pattern and an exclusion pattern, any item that matches the exclusion pattern isn't indexed.
	//
	// The exclusion pattern is applied to the file name.
	ExclusionPatterns *[]*string `json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Microsoft OneDrive fields to custom fields in the Amazon Kendra index.
	//
	// You must first create the index fields before you map OneDrive fields.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns.
	//
	// Documents that match the pattern are included in the index. Documents that don't match the pattern are excluded from the index. If a document matches both an inclusion pattern and an exclusion pattern, the document is not included in the index.
	//
	// The exclusion pattern is applied to the file name.
	InclusionPatterns *[]*string `json:"inclusionPatterns" yaml:"inclusionPatterns"`
}

// User accounts whose documents should be indexed.
//
// TODO: EXAMPLE
//
type CfnDataSource_OneDriveUsersProperty struct {
	// A list of users whose documents should be indexed.
	//
	// Specify the user names in email format, for example, `username@tenantdomain` . If you need to index the documents of more than 100 users, use the `OneDriveUserS3Path` field to specify the location of a file containing a list of users.
	OneDriveUserList *[]*string `json:"oneDriveUserList" yaml:"oneDriveUserList"`
	// The S3 bucket location of a file containing a list of users whose documents should be indexed.
	OneDriveUserS3Path interface{} `json:"oneDriveUserS3Path" yaml:"oneDriveUserS3Path"`
}

// Provides the configuration information for a web proxy to connect to website hosts.
//
// TODO: EXAMPLE
//
type CfnDataSource_ProxyConfigurationProperty struct {
	// The name of the website host you want to connect to via a web proxy server.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	Host *string `json:"host" yaml:"host"`
	// The port number of the website host you want to connect to via a web proxy server.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	Port *float64 `json:"port" yaml:"port"`
	// Your secret ARN, which you can create in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
	//
	// The credentials are optional. You use a secret if web proxy credentials are required to connect to a website host. Amazon Kendra currently support basic authentication to connect to a web proxy server. The secret stores your credentials.
	Credentials *string `json:"credentials" yaml:"credentials"`
}

// Provides configuration information for a data source to index documents in an Amazon S3 bucket.
//
// TODO: EXAMPLE
//
type CfnDataSource_S3DataSourceConfigurationProperty struct {
	// The name of the bucket that contains the documents.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// Provides the path to the S3 bucket that contains the user context filtering files for the data source.
	//
	// For the format of the file, see [Access control for S3 data sources](https://docs.aws.amazon.com/kendra/latest/dg/s3-acl.html) .
	AccessControlListConfiguration interface{} `json:"accessControlListConfiguration" yaml:"accessControlListConfiguration"`
	// Specifies document metadata files that contain information such as the document access control information, source URI, document author, and custom attributes.
	//
	// Each metadata file contains metadata about a single document.
	DocumentsMetadataConfiguration interface{} `json:"documentsMetadataConfiguration" yaml:"documentsMetadataConfiguration"`
	// A list of glob patterns for documents that should not be indexed.
	//
	// If a document that matches an inclusion prefix or inclusion pattern also matches an exclusion pattern, the document is not indexed.
	//
	// Some [examples](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) are:
	//
	// - **.png , *.jpg* will exclude all PNG and JPEG image files in a directory (files with the extensions .png and .jpg).
	// - **internal** will exclude all files in a directory that contain 'internal' in the file name, such as 'internal', 'internal_only', 'company_internal'.
	// - *** /*internal** will exclude all internal-related files in a directory and its subdirectories.
	ExclusionPatterns *[]*string `json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of glob patterns for documents that should be indexed.
	//
	// If a document that matches an inclusion pattern also matches an exclusion pattern, the document is not indexed.
	//
	// Some [examples](https://docs.aws.amazon.com/cli/latest/reference/s3/#use-of-exclude-and-include-filters) are:
	//
	// - **.txt* will include all text files in a directory (files with the extension .txt).
	// - *** /*.txt* will include all text files in a directory and its subdirectories.
	// - **tax** will include all files in a directory that contain 'tax' in the file name, such as 'tax', 'taxes', 'income_tax'.
	InclusionPatterns *[]*string `json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// A list of S3 prefixes for the documents that should be included in the index.
	InclusionPrefixes *[]*string `json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

// Information required to find a specific file in an Amazon S3 bucket.
//
// TODO: EXAMPLE
//
type CfnDataSource_S3PathProperty struct {
	// The name of the S3 bucket that contains the file.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The name of the file.
	Key *string `json:"key" yaml:"key"`
}

// Defines configuration for syncing a Salesforce chatter feed.
//
// The contents of the object comes from the Salesforce FeedItem table.
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceChatterFeedConfigurationProperty struct {
	// The name of the column in the Salesforce FeedItem table that contains the content to index.
	//
	// Typically this is the `Body` column.
	DocumentDataFieldName *string `json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the column in the Salesforce FeedItem table that contains the title of the document.
	//
	// This is typically the `Title` column.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps fields from a Salesforce chatter feed into Amazon Kendra index fields.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// Filters the documents in the feed based on status of the user.
	//
	// When you specify `ACTIVE_USERS` only documents from users who have an active account are indexed. When you specify `STANDARD_USER` only documents for Salesforce standard users are documented. You can specify both.
	IncludeFilterTypes *[]*string `json:"includeFilterTypes" yaml:"includeFilterTypes"`
}

// Provides configuration information for connecting to a Salesforce data source.
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key/value pairs required to connect to your Salesforce instance.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - authenticationUrl - The OAUTH endpoint that Amazon Kendra connects to get an OAUTH token.
	// - consumerKey - The application public key generated when you created your Salesforce application.
	// - consumerSecret - The application private key generated when you created your Salesforce application.
	// - password - The password associated with the user logging in to the Salesforce instance.
	// - securityToken - The token associated with the user account logging in to the Salesforce instance.
	// - username - The user name of the user logging in to the Salesforce instance.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The instance URL for the Salesforce site that you want to index.
	ServerUrl *string `json:"serverUrl" yaml:"serverUrl"`
	// Specifies configuration information for Salesforce chatter feeds.
	ChatterFeedConfiguration interface{} `json:"chatterFeedConfiguration" yaml:"chatterFeedConfiguration"`
	// Indicates whether Amazon Kendra should index attachments to Salesforce objects.
	CrawlAttachments interface{} `json:"crawlAttachments" yaml:"crawlAttachments"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an exclusion pattern and an inclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the name of the attached file.
	ExcludeAttachmentFilePatterns *[]*string `json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion pattern and an exclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the name of the attached file.
	IncludeAttachmentFilePatterns *[]*string `json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
	// Specifies configuration information for the knowledge article types that Amazon Kendra indexes.
	//
	// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
	KnowledgeArticleConfiguration interface{} `json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Provides configuration information for processing attachments to Salesforce standard objects.
	StandardObjectAttachmentConfiguration interface{} `json:"standardObjectAttachmentConfiguration" yaml:"standardObjectAttachmentConfiguration"`
	// Specifies the Salesforce standard objects that Amazon Kendra indexes.
	StandardObjectConfigurations interface{} `json:"standardObjectConfigurations" yaml:"standardObjectConfigurations"`
}

// Provides configuration information for indexing Salesforce custom articles.
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceCustomKnowledgeArticleTypeConfigurationProperty struct {
	// The name of the field in the custom knowledge article that contains the document data to index.
	DocumentDataFieldName *string `json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the configuration.
	Name *string `json:"name" yaml:"name"`
	// The name of the field in the custom knowledge article that contains the document title.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// One or more objects that map fields in the custom knowledge article to fields in the Amazon Kendra index.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
}

// Specifies configuration information for the knowledge article types that Amazon Kendra indexes.
//
// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceKnowledgeArticleConfigurationProperty struct {
	// Specifies the document states that should be included when Amazon Kendra indexes knowledge articles.
	//
	// You must specify at least one state.
	IncludedStates *[]*string `json:"includedStates" yaml:"includedStates"`
	// Provides configuration information for custom Salesforce knowledge articles.
	CustomKnowledgeArticleTypeConfigurations interface{} `json:"customKnowledgeArticleTypeConfigurations" yaml:"customKnowledgeArticleTypeConfigurations"`
	// Provides configuration information for standard Salesforce knowledge articles.
	StandardKnowledgeArticleTypeConfiguration interface{} `json:"standardKnowledgeArticleTypeConfiguration" yaml:"standardKnowledgeArticleTypeConfiguration"`
}

// Provides configuration information for standard Salesforce knowledge articles.
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceStandardKnowledgeArticleTypeConfigurationProperty struct {
	// The name of the field that contains the document data to index.
	DocumentDataFieldName *string `json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the field that contains the document title.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// One or more objects that map fields in the knowledge article to Amazon Kendra index fields.
	//
	// The index field must exist before you can map a Salesforce field to it.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
}

// Provides configuration information for processing attachments to Salesforce standard objects.
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceStandardObjectAttachmentConfigurationProperty struct {
	// The name of the field used for the document title.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// One or more objects that map fields in attachments to Amazon Kendra index fields.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
}

// Specifies configuration information for indexing a single standard object.
//
// TODO: EXAMPLE
//
type CfnDataSource_SalesforceStandardObjectConfigurationProperty struct {
	// The name of the field in the standard object table that contains the document contents.
	DocumentDataFieldName *string `json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the standard object.
	Name *string `json:"name" yaml:"name"`
	// The name of the field in the standard object table that contains the document title.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// One or more objects that map fields in the standard object to Amazon Kendra index fields.
	//
	// The index field must exist before you can map a Salesforce field to it.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
}

// Provides configuration information required to connect to a ServiceNow data source.
//
// TODO: EXAMPLE
//
type CfnDataSource_ServiceNowConfigurationProperty struct {
	// The ServiceNow instance that the data source connects to.
	//
	// The host endpoint should look like the following: `{instance}.service-now.com.`
	HostUrl *string `json:"hostUrl" yaml:"hostUrl"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the user name and password required to connect to the ServiceNow instance.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The identifier of the release that the ServiceNow host is running.
	//
	// If the host is not running the `LONDON` release, use `OTHERS` .
	ServiceNowBuildVersion *string `json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
	// Determines the type of authentication used to connect to the ServiceNow instance.
	//
	// If you choose `HTTP_BASIC` , Amazon Kendra is authenticated using the user name and password provided in the AWS Secrets Manager secret in the `SecretArn` field. When you choose `OAUTH2` , Amazon Kendra is authenticated using the OAuth token and secret provided in the Secrets Manager secret, and the user name and password are used to determine which information Amazon Kendra has access to.
	//
	// When you use `OAUTH2` authentication, you must generate a token and a client secret using the ServiceNow console. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	AuthenticationType *string `json:"authenticationType" yaml:"authenticationType"`
	// Provides configuration information for crawling knowledge articles in the ServiceNow site.
	KnowledgeArticleConfiguration interface{} `json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Provides configuration information for crawling service catalogs in the ServiceNow site.
	ServiceCatalogConfiguration interface{} `json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
}

// Provides configuration information for crawling knowledge articles in the ServiceNow site.
//
// TODO: EXAMPLE
//
type CfnDataSource_ServiceNowKnowledgeArticleConfigurationProperty struct {
	// The name of the ServiceNow field that is mapped to the index document contents field in the Amazon Kendra index.
	DocumentDataFieldName *string `json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Indicates whether Amazon Kendra should index attachments to knowledge articles.
	CrawlAttachments interface{} `json:"crawlAttachments" yaml:"crawlAttachments"`
	// The name of the ServiceNow field that is mapped to the index document title field.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// List of regular expressions applied to knowledge articles.
	//
	// Items that don't match the inclusion pattern are not indexed. The regex is applied to the field specified in the `PatternTargetField`
	ExcludeAttachmentFilePatterns *[]*string `json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Mapping between ServiceNow fields and Amazon Kendra index fields.
	//
	// You must create the index field before you map the field.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// A query that selects the knowledge articles to index.
	//
	// The query can return articles from multiple knowledge bases, and the knowledge bases can be public or private.
	//
	// The query string must be one generated by the ServiceNow console. For more information, see [Specifying documents to index with a query](https://docs.aws.amazon.com/kendra/latest/dg/servicenow-query.html) .
	FilterQuery *string `json:"filterQuery" yaml:"filterQuery"`
	// List of regular expressions applied to knowledge articles.
	//
	// Items that don't match the inclusion pattern are not indexed. The regex is applied to the field specified in the `PatternTargetField` .
	IncludeAttachmentFilePatterns *[]*string `json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
}

// Provides configuration information for crawling service catalog items in the ServiceNow site.
//
// TODO: EXAMPLE
//
type CfnDataSource_ServiceNowServiceCatalogConfigurationProperty struct {
	// The name of the ServiceNow field that is mapped to the index document contents field in the Amazon Kendra index.
	DocumentDataFieldName *string `json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Indicates whether Amazon Kendra should crawl attachments to the service catalog items.
	CrawlAttachments interface{} `json:"crawlAttachments" yaml:"crawlAttachments"`
	// The name of the ServiceNow field that is mapped to the index document title field.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an exclusion pattern and an inclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the file name of the attachment.
	ExcludeAttachmentFilePatterns *[]*string `json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Mapping between ServiceNow fields and Amazon Kendra index fields.
	//
	// You must create the index field before you map the field.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an exclusion pattern and an inclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the file name of the attachment.
	IncludeAttachmentFilePatterns *[]*string `json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
}

// Provides configuration information for connecting to a Microsoft SharePoint data source.
//
// TODO: EXAMPLE
//
type CfnDataSource_SharePointConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of credentials stored in AWS Secrets Manager .
	//
	// The credentials should be a user/password pair. If you use SharePoint Server, you also need to provide the sever domain name as part of the credentials. For more information, see [Using a Microsoft SharePoint Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-sharepoint.html) . For more information about AWS Secrets Manager see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	SecretArn *string `json:"secretArn" yaml:"secretArn"`
	// The version of Microsoft SharePoint that you are using as a data source.
	SharePointVersion *string `json:"sharePointVersion" yaml:"sharePointVersion"`
	// The URLs of the Microsoft SharePoint site that contains the documents that should be indexed.
	Urls *[]*string `json:"urls" yaml:"urls"`
	// `TRUE` to include attachments to documents stored in your Microsoft SharePoint site in the index;
	//
	// otherwise, `FALSE` .
	CrawlAttachments interface{} `json:"crawlAttachments" yaml:"crawlAttachments"`
	// A Boolean value that specifies whether local groups are disabled ( `True` ) or enabled ( `False` ).
	DisableLocalGroups interface{} `json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// The Microsoft SharePoint attribute field that contains the title of the document.
	DocumentTitleFieldName *string `json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an exclusion pattern and an inclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the display URL of the SharePoint document.
	ExclusionPatterns *[]*string `json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Microsoft SharePoint attributes to custom fields in the Amazon Kendra index.
	//
	// You must first create the index fields using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation before you map SharePoint attributes. For more information, see [Mapping Data Source Fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) .
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion pattern and an exclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the display URL of the SharePoint document.
	InclusionPatterns *[]*string `json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Information required to find a specific file in an Amazon S3 bucket.
	SslCertificateS3Path interface{} `json:"sslCertificateS3Path" yaml:"sslCertificateS3Path"`
	// Set to `TRUE` to use the Microsoft SharePoint change log to determine the documents that need to be updated in the index.
	//
	// Depending on the size of the SharePoint change log, it may take longer for Amazon Kendra to use the change log than it takes it to determine the changed documents using the Amazon Kendra document crawler.
	UseChangeLog interface{} `json:"useChangeLog" yaml:"useChangeLog"`
	// Provides information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// Provides information that configures Amazon Kendra to use a SQL database.
//
// TODO: EXAMPLE
//
type CfnDataSource_SqlConfigurationProperty struct {
	// Determines whether Amazon Kendra encloses SQL identifiers for tables and column names in double quotes (") when making a database query.
	//
	// You can set the value to `DOUBLE_QUOTES` or `NONE` .
	//
	// By default, Amazon Kendra passes SQL identifiers the way that they are entered into the data source configuration. It does not change the case of identifiers or enclose them in quotes.
	//
	// PostgreSQL internally converts uppercase characters to lower case characters in identifiers unless they are quoted. Choosing this option encloses identifiers in quotes so that PostgreSQL does not convert the character's case.
	//
	// For MySQL databases, you must enable the ansi_quotes option when you set this field to `DOUBLE_QUOTES` .
	QueryIdentifiersEnclosingOption *string `json:"queryIdentifiersEnclosingOption" yaml:"queryIdentifiersEnclosingOption"`
}

// Provides the configuration information to connect to websites that require user authentication.
//
// TODO: EXAMPLE
//
type CfnDataSource_WebCrawlerAuthenticationConfigurationProperty struct {
	// The list of configuration information that's required to connect to and crawl a website host using basic authentication credentials.
	//
	// The list includes the name and port number of the website host.
	BasicAuthentication interface{} `json:"basicAuthentication" yaml:"basicAuthentication"`
}

// Provides the configuration information to connect to websites that require basic user authentication.
//
// TODO: EXAMPLE
//
type CfnDataSource_WebCrawlerBasicAuthenticationProperty struct {
	// Your secret ARN, which you can create in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html).
	//
	// You use a secret if basic authentication credentials are required to connect to a website. The secret stores your credentials of user name and password.
	Credentials *string `json:"credentials" yaml:"credentials"`
	// The name of the website host you want to connect to using authentication credentials.
	//
	// For example, the host name of https://a.example.com/page1.html is "a.example.com".
	Host *string `json:"host" yaml:"host"`
	// The port number of the website host you want to connect to using authentication credentials.
	//
	// For example, the port for https://a.example.com/page1.html is 443, the standard port for HTTPS.
	Port *float64 `json:"port" yaml:"port"`
}

// TODO: EXAMPLE
//
type CfnDataSource_WebCrawlerConfigurationProperty struct {
	// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
	//
	// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
	//
	// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
	//
	// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use Amazon Kendra Web Crawler to index your own webpages, or webpages that you have authorization to index.*
	Urls interface{} `json:"urls" yaml:"urls"`
	// Provides configuration information required to connect to websites using authentication.
	//
	// You can connect to websites using basic authentication of user name and password.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS. You use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) to store your authentication credentials.
	AuthenticationConfiguration interface{} `json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Specifies the number of levels in a website that you want to crawl.
	//
	// The first level begins from the website seed or starting point URL. For example, if a website has 3 levels – index level (i.e. seed in this example), sections level, and subsections level – and you are only interested in crawling information up to the sections level (i.e. levels 0-1), you can set your depth to 1.
	//
	// The default crawl depth is set to 2.
	CrawlDepth *float64 `json:"crawlDepth" yaml:"crawlDepth"`
	// The maximum size (in MB) of a webpage or attachment to crawl.
	//
	// Files larger than this size (in MB) are skipped/not crawled.
	//
	// The default maximum size of a webpage or attachment is set to 50 MB.
	MaxContentSizePerPageInMegaBytes *float64 `json:"maxContentSizePerPageInMegaBytes" yaml:"maxContentSizePerPageInMegaBytes"`
	// The maximum number of URLs on a webpage to include when crawling a website. This number is per webpage.
	//
	// As a website’s webpages are crawled, any URLs the webpages link to are also crawled. URLs on a webpage are crawled in order of appearance.
	//
	// The default maximum links per page is 100.
	MaxLinksPerPage *float64 `json:"maxLinksPerPage" yaml:"maxLinksPerPage"`
	// The maximum number of URLs crawled per website host per minute.
	//
	// A minimum of one URL is required.
	//
	// The default maximum number of URLs crawled per website host per minute is 300.
	MaxUrlsPerMinuteCrawlRate *float64 `json:"maxUrlsPerMinuteCrawlRate" yaml:"maxUrlsPerMinuteCrawlRate"`
	// Provides configuration information required to connect to your internal websites via a web proxy.
	//
	// You must provide the website host name and port number. For example, the host name of https://a.example.com/page1.html is "a.example.com" and the port is 443, the standard port for HTTPS.
	//
	// Web proxy credentials are optional and you can use them to connect to a web proxy server that requires basic authentication. To store web proxy credentials, you use a secret in [AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) .
	ProxyConfiguration interface{} `json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The regular expression pattern to exclude certain URLs to crawl.
	//
	// If there is a regular expression pattern to include certain URLs that conflicts with the exclude pattern, the exclude pattern takes precedence.
	UrlExclusionPatterns *[]*string `json:"urlExclusionPatterns" yaml:"urlExclusionPatterns"`
	// The regular expression pattern to include certain URLs to crawl.
	//
	// If there is a regular expression pattern to exclude certain URLs that conflicts with the include pattern, the exclude pattern takes precedence.
	UrlInclusionPatterns *[]*string `json:"urlInclusionPatterns" yaml:"urlInclusionPatterns"`
}

// Provides the configuration information of the seed or starting point URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// TODO: EXAMPLE
//
type CfnDataSource_WebCrawlerSeedUrlConfigurationProperty struct {
	// The list of seed or starting point URLs of the websites you want to crawl.
	//
	// The list can include a maximum of 100 seed URLs.
	SeedUrls *[]*string `json:"seedUrls" yaml:"seedUrls"`
	// You can choose one of the following modes:.
	//
	// - `HOST_ONLY` – crawl only the website host names. For example, if the seed URL is "abc.example.com", then only URLs with host name "abc.example.com" are crawled.
	// - `SUBDOMAINS` – crawl the website host names with subdomains. For example, if the seed URL is "abc.example.com", then "a.abc.example.com" and "b.abc.example.com" are also crawled.
	// - `EVERYTHING` – crawl the website host names with subdomains and other domains that the webpages link to.
	//
	// The default mode is set to `HOST_ONLY` .
	WebCrawlerMode *string `json:"webCrawlerMode" yaml:"webCrawlerMode"`
}

// Provides the configuration information of the sitemap URLs to crawl.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// TODO: EXAMPLE
//
type CfnDataSource_WebCrawlerSiteMapsConfigurationProperty struct {
	// The list of sitemap URLs of the websites you want to crawl.
	//
	// The list can include a maximum of three sitemap URLs.
	SiteMaps *[]*string `json:"siteMaps" yaml:"siteMaps"`
}

// Specifies the seed or starting point URLs of the websites or the sitemap URLs of the websites you want to crawl.
//
// You can include website subdomains. You can list up to 100 seed URLs and up to three sitemap URLs.
//
// You can only crawl websites that use the secure communication protocol, Hypertext Transfer Protocol Secure (HTTPS). If you receive an error when crawling a website, it could be that the website is blocked from crawling.
//
// *When selecting websites to index, you must adhere to the [Amazon Acceptable Use Policy](https://docs.aws.amazon.com/aup/) and all other Amazon terms. Remember that you must only use the Amazon Kendra web crawler to index your own webpages, or webpages that you have authorization to index.*
//
// TODO: EXAMPLE
//
type CfnDataSource_WebCrawlerUrlsProperty struct {
	// Provides the configuration of the seed or starting point URLs of the websites you want to crawl.
	//
	// You can choose to crawl only the website host names, or the website host names with subdomains, or the website host names with subdomains and other domains that the webpages link to.
	//
	// You can list up to 100 seed URLs.
	SeedUrlConfiguration interface{} `json:"seedUrlConfiguration" yaml:"seedUrlConfiguration"`
	// Provides the configuration of the sitemap URLs of the websites you want to crawl.
	//
	// Only URLs belonging to the same website host names are crawled. You can list up to three sitemap URLs.
	SiteMapsConfiguration interface{} `json:"siteMapsConfiguration" yaml:"siteMapsConfiguration"`
}

// Provides the configuration information to connect to Amazon WorkDocs as your data source.
//
// Amazon WorkDocs connector is available in Oregon, North Virginia, Sydney, Singapore and Ireland regions.
//
// TODO: EXAMPLE
//
type CfnDataSource_WorkDocsConfigurationProperty struct {
	// The identifier of the directory corresponding to your Amazon WorkDocs site repository.
	//
	// You can find the organization ID in the [AWS Directory Service](https://docs.aws.amazon.com/directoryservicev2/) by going to *Active Directory* , then *Directories* . Your Amazon WorkDocs site directory has an ID, which is the organization ID. You can also set up a new Amazon WorkDocs directory in the AWS Directory Service console and enable a Amazon WorkDocs site for the directory in the Amazon WorkDocs console.
	OrganizationId *string `json:"organizationId" yaml:"organizationId"`
	// `TRUE` to include comments on documents in your index.
	//
	// Including comments in your index means each comment is a document that can be searched on.
	//
	// The default is set to `FALSE` .
	CrawlComments interface{} `json:"crawlComments" yaml:"crawlComments"`
	// A list of regular expression patterns to exclude certain files in your Amazon WorkDocs site repository.
	//
	// Files that match the patterns are excluded from the index. Files that don’t match the patterns are included in the index. If a file matches both an inclusion pattern and an exclusion pattern, the exclusion pattern takes precedence and the file isn’t included in the index.
	ExclusionPatterns *[]*string `json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Amazon WorkDocs field names to custom index field names in Amazon Kendra.
	//
	// You must first create the custom index fields using the `UpdateIndex` operation before you map to Amazon WorkDocs fields. For more information, see [Mapping Data Source Fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Amazon WorkDocs data source field names need to exist in your Amazon WorkDocs custom metadata.
	FieldMappings interface{} `json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain files in your Amazon WorkDocs site repository.
	//
	// Files that match the patterns are included in the index. Files that don't match the patterns are excluded from the index. If a file matches both an inclusion pattern and an exclusion pattern, the exclusion pattern takes precedence and the file isn’t included in the index.
	InclusionPatterns *[]*string `json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// `TRUE` to use the change logs to update documents in your index instead of scanning all documents.
	//
	// If you are syncing your Amazon WorkDocs data source with your index for the first time, all documents are scanned. After your first sync, you can use the change logs to update your documents in your index for future syncs.
	//
	// The default is set to `FALSE` .
	UseChangeLog interface{} `json:"useChangeLog" yaml:"useChangeLog"`
}

// Properties for defining a `CfnDataSource`.
//
// TODO: EXAMPLE
//
type CfnDataSourceProps struct {
	// The identifier of the index that should be associated with this data source.
	IndexId *string `json:"indexId" yaml:"indexId"`
	// The name of the data source.
	Name *string `json:"name" yaml:"name"`
	// The type of the data source.
	Type *string `json:"type" yaml:"type"`
	// Configuration information for an Amazon Kendra data source.
	//
	// The contents of the configuration depend on the type of data source. You can only specify one type of data source in the configuration. Choose from one of the following data sources.
	//
	// - Amazon S3
	// - Confluence
	// - Custom
	// - Database
	// - Microsoft OneDrive
	// - Microsoft SharePoint
	// - Salesforce
	// - ServiceNow
	//
	// You can't specify the `Configuration` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `Configuration` parameter is required for all other data sources.
	DataSourceConfiguration interface{} `json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
	// A description of the data source.
	Description *string `json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source.
	//
	// You can't specify the `RoleArn` parameter when the `Type` parameter is set to `CUSTOM` .
	//
	// The `RoleArn` parameter is required for all other data sources.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Sets the frequency that Amazon Kendra checks the documents in your data source and updates the index.
	//
	// If you don't set a schedule, Amazon Kendra doesn't periodically update the index.
	Schedule *string `json:"schedule" yaml:"schedule"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Kendra::Faq`.
//
// Specifies an new set of frequently asked question (FAQ) questions and answers.
//
// TODO: EXAMPLE
//
type CfnFaq interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	FileFormat() *string
	SetFileFormat(val *string)
	IndexId() *string
	SetIndexId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	S3Path() interface{}
	SetS3Path(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnFaq
type jsiiProxy_CfnFaq struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFaq) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) FileFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) S3Path() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFaq) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Faq`.
func NewCfnFaq(scope constructs.Construct, id *string, props *CfnFaqProps) CfnFaq {
	_init_.Initialize()

	j := jsiiProxy_CfnFaq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Faq`.
func NewCfnFaq_Override(c CfnFaq, scope constructs.Construct, id *string, props *CfnFaqProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFaq) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetFileFormat(val *string) {
	_jsii_.Set(
		j,
		"fileFormat",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetIndexId(val *string) {
	_jsii_.Set(
		j,
		"indexId",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnFaq) SetS3Path(val interface{}) {
	_jsii_.Set(
		j,
		"s3Path",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFaq_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFaq_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
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
func CfnFaq_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFaq_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnFaq",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFaq) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFaq) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFaq) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFaq) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFaq) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFaq) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnFaq) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFaq) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFaq) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFaq) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnFaq) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFaq) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFaq) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFaq) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFaq) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information required to find a specific file in an Amazon S3 bucket.
//
// TODO: EXAMPLE
//
type CfnFaq_S3PathProperty struct {
	// The name of the S3 bucket that contains the file.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The name of the file.
	Key *string `json:"key" yaml:"key"`
}

// Properties for defining a `CfnFaq`.
//
// TODO: EXAMPLE
//
type CfnFaqProps struct {
	// The identifier of the index that contains the FAQ.
	IndexId *string `json:"indexId" yaml:"indexId"`
	// The name that you assigned the FAQ when you created or updated the FAQ.
	Name *string `json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket that contains the FAQ.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The Amazon Simple Storage Service (Amazon S3) location of the FAQ input data.
	S3Path interface{} `json:"s3Path" yaml:"s3Path"`
	// A description of the FAQ.
	Description *string `json:"description" yaml:"description"`
	// The format of the input file.
	//
	// You can choose between a basic CSV format, a CSV format that includes customs attributes in a header, and a JSON format that includes custom attributes.
	//
	// The format must match the format of the file stored in the S3 bucket identified in the S3Path parameter.
	//
	// Valid values are:
	//
	// - `CSV`
	// - `CSV_WITH_HEADER`
	// - `JSON`
	FileFormat *string `json:"fileFormat" yaml:"fileFormat"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Kendra::Index`.
//
// Specifies a new Amazon Kendra index. And index is a collection of documents and associated metadata that you want to search for relevant documents.
//
// Once the index is active you can add documents to your index using the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation or using one of the supported data sources.
//
// TODO: EXAMPLE
//
type CfnIndex interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrId() *string
	CapacityUnits() interface{}
	SetCapacityUnits(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DocumentMetadataConfigurations() interface{}
	SetDocumentMetadataConfigurations(val interface{})
	Edition() *string
	SetEdition(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	ServerSideEncryptionConfiguration() interface{}
	SetServerSideEncryptionConfiguration(val interface{})
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UserContextPolicy() *string
	SetUserContextPolicy(val *string)
	UserTokenConfigurations() interface{}
	SetUserTokenConfigurations(val interface{})
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

// The jsii proxy struct for CfnIndex
type jsiiProxy_CfnIndex struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIndex) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CapacityUnits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) DocumentMetadataConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"documentMetadataConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) ServerSideEncryptionConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverSideEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserContextPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userContextPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIndex) UserTokenConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userTokenConfigurations",
		&returns,
	)
	return returns
}


// Create a new `AWS::Kendra::Index`.
func NewCfnIndex(scope constructs.Construct, id *string, props *CfnIndexProps) CfnIndex {
	_init_.Initialize()

	j := jsiiProxy_CfnIndex{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Kendra::Index`.
func NewCfnIndex_Override(c CfnIndex, scope constructs.Construct, id *string, props *CfnIndexProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIndex) SetCapacityUnits(val interface{}) {
	_jsii_.Set(
		j,
		"capacityUnits",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetDocumentMetadataConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"documentMetadataConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetEdition(val *string) {
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetServerSideEncryptionConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"serverSideEncryptionConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetUserContextPolicy(val *string) {
	_jsii_.Set(
		j,
		"userContextPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnIndex) SetUserTokenConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"userTokenConfigurations",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnIndex_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnIndex_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
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
func CfnIndex_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIndex_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kendra.CfnIndex",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnIndex) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIndex) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIndex) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIndex) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnIndex) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnIndex) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnIndex) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIndex) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIndex) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIndex) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnIndex) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIndex) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIndex) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnIndex) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIndex) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies capacity units configured for your enterprise edition index.
//
// You can add and remove capacity units to tune an index to your requirements.
//
// TODO: EXAMPLE
//
type CfnIndex_CapacityUnitsConfigurationProperty struct {
	// The amount of extra query capacity for an index and [GetQuerySuggestions](https://docs.aws.amazon.com/kendra/latest/dg/API_GetQuerySuggestions.html) capacity.
	//
	// A single extra capacity unit for an index provides 0.1 queries per second or approximately 8,000 queries per day.
	//
	// `GetQuerySuggestions` capacity is five times the provisioned query capacity for an index, or the base capacity of 2.5 calls per second, whichever is higher. For example, the base capacity for an index is 0.1 queries per second, and `GetQuerySuggestions` capacity has a base of 2.5 calls per second. If you add another 0.1 queries per second to total 0.2 queries per second for an index, the `GetQuerySuggestions` capacity is 2.5 calls per second (higher than five times 0.2 queries per second).
	QueryCapacityUnits *float64 `json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// The amount of extra storage capacity for an index.
	//
	// A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first.
	StorageCapacityUnits *float64 `json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

// Specifies the properties of a custom index field.
//
// TODO: EXAMPLE
//
type CfnIndex_DocumentMetadataConfigurationProperty struct {
	// The name of the index field.
	Name *string `json:"name" yaml:"name"`
	// The data type of the index field.
	Type *string `json:"type" yaml:"type"`
	// Provides manual tuning parameters to determine how the field affects the search results.
	Relevance interface{} `json:"relevance" yaml:"relevance"`
	// Provides information about how the field is used during a search.
	Search interface{} `json:"search" yaml:"search"`
}

// Configuration information for the JSON token type.
//
// TODO: EXAMPLE
//
type CfnIndex_JsonTokenTypeConfigurationProperty struct {
	// The group attribute field.
	GroupAttributeField *string `json:"groupAttributeField" yaml:"groupAttributeField"`
	// The user name attribute field.
	UserNameAttributeField *string `json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

// Configuration information for the JWT token type.
//
// TODO: EXAMPLE
//
type CfnIndex_JwtTokenTypeConfigurationProperty struct {
	// The location of the key.
	KeyLocation *string `json:"keyLocation" yaml:"keyLocation"`
	// The regular expression that identifies the claim.
	ClaimRegex *string `json:"claimRegex" yaml:"claimRegex"`
	// The group attribute field.
	GroupAttributeField *string `json:"groupAttributeField" yaml:"groupAttributeField"`
	// The issuer of the token.
	Issuer *string `json:"issuer" yaml:"issuer"`
	// The Amazon Resource Name (arn) of the secret.
	SecretManagerArn *string `json:"secretManagerArn" yaml:"secretManagerArn"`
	// The signing key URL.
	Url *string `json:"url" yaml:"url"`
	// The user name attribute field.
	UserNameAttributeField *string `json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

// Provides information for manually tuning the relevance of a field in a search.
//
// When a query includes terms that match the field, the results are given a boost in the response based on these tuning parameters.
//
// TODO: EXAMPLE
//
type CfnIndex_RelevanceProperty struct {
	// Specifies the time period that the boost applies to.
	//
	// For example, to make the boost apply to documents with the field value within the last month, you would use "2628000s". Once the field value is beyond the specified range, the effect of the boost drops off. The higher the importance, the faster the effect drops off. If you don't specify a value, the default is 3 months. The value of the field is a numeric string followed by the character "s", for example "86400s" for one day, or "604800s" for one week.
	//
	// Only applies to `DATE` fields.
	Duration *string `json:"duration" yaml:"duration"`
	// Indicates that this field determines how "fresh" a document is.
	//
	// For example, if document 1 was created on November 5, and document 2 was created on October 31, document 1 is "fresher" than document 2. You can only set the `Freshness` field on one `DATE` type field. Only applies to `DATE` fields.
	Freshness interface{} `json:"freshness" yaml:"freshness"`
	// The relative importance of the field in the search.
	//
	// Larger numbers provide more of a boost than smaller numbers.
	Importance *float64 `json:"importance" yaml:"importance"`
	// Determines how values should be interpreted.
	//
	// When the `RankOrder` field is `ASCENDING` , higher numbers are better. For example, a document with a rating score of 10 is higher ranking than a document with a rating score of 1.
	//
	// When the `RankOrder` field is `DESCENDING` , lower numbers are better. For example, in a task tracking application, a priority 1 task is more important than a priority 5 task.
	//
	// Only applies to `LONG` and `DOUBLE` fields.
	RankOrder *string `json:"rankOrder" yaml:"rankOrder"`
	// An array of key-value pairs that contains an array of values that should be given a different boost when they appear in the search result list.
	//
	// For example, if you are boosting query terms that match the department field in the result, query terms that match the department field are boosted in the result. You can add entries from the department field to boost documents with those values higher.
	//
	// For example, you can add entries to the map with names of departments. If you add "HR", 5 and "Legal",3 those departments are given special attention when they appear in the metadata of a document.
	ValueImportanceItems interface{} `json:"valueImportanceItems" yaml:"valueImportanceItems"`
}

// Provides information about how a custom index field is used during a search.
//
// TODO: EXAMPLE
//
type CfnIndex_SearchProperty struct {
	// Determines whether the field is returned in the query response.
	//
	// The default is `true` .
	Displayable interface{} `json:"displayable" yaml:"displayable"`
	// Indicates that the field can be used to create search facets, a count of results for each value in the field.
	//
	// The default is `false` .
	Facetable interface{} `json:"facetable" yaml:"facetable"`
	// Determines whether the field is used in the search.
	//
	// If the `Searchable` field is `true` , you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is `true` for string fields and `false` for number and date fields.
	Searchable interface{} `json:"searchable" yaml:"searchable"`
	// Indicates that the field can be used to sort the search results.
	//
	// The default is `false` .
	Sortable interface{} `json:"sortable" yaml:"sortable"`
}

// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt data indexed by Amazon Kendra.
//
// We suggest that you use a CMK from your account to help secure your index. Amazon Kendra doesn't support asymmetric CMKs.
//
// TODO: EXAMPLE
//
type CfnIndex_ServerSideEncryptionConfigurationProperty struct {
	// The identifier of the AWS KMS customer master key (CMK).
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
}

// Provides configuration information for a token configuration.
//
// TODO: EXAMPLE
//
type CfnIndex_UserTokenConfigurationProperty struct {
	// Information about the JSON token type configuration.
	JsonTokenTypeConfiguration interface{} `json:"jsonTokenTypeConfiguration" yaml:"jsonTokenTypeConfiguration"`
	// Information about the JWT token type configuration.
	JwtTokenTypeConfiguration interface{} `json:"jwtTokenTypeConfiguration" yaml:"jwtTokenTypeConfiguration"`
}

// Specifies a key-value pair that determines the search boost value that a document receives when the key is part of the metadata of a document.
//
// TODO: EXAMPLE
//
type CfnIndex_ValueImportanceItemProperty struct {
	// The document metadata value that receives the search boost.
	Key *string `json:"key" yaml:"key"`
	// The boost value that a document receives when the key is part of the metadata of a document.
	Value *float64 `json:"value" yaml:"value"`
}

// Properties for defining a `CfnIndex`.
//
// TODO: EXAMPLE
//
type CfnIndexProps struct {
	// Indicates whether the index is a enterprise edition index or a developer edition index.
	//
	// Valid values are `DEVELOPER_EDITION` and `ENTERPRISE_EDITION` .
	Edition *string `json:"edition" yaml:"edition"`
	// The name of the index.
	Name *string `json:"name" yaml:"name"`
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the role used when you use the [BatchPutDocument](https://docs.aws.amazon.com/kendra/latest/dg/BatchPutDocument.html) operation to index documents from an Amazon S3 bucket.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Specifies capacity units configured for your index.
	//
	// You can add and remove capacity units to tune an index to your requirements. You can set capacity units only for Enterprise edition indexes.
	CapacityUnits interface{} `json:"capacityUnits" yaml:"capacityUnits"`
	// A description of the index.
	Description *string `json:"description" yaml:"description"`
	// Specifies the properties of an index field.
	//
	// You can add either a custom or a built-in field. You can add and remove built-in fields at any time. When a built-in field is removed it's configuration reverts to the default for the field. Custom fields can't be removed from an index after they are added.
	DocumentMetadataConfigurations interface{} `json:"documentMetadataConfigurations" yaml:"documentMetadataConfigurations"`
	// The identifier of the AWS KMS customer managed key (CMK) to use to encrypt data indexed by Amazon Kendra.
	//
	// Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration interface{} `json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The user context policy.
	//
	// ATTRIBUTE_FILTER
	//
	// - All indexed content is searchable and displayable for all users. If there is an access control list, it is ignored. You can filter on user and group attributes.
	//
	// USER_TOKEN
	//
	// - Enables SSO and token-based user access control. All documents with no access control and all documents accessible to the user will be searchable and displayable.
	UserContextPolicy *string `json:"userContextPolicy" yaml:"userContextPolicy"`
	// Defines the type of user token used for the index.
	UserTokenConfigurations interface{} `json:"userTokenConfigurations" yaml:"userTokenConfigurations"`
}

