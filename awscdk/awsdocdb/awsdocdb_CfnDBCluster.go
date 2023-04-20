package awsdocdb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdocdb/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::DocDB::DBCluster`.
//
// The `AWS::DocDB::DBCluster` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBCluster. Amazon DocumentDB is a fully managed, MongoDB-compatible document database engine. For more information, see [DBCluster](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBCluster.html) in the *Amazon DocumentDB Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBCluster := awscdk.Aws_docdb.NewCfnDBCluster(this, jsii.String("MyCfnDBCluster"), &CfnDBClusterProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	BackupRetentionPeriod: jsii.Number(123),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	DeletionProtection: jsii.Boolean(false),
//   	EnableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MasterUsername: jsii.String("masterUsername"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	Port: jsii.Number(123),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	RestoreToTime: jsii.String("restoreToTime"),
//   	RestoreType: jsii.String("restoreType"),
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	StorageEncrypted: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseLatestRestorableTime: jsii.Boolean(false),
//   	VpcSecurityGroupIds: []*string{
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
	// `AWS::DocDB::DBCluster.RestoreToTime`.
	RestoreToTime() *string
	SetRestoreToTime(val *string)
	// `AWS::DocDB::DBCluster.RestoreType`.
	RestoreType() *string
	SetRestoreType(val *string)
	// The identifier for the snapshot or cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a cluster snapshot. However, you can use only the ARN to specify a snapshot.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing snapshot.
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	// `AWS::DocDB::DBCluster.SourceDBClusterIdentifier`.
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
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
	// `AWS::DocDB::DBCluster.UseLatestRestorableTime`.
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
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

func (j *jsiiProxy_CfnDBCluster) RestoreToTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreToTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBCluster) RestoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreType",
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

func (j *jsiiProxy_CfnDBCluster) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
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

func (j *jsiiProxy_CfnDBCluster) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
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

	if err := validateNewCfnDBClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnDBCluster)SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetBackupRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbClusterParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDbSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEnableCloudwatchLogsExports(val *[]*string) {
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetEngineVersion(val *string) {
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPreferredBackupWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetRestoreToTime(val *string) {
	_jsii_.Set(
		j,
		"restoreToTime",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetRestoreType(val *string) {
	_jsii_.Set(
		j,
		"restoreType",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetSourceDbClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_CfnDBCluster)SetVpcSecurityGroupIds(val *[]*string) {
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

	if err := validateCfnDBCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnDBCluster_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnDBCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDBCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDBCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDBCluster) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnDBCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
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
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
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
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
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
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

