package awsredshiftserverless

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshiftserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A collection of database objects and users.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var namespaceResourcePolicy interface{}
//
//   cfnNamespace := awscdk.Aws_redshiftserverless.NewCfnNamespace(this, jsii.String("MyCfnNamespace"), &CfnNamespaceProps{
//   	NamespaceName: jsii.String("namespaceName"),
//
//   	// the properties below are optional
//   	AdminPasswordSecretKmsKeyId: jsii.String("adminPasswordSecretKmsKeyId"),
//   	AdminUsername: jsii.String("adminUsername"),
//   	AdminUserPassword: jsii.String("adminUserPassword"),
//   	DbName: jsii.String("dbName"),
//   	DefaultIamRoleArn: jsii.String("defaultIamRoleArn"),
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	FinalSnapshotRetentionPeriod: jsii.Number(123),
//   	IamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	LogExports: []*string{
//   		jsii.String("logExports"),
//   	},
//   	ManageAdminPassword: jsii.Boolean(false),
//   	NamespaceResourcePolicy: namespaceResourcePolicy,
//   	RedshiftIdcApplicationArn: jsii.String("redshiftIdcApplicationArn"),
//   	SnapshotCopyConfigurations: []interface{}{
//   		&SnapshotCopyConfigurationProperty{
//   			DestinationRegion: jsii.String("destinationRegion"),
//
//   			// the properties below are optional
//   			DestinationKmsKeyId: jsii.String("destinationKmsKeyId"),
//   			SnapshotRetentionPeriod: jsii.Number(123),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-namespace.html
//
type CfnNamespace interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret.
	AdminPasswordSecretKmsKeyId() *string
	SetAdminPasswordSecretKmsKeyId(val *string)
	// The username of the administrator for the primary database created in the namespace.
	AdminUsername() *string
	SetAdminUsername(val *string)
	// The password of the administrator for the primary database created in the namespace.
	AdminUserPassword() *string
	SetAdminUserPassword(val *string)
	AttrNamespace() awscdk.IResolvable
	// The username of the administrator for the first database created in the namespace.
	AttrNamespaceAdminUsername() *string
	// The date of when the namespace was created.
	AttrNamespaceCreationDate() *string
	// The name of the first database created in the namespace.
	AttrNamespaceDbName() *string
	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.
	AttrNamespaceDefaultIamRoleArn() *string
	// A list of IAM roles to associate with the namespace.
	AttrNamespaceIamRoles() *[]*string
	// The ID of the AWS Key Management Service key used to encrypt your data.
	AttrNamespaceKmsKeyId() *string
	// The types of logs the namespace can export.
	//
	// Available export types are `User log` , `Connection log` , and `User activity log` .
	AttrNamespaceLogExports() *[]*string
	// The Amazon Resource Name (ARN) associated with a namespace.
	AttrNamespaceNamespaceArn() *string
	// The unique identifier of a namespace.
	AttrNamespaceNamespaceId() *string
	// The name of the namespace.
	//
	// Must be between 3-64 alphanumeric characters in lowercase, and it cannot be a reserved word. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com//redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	AttrNamespaceNamespaceName() *string
	// The status of the namespace.
	AttrNamespaceStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the primary database created in the namespace.
	DbName() *string
	SetDbName(val *string)
	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace.
	DefaultIamRoleArn() *string
	SetDefaultIamRoleArn(val *string)
	// The name of the snapshot to be created before the namespace is deleted.
	FinalSnapshotName() *string
	SetFinalSnapshotName(val *string)
	// How long to retain the final snapshot.
	FinalSnapshotRetentionPeriod() *float64
	SetFinalSnapshotRetentionPeriod(val *float64)
	// A list of IAM roles to associate with the namespace.
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	// The ID of the AWS Key Management Service key used to encrypt your data.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The types of logs the namespace can export.
	LogExports() *[]*string
	SetLogExports(val *[]*string)
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
	// If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials.
	ManageAdminPassword() interface{}
	SetManageAdminPassword(val interface{})
	// The name of the namespace.
	NamespaceName() *string
	SetNamespaceName(val *string)
	// The resource policy that will be attached to the namespace.
	NamespaceResourcePolicy() interface{}
	SetNamespaceResourcePolicy(val interface{})
	// The tree node.
	Node() constructs.Node
	// The ARN for the Redshift application that integrates with IAM Identity Center.
	RedshiftIdcApplicationArn() *string
	SetRedshiftIdcApplicationArn(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The snapshot copy configurations for the namespace.
	SnapshotCopyConfigurations() interface{}
	SetSnapshotCopyConfigurations(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The map of the key-value pairs used to tag the namespace.
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

// The jsii proxy struct for CfnNamespace
type jsiiProxy_CfnNamespace struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnNamespace) AdminPasswordSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AdminUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespace() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceAdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceAdminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceCreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceDbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceDbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceDefaultIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceDefaultIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceIamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrNamespaceIamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceLogExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrNamespaceLogExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) AttrNamespaceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNamespaceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) DefaultIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) FinalSnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) FinalSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"finalSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) LogExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) ManageAdminPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) NamespaceResourcePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceResourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) RedshiftIdcApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftIdcApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) SnapshotCopyConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotCopyConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) TagsRaw() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamespace) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnNamespace(scope constructs.Construct, id *string, props *CfnNamespaceProps) CfnNamespace {
	_init_.Initialize()

	if err := validateNewCfnNamespaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnNamespace_Override(c CfnNamespace, scope constructs.Construct, id *string, props *CfnNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnNamespace)SetAdminPasswordSecretKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"adminPasswordSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetAdminUsername(val *string) {
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetAdminUserPassword(val *string) {
	_jsii_.Set(
		j,
		"adminUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetDbName(val *string) {
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetDefaultIamRoleArn(val *string) {
	_jsii_.Set(
		j,
		"defaultIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetFinalSnapshotName(val *string) {
	_jsii_.Set(
		j,
		"finalSnapshotName",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetFinalSnapshotRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"finalSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetIamRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetLogExports(val *[]*string) {
	_jsii_.Set(
		j,
		"logExports",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetManageAdminPassword(val interface{}) {
	if err := j.validateSetManageAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageAdminPassword",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetNamespaceName(val *string) {
	if err := j.validateSetNamespaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceName",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetNamespaceResourcePolicy(val interface{}) {
	_jsii_.Set(
		j,
		"namespaceResourcePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetRedshiftIdcApplicationArn(val *string) {
	_jsii_.Set(
		j,
		"redshiftIdcApplicationArn",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetSnapshotCopyConfigurations(val interface{}) {
	if err := j.validateSetSnapshotCopyConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCopyConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnNamespace)SetTagsRaw(val *[]*awscdk.CfnTag) {
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
func CfnNamespace_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNamespace_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnNamespace_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNamespace_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
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
func CfnNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNamespace_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshiftserverless.CfnNamespace",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNamespace) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnNamespace) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNamespace) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNamespace) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnNamespace) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnNamespace) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnNamespace) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnNamespace) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnNamespace) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnNamespace) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnNamespace) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNamespace) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNamespace) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnNamespace) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnNamespace) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnNamespace) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnNamespace) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnNamespace) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

