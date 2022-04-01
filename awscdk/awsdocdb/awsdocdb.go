package awsdocdb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdocdb/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

// Backup configuration for DocumentDB databases.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//
//   var duration duration
//   backupProps := &backupProps{
//   	retention: duration,
//
//   	// the properties below are optional
//   	preferredWindow: jsii.String("preferredWindow"),
//   }
//
// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
//
// Experimental.
type BackupProps struct {
	// How many days to retain the backup.
	// Experimental.
	Retention awscdk.Duration `json:"retention" yaml:"retention"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'.
	// Experimental.
	PreferredWindow *string `json:"preferredWindow" yaml:"preferredWindow"`
}

// A CloudFormation `AWS::DocDB::DBCluster`.
//
// The `AWS::DocDB::DBCluster` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBCluster. Amazon DocumentDB is a fully managed, MongoDB-compatible document database engine. For more information, see [DBCluster](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html) in the *Amazon DocumentDB Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   cfnDBCluster := docdb.NewCfnDBCluster(this, jsii.String("MyCfnDBCluster"), &cfnDBClusterProps{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	backupRetentionPeriod: jsii.Number(123),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deletionProtection: jsii.Boolean(false),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	engineVersion: jsii.String("engineVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	port: jsii.Number(123),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	storageEncrypted: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   })
//
type CfnDBCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The resource id for the cluster;
	//
	// for example: `cluster-ABCD1234EFGH5678IJKL90MNOP` . The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.
	AttrClusterResourceId() *string
	// The connection endpoint for the cluster, such as `sample-cluster.cluster-cozrlsfrcjoc.us-east-1.docdb.amazonaws.com` .
	AttrEndpoint() *string
	// The port number on which the cluster accepts connections.
	//
	// For example: `27017` .
	AttrPort() *string
	// The reader endpoint for the cluster.
	//
	// For example: `sample-cluster.cluster-ro-cozrlsfrcjoc.us-east-1.docdb.amazonaws.com` .
	AttrReadEndpoint() *string
	// A list of Amazon EC2 Availability Zones that instances in the cluster can be created in.
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	// The number of days for which automated backups are retained. You must specify a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35.
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// `AWS::DocDB::DBCluster.CopyTagsToSnapshot`.
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster`.
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	// The name of the cluster parameter group to associate with this cluster.
	DbClusterParameterGroupName() *string
	SetDbClusterParameterGroupName(val *string)
	// A subnet group to associate with this cluster.
	//
	// Constraints: Must match the name of an existing `DBSubnetGroup` . Must not be default.
	//
	// Example: `mySubnetgroup`.
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	// Protects clusters from being accidentally deleted.
	//
	// If enabled, the cluster cannot be deleted unless it is modified and `DeletionProtection` is disabled.
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	// The list of log types that need to be enabled for exporting to Amazon CloudWatch Logs.
	//
	// You can enable audit logs or profiler logs. For more information, see [Auditing Amazon DocumentDB Events](https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html) and [Profiling Amazon DocumentDB Operations](https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html) .
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	// The version number of the database engine to use.
	//
	// The `--engine-version` will default to the latest major engine version. For production workloads, we recommend explicitly declaring this parameter with the intended major engine version.
	EngineVersion() *string
	SetEngineVersion(val *string)
	// The AWS KMS key identifier for an encrypted cluster.
	//
	// The AWS KMS key identifier is the Amazon Resource Name (ARN) for the AWS KMS encryption key. If you are creating a cluster using the same AWS account that owns the AWS KMS encryption key that is used to encrypt the new cluster, you can use the AWS KMS key alias instead of the ARN for the AWS KMS encryption key.
	//
	// If an encryption key is not specified in `KmsKeyId` :
	//
	// - If the `StorageEncrypted` parameter is `true` , Amazon DocumentDB uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Regions .
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the master user for the cluster.
	//
	// Constraints:
	//
	// - Must be from 1 to 63 letters or numbers.
	// - The first character must be a letter.
	// - Cannot be a reserved word for the chosen database engine.
	MasterUsername() *string
	SetMasterUsername(val *string)
	// The password for the master database user.
	//
	// This password can contain any printable ASCII character except forward slash (/), double quote ("), or the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 100 characters.
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Specifies the port that the database engine is listening on.
	Port() *float64
	SetPort(val *float64)
	// The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region .
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The identifier for the snapshot or cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a cluster snapshot. However, you can use only the ARN to specify a snapshot.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing snapshot.
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Specifies whether the cluster is encrypted.
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	// The tags to be assigned to the cluster.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A list of EC2 VPC security groups to associate with this cluster.
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBCluster
type jsiiProxy_CfnDBCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBCluster) AttrClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AttrReadEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::DocDB::DBCluster`.
func NewCfnDBCluster(scope awscdk.Construct, id *string, props *CfnDBClusterProps) CfnDBCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnDBCluster{}

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DocDB::DBCluster`.
func NewCfnDBCluster_Override(c CfnDBCluster, scope awscdk.Construct, id *string, props *CfnDBClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetBackupRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetCopyTagsToSnapshot(val interface{}) {
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDbClusterParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetDeletionProtection(val interface{}) {
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEnableCloudwatchLogsExports(val *[]*string) {
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetStorageEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster) SetVpcSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDBCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDBCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDBCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_docdb.CfnDBCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBCluster) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDBCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::DocDB::DBClusterParameterGroup`.
//
// The `AWS::DocDB::DBClusterParameterGroup` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBClusterParameterGroup. For more information, see [DBClusterParameterGroup](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterParameterGroup.html) in the *Amazon DocumentDB Developer Guide* .
//
// Parameters in a cluster parameter group apply to all of the instances in a cluster.
//
// A cluster parameter group is initially created with the default parameters for the database engine used by instances in the cluster. To provide custom values for any of the parameters, you must modify the group after you create it. After you create a DB cluster parameter group, you must associate it with your cluster. For the new cluster parameter group and associated settings to take effect, you must then reboot the DB instances in the cluster without failover.
//
// > After you create a cluster parameter group, you should wait at least 5 minutes before creating your first cluster that uses that cluster parameter group as the default parameter group. This allows Amazon DocumentDB to fully complete the create action before the cluster parameter group is used as the default for a new cluster. This step is especially important for parameters that are critical when creating the default database for a cluster, such as the character set for the default database defined by the `character_set_database` parameter.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//
//   var parameters interface{}
//   cfnDBClusterParameterGroup := docdb.NewCfnDBClusterParameterGroup(this, jsii.String("MyCfnDBClusterParameterGroup"), &cfnDBClusterParameterGroupProps{
//   	description: jsii.String("description"),
//   	family: jsii.String("family"),
//   	parameters: parameters,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDBClusterParameterGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The description for the cluster parameter group.
	Description() *string
	SetDescription(val *string)
	// The cluster parameter group family name.
	Family() *string
	SetFamily(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	// - Must not match the name of an existing `DBClusterParameterGroup` .
	//
	// > This value is stored as a lowercase string.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Provides a list of parameters for the cluster parameter group.
	Parameters() interface{}
	SetParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags to be assigned to the cluster parameter group.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBClusterParameterGroup
type jsiiProxy_CfnDBClusterParameterGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DocDB::DBClusterParameterGroup`.
func NewCfnDBClusterParameterGroup(scope awscdk.Construct, id *string, props *CfnDBClusterParameterGroupProps) CfnDBClusterParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBClusterParameterGroup{}

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBClusterParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DocDB::DBClusterParameterGroup`.
func NewCfnDBClusterParameterGroup_Override(c CfnDBClusterParameterGroup, scope awscdk.Construct, id *string, props *CfnDBClusterParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBClusterParameterGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetFamily(val *string) {
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDBClusterParameterGroup) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDBClusterParameterGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBClusterParameterGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDBClusterParameterGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBClusterParameterGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDBClusterParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBClusterParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBClusterParameterGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_docdb.CfnDBClusterParameterGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBClusterParameterGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBClusterParameterGroup`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//
//   var parameters interface{}
//   cfnDBClusterParameterGroupProps := &cfnDBClusterParameterGroupProps{
//   	description: jsii.String("description"),
//   	family: jsii.String("family"),
//   	parameters: parameters,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBClusterParameterGroupProps struct {
	// The description for the cluster parameter group.
	Description *string `json:"description" yaml:"description"`
	// The cluster parameter group family name.
	Family *string `json:"family" yaml:"family"`
	// Provides a list of parameters for the cluster parameter group.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// The name of the DB cluster parameter group.
	//
	// Constraints:
	//
	// - Must not match the name of an existing `DBClusterParameterGroup` .
	//
	// > This value is stored as a lowercase string.
	Name *string `json:"name" yaml:"name"`
	// The tags to be assigned to the cluster parameter group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnDBCluster`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   cfnDBClusterProps := &cfnDBClusterProps{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	backupRetentionPeriod: jsii.Number(123),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deletionProtection: jsii.Boolean(false),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	engineVersion: jsii.String("engineVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	port: jsii.Number(123),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	storageEncrypted: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnDBClusterProps struct {
	// A list of Amazon EC2 Availability Zones that instances in the cluster can be created in.
	AvailabilityZones *[]*string `json:"availabilityZones" yaml:"availabilityZones"`
	// The number of days for which automated backups are retained. You must specify a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35.
	BackupRetentionPeriod *float64 `json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// `AWS::DocDB::DBCluster.CopyTagsToSnapshot`.
	CopyTagsToSnapshot interface{} `json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster`.
	DbClusterIdentifier *string `json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The name of the cluster parameter group to associate with this cluster.
	DbClusterParameterGroupName *string `json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// A subnet group to associate with this cluster.
	//
	// Constraints: Must match the name of an existing `DBSubnetGroup` . Must not be default.
	//
	// Example: `mySubnetgroup`.
	DbSubnetGroupName *string `json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Protects clusters from being accidentally deleted.
	//
	// If enabled, the cluster cannot be deleted unless it is modified and `DeletionProtection` is disabled.
	DeletionProtection interface{} `json:"deletionProtection" yaml:"deletionProtection"`
	// The list of log types that need to be enabled for exporting to Amazon CloudWatch Logs.
	//
	// You can enable audit logs or profiler logs. For more information, see [Auditing Amazon DocumentDB Events](https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html) and [Profiling Amazon DocumentDB Operations](https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html) .
	EnableCloudwatchLogsExports *[]*string `json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// The version number of the database engine to use.
	//
	// The `--engine-version` will default to the latest major engine version. For production workloads, we recommend explicitly declaring this parameter with the intended major engine version.
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// The AWS KMS key identifier for an encrypted cluster.
	//
	// The AWS KMS key identifier is the Amazon Resource Name (ARN) for the AWS KMS encryption key. If you are creating a cluster using the same AWS account that owns the AWS KMS encryption key that is used to encrypt the new cluster, you can use the AWS KMS key alias instead of the ARN for the AWS KMS encryption key.
	//
	// If an encryption key is not specified in `KmsKeyId` :
	//
	// - If the `StorageEncrypted` parameter is `true` , Amazon DocumentDB uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Regions .
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the master user for the cluster.
	//
	// Constraints:
	//
	// - Must be from 1 to 63 letters or numbers.
	// - The first character must be a letter.
	// - Cannot be a reserved word for the chosen database engine.
	MasterUsername *string `json:"masterUsername" yaml:"masterUsername"`
	// The password for the master database user.
	//
	// This password can contain any printable ASCII character except forward slash (/), double quote ("), or the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 100 characters.
	MasterUserPassword *string `json:"masterUserPassword" yaml:"masterUserPassword"`
	// Specifies the port that the database engine is listening on.
	Port *float64 `json:"port" yaml:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region .
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The identifier for the snapshot or cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a cluster snapshot. However, you can use only the ARN to specify a snapshot.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing snapshot.
	SnapshotIdentifier *string `json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Specifies whether the cluster is encrypted.
	StorageEncrypted interface{} `json:"storageEncrypted" yaml:"storageEncrypted"`
	// The tags to be assigned to the cluster.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// A list of EC2 VPC security groups to associate with this cluster.
	VpcSecurityGroupIds *[]*string `json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::DocDB::DBInstance`.
//
// The `AWS::DocDB::DBInstance` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBInstance. For more information, see [DBInstance](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html) in the *Amazon DocumentDB Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   cfnDBInstance := docdb.NewCfnDBInstance(this, jsii.String("MyCfnDBInstance"), &cfnDBInstanceProps{
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbInstanceClass: jsii.String("dbInstanceClass"),
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	enablePerformanceInsights: jsii.Boolean(false),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDBInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The connection endpoint for the instance.
	//
	// For example: `sample-cluster.cluster-abcdefghijkl.us-east-1.docdb.amazonaws.com` .
	AttrEndpoint() *string
	// The port number on which the database accepts connections, such as `27017` .
	AttrPort() *string
	// This parameter does not apply to Amazon DocumentDB.
	//
	// Amazon DocumentDB does not perform minor version upgrades regardless of the value set.
	//
	// Default: `false`.
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	// The Amazon EC2 Availability Zone that the instance is created in.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region .
	//
	// Example: `us-east-1d`.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The identifier of the cluster that the instance will belong to.
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	// The compute and memory capacity of the instance;
	//
	// for example, `db.m4.large` . If you change the class of an instance there can be some interruption in the cluster's service.
	DbInstanceClass() *string
	SetDbInstanceClass(val *string)
	// The instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `mydbinstance`.
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
	// `AWS::DocDB::DBInstance.EnablePerformanceInsights`.
	EnablePerformanceInsights() interface{}
	SetEnablePerformanceInsights(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The tags to be assigned to the instance.
	//
	// You can assign up to 10 tags to an instance.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBInstance
type jsiiProxy_CfnDBInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBInstance) AttrEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AttrPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) EnablePerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DocDB::DBInstance`.
func NewCfnDBInstance(scope awscdk.Construct, id *string, props *CfnDBInstanceProps) CfnDBInstance {
	_init_.Initialize()

	j := jsiiProxy_CfnDBInstance{}

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DocDB::DBInstance`.
func NewCfnDBInstance_Override(c CfnDBInstance, scope awscdk.Construct, id *string, props *CfnDBInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAutoMinorVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbInstanceClass(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceClass",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetDbInstanceIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetEnablePerformanceInsights(val interface{}) {
	_jsii_.Set(
		j,
		"enablePerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_CfnDBInstance) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDBInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDBInstance_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBInstance",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDBInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_docdb.CfnDBInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBInstance) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBInstance) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDBInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBInstance) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBInstance) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBInstance) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBInstance`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   cfnDBInstanceProps := &cfnDBInstanceProps{
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbInstanceClass: jsii.String("dbInstanceClass"),
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	dbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	enablePerformanceInsights: jsii.Boolean(false),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBInstanceProps struct {
	// The identifier of the cluster that the instance will belong to.
	DbClusterIdentifier *string `json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of the instance;
	//
	// for example, `db.m4.large` . If you change the class of an instance there can be some interruption in the cluster's service.
	DbInstanceClass *string `json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// This parameter does not apply to Amazon DocumentDB.
	//
	// Amazon DocumentDB does not perform minor version upgrades regardless of the value set.
	//
	// Default: `false`.
	AutoMinorVersionUpgrade interface{} `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Amazon EC2 Availability Zone that the instance is created in.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region .
	//
	// Example: `us-east-1d`.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `mydbinstance`.
	DbInstanceIdentifier *string `json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// `AWS::DocDB::DBInstance.EnablePerformanceInsights`.
	EnablePerformanceInsights interface{} `json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The tags to be assigned to the instance.
	//
	// You can assign up to 10 tags to an instance.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::DocDB::DBSubnetGroup`.
//
// The `AWS::DocDB::DBSubnetGroup` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBSubnetGroup. subnet groups must contain at least one subnet in at least two Availability Zones in the AWS Region . For more information, see [DBSubnetGroup](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBSubnetGroup.html) in the *Amazon DocumentDB Developer Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   cfnDBSubnetGroup := docdb.NewCfnDBSubnetGroup(this, jsii.String("MyCfnDBSubnetGroup"), &cfnDBSubnetGroupProps{
//   	dbSubnetGroupDescription: jsii.String("dbSubnetGroupDescription"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDBSubnetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The description for the subnet group.
	DbSubnetGroupDescription() *string
	SetDbSubnetGroupDescription(val *string)
	// The name for the subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens. Must not be default.
	//
	// Example: `mySubnetgroup`.
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The Amazon EC2 subnet IDs for the subnet group.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// The tags to be assigned to the subnet group.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
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
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDBSubnetGroup
type jsiiProxy_CfnDBSubnetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDBSubnetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) DbSubnetGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBSubnetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DocDB::DBSubnetGroup`.
func NewCfnDBSubnetGroup(scope awscdk.Construct, id *string, props *CfnDBSubnetGroupProps) CfnDBSubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnDBSubnetGroup{}

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBSubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DocDB::DBSubnetGroup`.
func NewCfnDBSubnetGroup_Override(c CfnDBSubnetGroup, scope awscdk.Construct, id *string, props *CfnDBSubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.CfnDBSubnetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDBSubnetGroup) SetDbSubnetGroupDescription(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupDescription",
		val,
	)
}

func (j *jsiiProxy_CfnDBSubnetGroup) SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBSubnetGroup) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnDBSubnetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBSubnetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDBSubnetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBSubnetGroup",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDBSubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.CfnDBSubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBSubnetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_docdb.CfnDBSubnetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnDBSubnetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDBSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDBSubnetGroup`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   cfnDBSubnetGroupProps := &cfnDBSubnetGroupProps{
//   	dbSubnetGroupDescription: jsii.String("dbSubnetGroupDescription"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDBSubnetGroupProps struct {
	// The description for the subnet group.
	DbSubnetGroupDescription *string `json:"dbSubnetGroupDescription" yaml:"dbSubnetGroupDescription"`
	// The Amazon EC2 subnet IDs for the subnet group.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The name for the subnet group. This value is stored as a lowercase string.
	//
	// Constraints: Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens. Must not be default.
	//
	// Example: `mySubnetgroup`.
	DbSubnetGroupName *string `json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// The tags to be assigned to the subnet group.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A cluster parameter group.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   clusterParameterGroup := docdb.NewClusterParameterGroup(this, jsii.String("MyClusterParameterGroup"), &clusterParameterGroupProps{
//   	family: jsii.String("family"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//
//   	// the properties below are optional
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	description: jsii.String("description"),
//   })
//
// Experimental.
type ClusterParameterGroup interface {
	awscdk.Resource
	IClusterParameterGroup
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The name of the parameter group.
	// Experimental.
	ParameterGroupName() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ClusterParameterGroup
type jsiiProxy_ClusterParameterGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IClusterParameterGroup
}

func (j *jsiiProxy_ClusterParameterGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClusterParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewClusterParameterGroup(scope constructs.Construct, id *string, props *ClusterParameterGroupProps) ClusterParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_ClusterParameterGroup{}

	_jsii_.Create(
		"monocdk.aws_docdb.ClusterParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewClusterParameterGroup_Override(c ClusterParameterGroup, scope constructs.Construct, id *string, props *ClusterParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.ClusterParameterGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

// Imports a parameter group.
// Experimental.
func ClusterParameterGroup_FromParameterGroupName(scope constructs.Construct, id *string, parameterGroupName *string) IClusterParameterGroup {
	_init_.Initialize()

	var returns IClusterParameterGroup

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.ClusterParameterGroup",
		"fromParameterGroupName",
		[]interface{}{scope, id, parameterGroupName},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ClusterParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.ClusterParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ClusterParameterGroup_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.ClusterParameterGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_ClusterParameterGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterParameterGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterParameterGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterParameterGroup) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterParameterGroup) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_ClusterParameterGroup) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterParameterGroup) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClusterParameterGroup) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_ClusterParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClusterParameterGroup) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a cluster parameter group.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   clusterParameterGroupProps := &clusterParameterGroupProps{
//   	family: jsii.String("family"),
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//
//   	// the properties below are optional
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	description: jsii.String("description"),
//   }
//
// Experimental.
type ClusterParameterGroupProps struct {
	// Database family of this parameter group.
	// Experimental.
	Family *string `json:"family" yaml:"family"`
	// The parameters in this parameter group.
	// Experimental.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The name of the cluster parameter group.
	// Experimental.
	DbClusterParameterGroupName *string `json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Description for this parameter group.
	// Experimental.
	Description *string `json:"description" yaml:"description"`
}

// Create a clustered database with a given number of instances.
//
// Example:
//   var vpc vpc
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	masterUser: &login{
//   		username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		excludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		secretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_R5, ec2.instanceSize_LARGE),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   	vpc: vpc,
//   })
//
// Experimental.
type DatabaseCluster interface {
	awscdk.Resource
	IDatabaseCluster
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// The resource id for the cluster;
	//
	// for example: cluster-ABCD1234EFGH5678IJKL90MNOP. The cluster ID uniquely
	// identifies the cluster and is used in things like IAM authentication policies.
	// Experimental.
	ClusterResourceIdentifier() *string
	// The connections object to implement IConnectable.
	// Experimental.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// Endpoints which address each individual replica.
	// Experimental.
	InstanceEndpoints() *[]Endpoint
	// Identifiers of the replicas.
	// Experimental.
	InstanceIdentifiers() *[]*string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The secret attached to this cluster.
	// Experimental.
	Secret() awssecretsmanager.ISecret
	// Security group identifier of this database.
	// Experimental.
	SecurityGroupId() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds the multi user rotation to this cluster.
	// Experimental.
	AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation
	// Adds the single user rotation of the master password to this cluster.
	// Experimental.
	AddRotationSingleUser(automaticallyAfter awscdk.Duration) awssecretsmanager.SecretRotation
	// Adds security groups to this cluster.
	// Experimental.
	AddSecurityGroups(securityGroups ...awsec2.ISecurityGroup)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Renders the secret attachment target specifications.
	// Experimental.
	AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for DatabaseCluster
type jsiiProxy_DatabaseCluster struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseCluster
}

func (j *jsiiProxy_DatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) ClusterResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseCluster(scope constructs.Construct, id *string, props *DatabaseClusterProps) DatabaseCluster {
	_init_.Initialize()

	j := jsiiProxy_DatabaseCluster{}

	_jsii_.Create(
		"monocdk.aws_docdb.DatabaseCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseCluster_Override(d DatabaseCluster, scope constructs.Construct, id *string, props *DatabaseClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.DatabaseCluster",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing DatabaseCluster from properties.
// Experimental.
func DatabaseCluster_FromDatabaseClusterAttributes(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) IDatabaseCluster {
	_init_.Initialize()

	var returns IDatabaseCluster

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseCluster",
		"fromDatabaseClusterAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DatabaseCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseCluster_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseCluster",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func DatabaseCluster_DEFAULT_NUM_INSTANCES() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"monocdk.aws_docdb.DatabaseCluster",
		"DEFAULT_NUM_INSTANCES",
		&returns,
	)
	return returns
}

func DatabaseCluster_DEFAULT_PORT() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"monocdk.aws_docdb.DatabaseCluster",
		"DEFAULT_PORT",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseCluster) AddRotationMultiUser(id *string, options *RotationMultiUserOptions) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationMultiUser",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) AddRotationSingleUser(automaticallyAfter awscdk.Duration) awssecretsmanager.SecretRotation {
	var returns awssecretsmanager.SecretRotation

	_jsii_.Invoke(
		d,
		"addRotationSingleUser",
		[]interface{}{automaticallyAfter},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) AddSecurityGroups(securityGroups ...awsec2.ISecurityGroup) {
	args := []interface{}{}
	for _, a := range securityGroups {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addSecurityGroups",
		args,
	)
}

func (d *jsiiProxy_DatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		d,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing cluster instance.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"
//
//   var securityGroup securityGroup
//   databaseClusterAttributes := &databaseClusterAttributes{
//   	clusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	instanceEndpointAddresses: []*string{
//   		jsii.String("instanceEndpointAddresses"),
//   	},
//   	instanceIdentifiers: []*string{
//   		jsii.String("instanceIdentifiers"),
//   	},
//   	port: jsii.Number(123),
//   	readerEndpointAddress: jsii.String("readerEndpointAddress"),
//   	securityGroup: securityGroup,
//   }
//
// Experimental.
type DatabaseClusterAttributes struct {
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// Identifier for the cluster.
	// Experimental.
	ClusterIdentifier *string `json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Endpoint addresses of individual instances.
	// Experimental.
	InstanceEndpointAddresses *[]*string `json:"instanceEndpointAddresses" yaml:"instanceEndpointAddresses"`
	// Identifier for the instances.
	// Experimental.
	InstanceIdentifiers *[]*string `json:"instanceIdentifiers" yaml:"instanceIdentifiers"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
	// Reader endpoint address.
	// Experimental.
	ReaderEndpointAddress *string `json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security group of the database cluster.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
}

// Properties for a new database cluster.
//
// Example:
//   var vpc vpc
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	masterUser: &login{
//   		username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		excludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		secretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_R5, ec2.instanceSize_LARGE),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   	vpc: vpc,
//   })
//
// Experimental.
type DatabaseClusterProps struct {
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// Username and password for the administrative user.
	// Experimental.
	MasterUser *Login `json:"masterUser" yaml:"masterUser"`
	// What subnets to run the DocumentDB instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
	//
	// Experimental.
	Backup *BackupProps `json:"backup" yaml:"backup"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudWatchLogsRetention awslogs.RetentionDays `json:"cloudWatchLogsRetention" yaml:"cloudWatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudWatchLogsRetentionRole awsiam.IRole `json:"cloudWatchLogsRetentionRole" yaml:"cloudWatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	// Experimental.
	DbClusterName *string `json:"dbClusterName" yaml:"dbClusterName"`
	// Specifies whether this cluster can be deleted.
	//
	// If deletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// deletionProtection is disabled. deletionProtection protects clusters from
	// being accidentally deleted.
	// Experimental.
	DeletionProtection *bool `json:"deletionProtection" yaml:"deletionProtection"`
	// What version of the database to start.
	// Experimental.
	EngineVersion *string `json:"engineVersion" yaml:"engineVersion"`
	// Whether the audit logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the audit log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html#event-auditing-enabling-auditing
	//
	// Experimental.
	ExportAuditLogsToCloudWatch *bool `json:"exportAuditLogsToCloudWatch" yaml:"exportAuditLogsToCloudWatch"`
	// Whether the profiler logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the profiler log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html#profiling.enable-profiling
	//
	// Experimental.
	ExportProfilerLogsToCloudWatch *bool `json:"exportProfilerLogsToCloudWatch" yaml:"exportProfilerLogsToCloudWatch"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Number of DocDB compute instances.
	// Experimental.
	Instances *float64 `json:"instances" yaml:"instances"`
	// The KMS key for storage encryption.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IClusterParameterGroup `json:"parameterGroup" yaml:"parameterGroup"`
	// The port the DocumentDB cluster will listen on.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-maintain.html#maintenance-window
	//
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup" yaml:"securityGroup"`
	// Whether to enable storage encryption.
	// Experimental.
	StorageEncrypted *bool `json:"storageEncrypted" yaml:"storageEncrypted"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
}

// A database instance.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"
//
//   var databaseCluster databaseCluster
//   var instanceType instanceType
//   databaseInstance := docdb.NewDatabaseInstance(this, jsii.String("MyDatabaseInstance"), &databaseInstanceProps{
//   	cluster: databaseCluster,
//   	instanceType: instanceType,
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	dbInstanceName: jsii.String("dbInstanceName"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   })
//
// Experimental.
type DatabaseInstance interface {
	awscdk.Resource
	IDatabaseInstance
	// The instance's database cluster.
	// Experimental.
	Cluster() IDatabaseCluster
	// The instance endpoint address.
	// Experimental.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	// Experimental.
	DbInstanceEndpointPort() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The instance arn.
	// Experimental.
	InstanceArn() *string
	// The instance endpoint.
	// Experimental.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	// Experimental.
	InstanceIdentifier() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for DatabaseInstance
type jsiiProxy_DatabaseInstance struct {
	internal.Type__awscdkResource
	jsiiProxy_IDatabaseInstance
}

func (j *jsiiProxy_DatabaseInstance) Cluster() IDatabaseCluster {
	var returns IDatabaseCluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseInstance(scope constructs.Construct, id *string, props *DatabaseInstanceProps) DatabaseInstance {
	_init_.Initialize()

	j := jsiiProxy_DatabaseInstance{}

	_jsii_.Create(
		"monocdk.aws_docdb.DatabaseInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseInstance_Override(d DatabaseInstance, scope constructs.Construct, id *string, props *DatabaseInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.DatabaseInstance",
		[]interface{}{scope, id, props},
		d,
	)
}

// Import an existing database instance.
// Experimental.
func DatabaseInstance_FromDatabaseInstanceAttributes(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) IDatabaseInstance {
	_init_.Initialize()

	var returns IDatabaseInstance

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseInstance",
		"fromDatabaseInstanceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DatabaseInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseInstance_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseInstance",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseInstance) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseInstance) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseInstance) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstance) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties that describe an existing instance.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   databaseInstanceAttributes := &databaseInstanceAttributes{
//   	instanceEndpointAddress: jsii.String("instanceEndpointAddress"),
//   	instanceIdentifier: jsii.String("instanceIdentifier"),
//   	port: jsii.Number(123),
//   }
//
// Experimental.
type DatabaseInstanceAttributes struct {
	// The endpoint address.
	// Experimental.
	InstanceEndpointAddress *string `json:"instanceEndpointAddress" yaml:"instanceEndpointAddress"`
	// The instance identifier.
	// Experimental.
	InstanceIdentifier *string `json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The database port.
	// Experimental.
	Port *float64 `json:"port" yaml:"port"`
}

// Construction properties for a DatabaseInstanceNew.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"
//
//   var databaseCluster databaseCluster
//   var instanceType instanceType
//   databaseInstanceProps := &databaseInstanceProps{
//   	cluster: databaseCluster,
//   	instanceType: instanceType,
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	dbInstanceName: jsii.String("dbInstanceName"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	removalPolicy: monocdk.removalPolicy_DESTROY,
//   }
//
// Experimental.
type DatabaseInstanceProps struct {
	// The DocumentDB database cluster the instance should launch into.
	// Experimental.
	Cluster IDatabaseCluster `json:"cluster" yaml:"cluster"`
	// The name of the compute and memory capacity classes.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType" yaml:"instanceType"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	// Experimental.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	// Experimental.
	DbInstanceName *string `json:"dbInstanceName" yaml:"dbInstanceName"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window.
	// Experimental.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
}

// A database secret.
//
// Example:
//   var cluster databaseCluster
//   myUserSecret := docdb.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &databaseSecretProps{
//   	username: jsii.String("myuser"),
//   	masterSecret: cluster.secret,
//   })
//   myUserSecretAttached := myUserSecret.attach(cluster) // Adds DB connections information in the secret
//
//   cluster.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
//   	 // Add rotation using the multi user scheme
//   	secret: myUserSecretAttached,
//   })
//
// Experimental.
type DatabaseSecret interface {
	awssecretsmanager.Secret
	// Provides an identifier for this secret for use in IAM policies.
	//
	// If there is a full ARN, this is just the ARN;
	// if we have a partial ARN -- due to either importing by secret name or partial ARN --
	// then we need to add a suffix to capture the full ARN's format.
	// Experimental.
	ArnForPolicies() *string
	// Experimental.
	AutoCreatePolicy() *bool
	// The customer-managed encryption key that is used to encrypt this secret, if any.
	//
	// When not specified, the default
	// KMS key for the account and region is being used.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The ARN of the secret in AWS Secrets Manager.
	//
	// Will return the full ARN if available, otherwise a partial arn.
	// For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.
	// Experimental.
	SecretArn() *string
	// The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.
	//
	// This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).
	// Experimental.
	SecretFullArn() *string
	// The name of the secret.
	//
	// For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
	// '@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.
	// Experimental.
	SecretName() *string
	// Retrieve the value of the stored secret as a `SecretValue`.
	// Experimental.
	SecretValue() awscdk.SecretValue
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Adds a replica region for the secret.
	// Experimental.
	AddReplicaRegion(region *string, encryptionKey awskms.IKey)
	// Adds a rotation schedule to the secret.
	// Experimental.
	AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule
	// Adds a target attachment to the secret.
	//
	// Returns: an AttachedSecret.
	// Deprecated: use `attach()` instead.
	AddTargetAttachment(id *string, options *awssecretsmanager.AttachedSecretOptions) awssecretsmanager.SecretTargetAttachment
	// Adds a statement to the IAM resource policy associated with this secret.
	//
	// If this secret was created in this stack, a resource policy will be
	// automatically created upon the first call to `addToResourcePolicy`. If
	// the secret is imported, then this is a no-op.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Attach a target to this secret.
	//
	// Returns: An attached secret.
	// Experimental.
	Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret
	// Denies the `DeleteSecret` action to all principals within the current account.
	// Experimental.
	DenyAccountRootDelete()
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants reading the secret value to some role.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant
	// Grants writing and updating the secret value to some role.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.
	// Experimental.
	SecretValueFromJson(jsonField *string) awscdk.SecretValue
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for DatabaseSecret
type jsiiProxy_DatabaseSecret struct {
	internal.Type__awssecretsmanagerSecret
}

func (j *jsiiProxy_DatabaseSecret) ArnForPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretFullArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretFullArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) SecretValue() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"secretValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewDatabaseSecret(scope constructs.Construct, id *string, props *DatabaseSecretProps) DatabaseSecret {
	_init_.Initialize()

	j := jsiiProxy_DatabaseSecret{}

	_jsii_.Create(
		"monocdk.aws_docdb.DatabaseSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDatabaseSecret_Override(d DatabaseSecret, scope constructs.Construct, id *string, props *DatabaseSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.DatabaseSecret",
		[]interface{}{scope, id, props},
		d,
	)
}

// Deprecated: use `fromSecretCompleteArn` or `fromSecretPartialArn`.
func DatabaseSecret_FromSecretArn(scope constructs.Construct, id *string, secretArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"fromSecretArn",
		[]interface{}{scope, id, secretArn},
		&returns,
	)

	return returns
}

// Import an existing secret into the Stack.
// Experimental.
func DatabaseSecret_FromSecretAttributes(scope constructs.Construct, id *string, attrs *awssecretsmanager.SecretAttributes) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"fromSecretAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports a secret by complete ARN.
//
// The complete ARN is the ARN with the Secrets Manager-supplied suffix.
// Experimental.
func DatabaseSecret_FromSecretCompleteArn(scope constructs.Construct, id *string, secretCompleteArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"fromSecretCompleteArn",
		[]interface{}{scope, id, secretCompleteArn},
		&returns,
	)

	return returns
}

// Imports a secret by secret name;
//
// the ARN of the Secret will be set to the secret name.
// A secret with this name must exist in the same account & region.
// Deprecated: use `fromSecretNameV2`.
func DatabaseSecret_FromSecretName(scope constructs.Construct, id *string, secretName *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"fromSecretName",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by secret name.
//
// A secret with this name must exist in the same account & region.
// Replaces the deprecated `fromSecretName`.
// Experimental.
func DatabaseSecret_FromSecretNameV2(scope constructs.Construct, id *string, secretName *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"fromSecretNameV2",
		[]interface{}{scope, id, secretName},
		&returns,
	)

	return returns
}

// Imports a secret by partial ARN.
//
// The partial ARN is the ARN without the Secrets Manager-supplied suffix.
// Experimental.
func DatabaseSecret_FromSecretPartialArn(scope constructs.Construct, id *string, secretPartialArn *string) awssecretsmanager.ISecret {
	_init_.Initialize()

	var returns awssecretsmanager.ISecret

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"fromSecretPartialArn",
		[]interface{}{scope, id, secretPartialArn},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DatabaseSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func DatabaseSecret_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_docdb.DatabaseSecret",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) AddReplicaRegion(region *string, encryptionKey awskms.IKey) {
	_jsii_.InvokeVoid(
		d,
		"addReplicaRegion",
		[]interface{}{region, encryptionKey},
	)
}

func (d *jsiiProxy_DatabaseSecret) AddRotationSchedule(id *string, options *awssecretsmanager.RotationScheduleOptions) awssecretsmanager.RotationSchedule {
	var returns awssecretsmanager.RotationSchedule

	_jsii_.Invoke(
		d,
		"addRotationSchedule",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) AddTargetAttachment(id *string, options *awssecretsmanager.AttachedSecretOptions) awssecretsmanager.SecretTargetAttachment {
	var returns awssecretsmanager.SecretTargetAttachment

	_jsii_.Invoke(
		d,
		"addTargetAttachment",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		d,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DatabaseSecret) Attach(target awssecretsmanager.ISecretAttachmentTarget) awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret

	_jsii_.Invoke(
		d,
		"attach",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) DenyAccountRootDelete() {
	_jsii_.InvokeVoid(
		d,
		"denyAccountRootDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecret) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GrantRead(grantee awsiam.IGrantable, versionStages *[]*string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantRead",
		[]interface{}{grantee, versionStages},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecret) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseSecret) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecret) SecretValueFromJson(jsonField *string) awscdk.SecretValue {
	var returns awscdk.SecretValue

	_jsii_.Invoke(
		d,
		"secretValueFromJson",
		[]interface{}{jsonField},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DatabaseSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecret) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a DatabaseSecret.
//
// Example:
//   var cluster databaseCluster
//   myUserSecret := docdb.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &databaseSecretProps{
//   	username: jsii.String("myuser"),
//   	masterSecret: cluster.secret,
//   })
//   myUserSecretAttached := myUserSecret.attach(cluster) // Adds DB connections information in the secret
//
//   cluster.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
//   	 // Add rotation using the multi user scheme
//   	secret: myUserSecretAttached,
//   })
//
// Experimental.
type DatabaseSecretProps struct {
	// The username.
	// Experimental.
	Username *string `json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	// Experimental.
	MasterSecret awssecretsmanager.ISecret `json:"masterSecret" yaml:"masterSecret"`
	// The physical name of the secret.
	// Experimental.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// Connection endpoint of a database cluster or instance.
//
// Consists of a combination of hostname and port.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import docdb "github.com/aws/aws-cdk-go/awscdk/aws_docdb"
//   endpoint := docdb.NewEndpoint(jsii.String("address"), jsii.Number(123))
//
// Experimental.
type Endpoint interface {
	// The hostname of the endpoint.
	// Experimental.
	Hostname() *string
	// The port number of the endpoint.
	//
	// This can potentially be a CDK token. If you need to embed the port in a string (e.g. instance user data script),
	// use {@link Endpoint.portAsString}.
	// Experimental.
	Port() *float64
	// The combination of ``HOSTNAME:PORT`` for this endpoint.
	// Experimental.
	SocketAddress() *string
	// Returns the port number as a string representation that can be used for embedding within other strings.
	//
	// This is intended to deal with CDK's token system. Numeric CDK tokens are not expanded when their string
	// representation is embedded in a string. This function returns the port either as an unresolved string token or
	// as a resolved string representation of the port value.
	//
	// Returns: An (un)resolved string representation of the endpoint's port number.
	// Experimental.
	PortAsString() *string
}

// The jsii proxy struct for Endpoint
type jsiiProxy_Endpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_Endpoint) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Endpoint) SocketAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"socketAddress",
		&returns,
	)
	return returns
}


// Constructs an Endpoint instance.
// Experimental.
func NewEndpoint(address *string, port *float64) Endpoint {
	_init_.Initialize()

	j := jsiiProxy_Endpoint{}

	_jsii_.Create(
		"monocdk.aws_docdb.Endpoint",
		[]interface{}{address, port},
		&j,
	)

	return &j
}

// Constructs an Endpoint instance.
// Experimental.
func NewEndpoint_Override(e Endpoint, address *string, port *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_docdb.Endpoint",
		[]interface{}{address, port},
		e,
	)
}

func (e *jsiiProxy_Endpoint) PortAsString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"portAsString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A parameter group.
// Experimental.
type IClusterParameterGroup interface {
	awscdk.IResource
	// The name of this parameter group.
	// Experimental.
	ParameterGroupName() *string
}

// The jsii proxy for IClusterParameterGroup
type jsiiProxy_IClusterParameterGroup struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IClusterParameterGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

// Create a clustered database with a given number of instances.
// Experimental.
type IDatabaseCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// Endpoints which address each individual replica.
	// Experimental.
	InstanceEndpoints() *[]Endpoint
	// Identifiers of the replicas.
	// Experimental.
	InstanceIdentifiers() *[]*string
	// The security group for this database cluster.
	// Experimental.
	SecurityGroupId() *string
}

// The jsii proxy for IDatabaseCluster
type jsiiProxy_IDatabaseCluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (i *jsiiProxy_IDatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDatabaseCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// A database instance.
// Experimental.
type IDatabaseInstance interface {
	awscdk.IResource
	// The instance endpoint address.
	// Experimental.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	// Experimental.
	DbInstanceEndpointPort() *string
	// The instance arn.
	// Experimental.
	InstanceArn() *string
	// The instance endpoint.
	// Experimental.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	// Experimental.
	InstanceIdentifier() *string
}

// The jsii proxy for IDatabaseInstance
type jsiiProxy_IDatabaseInstance struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}

// Login credentials for a database cluster.
//
// Example:
//   var vpc vpc
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	masterUser: &login{
//   		username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		excludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		secretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_R5, ec2.instanceSize_LARGE),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   	vpc: vpc,
//   })
//
// Experimental.
type Login struct {
	// Username.
	// Experimental.
	Username *string `json:"username" yaml:"username"`
	// Specifies characters to not include in generated passwords.
	// Experimental.
	ExcludeCharacters *string `json:"excludeCharacters" yaml:"excludeCharacters"`
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey" yaml:"kmsKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	Password awscdk.SecretValue `json:"password" yaml:"password"`
	// The physical name of the secret, that will be generated.
	// Experimental.
	SecretName *string `json:"secretName" yaml:"secretName"`
}

// Options to add the multi user rotation.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//
//   var myImportedSecret secret
//   var cluster databaseCluster
//
//   cluster.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
//   	secret: myImportedSecret,
//   })
//
// Experimental.
type RotationMultiUserOptions struct {
	// The secret to rotate.
	//
	// It must be a JSON string with the following format:
	// ```
	// {
	//    "engine": <required: must be set to 'mongo'>,
	//    "host": <required: instance host name>,
	//    "username": <required: username>,
	//    "password": <required: password>,
	//    "dbname": <optional: database name>,
	//    "port": <optional: if not specified, default port 27017 will be used>,
	//    "masterarn": <required: the arn of the master secret which will be used to create users/change passwords>
	//    "ssl": <optional: if not specified, defaults to false. This must be true if being used for DocumentDB rotations
	//           where the cluster has TLS enabled>
	// }
	// ```.
	// Experimental.
	Secret awssecretsmanager.ISecret `json:"secret" yaml:"secret"`
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	// Experimental.
	AutomaticallyAfter awscdk.Duration `json:"automaticallyAfter" yaml:"automaticallyAfter"`
}

