package awslightsail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Lightsail::Database`.
//
// The `AWS::Lightsail::Database` resource specifies an Amazon Lightsail database.
//
// TODO: EXAMPLE
//
type CfnDatabase interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrDatabaseArn() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	BackupRetention() interface{}
	SetBackupRetention(val interface{})
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	MasterDatabaseName() *string
	SetMasterDatabaseName(val *string)
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	Node() constructs.Node
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	Ref() *string
	RelationalDatabaseBlueprintId() *string
	SetRelationalDatabaseBlueprintId(val *string)
	RelationalDatabaseBundleId() *string
	SetRelationalDatabaseBundleId(val *string)
	RelationalDatabaseName() *string
	SetRelationalDatabaseName(val *string)
	RelationalDatabaseParameters() interface{}
	SetRelationalDatabaseParameters(val interface{})
	RotateMasterUserPassword() interface{}
	SetRotateMasterUserPassword(val interface{})
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

// The jsii proxy struct for CfnDatabase
type jsiiProxy_CfnDatabase struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatabase) AttrDatabaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatabaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) BackupRetention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) MasterDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseBlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBlueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseBundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseBundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relationalDatabaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RelationalDatabaseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalDatabaseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) RotateMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Database`.
func NewCfnDatabase(scope constructs.Construct, id *string, props *CfnDatabaseProps) CfnDatabase {
	_init_.Initialize()

	j := jsiiProxy_CfnDatabase{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Database`.
func NewCfnDatabase_Override(c CfnDatabase, scope constructs.Construct, id *string, props *CfnDatabaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatabase) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetBackupRetention(val interface{}) {
	_jsii_.Set(
		j,
		"backupRetention",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetCaCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetMasterDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"masterDatabaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseBlueprintId(val *string) {
	_jsii_.Set(
		j,
		"relationalDatabaseBlueprintId",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseBundleId(val *string) {
	_jsii_.Set(
		j,
		"relationalDatabaseBundleId",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"relationalDatabaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRelationalDatabaseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"relationalDatabaseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetRotateMasterUserPassword(val interface{}) {
	_jsii_.Set(
		j,
		"rotateMasterUserPassword",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatabase_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatabase_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
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
func CfnDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatabase_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnDatabase",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDatabase) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDatabase) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDatabase) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDatabase) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDatabase) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDatabase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDatabase) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDatabase) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDatabase) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatabase) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDatabase) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `RelationalDatabaseParameter` is a property of the [AWS::Lightsail::Database](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-database.html) resource. It describes parameters for the database.
//
// TODO: EXAMPLE
//
type CfnDatabase_RelationalDatabaseParameterProperty struct {
	// The valid range of values for the parameter.
	AllowedValues *string `json:"allowedValues" yaml:"allowedValues"`
	// Indicates when parameter updates are applied.
	//
	// Can be `immediate` or `pending-reboot` .
	ApplyMethod *string `json:"applyMethod" yaml:"applyMethod"`
	// Specifies the engine-specific parameter type.
	ApplyType *string `json:"applyType" yaml:"applyType"`
	// The valid data type of the parameter.
	DataType *string `json:"dataType" yaml:"dataType"`
	// A description of the parameter.
	Description *string `json:"description" yaml:"description"`
	// A Boolean value indicating whether the parameter can be modified.
	IsModifiable interface{} `json:"isModifiable" yaml:"isModifiable"`
	// The name of the parameter.
	ParameterName *string `json:"parameterName" yaml:"parameterName"`
	// The value for the parameter.
	ParameterValue *string `json:"parameterValue" yaml:"parameterValue"`
}

// Properties for defining a `CfnDatabase`.
//
// TODO: EXAMPLE
//
type CfnDatabaseProps struct {
	// The meaning of this parameter differs according to the database engine you use.
	//
	// *MySQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, no database is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-64 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in MySQL, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , and [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// The name of the database to create when the Lightsail database resource is created. If this parameter isn't specified, a database named `postgres` is created in the database resource.
	//
	// Constraints:
	//
	// - Must contain 1-63 letters or numbers.
	// - Must begin with a letter. Subsequent characters can be letters, underscores, or numbers (0-9).
	// - Can't be a word reserved by the specified database engine.
	//
	// For more information about reserved words in PostgreSQL, see the SQL Key Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	MasterDatabaseName *string `json:"masterDatabaseName" yaml:"masterDatabaseName"`
	// The name for the primary user.
	//
	// *MySQL*
	//
	// Constraints:
	//
	// - Required for MySQL.
	// - Must be 1-16 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [MySQL 5.6](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.6/en/keywords.html) , [MySQL 5.7](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/5.7/en/keywords.html) , or [MySQL 8.0](https://docs.aws.amazon.com/https://dev.mysql.com/doc/refman/8.0/en/keywords.html) .
	//
	// *PostgreSQL*
	//
	// Constraints:
	//
	// - Required for PostgreSQL.
	// - Must be 1-63 letters or numbers. Can contain underscores.
	// - First character must be a letter.
	// - Can't be a reserved word for the chosen database engine.
	//
	// For more information about reserved words in MySQL 5.6 or 5.7, see the Keywords and Reserved Words articles for [PostgreSQL 9.6](https://docs.aws.amazon.com/https://www.postgresql.org/docs/9.6/sql-keywords-appendix.html) , [PostgreSQL 10](https://docs.aws.amazon.com/https://www.postgresql.org/docs/10/sql-keywords-appendix.html) , [PostgreSQL 11](https://docs.aws.amazon.com/https://www.postgresql.org/docs/11/sql-keywords-appendix.html) , and [PostgreSQL 12](https://docs.aws.amazon.com/https://www.postgresql.org/docs/12/sql-keywords-appendix.html) .
	MasterUsername *string `json:"masterUsername" yaml:"masterUsername"`
	// The blueprint ID for the database (for example, `mysql_8_0` ).
	RelationalDatabaseBlueprintId *string `json:"relationalDatabaseBlueprintId" yaml:"relationalDatabaseBlueprintId"`
	// The bundle ID for the database (for example, `medium_1_0` ).
	RelationalDatabaseBundleId *string `json:"relationalDatabaseBundleId" yaml:"relationalDatabaseBundleId"`
	// The name of the instance.
	RelationalDatabaseName *string `json:"relationalDatabaseName" yaml:"relationalDatabaseName"`
	// The Availability Zone for the database.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// A Boolean value indicating whether automated backup retention is enabled for the database.
	BackupRetention interface{} `json:"backupRetention" yaml:"backupRetention"`
	// The certificate associated with the database.
	CaCertificateIdentifier *string `json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// The password for the primary user of the database.
	//
	// The password can include any printable ASCII character except the following: /, ", or @. It cannot contain spaces.
	//
	// > The `MasterUserPassword` and `RotateMasterUserPassword` parameters cannot be used together in the same template.
	//
	// *MySQL*
	//
	// Constraints: Must contain 8-41 characters.
	//
	// *PostgreSQL*
	//
	// Constraints: Must contain 8-128 characters.
	MasterUserPassword *string `json:"masterUserPassword" yaml:"masterUserPassword"`
	// The daily time range during which automated backups are created for the database (for example, `16:00-16:30` ).
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur for the database, formatted as follows: `ddd:hh24:mi-ddd:hh24:mi` .
	//
	// For example, `Tue:17:00-Tue:17:30` .
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A Boolean value indicating whether the database is accessible to anyone on the internet.
	PubliclyAccessible interface{} `json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An array of parameters for the database.
	RelationalDatabaseParameters interface{} `json:"relationalDatabaseParameters" yaml:"relationalDatabaseParameters"`
	// A Boolean value indicating whether to change the primary user password to a new, strong password generated by Lightsail .
	//
	// > The `RotateMasterUserPassword` and `MasterUserPassword` parameters cannot be used together in the same template.
	RotateMasterUserPassword interface{} `json:"rotateMasterUserPassword" yaml:"rotateMasterUserPassword"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Disk`.
//
// The `AWS::Lightsail::Disk` resource specifies a disk that can be attached to an Amazon Lightsail instance that is in the same AWS Region and Availability Zone.
//
// TODO: EXAMPLE
//
type CfnDisk interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AddOns() interface{}
	SetAddOns(val interface{})
	AttrAttachedTo() *string
	AttrAttachmentState() *string
	AttrDiskArn() *string
	AttrIops() *float64
	AttrIsAttached() awscdk.IResolvable
	AttrPath() *string
	AttrResourceType() *string
	AttrState() *string
	AttrSupportCode() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DiskName() *string
	SetDiskName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SizeInGb() *float64
	SetSizeInGb(val *float64)
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

// The jsii proxy struct for CfnDisk
type jsiiProxy_CfnDisk struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDisk) AddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrAttachedTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrAttachmentState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAttachmentState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrDiskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDiskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrIsAttached() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AttrSupportCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSupportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) DiskName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) SizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDisk) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Disk`.
func NewCfnDisk(scope constructs.Construct, id *string, props *CfnDiskProps) CfnDisk {
	_init_.Initialize()

	j := jsiiProxy_CfnDisk{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Disk`.
func NewCfnDisk_Override(c CfnDisk, scope constructs.Construct, id *string, props *CfnDiskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDisk) SetAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"addOns",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetDiskName(val *string) {
	_jsii_.Set(
		j,
		"diskName",
		val,
	)
}

func (j *jsiiProxy_CfnDisk) SetSizeInGb(val *float64) {
	_jsii_.Set(
		j,
		"sizeInGb",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDisk_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDisk_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
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
func CfnDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDisk_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnDisk",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDisk) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDisk) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDisk) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDisk) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDisk) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDisk) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDisk) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDisk) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDisk) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDisk) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDisk) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDisk) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDisk) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDisk) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AddOn` is a property of the [AWS::Lightsail::Disk](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-disk.html) resource. It describes the add-ons for a disk.
//
// TODO: EXAMPLE
//
type CfnDisk_AddOnProperty struct {
	// The add-on type (for example, `AutoSnapshot` ).
	//
	// > `AutoSnapshot` is the only add-on that can be enabled for a disk.
	AddOnType *string `json:"addOnType" yaml:"addOnType"`
	// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
	AutoSnapshotAddOnRequest interface{} `json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// The status of the add-on.
	//
	// Valid Values: `Enabled` | `Disabled`
	Status *string `json:"status" yaml:"status"`
}

// `AutoSnapshotAddOn` is a property of the [AddOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html) property. It describes the automatic snapshot add-on for a disk.
//
// TODO: EXAMPLE
//
type CfnDisk_AutoSnapshotAddOnProperty struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Constraints:
	//
	// - Must be in `HH:00` format, and in an hourly increment.
	// - Specified in Coordinated Universal Time (UTC).
	// - The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTimeOfDay *string `json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

// Properties for defining a `CfnDisk`.
//
// TODO: EXAMPLE
//
type CfnDiskProps struct {
	// The name of the disk.
	DiskName *string `json:"diskName" yaml:"diskName"`
	// The size of the disk in GB.
	SizeInGb *float64 `json:"sizeInGb" yaml:"sizeInGb"`
	// An array of add-ons for the disk.
	//
	// > If the disk has an add-on enabled when performing a delete disk request, the add-on is automatically disabled before the disk is deleted.
	AddOns interface{} `json:"addOns" yaml:"addOns"`
	// The AWS Region and Availability Zone location for the disk (for example, `us-east-1a` ).
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Lightsail::Instance`.
//
// The `AWS::Lightsail::Instance` resource specifies an Amazon Lightsail instance.
//
// TODO: EXAMPLE
//
type CfnInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AddOns() interface{}
	SetAddOns(val interface{})
	AttrHardwareCpuCount() *float64
	AttrHardwareRamSizeInGb() *float64
	AttrInstanceArn() *string
	AttrIsStaticIp() awscdk.IResolvable
	AttrLocationAvailabilityZone() *string
	AttrLocationRegionName() *string
	AttrNetworkingMonthlyTransferGbPerMonthAllocated() *string
	AttrPrivateIpAddress() *string
	AttrPublicIpAddress() *string
	AttrResourceType() *string
	AttrSshKeyName() *string
	AttrStateCode() *float64
	AttrStateName() *string
	AttrSupportCode() *string
	AttrUserName() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	BlueprintId() *string
	SetBlueprintId(val *string)
	BundleId() *string
	SetBundleId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Hardware() interface{}
	SetHardware(val interface{})
	InstanceName() *string
	SetInstanceName(val *string)
	KeyPairName() *string
	SetKeyPairName(val *string)
	LogicalId() *string
	Networking() interface{}
	SetNetworking(val interface{})
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UserData() *string
	SetUserData(val *string)
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

// The jsii proxy struct for CfnInstance
type jsiiProxy_CfnInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstance) AddOns() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addOns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrHardwareCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrHardwareCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrHardwareRamSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrHardwareRamSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrIsStaticIp() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsStaticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrLocationAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrLocationRegionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLocationRegionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrNetworkingMonthlyTransferGbPerMonthAllocated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkingMonthlyTransferGbPerMonthAllocated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPrivateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrPublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPublicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrSshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrStateCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrStateCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrStateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrSupportCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSupportCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AttrUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BlueprintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blueprintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) BundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Hardware() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardware",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) InstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) KeyPairName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPairName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Networking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstance) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::Instance`.
func NewCfnInstance(scope constructs.Construct, id *string, props *CfnInstanceProps) CfnInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::Instance`.
func NewCfnInstance_Override(c CfnInstance, scope constructs.Construct, id *string, props *CfnInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstance) SetAddOns(val interface{}) {
	_jsii_.Set(
		j,
		"addOns",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBlueprintId(val *string) {
	_jsii_.Set(
		j,
		"blueprintId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetBundleId(val *string) {
	_jsii_.Set(
		j,
		"bundleId",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetHardware(val interface{}) {
	_jsii_.Set(
		j,
		"hardware",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetInstanceName(val *string) {
	_jsii_.Set(
		j,
		"instanceName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetKeyPairName(val *string) {
	_jsii_.Set(
		j,
		"keyPairName",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetNetworking(val interface{}) {
	_jsii_.Set(
		j,
		"networking",
		val,
	)
}

func (j *jsiiProxy_CfnInstance) SetUserData(val *string) {
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
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
func CfnInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnInstance) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnInstance) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnInstance) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnInstance) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnInstance) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnInstance) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `AddOn` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the add-ons for an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_AddOnProperty struct {
	// The add-on type (for example, `AutoSnapshot` ).
	//
	// > `AutoSnapshot` is the only add-on that can be enabled for an instance.
	AddOnType *string `json:"addOnType" yaml:"addOnType"`
	// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
	AutoSnapshotAddOnRequest interface{} `json:"autoSnapshotAddOnRequest" yaml:"autoSnapshotAddOnRequest"`
	// The status of the add-on.
	//
	// Valid Values: `Enabled` | `Disabled`
	Status *string `json:"status" yaml:"status"`
}

// `AutoSnapshotAddOn` is a property of the [AddOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-addon.html) property. It describes the automatic snapshot add-on for an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_AutoSnapshotAddOnProperty struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Constraints:
	//
	// - Must be in `HH:00` format, and in an hourly increment.
	// - Specified in Coordinated Universal Time (UTC).
	// - The snapshot will be automatically created between the time specified and up to 45 minutes after.
	SnapshotTimeOfDay *string `json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

// `Disk` is a property of the [Hardware](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-hardware.html) property. It describes a disk attached to an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_DiskProperty struct {
	// The unique name of the disk.
	DiskName *string `json:"diskName" yaml:"diskName"`
	// The disk path.
	Path *string `json:"path" yaml:"path"`
	// The resources to which the disk is attached.
	AttachedTo *string `json:"attachedTo" yaml:"attachedTo"`
	// (Deprecated) The attachment state of the disk.
	//
	// > In releases prior to November 14, 2017, this parameter returned `attached` for system disks in the API response. It is now deprecated, but still included in the response. Use `isAttached` instead.
	AttachmentState *string `json:"attachmentState" yaml:"attachmentState"`
	// The input/output operations per second (IOPS) of the disk.
	Iops *float64 `json:"iops" yaml:"iops"`
	// A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it).
	IsSystemDisk interface{} `json:"isSystemDisk" yaml:"isSystemDisk"`
	// The size of the disk in GB.
	SizeInGb *string `json:"sizeInGb" yaml:"sizeInGb"`
}

// `Hardware` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
//
// TODO: EXAMPLE
//
type CfnInstance_HardwareProperty struct {
	// The number of vCPUs the instance has.
	//
	// > The `CpuCount` property is read-only and should not be specified in a create instance or update instance request.
	CpuCount *float64 `json:"cpuCount" yaml:"cpuCount"`
	// The disks attached to the instance.
	//
	// The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	Disks interface{} `json:"disks" yaml:"disks"`
	// The amount of RAM in GB on the instance (for example, `1.0` ).
	//
	// > The `RamSizeInGb` property is read-only and should not be specified in a create instance or update instance request.
	RamSizeInGb *float64 `json:"ramSizeInGb" yaml:"ramSizeInGb"`
}

// `Location` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the location for an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_LocationProperty struct {
	// The Availability Zone for the instance.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the AWS Region for the instance.
	RegionName *string `json:"regionName" yaml:"regionName"`
}

// `MonthlyTransfer` is a property of the [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html) property. It describes the amount of allocated monthly data transfer (in GB) for an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_MonthlyTransferProperty struct {
	// The amount of allocated monthly data transfer (in GB) for an instance.
	GbPerMonthAllocated *string `json:"gbPerMonthAllocated" yaml:"gbPerMonthAllocated"`
}

// `Networking` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the public ports and the monthly amount of data transfer allocated for the instance.
//
// TODO: EXAMPLE
//
type CfnInstance_NetworkingProperty struct {
	// An array of ports to open on the instance.
	Ports interface{} `json:"ports" yaml:"ports"`
	// The monthly amount of data transfer, in GB, allocated for the instance.
	MonthlyTransfer *float64 `json:"monthlyTransfer" yaml:"monthlyTransfer"`
}

// `Port` is a property of the [Networking](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instance-networking.html) property. It describes information about ports for an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_PortProperty struct {
	// The access direction ( `inbound` or `outbound` ).
	//
	// > Lightsail currently supports only `inbound` access direction.
	AccessDirection *string `json:"accessDirection" yaml:"accessDirection"`
	// The location from which access is allowed.
	//
	// For example, `Anywhere (0.0.0.0/0)` , or `Custom` if a specific IP address or range of IP addresses is allowed.
	AccessFrom *string `json:"accessFrom" yaml:"accessFrom"`
	// The type of access ( `Public` or `Private` ).
	AccessType *string `json:"accessType" yaml:"accessType"`
	// An alias that defines access for a preconfigured range of IP addresses.
	//
	// The only alias currently supported is `lightsail-connect` , which allows IP addresses of the browser-based RDP/SSH client in the Lightsail console to connect to your instance.
	CidrListAliases *[]*string `json:"cidrListAliases" yaml:"cidrListAliases"`
	// The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol.
	//
	// > The `ipv6Cidrs` parameter lists the IPv6 addresses that are allowed to connect to an instance.
	//
	// Examples:
	//
	// - To allow the IP address `192.0.2.44` , specify `192.0.2.44` or `192.0.2.44/32` .
	// - To allow the IP addresses `192.0.2.0` to `192.0.2.255` , specify `192.0.2.0/24` .
	Cidrs *[]*string `json:"cidrs" yaml:"cidrs"`
	// The common name of the port information.
	CommonName *string `json:"commonName" yaml:"commonName"`
	// The first port in a range of open ports on an instance.
	//
	// Allowed ports:
	//
	// - TCP and UDP - `0` to `65535`
	// - ICMP - The ICMP type for IPv4 addresses. For example, specify `8` as the `fromPort` (ICMP type), and `-1` as the `toPort` (ICMP code), to enable ICMP Ping.
	// - ICMPv6 - The ICMP type for IPv6 addresses. For example, specify `128` as the `fromPort` (ICMPv6 type), and `0` as `toPort` (ICMPv6 code).
	FromPort *float64 `json:"fromPort" yaml:"fromPort"`
	// The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol.
	//
	// Only devices with an IPv6 address can connect to an instance through IPv6; otherwise, IPv4 should be used.
	//
	// > The `cidrs` parameter lists the IPv4 addresses that are allowed to connect to an instance.
	Ipv6Cidrs *[]*string `json:"ipv6Cidrs" yaml:"ipv6Cidrs"`
	// The IP protocol name.
	//
	// The name can be one of the following:
	//
	// - `tcp` - Transmission Control Protocol (TCP) provides reliable, ordered, and error-checked delivery of streamed data between applications running on hosts communicating by an IP network. If you have an application that doesn't require reliable data stream service, use UDP instead.
	// - `all` - All transport layer protocol types.
	// - `udp` - With User Datagram Protocol (UDP), computer applications can send messages (or datagrams) to other hosts on an Internet Protocol (IP) network. Prior communications are not required to set up transmission channels or data paths. Applications that don't require reliable data stream service can use UDP, which provides a connectionless datagram service that emphasizes reduced latency over reliability. If you do require reliable data stream service, use TCP instead.
	// - `icmp` - Internet Control Message Protocol (ICMP) is used to send error messages and operational information indicating success or failure when communicating with an instance. For example, an error is indicated when an instance could not be reached. When you specify `icmp` as the `protocol` , you must specify the ICMP type using the `fromPort` parameter, and ICMP code using the `toPort` parameter.
	Protocol *string `json:"protocol" yaml:"protocol"`
	// The last port in a range of open ports on an instance.
	//
	// Allowed ports:
	//
	// - TCP and UDP - `0` to `65535`
	// - ICMP - The ICMP code for IPv4 addresses. For example, specify `8` as the `fromPort` (ICMP type), and `-1` as the `toPort` (ICMP code), to enable ICMP Ping.
	// - ICMPv6 - The ICMP code for IPv6 addresses. For example, specify `128` as the `fromPort` (ICMPv6 type), and `0` as `toPort` (ICMPv6 code).
	ToPort *float64 `json:"toPort" yaml:"toPort"`
}

// `State` is a property of the [AWS::Lightsail::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-instance.html) resource. It describes the status code and the state (for example, `running` ) of an instance.
//
// TODO: EXAMPLE
//
type CfnInstance_StateProperty struct {
	// The status code of the instance.
	Code *float64 `json:"code" yaml:"code"`
	// The state of the instance (for example, `running` or `pending` ).
	Name *string `json:"name" yaml:"name"`
}

// Properties for defining a `CfnInstance`.
//
// TODO: EXAMPLE
//
type CfnInstanceProps struct {
	// The blueprint ID for the instance (for example, `os_amlinux_2016_03` ).
	BlueprintId *string `json:"blueprintId" yaml:"blueprintId"`
	// The bundle ID for the instance (for example, `micro_1_0` ).
	BundleId *string `json:"bundleId" yaml:"bundleId"`
	// The name of the instance.
	InstanceName *string `json:"instanceName" yaml:"instanceName"`
	// An array of add-ons for the instance.
	//
	// > If the instance has an add-on enabled when performing a delete instance request, the add-on is automatically disabled before the instance is deleted.
	AddOns interface{} `json:"addOns" yaml:"addOns"`
	// The Availability Zone for the instance.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The hardware properties for the instance, such as the vCPU count, attached disks, and amount of RAM.
	//
	// > The instance restarts when performing an attach disk or detach disk request. This resets the public IP address of your instance if a static IP isn't attached to it.
	Hardware interface{} `json:"hardware" yaml:"hardware"`
	// The name of the key pair to use for the instance.
	//
	// If no key pair name is specified, the Regional Lightsail default key pair is used.
	KeyPairName *string `json:"keyPairName" yaml:"keyPairName"`
	// The public ports and the monthly amount of data transfer allocated for the instance.
	Networking interface{} `json:"networking" yaml:"networking"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	//
	// > The `Value` of `Tags` is optional for Lightsail resources.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// The optional launch script for the instance.
	//
	// Specify a launch script to configure an instance with additional user data. For example, you might want to specify `apt-get -y update` as a launch script.
	//
	// > Depending on the blueprint of your instance, the command to get software on your instance varies. Amazon Linux and CentOS use `yum` , Debian and Ubuntu use `apt-get` , and FreeBSD uses `pkg` .
	UserData *string `json:"userData" yaml:"userData"`
}

// A CloudFormation `AWS::Lightsail::StaticIp`.
//
// The `AWS::Lightsail::StaticIp` resource specifies a static IP that can be attached to an Amazon Lightsail instance that is in the same AWS Region and Availability Zone.
//
// TODO: EXAMPLE
//
type CfnStaticIp interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttachedTo() *string
	SetAttachedTo(val *string)
	AttrIpAddress() *string
	AttrIsAttached() awscdk.IResolvable
	AttrStaticIpArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	StaticIpName() *string
	SetStaticIpName(val *string)
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

// The jsii proxy struct for CfnStaticIp
type jsiiProxy_CfnStaticIp struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStaticIp) AttachedTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachedTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) AttrIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) AttrIsAttached() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrIsAttached",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) AttrStaticIpArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStaticIpArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) StaticIpName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"staticIpName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStaticIp) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lightsail::StaticIp`.
func NewCfnStaticIp(scope constructs.Construct, id *string, props *CfnStaticIpProps) CfnStaticIp {
	_init_.Initialize()

	j := jsiiProxy_CfnStaticIp{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lightsail::StaticIp`.
func NewCfnStaticIp_Override(c CfnStaticIp, scope constructs.Construct, id *string, props *CfnStaticIpProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStaticIp) SetAttachedTo(val *string) {
	_jsii_.Set(
		j,
		"attachedTo",
		val,
	)
}

func (j *jsiiProxy_CfnStaticIp) SetStaticIpName(val *string) {
	_jsii_.Set(
		j,
		"staticIpName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStaticIp_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStaticIp_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
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
func CfnStaticIp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStaticIp_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lightsail.CfnStaticIp",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnStaticIp) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnStaticIp) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnStaticIp) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnStaticIp) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnStaticIp) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnStaticIp) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnStaticIp) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnStaticIp) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnStaticIp) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStaticIp) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnStaticIp) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStaticIp) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnStaticIp) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnStaticIp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStaticIp) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnStaticIp`.
//
// TODO: EXAMPLE
//
type CfnStaticIpProps struct {
	// The name of the static IP.
	StaticIpName *string `json:"staticIpName" yaml:"staticIpName"`
	// The instance that the static IP is attached to.
	AttachedTo *string `json:"attachedTo" yaml:"attachedTo"`
}

