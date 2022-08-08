package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Redshift::Cluster`.
//
// Specifies a cluster. A *cluster* is a fully managed data warehouse that consists of a set of compute nodes.
//
// To create a cluster in Virtual Private Cloud (VPC), you must provide a cluster subnet group name. The cluster subnet group identifies the subnets of your VPC that Amazon Redshift uses when creating the cluster. For more information about managing clusters, go to [Amazon Redshift Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCluster := awscdk.Aws_redshift.NewCfnCluster(this, jsii.String("MyCfnCluster"), &cfnClusterProps{
//   	clusterType: jsii.String("clusterType"),
//   	dbName: jsii.String("dbName"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	nodeType: jsii.String("nodeType"),
//
//   	// the properties below are optional
//   	allowVersionUpgrade: jsii.Boolean(false),
//   	aquaConfigurationStatus: jsii.String("aquaConfigurationStatus"),
//   	automatedSnapshotRetentionPeriod: jsii.Number(123),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	availabilityZoneRelocation: jsii.Boolean(false),
//   	availabilityZoneRelocationStatus: jsii.String("availabilityZoneRelocationStatus"),
//   	classic: jsii.Boolean(false),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	clusterParameterGroupName: jsii.String("clusterParameterGroupName"),
//   	clusterSecurityGroups: []*string{
//   		jsii.String("clusterSecurityGroups"),
//   	},
//   	clusterSubnetGroupName: jsii.String("clusterSubnetGroupName"),
//   	clusterVersion: jsii.String("clusterVersion"),
//   	deferMaintenance: jsii.Boolean(false),
//   	deferMaintenanceDuration: jsii.Number(123),
//   	deferMaintenanceEndTime: jsii.String("deferMaintenanceEndTime"),
//   	deferMaintenanceStartTime: jsii.String("deferMaintenanceStartTime"),
//   	destinationRegion: jsii.String("destinationRegion"),
//   	elasticIp: jsii.String("elasticIp"),
//   	encrypted: jsii.Boolean(false),
//   	enhancedVpcRouting: jsii.Boolean(false),
//   	hsmClientCertificateIdentifier: jsii.String("hsmClientCertificateIdentifier"),
//   	hsmConfigurationIdentifier: jsii.String("hsmConfigurationIdentifier"),
//   	iamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	loggingProperties: &loggingPropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	maintenanceTrackName: jsii.String("maintenanceTrackName"),
//   	manualSnapshotRetentionPeriod: jsii.Number(123),
//   	numberOfNodes: jsii.Number(123),
//   	ownerAccount: jsii.String("ownerAccount"),
//   	port: jsii.Number(123),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	resourceAction: jsii.String("resourceAction"),
//   	revisionTarget: jsii.String("revisionTarget"),
//   	rotateEncryptionKey: jsii.Boolean(false),
//   	snapshotClusterIdentifier: jsii.String("snapshotClusterIdentifier"),
//   	snapshotCopyGrantName: jsii.String("snapshotCopyGrantName"),
//   	snapshotCopyManual: jsii.Boolean(false),
//   	snapshotCopyRetentionPeriod: jsii.Number(123),
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
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
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// If `true` , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster.
	//
	// When a new major version of the Amazon Redshift engine is released, you can request that the service automatically apply upgrades during the maintenance window to the Amazon Redshift engine that is running on your cluster.
	//
	// Default: `true`.
	AllowVersionUpgrade() interface{}
	SetAllowVersionUpgrade(val interface{})
	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) when it is created.
	//
	// Possible values include the following.
	//
	// - enabled - Use AQUA if it is available for the current AWS Region and Amazon Redshift node type.
	// - disabled - Don't use AQUA.
	// - auto - Amazon Redshift determines whether to use AQUA.
	AquaConfigurationStatus() *string
	SetAquaConfigurationStatus(val *string)
	AttrDeferMaintenanceIdentifier() *string
	// The connection endpoint for the Amazon Redshift cluster.
	//
	// For example: `examplecluster.cg034hpkmmjt.us-east-1.redshift.amazonaws.com` .
	AttrEndpointAddress() *string
	// The port number on which the Amazon Redshift cluster accepts connections.
	//
	// For example: `5439` .
	AttrEndpointPort() *string
	// A unique identifier for the cluster.
	//
	// You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. The identifier also appears in the Amazon Redshift console.
	//
	// Example: `myexamplecluster`.
	AttrId() *string
	// The number of days that automated snapshots are retained.
	//
	// If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with [CreateClusterSnapshot](https://docs.aws.amazon.com/redshift/latest/APIReference/API_CreateClusterSnapshot.html) in the *Amazon Redshift API Reference* .
	//
	// Default: `1`
	//
	// Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod() *float64
	SetAutomatedSnapshotRetentionPeriod(val *float64)
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster.
	//
	// For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
	//
	// Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint.
	//
	// Example: `us-east-2d`
	//
	// Constraint: The specified Availability Zone must be in the same region as the current endpoint.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster is created.
	AvailabilityZoneRelocation() interface{}
	SetAvailabilityZoneRelocation(val interface{})
	// Describes the status of the Availability Zone relocation operation.
	AvailabilityZoneRelocationStatus() *string
	SetAvailabilityZoneRelocationStatus(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A boolean value indicating whether the resize operation is using the classic resize process.
	//
	// If you don't provide this parameter or set the value to `false` , the resize type is elastic.
	Classic() interface{}
	SetClassic(val interface{})
	// A unique identifier for the cluster.
	//
	// You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. The identifier also appears in the Amazon Redshift console.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 alphanumeric characters or hyphens.
	// - Alphabetic characters must be lowercase.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	// - Must be unique for all clusters within an AWS account .
	//
	// Example: `myexamplecluster`.
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	// The name of the parameter group to be associated with this cluster.
	//
	// Default: The default Amazon Redshift cluster parameter group. For information about the default parameter group, go to [Working with Amazon Redshift Parameter Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
	//
	// Constraints:
	//
	// - Must be 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	ClusterParameterGroupName() *string
	SetClusterParameterGroupName(val *string)
	// A list of security groups to be associated with this cluster.
	//
	// Default: The default cluster security group for Amazon Redshift.
	ClusterSecurityGroups() *[]*string
	SetClusterSecurityGroups(val *[]*string)
	// The name of a cluster subnet group to be associated with this cluster.
	//
	// If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName() *string
	SetClusterSubnetGroupName(val *string)
	// The type of the cluster. When cluster type is specified as.
	//
	// - `single-node` , the *NumberOfNodes* parameter is not required.
	// - `multi-node` , the *NumberOfNodes* parameter is required.
	//
	// Valid Values: `multi-node` | `single-node`
	//
	// Default: `multi-node`.
	ClusterType() *string
	SetClusterType(val *string)
	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	//
	// The version selected runs on all the nodes in the cluster.
	//
	// Constraints: Only version 1.0 is currently available.
	//
	// Example: `1.0`
	ClusterVersion() *string
	SetClusterVersion(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the first database to be created when the cluster is created.
	//
	// To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database. For more information, go to [Create a Database](https://docs.aws.amazon.com/redshift/latest/dg/t_creating_database.html) in the Amazon Redshift Database Developer Guide.
	//
	// Default: `dev`
	//
	// Constraints:
	//
	// - Must contain 1 to 64 alphanumeric characters.
	// - Must contain only lowercase letters.
	// - Cannot be a word that is reserved by the service. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	DbName() *string
	SetDbName(val *string)
	// `AWS::Redshift::Cluster.DeferMaintenance`.
	DeferMaintenance() interface{}
	SetDeferMaintenance(val interface{})
	// `AWS::Redshift::Cluster.DeferMaintenanceDuration`.
	DeferMaintenanceDuration() *float64
	SetDeferMaintenanceDuration(val *float64)
	// `AWS::Redshift::Cluster.DeferMaintenanceEndTime`.
	DeferMaintenanceEndTime() *string
	SetDeferMaintenanceEndTime(val *string)
	// `AWS::Redshift::Cluster.DeferMaintenanceStartTime`.
	DeferMaintenanceStartTime() *string
	SetDeferMaintenanceStartTime(val *string)
	// The destination region that snapshots are automatically copied to when cross-region snapshot copy is enabled.
	DestinationRegion() *string
	SetDestinationRegion(val *string)
	// The Elastic IP (EIP) address for the cluster.
	//
	// Constraints: The cluster must be provisioned in EC2-VPC and publicly-accessible through an Internet gateway. For more information about provisioning clusters in EC2-VPC, go to [Supported Platforms to Launch Your Cluster](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms) in the Amazon Redshift Cluster Management Guide.
	ElasticIp() *string
	SetElasticIp(val *string)
	// If `true` , the data in the cluster is encrypted at rest.
	//
	// Default: false.
	Encrypted() interface{}
	SetEncrypted(val interface{})
	// An option that specifies whether to create the cluster with enhanced VPC routing enabled.
	//
	// To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see [Enhanced VPC Routing](https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html) in the Amazon Redshift Cluster Management Guide.
	//
	// If this option is `true` , enhanced VPC routing is enabled.
	//
	// Default: false.
	EnhancedVpcRouting() interface{}
	SetEnhancedVpcRouting(val interface{})
	// Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier() *string
	SetHsmClientCertificateIdentifier(val *string)
	// Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier() *string
	SetHsmConfigurationIdentifier(val *string)
	// A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services.
	//
	// You must supply the IAM roles in their Amazon Resource Name (ARN) format.
	//
	// The maximum number of IAM roles that you can associate is subject to a quota. For more information, go to [Quotas and limits](https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html) in the *Amazon Redshift Cluster Management Guide* .
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
	LoggingProperties() interface{}
	SetLoggingProperties(val interface{})
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
	// An optional parameter for the name of the maintenance track for the cluster.
	//
	// If you don't provide a maintenance track name, the cluster is assigned to the `current` track.
	MaintenanceTrackName() *string
	SetMaintenanceTrackName(val *string)
	// The default number of days to retain a manual snapshot.
	//
	// If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod() *float64
	SetManualSnapshotRetentionPeriod(val *float64)
	// The user name associated with the admin user account for the cluster that is being created.
	//
	// Constraints:
	//
	// - Must be 1 - 128 alphanumeric characters. The user name can't be `PUBLIC` .
	// - First character must be a letter.
	// - Cannot be a reserved word. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	MasterUsername() *string
	SetMasterUsername(val *string)
	// The password associated with the admin user account for the cluster that is being created.
	//
	// Constraints:
	//
	// - Must be between 8 and 64 characters in length.
	// - Must contain at least one uppercase letter.
	// - Must contain at least one lowercase letter.
	// - Must contain one number.
	// - Can be any printable ASCII character (ASCII code 33-126) except ' (single quote), " (double quote), \, /, or @.
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// The tree node.
	Node() constructs.Node
	// The node type to be provisioned for the cluster.
	//
	// For information about node types, go to [Working with Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes) in the *Amazon Redshift Cluster Management Guide* .
	//
	// Valid Values: `ds2.xlarge` | `ds2.8xlarge` | `dc1.large` | `dc1.8xlarge` | `dc2.large` | `dc2.8xlarge` | `ra3.xlplus` | `ra3.4xlarge` | `ra3.16xlarge`
	NodeType() *string
	SetNodeType(val *string)
	// The number of compute nodes in the cluster.
	//
	// This parameter is required when the *ClusterType* parameter is specified as `multi-node` .
	//
	// For information about determining how many nodes you need, go to [Working with Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes) in the *Amazon Redshift Cluster Management Guide* .
	//
	// If you don't specify this parameter, you get a single-node cluster. When requesting a multi-node cluster, you must specify the number of nodes that you want in the cluster.
	//
	// Default: `1`
	//
	// Constraints: Value must be at least 1 and no more than 100.
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	// The AWS account used to create or copy the snapshot.
	//
	// Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount() *string
	SetOwnerAccount(val *string)
	// The port number on which the cluster accepts incoming connections.
	//
	// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections.
	//
	// Default: `5439`
	//
	// Valid Values: `1150-65535`.
	Port() *float64
	SetPort(val *float64)
	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// Default: A 30-minute window selected at random from an 8-hour block of time per region, occurring on a random day of the week. For more information about the time blocks for each region, see [Maintenance Windows](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows) in Amazon Redshift Cluster Management Guide.
	//
	// Valid Days: Mon | Tue | Wed | Thu | Fri | Sat | Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	// If `true` , the cluster can be accessed from a public network.
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::Redshift::Cluster.ResourceAction`.
	ResourceAction() *string
	SetResourceAction(val *string)
	// `AWS::Redshift::Cluster.RevisionTarget`.
	RevisionTarget() *string
	SetRevisionTarget(val *string)
	// `AWS::Redshift::Cluster.RotateEncryptionKey`.
	RotateEncryptionKey() interface{}
	SetRotateEncryptionKey(val interface{})
	// The name of the cluster the source snapshot was created from.
	//
	// This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier() *string
	SetSnapshotClusterIdentifier(val *string)
	// The name of the snapshot copy grant.
	SnapshotCopyGrantName() *string
	SetSnapshotCopyGrantName(val *string)
	// Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
	SnapshotCopyManual() interface{}
	SetSnapshotCopyManual(val interface{})
	// The number of days to retain automated snapshots in the destination AWS Region after they are copied from the source AWS Region .
	//
	// By default, this only changes the retention period of copied automated snapshots.
	//
	// If you decrease the retention period for automated snapshots that are copied to a destination AWS Region , Amazon Redshift deletes any existing automated snapshots that were copied to the destination AWS Region and that fall outside of the new retention period.
	//
	// Constraints: Must be at least 1 and no more than 35 for automated snapshots.
	//
	// If you specify the `manual` option, only newly copied manual snapshots will have the new retention period.
	//
	// If you specify the value of -1 newly copied manual snapshots are retained indefinitely.
	//
	// Constraints: The number of days must be either -1 or an integer between 1 and 3,653 for manual snapshots.
	SnapshotCopyRetentionPeriod() *float64
	SetSnapshotCopyRetentionPeriod(val *float64)
	// The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
	//
	// Example: `my-snapshot-id`.
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of tag instances.
	Tags() awscdk.TagManager
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
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	//
	// Default: The default VPC security group is associated with the cluster.
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) AllowVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AquaConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrDeferMaintenanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDeferMaintenanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AutomatedSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AvailabilityZoneRelocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AvailabilityZoneRelocationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRelocationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Classic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"classic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) DeferMaintenance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deferMaintenance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) DeferMaintenanceDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deferMaintenanceDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) DeferMaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) DeferMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deferMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) DestinationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EnhancedVpcRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) HsmClientCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmClientCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) HsmConfigurationIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmConfigurationIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LoggingProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) MaintenanceTrackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ManualSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) OwnerAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ResourceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) RevisionTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) RotateEncryptionKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rotateEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SnapshotClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SnapshotCopyGrantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotCopyGrantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SnapshotCopyManual() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotCopyManual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SnapshotCopyRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotCopyRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::Cluster`.
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetAllowVersionUpgrade(val interface{}) {
	_jsii_.Set(
		j,
		"allowVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAquaConfigurationStatus(val *string) {
	_jsii_.Set(
		j,
		"aquaConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAutomatedSnapshotRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"automatedSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAvailabilityZoneRelocation(val interface{}) {
	_jsii_.Set(
		j,
		"availabilityZoneRelocation",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAvailabilityZoneRelocationStatus(val *string) {
	_jsii_.Set(
		j,
		"availabilityZoneRelocationStatus",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClassic(val interface{}) {
	_jsii_.Set(
		j,
		"classic",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"clusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"clusterSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"clusterSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterType(val *string) {
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetClusterVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDbName(val *string) {
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDeferMaintenance(val interface{}) {
	_jsii_.Set(
		j,
		"deferMaintenance",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDeferMaintenanceDuration(val *float64) {
	_jsii_.Set(
		j,
		"deferMaintenanceDuration",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDeferMaintenanceEndTime(val *string) {
	_jsii_.Set(
		j,
		"deferMaintenanceEndTime",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDeferMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"deferMaintenanceStartTime",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetDestinationRegion(val *string) {
	_jsii_.Set(
		j,
		"destinationRegion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetElasticIp(val *string) {
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEncrypted(val interface{}) {
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEnhancedVpcRouting(val interface{}) {
	_jsii_.Set(
		j,
		"enhancedVpcRouting",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetHsmClientCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"hsmClientCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetHsmConfigurationIdentifier(val *string) {
	_jsii_.Set(
		j,
		"hsmConfigurationIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetIamRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLoggingProperties(val interface{}) {
	_jsii_.Set(
		j,
		"loggingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetMaintenanceTrackName(val *string) {
	_jsii_.Set(
		j,
		"maintenanceTrackName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetManualSnapshotRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"manualSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetMasterUsername(val *string) {
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetMasterUserPassword(val *string) {
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetNodeType(val *string) {
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetNumberOfNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetOwnerAccount(val *string) {
	_jsii_.Set(
		j,
		"ownerAccount",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetPubliclyAccessible(val interface{}) {
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetResourceAction(val *string) {
	_jsii_.Set(
		j,
		"resourceAction",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetRevisionTarget(val *string) {
	_jsii_.Set(
		j,
		"revisionTarget",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetRotateEncryptionKey(val interface{}) {
	_jsii_.Set(
		j,
		"rotateEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSnapshotClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSnapshotCopyGrantName(val *string) {
	_jsii_.Set(
		j,
		"snapshotCopyGrantName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSnapshotCopyManual(val interface{}) {
	_jsii_.Set(
		j,
		"snapshotCopyManual",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSnapshotCopyRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"snapshotCopyRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetVpcSecurityGroupIds(val *[]*string) {
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes a connection endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	port: jsii.String("port"),
//   }
//
type CfnCluster_EndpointProperty struct {
	// The DNS address of the Cluster.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port that the database engine is listening on.
	Port *string `field:"optional" json:"port" yaml:"port"`
}

// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingPropertiesProperty := &loggingPropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   }
//
type CfnCluster_LoggingPropertiesProperty struct {
	// The name of an existing S3 bucket where the log files are to be stored.
	//
	// Constraints:
	//
	// - Must be in the same region as the cluster
	// - The cluster must have read bucket and put object permissions.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The prefix applied to the log file names.
	//
	// Constraints:
	//
	// - Cannot exceed 512 characters
	// - Cannot contain spaces( ), double quotes ("), single quotes ('), a backslash (\), or control characters. The hexadecimal codes for invalid characters are:
	//
	// - x00 to x20
	// - x22
	// - x27
	// - x5c
	// - x7f or larger.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

// A CloudFormation `AWS::Redshift::ClusterParameterGroup`.
//
// Describes a parameter group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterParameterGroup := awscdk.Aws_redshift.NewCfnClusterParameterGroup(this, jsii.String("MyCfnClusterParameterGroup"), &cfnClusterParameterGroupProps{
//   	description: jsii.String("description"),
//   	parameterGroupFamily: jsii.String("parameterGroupFamily"),
//
//   	// the properties below are optional
//   	parameters: []interface{}{
//   		&parameterProperty{
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnClusterParameterGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The description of the parameter group.
	Description() *string
	SetDescription(val *string)
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
	// The name of the cluster parameter group family that this cluster parameter group is compatible with.
	ParameterGroupFamily() *string
	SetParameterGroupFamily(val *string)
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// For each parameter to be modified, you must supply at least the parameter name and parameter value; other name-value pairs of the parameter are optional.
	//
	// For the workload management (WLM) configuration, you must supply all the name-value pairs in the wlm_json_configuration parameter.
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
	// The list of tags for the cluster parameter group.
	Tags() awscdk.TagManager
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnClusterParameterGroup
type jsiiProxy_CfnClusterParameterGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClusterParameterGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterParameterGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::ClusterParameterGroup`.
func NewCfnClusterParameterGroup(scope constructs.Construct, id *string, props *CfnClusterParameterGroupProps) CfnClusterParameterGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterParameterGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterParameterGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::ClusterParameterGroup`.
func NewCfnClusterParameterGroup_Override(c CfnClusterParameterGroup, scope constructs.Construct, id *string, props *CfnClusterParameterGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterParameterGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClusterParameterGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnClusterParameterGroup) SetParameterGroupFamily(val *string) {
	_jsii_.Set(
		j,
		"parameterGroupFamily",
		val,
	)
}

func (j *jsiiProxy_CfnClusterParameterGroup) SetParameters(val interface{}) {
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
func CfnClusterParameterGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterParameterGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClusterParameterGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterParameterGroup",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnClusterParameterGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterParameterGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterParameterGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnClusterParameterGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterParameterGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterParameterGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterParameterGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClusterParameterGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterParameterGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterParameterGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterParameterGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes a parameter in a cluster parameter group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterProperty := &parameterProperty{
//   	parameterName: jsii.String("parameterName"),
//   	parameterValue: jsii.String("parameterValue"),
//   }
//
type CfnClusterParameterGroup_ParameterProperty struct {
	// The name of the parameter.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The value of the parameter.
	//
	// If `ParameterName` is `wlm_json_configuration` , then the maximum size of `ParameterValue` is 8000 characters.
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

// Properties for defining a `CfnClusterParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterParameterGroupProps := &cfnClusterParameterGroupProps{
//   	description: jsii.String("description"),
//   	parameterGroupFamily: jsii.String("parameterGroupFamily"),
//
//   	// the properties below are optional
//   	parameters: []interface{}{
//   		&parameterProperty{
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterParameterGroupProps struct {
	// The description of the parameter group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the cluster parameter group family that this cluster parameter group is compatible with.
	ParameterGroupFamily *string `field:"required" json:"parameterGroupFamily" yaml:"parameterGroupFamily"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// For each parameter to be modified, you must supply at least the parameter name and parameter value; other name-value pairs of the parameter are optional.
	//
	// For the workload management (WLM) configuration, you must supply all the name-value pairs in the wlm_json_configuration parameter.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The list of tags for the cluster parameter group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	clusterType: jsii.String("clusterType"),
//   	dbName: jsii.String("dbName"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	nodeType: jsii.String("nodeType"),
//
//   	// the properties below are optional
//   	allowVersionUpgrade: jsii.Boolean(false),
//   	aquaConfigurationStatus: jsii.String("aquaConfigurationStatus"),
//   	automatedSnapshotRetentionPeriod: jsii.Number(123),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	availabilityZoneRelocation: jsii.Boolean(false),
//   	availabilityZoneRelocationStatus: jsii.String("availabilityZoneRelocationStatus"),
//   	classic: jsii.Boolean(false),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	clusterParameterGroupName: jsii.String("clusterParameterGroupName"),
//   	clusterSecurityGroups: []*string{
//   		jsii.String("clusterSecurityGroups"),
//   	},
//   	clusterSubnetGroupName: jsii.String("clusterSubnetGroupName"),
//   	clusterVersion: jsii.String("clusterVersion"),
//   	deferMaintenance: jsii.Boolean(false),
//   	deferMaintenanceDuration: jsii.Number(123),
//   	deferMaintenanceEndTime: jsii.String("deferMaintenanceEndTime"),
//   	deferMaintenanceStartTime: jsii.String("deferMaintenanceStartTime"),
//   	destinationRegion: jsii.String("destinationRegion"),
//   	elasticIp: jsii.String("elasticIp"),
//   	encrypted: jsii.Boolean(false),
//   	enhancedVpcRouting: jsii.Boolean(false),
//   	hsmClientCertificateIdentifier: jsii.String("hsmClientCertificateIdentifier"),
//   	hsmConfigurationIdentifier: jsii.String("hsmConfigurationIdentifier"),
//   	iamRoles: []*string{
//   		jsii.String("iamRoles"),
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	loggingProperties: &loggingPropertiesProperty{
//   		bucketName: jsii.String("bucketName"),
//
//   		// the properties below are optional
//   		s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   	},
//   	maintenanceTrackName: jsii.String("maintenanceTrackName"),
//   	manualSnapshotRetentionPeriod: jsii.Number(123),
//   	numberOfNodes: jsii.Number(123),
//   	ownerAccount: jsii.String("ownerAccount"),
//   	port: jsii.Number(123),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	resourceAction: jsii.String("resourceAction"),
//   	revisionTarget: jsii.String("revisionTarget"),
//   	rotateEncryptionKey: jsii.Boolean(false),
//   	snapshotClusterIdentifier: jsii.String("snapshotClusterIdentifier"),
//   	snapshotCopyGrantName: jsii.String("snapshotCopyGrantName"),
//   	snapshotCopyManual: jsii.Boolean(false),
//   	snapshotCopyRetentionPeriod: jsii.Number(123),
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
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
type CfnClusterProps struct {
	// The type of the cluster. When cluster type is specified as.
	//
	// - `single-node` , the *NumberOfNodes* parameter is not required.
	// - `multi-node` , the *NumberOfNodes* parameter is required.
	//
	// Valid Values: `multi-node` | `single-node`
	//
	// Default: `multi-node`.
	ClusterType *string `field:"required" json:"clusterType" yaml:"clusterType"`
	// The name of the first database to be created when the cluster is created.
	//
	// To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database. For more information, go to [Create a Database](https://docs.aws.amazon.com/redshift/latest/dg/t_creating_database.html) in the Amazon Redshift Database Developer Guide.
	//
	// Default: `dev`
	//
	// Constraints:
	//
	// - Must contain 1 to 64 alphanumeric characters.
	// - Must contain only lowercase letters.
	// - Cannot be a word that is reserved by the service. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	DbName *string `field:"required" json:"dbName" yaml:"dbName"`
	// The user name associated with the admin user account for the cluster that is being created.
	//
	// Constraints:
	//
	// - Must be 1 - 128 alphanumeric characters. The user name can't be `PUBLIC` .
	// - First character must be a letter.
	// - Cannot be a reserved word. A list of reserved words can be found in [Reserved Words](https://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html) in the Amazon Redshift Database Developer Guide.
	MasterUsername *string `field:"required" json:"masterUsername" yaml:"masterUsername"`
	// The password associated with the admin user account for the cluster that is being created.
	//
	// Constraints:
	//
	// - Must be between 8 and 64 characters in length.
	// - Must contain at least one uppercase letter.
	// - Must contain at least one lowercase letter.
	// - Must contain one number.
	// - Can be any printable ASCII character (ASCII code 33-126) except ' (single quote), " (double quote), \, /, or @.
	MasterUserPassword *string `field:"required" json:"masterUserPassword" yaml:"masterUserPassword"`
	// The node type to be provisioned for the cluster.
	//
	// For information about node types, go to [Working with Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes) in the *Amazon Redshift Cluster Management Guide* .
	//
	// Valid Values: `ds2.xlarge` | `ds2.8xlarge` | `dc1.large` | `dc1.8xlarge` | `dc2.large` | `dc2.8xlarge` | `ra3.xlplus` | `ra3.4xlarge` | `ra3.16xlarge`
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// If `true` , major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster.
	//
	// When a new major version of the Amazon Redshift engine is released, you can request that the service automatically apply upgrades during the maintenance window to the Amazon Redshift engine that is running on your cluster.
	//
	// Default: `true`.
	AllowVersionUpgrade interface{} `field:"optional" json:"allowVersionUpgrade" yaml:"allowVersionUpgrade"`
	// The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) when it is created.
	//
	// Possible values include the following.
	//
	// - enabled - Use AQUA if it is available for the current AWS Region and Amazon Redshift node type.
	// - disabled - Don't use AQUA.
	// - auto - Amazon Redshift determines whether to use AQUA.
	AquaConfigurationStatus *string `field:"optional" json:"aquaConfigurationStatus" yaml:"aquaConfigurationStatus"`
	// The number of days that automated snapshots are retained.
	//
	// If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want with [CreateClusterSnapshot](https://docs.aws.amazon.com/redshift/latest/APIReference/API_CreateClusterSnapshot.html) in the *Amazon Redshift API Reference* .
	//
	// Default: `1`
	//
	// Constraints: Must be a value from 0 to 35.
	AutomatedSnapshotRetentionPeriod *float64 `field:"optional" json:"automatedSnapshotRetentionPeriod" yaml:"automatedSnapshotRetentionPeriod"`
	// The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster.
	//
	// For example, if you have several EC2 instances running in a specific Availability Zone, then you might want the cluster to be provisioned in the same zone in order to decrease network latency.
	//
	// Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint.
	//
	// Example: `us-east-2d`
	//
	// Constraint: The specified Availability Zone must be in the same region as the current endpoint.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster is created.
	AvailabilityZoneRelocation interface{} `field:"optional" json:"availabilityZoneRelocation" yaml:"availabilityZoneRelocation"`
	// Describes the status of the Availability Zone relocation operation.
	AvailabilityZoneRelocationStatus *string `field:"optional" json:"availabilityZoneRelocationStatus" yaml:"availabilityZoneRelocationStatus"`
	// A boolean value indicating whether the resize operation is using the classic resize process.
	//
	// If you don't provide this parameter or set the value to `false` , the resize type is elastic.
	Classic interface{} `field:"optional" json:"classic" yaml:"classic"`
	// A unique identifier for the cluster.
	//
	// You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. The identifier also appears in the Amazon Redshift console.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 alphanumeric characters or hyphens.
	// - Alphabetic characters must be lowercase.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	// - Must be unique for all clusters within an AWS account .
	//
	// Example: `myexamplecluster`.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the parameter group to be associated with this cluster.
	//
	// Default: The default Amazon Redshift cluster parameter group. For information about the default parameter group, go to [Working with Amazon Redshift Parameter Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
	//
	// Constraints:
	//
	// - Must be 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	ClusterParameterGroupName *string `field:"optional" json:"clusterParameterGroupName" yaml:"clusterParameterGroupName"`
	// A list of security groups to be associated with this cluster.
	//
	// Default: The default cluster security group for Amazon Redshift.
	ClusterSecurityGroups *[]*string `field:"optional" json:"clusterSecurityGroups" yaml:"clusterSecurityGroups"`
	// The name of a cluster subnet group to be associated with this cluster.
	//
	// If this parameter is not provided the resulting cluster will be deployed outside virtual private cloud (VPC).
	ClusterSubnetGroupName *string `field:"optional" json:"clusterSubnetGroupName" yaml:"clusterSubnetGroupName"`
	// The version of the Amazon Redshift engine software that you want to deploy on the cluster.
	//
	// The version selected runs on all the nodes in the cluster.
	//
	// Constraints: Only version 1.0 is currently available.
	//
	// Example: `1.0`
	ClusterVersion *string `field:"optional" json:"clusterVersion" yaml:"clusterVersion"`
	// `AWS::Redshift::Cluster.DeferMaintenance`.
	DeferMaintenance interface{} `field:"optional" json:"deferMaintenance" yaml:"deferMaintenance"`
	// `AWS::Redshift::Cluster.DeferMaintenanceDuration`.
	DeferMaintenanceDuration *float64 `field:"optional" json:"deferMaintenanceDuration" yaml:"deferMaintenanceDuration"`
	// `AWS::Redshift::Cluster.DeferMaintenanceEndTime`.
	DeferMaintenanceEndTime *string `field:"optional" json:"deferMaintenanceEndTime" yaml:"deferMaintenanceEndTime"`
	// `AWS::Redshift::Cluster.DeferMaintenanceStartTime`.
	DeferMaintenanceStartTime *string `field:"optional" json:"deferMaintenanceStartTime" yaml:"deferMaintenanceStartTime"`
	// The destination region that snapshots are automatically copied to when cross-region snapshot copy is enabled.
	DestinationRegion *string `field:"optional" json:"destinationRegion" yaml:"destinationRegion"`
	// The Elastic IP (EIP) address for the cluster.
	//
	// Constraints: The cluster must be provisioned in EC2-VPC and publicly-accessible through an Internet gateway. For more information about provisioning clusters in EC2-VPC, go to [Supported Platforms to Launch Your Cluster](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms) in the Amazon Redshift Cluster Management Guide.
	ElasticIp *string `field:"optional" json:"elasticIp" yaml:"elasticIp"`
	// If `true` , the data in the cluster is encrypted at rest.
	//
	// Default: false.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// An option that specifies whether to create the cluster with enhanced VPC routing enabled.
	//
	// To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see [Enhanced VPC Routing](https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html) in the Amazon Redshift Cluster Management Guide.
	//
	// If this option is `true` , enhanced VPC routing is enabled.
	//
	// Default: false.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.
	HsmClientCertificateIdentifier *string `field:"optional" json:"hsmClientCertificateIdentifier" yaml:"hsmClientCertificateIdentifier"`
	// Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
	HsmConfigurationIdentifier *string `field:"optional" json:"hsmConfigurationIdentifier" yaml:"hsmConfigurationIdentifier"`
	// A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services.
	//
	// You must supply the IAM roles in their Amazon Resource Name (ARN) format.
	//
	// The maximum number of IAM roles that you can associate is subject to a quota. For more information, go to [Quotas and limits](https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html) in the *Amazon Redshift Cluster Management Guide* .
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies logging information, such as queries and connection attempts, for the specified Amazon Redshift cluster.
	LoggingProperties interface{} `field:"optional" json:"loggingProperties" yaml:"loggingProperties"`
	// An optional parameter for the name of the maintenance track for the cluster.
	//
	// If you don't provide a maintenance track name, the cluster is assigned to the `current` track.
	MaintenanceTrackName *string `field:"optional" json:"maintenanceTrackName" yaml:"maintenanceTrackName"`
	// The default number of days to retain a manual snapshot.
	//
	// If the value is -1, the snapshot is retained indefinitely. This setting doesn't change the retention period of existing snapshots.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *float64 `field:"optional" json:"manualSnapshotRetentionPeriod" yaml:"manualSnapshotRetentionPeriod"`
	// The number of compute nodes in the cluster.
	//
	// This parameter is required when the *ClusterType* parameter is specified as `multi-node` .
	//
	// For information about determining how many nodes you need, go to [Working with Clusters](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#how-many-nodes) in the *Amazon Redshift Cluster Management Guide* .
	//
	// If you don't specify this parameter, you get a single-node cluster. When requesting a multi-node cluster, you must specify the number of nodes that you want in the cluster.
	//
	// Default: `1`
	//
	// Constraints: Value must be at least 1 and no more than 100.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// The AWS account used to create or copy the snapshot.
	//
	// Required if you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// The port number on which the cluster accepts incoming connections.
	//
	// The cluster is accessible only via the JDBC and ODBC connection strings. Part of the connection string requires the port on which the cluster will listen for incoming connections.
	//
	// Default: `5439`
	//
	// Valid Values: `1150-65535`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The weekly time range (in UTC) during which automated cluster maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// Default: A 30-minute window selected at random from an 8-hour block of time per region, occurring on a random day of the week. For more information about the time blocks for each region, see [Maintenance Windows](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#rs-maintenance-windows) in Amazon Redshift Cluster Management Guide.
	//
	// Valid Days: Mon | Tue | Wed | Thu | Fri | Sat | Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// If `true` , the cluster can be accessed from a public network.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// `AWS::Redshift::Cluster.ResourceAction`.
	ResourceAction *string `field:"optional" json:"resourceAction" yaml:"resourceAction"`
	// `AWS::Redshift::Cluster.RevisionTarget`.
	RevisionTarget *string `field:"optional" json:"revisionTarget" yaml:"revisionTarget"`
	// `AWS::Redshift::Cluster.RotateEncryptionKey`.
	RotateEncryptionKey interface{} `field:"optional" json:"rotateEncryptionKey" yaml:"rotateEncryptionKey"`
	// The name of the cluster the source snapshot was created from.
	//
	// This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
	SnapshotClusterIdentifier *string `field:"optional" json:"snapshotClusterIdentifier" yaml:"snapshotClusterIdentifier"`
	// The name of the snapshot copy grant.
	SnapshotCopyGrantName *string `field:"optional" json:"snapshotCopyGrantName" yaml:"snapshotCopyGrantName"`
	// Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
	SnapshotCopyManual interface{} `field:"optional" json:"snapshotCopyManual" yaml:"snapshotCopyManual"`
	// The number of days to retain automated snapshots in the destination AWS Region after they are copied from the source AWS Region .
	//
	// By default, this only changes the retention period of copied automated snapshots.
	//
	// If you decrease the retention period for automated snapshots that are copied to a destination AWS Region , Amazon Redshift deletes any existing automated snapshots that were copied to the destination AWS Region and that fall outside of the new retention period.
	//
	// Constraints: Must be at least 1 and no more than 35 for automated snapshots.
	//
	// If you specify the `manual` option, only newly copied manual snapshots will have the new retention period.
	//
	// If you specify the value of -1 newly copied manual snapshots are retained indefinitely.
	//
	// Constraints: The number of days must be either -1 or an integer between 1 and 3,653 for manual snapshots.
	SnapshotCopyRetentionPeriod *float64 `field:"optional" json:"snapshotCopyRetentionPeriod" yaml:"snapshotCopyRetentionPeriod"`
	// The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
	//
	// Example: `my-snapshot-id`.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// A list of tag instances.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	//
	// Default: The default VPC security group is associated with the cluster.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

// A CloudFormation `AWS::Redshift::ClusterSecurityGroup`.
//
// Specifies a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters.
//
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroup := awscdk.Aws_redshift.NewCfnClusterSecurityGroup(this, jsii.String("MyCfnClusterSecurityGroup"), &cfnClusterSecurityGroupProps{
//   	description: jsii.String("description"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnClusterSecurityGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description for the security group.
	Description() *string
	SetDescription(val *string)
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
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this security group.
	//
	// Use tags to manage your resources.
	Tags() awscdk.TagManager
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnClusterSecurityGroup
type jsiiProxy_CfnClusterSecurityGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClusterSecurityGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::ClusterSecurityGroup`.
func NewCfnClusterSecurityGroup(scope constructs.Construct, id *string, props *CfnClusterSecurityGroupProps) CfnClusterSecurityGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterSecurityGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::ClusterSecurityGroup`.
func NewCfnClusterSecurityGroup_Override(c CfnClusterSecurityGroup, scope constructs.Construct, id *string, props *CfnClusterSecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClusterSecurityGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnClusterSecurityGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClusterSecurityGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroup",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnClusterSecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterSecurityGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A CloudFormation `AWS::Redshift::ClusterSecurityGroupIngress`.
//
// Adds an inbound (ingress) rule to an Amazon Redshift security group. Depending on whether the application accessing your cluster is running on the Internet or an Amazon EC2 instance, you can authorize inbound access to either a Classless Interdomain Routing (CIDR)/Internet Protocol (IP) range or to an Amazon EC2 security group. You can add as many as 20 ingress rules to an Amazon Redshift security group.
//
// If you authorize access to an Amazon EC2 security group, specify *EC2SecurityGroupName* and *EC2SecurityGroupOwnerId* . The Amazon EC2 security group and Amazon Redshift cluster must be in the same AWS Region .
//
// If you authorize access to a CIDR/IP address range, specify *CIDRIP* . For an overview of CIDR blocks, see the Wikipedia article on [Classless Inter-Domain Routing](https://docs.aws.amazon.com/http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) .
//
// You must also associate the security group with a cluster so that clients running on these IP addresses or the EC2 instance are authorized to connect to the cluster. For information about managing security groups, go to [Working with Security Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroupIngress := awscdk.Aws_redshift.NewCfnClusterSecurityGroupIngress(this, jsii.String("MyCfnClusterSecurityGroupIngress"), &cfnClusterSecurityGroupIngressProps{
//   	clusterSecurityGroupName: jsii.String("clusterSecurityGroupName"),
//
//   	// the properties below are optional
//   	cidrip: jsii.String("cidrip"),
//   	ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   })
//
type CfnClusterSecurityGroupIngress interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The IP range to be added the Amazon Redshift security group.
	Cidrip() *string
	SetCidrip(val *string)
	// The name of the security group to which the ingress rule is added.
	ClusterSecurityGroupName() *string
	SetClusterSecurityGroupName(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The EC2 security group to be added the Amazon Redshift security group.
	Ec2SecurityGroupName() *string
	SetEc2SecurityGroupName(val *string)
	// The AWS account number of the owner of the security group specified by the *EC2SecurityGroupName* parameter.
	//
	// The AWS Access Key ID is not an acceptable value.
	//
	// Example: `111122223333`
	//
	// Conditional. If you specify the `EC2SecurityGroupName` property, you must specify this property.
	Ec2SecurityGroupOwnerId() *string
	SetEc2SecurityGroupOwnerId(val *string)
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnClusterSecurityGroupIngress
type jsiiProxy_CfnClusterSecurityGroupIngress struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) Cidrip() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) ClusterSecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) Ec2SecurityGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2SecurityGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) Ec2SecurityGroupOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2SecurityGroupOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::ClusterSecurityGroupIngress`.
func NewCfnClusterSecurityGroupIngress(scope constructs.Construct, id *string, props *CfnClusterSecurityGroupIngressProps) CfnClusterSecurityGroupIngress {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterSecurityGroupIngress{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroupIngress",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::ClusterSecurityGroupIngress`.
func NewCfnClusterSecurityGroupIngress_Override(c CfnClusterSecurityGroupIngress, scope constructs.Construct, id *string, props *CfnClusterSecurityGroupIngressProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroupIngress",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) SetCidrip(val *string) {
	_jsii_.Set(
		j,
		"cidrip",
		val,
	)
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) SetClusterSecurityGroupName(val *string) {
	_jsii_.Set(
		j,
		"clusterSecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) SetEc2SecurityGroupName(val *string) {
	_jsii_.Set(
		j,
		"ec2SecurityGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnClusterSecurityGroupIngress) SetEc2SecurityGroupOwnerId(val *string) {
	_jsii_.Set(
		j,
		"ec2SecurityGroupOwnerId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnClusterSecurityGroupIngress_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroupIngress",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClusterSecurityGroupIngress_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroupIngress",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnClusterSecurityGroupIngress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroupIngress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterSecurityGroupIngress_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnClusterSecurityGroupIngress",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSecurityGroupIngress) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnClusterSecurityGroupIngress`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroupIngressProps := &cfnClusterSecurityGroupIngressProps{
//   	clusterSecurityGroupName: jsii.String("clusterSecurityGroupName"),
//
//   	// the properties below are optional
//   	cidrip: jsii.String("cidrip"),
//   	ec2SecurityGroupName: jsii.String("ec2SecurityGroupName"),
//   	ec2SecurityGroupOwnerId: jsii.String("ec2SecurityGroupOwnerId"),
//   }
//
type CfnClusterSecurityGroupIngressProps struct {
	// The name of the security group to which the ingress rule is added.
	ClusterSecurityGroupName *string `field:"required" json:"clusterSecurityGroupName" yaml:"clusterSecurityGroupName"`
	// The IP range to be added the Amazon Redshift security group.
	Cidrip *string `field:"optional" json:"cidrip" yaml:"cidrip"`
	// The EC2 security group to be added the Amazon Redshift security group.
	Ec2SecurityGroupName *string `field:"optional" json:"ec2SecurityGroupName" yaml:"ec2SecurityGroupName"`
	// The AWS account number of the owner of the security group specified by the *EC2SecurityGroupName* parameter.
	//
	// The AWS Access Key ID is not an acceptable value.
	//
	// Example: `111122223333`
	//
	// Conditional. If you specify the `EC2SecurityGroupName` property, you must specify this property.
	Ec2SecurityGroupOwnerId *string `field:"optional" json:"ec2SecurityGroupOwnerId" yaml:"ec2SecurityGroupOwnerId"`
}

// Properties for defining a `CfnClusterSecurityGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSecurityGroupProps := &cfnClusterSecurityGroupProps{
//   	description: jsii.String("description"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterSecurityGroupProps struct {
	// A description for the security group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this security group.
	//
	// Use tags to manage your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Redshift::ClusterSubnetGroup`.
//
// Specifies an Amazon Redshift subnet group. You must provide a list of one or more subnets in your existing Amazon Virtual Private Cloud ( Amazon VPC ) when creating Amazon Redshift subnet group.
//
// For information about subnet groups, go to [Amazon Redshift Cluster Subnet Groups](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html) in the *Amazon Redshift Cluster Management Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSubnetGroup := awscdk.Aws_redshift.NewCfnClusterSubnetGroup(this, jsii.String("MyCfnClusterSubnetGroup"), &cfnClusterSubnetGroupProps{
//   	description: jsii.String("description"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnClusterSubnetGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A description for the subnet group.
	Description() *string
	SetDescription(val *string)
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
	// An array of VPC subnet IDs.
	//
	// A maximum of 20 subnets can be modified in a single request.
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this subnet group.
	//
	// Use tags to manage your resources.
	Tags() awscdk.TagManager
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnClusterSubnetGroup
type jsiiProxy_CfnClusterSubnetGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnClusterSubnetGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterSubnetGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::ClusterSubnetGroup`.
func NewCfnClusterSubnetGroup(scope constructs.Construct, id *string, props *CfnClusterSubnetGroupProps) CfnClusterSubnetGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterSubnetGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterSubnetGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::ClusterSubnetGroup`.
func NewCfnClusterSubnetGroup_Override(c CfnClusterSubnetGroup, scope constructs.Construct, id *string, props *CfnClusterSubnetGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnClusterSubnetGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnClusterSubnetGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnClusterSubnetGroup) SetSubnetIds(val *[]*string) {
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
func CfnClusterSubnetGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSubnetGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnClusterSubnetGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSubnetGroup",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnClusterSubnetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnClusterSubnetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterSubnetGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnClusterSubnetGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterSubnetGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSubnetGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSubnetGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnClusterSubnetGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSubnetGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSubnetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnClusterSubnetGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnClusterSubnetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterSubnetGroupProps := &cfnClusterSubnetGroupProps{
//   	description: jsii.String("description"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterSubnetGroupProps struct {
	// A description for the subnet group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// An array of VPC subnet IDs.
	//
	// A maximum of 20 subnets can be modified in a single request.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Specifies an arbitrary set of tags (key–value pairs) to associate with this subnet group.
	//
	// Use tags to manage your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Redshift::EndpointAccess`.
//
// Creates a Redshift-managed VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAccess := awscdk.Aws_redshift.NewCfnEndpointAccess(this, jsii.String("MyCfnEndpointAccess"), &cfnEndpointAccessProps{
//   	endpointName: jsii.String("endpointName"),
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	resourceOwner: jsii.String("resourceOwner"),
//   	subnetGroupName: jsii.String("subnetGroupName"),
//   })
//
type CfnEndpointAccess interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The DNS address of the endpoint.
	AttrAddress() *string
	// The time (UTC) that the endpoint was created.
	AttrEndpointCreateTime() *string
	// The status of the endpoint.
	AttrEndpointStatus() *string
	// The port number on which the cluster accepts incoming connections.
	AttrPort() *float64
	AttrVpcSecurityGroups() awscdk.IResolvable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The cluster identifier of the cluster associated with the endpoint.
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The name of the endpoint.
	EndpointName() *string
	SetEndpointName(val *string)
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
	// The AWS account ID of the owner of the cluster.
	ResourceOwner() *string
	SetResourceOwner(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName() *string
	SetSubnetGroupName(val *string)
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
	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnEndpointAccess
type jsiiProxy_CfnEndpointAccess struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpointAccess) AttrAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) AttrEndpointCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) AttrEndpointStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEndpointStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) AttrPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) AttrVpcSecurityGroups() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrVpcSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) EndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) ResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAccess) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::EndpointAccess`.
func NewCfnEndpointAccess(scope constructs.Construct, id *string, props *CfnEndpointAccessProps) CfnEndpointAccess {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpointAccess{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnEndpointAccess",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::EndpointAccess`.
func NewCfnEndpointAccess_Override(c CfnEndpointAccess, scope constructs.Construct, id *string, props *CfnEndpointAccessProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnEndpointAccess",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpointAccess) SetClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAccess) SetEndpointName(val *string) {
	_jsii_.Set(
		j,
		"endpointName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAccess) SetResourceOwner(val *string) {
	_jsii_.Set(
		j,
		"resourceOwner",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAccess) SetSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"subnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAccess) SetVpcSecurityGroupIds(val *[]*string) {
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
func CfnEndpointAccess_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEndpointAccess",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEndpointAccess_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEndpointAccess",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnEndpointAccess_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEndpointAccess",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointAccess_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnEndpointAccess",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointAccess) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAccess) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAccess) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpointAccess) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAccess) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAccess) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAccess) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The security groups associated with the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSecurityGroupProperty := &vpcSecurityGroupProperty{
//   	status: jsii.String("status"),
//   	vpcSecurityGroupId: jsii.String("vpcSecurityGroupId"),
//   }
//
type CfnEndpointAccess_VpcSecurityGroupProperty struct {
	// The status of the endpoint.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The identifier of the VPC security group.
	VpcSecurityGroupId *string `field:"optional" json:"vpcSecurityGroupId" yaml:"vpcSecurityGroupId"`
}

// Properties for defining a `CfnEndpointAccess`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAccessProps := &cfnEndpointAccessProps{
//   	endpointName: jsii.String("endpointName"),
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//
//   	// the properties below are optional
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   	resourceOwner: jsii.String("resourceOwner"),
//   	subnetGroupName: jsii.String("subnetGroupName"),
//   }
//
type CfnEndpointAccessProps struct {
	// The name of the endpoint.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	VpcSecurityGroupIds *[]*string `field:"required" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// The cluster identifier of the cluster associated with the endpoint.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The AWS account ID of the owner of the cluster.
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
}

// A CloudFormation `AWS::Redshift::EndpointAuthorization`.
//
// Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAuthorization := awscdk.Aws_redshift.NewCfnEndpointAuthorization(this, jsii.String("MyCfnEndpointAuthorization"), &cfnEndpointAuthorizationProps{
//   	account: jsii.String("account"),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	force: jsii.Boolean(false),
//   	vpcIds: []*string{
//   		jsii.String("vpcIds"),
//   	},
//   })
//
type CfnEndpointAuthorization interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The A AWS account ID of either the cluster owner (grantor) or grantee.
	//
	// If `Grantee` parameter is true, then the `Account` value is of the grantor.
	Account() *string
	SetAccount(val *string)
	// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
	AttrAllowedAllVpCs() awscdk.IResolvable
	// The VPCs allowed access to the cluster.
	AttrAllowedVpCs() *[]*string
	// The time (UTC) when the authorization was created.
	AttrAuthorizeTime() *string
	// The status of the cluster.
	AttrClusterStatus() *string
	// The number of Redshift-managed VPC endpoints created for the authorization.
	AttrEndpointCount() *float64
	// The AWS account ID of the grantee of the cluster.
	AttrGrantee() *string
	// The AWS account ID of the cluster owner.
	AttrGrantor() *string
	// The status of the authorization action.
	AttrStatus() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The cluster identifier.
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Indicates whether to force the revoke action.
	//
	// If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	Force() interface{}
	SetForce(val interface{})
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
	// The virtual private cloud (VPC) identifiers to grant access to.
	VpcIds() *[]*string
	SetVpcIds(val *[]*string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnEndpointAuthorization
type jsiiProxy_CfnEndpointAuthorization struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEndpointAuthorization) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrAllowedAllVpCs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAllowedAllVpCs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrAllowedVpCs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrAllowedVpCs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrAuthorizeTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAuthorizeTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrClusterStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrClusterStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrEndpointCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrEndpointCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrGrantee() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGrantee",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrGrantor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGrantor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) Force() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"force",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEndpointAuthorization) VpcIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcIds",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::EndpointAuthorization`.
func NewCfnEndpointAuthorization(scope constructs.Construct, id *string, props *CfnEndpointAuthorizationProps) CfnEndpointAuthorization {
	_init_.Initialize()

	j := jsiiProxy_CfnEndpointAuthorization{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnEndpointAuthorization",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::EndpointAuthorization`.
func NewCfnEndpointAuthorization_Override(c CfnEndpointAuthorization, scope constructs.Construct, id *string, props *CfnEndpointAuthorizationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnEndpointAuthorization",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEndpointAuthorization) SetAccount(val *string) {
	_jsii_.Set(
		j,
		"account",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAuthorization) SetClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAuthorization) SetForce(val interface{}) {
	_jsii_.Set(
		j,
		"force",
		val,
	)
}

func (j *jsiiProxy_CfnEndpointAuthorization) SetVpcIds(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcIds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEndpointAuthorization_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEndpointAuthorization",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEndpointAuthorization_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEndpointAuthorization",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnEndpointAuthorization_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEndpointAuthorization",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEndpointAuthorization_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnEndpointAuthorization",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorization) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorization) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorization) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEndpointAuthorization) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorization) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorization) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEndpointAuthorization) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEndpointAuthorization`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointAuthorizationProps := &cfnEndpointAuthorizationProps{
//   	account: jsii.String("account"),
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	force: jsii.Boolean(false),
//   	vpcIds: []*string{
//   		jsii.String("vpcIds"),
//   	},
//   }
//
type CfnEndpointAuthorizationProps struct {
	// The A AWS account ID of either the cluster owner (grantor) or grantee.
	//
	// If `Grantee` parameter is true, then the `Account` value is of the grantor.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The cluster identifier.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Indicates whether to force the revoke action.
	//
	// If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// The virtual private cloud (VPC) identifiers to grant access to.
	VpcIds *[]*string `field:"optional" json:"vpcIds" yaml:"vpcIds"`
}

// A CloudFormation `AWS::Redshift::EventSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSubscription := awscdk.Aws_redshift.NewCfnEventSubscription(this, jsii.String("MyCfnEventSubscription"), &cfnEventSubscriptionProps{
//   	subscriptionName: jsii.String("subscriptionName"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	eventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	severity: jsii.String("severity"),
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   	sourceIds: []*string{
//   		jsii.String("sourceIds"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnEventSubscription interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The AWS account associated with the Amazon Redshift event notification subscription.
	AttrCustomerAwsId() *string
	// The name of the Amazon Redshift event notification subscription.
	AttrCustSubscriptionId() *string
	// The list of Amazon Redshift event categories specified in the event notification subscription.
	//
	// Values: Configuration, Management, Monitoring, Security, Pending.
	AttrEventCategoriesList() *[]*string
	// A list of the sources that publish events to the Amazon Redshift event notification subscription.
	AttrSourceIdsList() *[]*string
	// The status of the Amazon Redshift event notification subscription.
	//
	// Constraints:
	//
	// - Can be one of the following: active | no-permission | topic-not-exist
	// - The status "no-permission" indicates that Amazon Redshift no longer has permission to post to the Amazon SNS topic. The status "topic-not-exist" indicates that the topic was deleted after the subscription was created.
	AttrStatus() *string
	// The date and time the Amazon Redshift event notification subscription was created.
	AttrSubscriptionCreationTime() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A boolean value;
	//
	// set to `true` to activate the subscription, and set to `false` to create the subscription but not activate it.
	Enabled() interface{}
	SetEnabled(val interface{})
	// Specifies the Amazon Redshift event categories to be published by the event notification subscription.
	//
	// Values: configuration, management, monitoring, security, pending.
	EventCategories() *[]*string
	SetEventCategories(val *[]*string)
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
	// Specifies the Amazon Redshift event severity to be published by the event notification subscription.
	//
	// Values: ERROR, INFO.
	Severity() *string
	SetSeverity(val *string)
	// The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.
	//
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	// A list of one or more identifiers of Amazon Redshift source objects.
	//
	// All of the objects must be of the same type as was specified in the source type parameter. The event subscription will return only events generated by the specified objects. If not specified, then events are returned for all objects within the source type specified.
	//
	// Example: my-cluster-1, my-cluster-2
	//
	// Example: my-snapshot-20131010.
	SourceIds() *[]*string
	SetSourceIds(val *[]*string)
	// The type of source that will be generating the events.
	//
	// For example, if you want to be notified of events generated by a cluster, you would set this parameter to cluster. If this value is not specified, events are returned for all Amazon Redshift objects in your AWS account . You must specify a source type in order to specify source IDs.
	//
	// Valid values: cluster, cluster-parameter-group, cluster-security-group, cluster-snapshot, and scheduled-action.
	SourceType() *string
	SetSourceType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The name of the event subscription to be created.
	//
	// Constraints:
	//
	// - Cannot be null, empty, or blank.
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	SubscriptionName() *string
	SetSubscriptionName(val *string)
	// A list of tag instances.
	Tags() awscdk.TagManager
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnEventSubscription
type jsiiProxy_CfnEventSubscription struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventSubscription) AttrCustomerAwsId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCustomerAwsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) AttrCustSubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCustSubscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) AttrEventCategoriesList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrEventCategoriesList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) AttrSourceIdsList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrSourceIdsList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) AttrSubscriptionCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSubscriptionCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) EventCategories() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eventCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) SubscriptionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSubscription) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::EventSubscription`.
func NewCfnEventSubscription(scope constructs.Construct, id *string, props *CfnEventSubscriptionProps) CfnEventSubscription {
	_init_.Initialize()

	j := jsiiProxy_CfnEventSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnEventSubscription",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::EventSubscription`.
func NewCfnEventSubscription_Override(c CfnEventSubscription, scope constructs.Construct, id *string, props *CfnEventSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnEventSubscription",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetEventCategories(val *[]*string) {
	_jsii_.Set(
		j,
		"eventCategories",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSeverity(val *string) {
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSnsTopicArn(val *string) {
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSourceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"sourceIds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSourceType(val *string) {
	_jsii_.Set(
		j,
		"sourceType",
		val,
	)
}

func (j *jsiiProxy_CfnEventSubscription) SetSubscriptionName(val *string) {
	_jsii_.Set(
		j,
		"subscriptionName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventSubscription_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEventSubscription",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventSubscription_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEventSubscription",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventSubscription_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnEventSubscription",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventSubscription) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEventSubscription) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEventSubscription) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEventSubscription) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEventSubscription) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventSubscription) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSubscription) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnEventSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventSubscriptionProps := &cfnEventSubscriptionProps{
//   	subscriptionName: jsii.String("subscriptionName"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	eventCategories: []*string{
//   		jsii.String("eventCategories"),
//   	},
//   	severity: jsii.String("severity"),
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   	sourceIds: []*string{
//   		jsii.String("sourceIds"),
//   	},
//   	sourceType: jsii.String("sourceType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEventSubscriptionProps struct {
	// The name of the event subscription to be created.
	//
	// Constraints:
	//
	// - Cannot be null, empty, or blank.
	// - Must contain from 1 to 255 alphanumeric characters or hyphens.
	// - First character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
	// A boolean value;
	//
	// set to `true` to activate the subscription, and set to `false` to create the subscription but not activate it.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specifies the Amazon Redshift event categories to be published by the event notification subscription.
	//
	// Values: configuration, management, monitoring, security, pending.
	EventCategories *[]*string `field:"optional" json:"eventCategories" yaml:"eventCategories"`
	// Specifies the Amazon Redshift event severity to be published by the event notification subscription.
	//
	// Values: ERROR, INFO.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic used to transmit the event notifications.
	//
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// A list of one or more identifiers of Amazon Redshift source objects.
	//
	// All of the objects must be of the same type as was specified in the source type parameter. The event subscription will return only events generated by the specified objects. If not specified, then events are returned for all objects within the source type specified.
	//
	// Example: my-cluster-1, my-cluster-2
	//
	// Example: my-snapshot-20131010.
	SourceIds *[]*string `field:"optional" json:"sourceIds" yaml:"sourceIds"`
	// The type of source that will be generating the events.
	//
	// For example, if you want to be notified of events generated by a cluster, you would set this parameter to cluster. If this value is not specified, events are returned for all Amazon Redshift objects in your AWS account . You must specify a source type in order to specify source IDs.
	//
	// Valid values: cluster, cluster-parameter-group, cluster-security-group, cluster-snapshot, and scheduled-action.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// A list of tag instances.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Redshift::ScheduledAction`.
//
// Creates a scheduled action. A scheduled action contains a schedule and an Amazon Redshift API action. For example, you can create a schedule of when to run the `ResizeCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledAction := awscdk.Aws_redshift.NewCfnScheduledAction(this, jsii.String("MyCfnScheduledAction"), &cfnScheduledActionProps{
//   	scheduledActionName: jsii.String("scheduledActionName"),
//
//   	// the properties below are optional
//   	enable: jsii.Boolean(false),
//   	endTime: jsii.String("endTime"),
//   	iamRole: jsii.String("iamRole"),
//   	schedule: jsii.String("schedule"),
//   	scheduledActionDescription: jsii.String("scheduledActionDescription"),
//   	startTime: jsii.String("startTime"),
//   	targetAction: &scheduledActionTypeProperty{
//   		pauseCluster: &pauseClusterMessageProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		resizeCluster: &resizeClusterMessageProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   			// the properties below are optional
//   			classic: jsii.Boolean(false),
//   			clusterType: jsii.String("clusterType"),
//   			nodeType: jsii.String("nodeType"),
//   			numberOfNodes: jsii.Number(123),
//   		},
//   		resumeCluster: &resumeClusterMessageProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   	},
//   })
//
type CfnScheduledAction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// List of times when the scheduled action will run.
	AttrNextInvocations() *[]*string
	// The state of the scheduled action.
	//
	// For example, `DISABLED` .
	AttrState() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// If true, the schedule is enabled.
	//
	// If false, the scheduled action does not trigger. For more information about `state` of the scheduled action, see `ScheduledAction` .
	Enable() interface{}
	SetEnable(val interface{})
	// The end time in UTC when the schedule is no longer active.
	//
	// After this time, the scheduled action does not trigger.
	EndTime() *string
	SetEndTime(val *string)
	// The IAM role to assume to run the scheduled action.
	//
	// This IAM role must have permission to run the Amazon Redshift API operation in the scheduled action. This IAM role must allow the Amazon Redshift scheduler (Principal scheduler.redshift.amazonaws.com) to assume permissions on your behalf. For more information about the IAM role to use with the Amazon Redshift scheduler, see [Using Identity-Based Policies for Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html) in the *Amazon Redshift Cluster Management Guide* .
	IamRole() *string
	SetIamRole(val *string)
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
	// The schedule for a one-time (at format) or recurring (cron format) scheduled action.
	//
	// Schedule invocations must be separated by at least one hour.
	//
	// Format of at expressions is " `at(yyyy-mm-ddThh:mm:ss)` ". For example, " `at(2016-03-04T17:27:00)` ".
	//
	// Format of cron expressions is " `cron(Minutes Hours Day-of-month Month Day-of-week Year)` ". For example, " `cron(0 10 ? * MON *)` ". For more information, see [Cron Expressions](https://docs.aws.amazon.com//AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch Events User Guide* .
	Schedule() *string
	SetSchedule(val *string)
	// The description of the scheduled action.
	ScheduledActionDescription() *string
	SetScheduledActionDescription(val *string)
	// The name of the scheduled action.
	ScheduledActionName() *string
	SetScheduledActionName(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The start time in UTC when the schedule is active.
	//
	// Before this time, the scheduled action does not trigger.
	StartTime() *string
	SetStartTime(val *string)
	// A JSON format string of the Amazon Redshift API operation with input parameters.
	//
	// " `{\"ResizeCluster\":{\"NodeType\":\"ds2.8xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}` ".
	TargetAction() interface{}
	SetTargetAction(val interface{})
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnScheduledAction
type jsiiProxy_CfnScheduledAction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScheduledAction) AttrNextInvocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrNextInvocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) IamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) ScheduledActionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledActionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) ScheduledActionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledActionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) TargetAction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledAction) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Redshift::ScheduledAction`.
func NewCfnScheduledAction(scope constructs.Construct, id *string, props *CfnScheduledActionProps) CfnScheduledAction {
	_init_.Initialize()

	j := jsiiProxy_CfnScheduledAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnScheduledAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::ScheduledAction`.
func NewCfnScheduledAction_Override(c CfnScheduledAction, scope constructs.Construct, id *string, props *CfnScheduledActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_redshift.CfnScheduledAction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetEnable(val interface{}) {
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetEndTime(val *string) {
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetIamRole(val *string) {
	_jsii_.Set(
		j,
		"iamRole",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetSchedule(val *string) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetScheduledActionDescription(val *string) {
	_jsii_.Set(
		j,
		"scheduledActionDescription",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetScheduledActionName(val *string) {
	_jsii_.Set(
		j,
		"scheduledActionName",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetStartTime(val *string) {
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledAction) SetTargetAction(val interface{}) {
	_jsii_.Set(
		j,
		"targetAction",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScheduledAction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnScheduledAction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScheduledAction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnScheduledAction",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnScheduledAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_redshift.CfnScheduledAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledAction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_redshift.CfnScheduledAction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScheduledAction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScheduledAction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScheduledAction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScheduledAction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScheduledAction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScheduledAction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScheduledAction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScheduledAction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScheduledAction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScheduledAction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledAction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Describes a pause cluster operation.
//
// For example, a scheduled action to run the `PauseCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pauseClusterMessageProperty := &pauseClusterMessageProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
type CfnScheduledAction_PauseClusterMessageProperty struct {
	// The identifier of the cluster to be paused.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

// Describes a resize cluster operation.
//
// For example, a scheduled action to run the `ResizeCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resizeClusterMessageProperty := &resizeClusterMessageProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	classic: jsii.Boolean(false),
//   	clusterType: jsii.String("clusterType"),
//   	nodeType: jsii.String("nodeType"),
//   	numberOfNodes: jsii.Number(123),
//   }
//
type CfnScheduledAction_ResizeClusterMessageProperty struct {
	// The unique identifier for the cluster to resize.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// A boolean value indicating whether the resize operation is using the classic resize process.
	//
	// If you don't provide this parameter or set the value to `false` , the resize type is elastic.
	Classic interface{} `field:"optional" json:"classic" yaml:"classic"`
	// The new cluster type for the specified cluster.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// The new node type for the nodes you are adding.
	//
	// If not specified, the cluster's current node type is used.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// The new number of nodes for the cluster.
	//
	// If not specified, the cluster's current number of nodes is used.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
}

// Describes a resume cluster operation.
//
// For example, a scheduled action to run the `ResumeCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resumeClusterMessageProperty := &resumeClusterMessageProperty{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//   }
//
type CfnScheduledAction_ResumeClusterMessageProperty struct {
	// The identifier of the cluster to be resumed.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

// The action type that specifies an Amazon Redshift API operation that is supported by the Amazon Redshift scheduler.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduledActionTypeProperty := &scheduledActionTypeProperty{
//   	pauseCluster: &pauseClusterMessageProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   	resizeCluster: &resizeClusterMessageProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   		// the properties below are optional
//   		classic: jsii.Boolean(false),
//   		clusterType: jsii.String("clusterType"),
//   		nodeType: jsii.String("nodeType"),
//   		numberOfNodes: jsii.Number(123),
//   	},
//   	resumeCluster: &resumeClusterMessageProperty{
//   		clusterIdentifier: jsii.String("clusterIdentifier"),
//   	},
//   }
//
type CfnScheduledAction_ScheduledActionTypeProperty struct {
	// An action that runs a `PauseCluster` API operation.
	PauseCluster interface{} `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// An action that runs a `ResizeCluster` API operation.
	ResizeCluster interface{} `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// An action that runs a `ResumeCluster` API operation.
	ResumeCluster interface{} `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

// Properties for defining a `CfnScheduledAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScheduledActionProps := &cfnScheduledActionProps{
//   	scheduledActionName: jsii.String("scheduledActionName"),
//
//   	// the properties below are optional
//   	enable: jsii.Boolean(false),
//   	endTime: jsii.String("endTime"),
//   	iamRole: jsii.String("iamRole"),
//   	schedule: jsii.String("schedule"),
//   	scheduledActionDescription: jsii.String("scheduledActionDescription"),
//   	startTime: jsii.String("startTime"),
//   	targetAction: &scheduledActionTypeProperty{
//   		pauseCluster: &pauseClusterMessageProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		resizeCluster: &resizeClusterMessageProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   			// the properties below are optional
//   			classic: jsii.Boolean(false),
//   			clusterType: jsii.String("clusterType"),
//   			nodeType: jsii.String("nodeType"),
//   			numberOfNodes: jsii.Number(123),
//   		},
//   		resumeCluster: &resumeClusterMessageProperty{
//   			clusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   	},
//   }
//
type CfnScheduledActionProps struct {
	// The name of the scheduled action.
	ScheduledActionName *string `field:"required" json:"scheduledActionName" yaml:"scheduledActionName"`
	// If true, the schedule is enabled.
	//
	// If false, the scheduled action does not trigger. For more information about `state` of the scheduled action, see `ScheduledAction` .
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// The end time in UTC when the schedule is no longer active.
	//
	// After this time, the scheduled action does not trigger.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The IAM role to assume to run the scheduled action.
	//
	// This IAM role must have permission to run the Amazon Redshift API operation in the scheduled action. This IAM role must allow the Amazon Redshift scheduler (Principal scheduler.redshift.amazonaws.com) to assume permissions on your behalf. For more information about the IAM role to use with the Amazon Redshift scheduler, see [Using Identity-Based Policies for Amazon Redshift](https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html) in the *Amazon Redshift Cluster Management Guide* .
	IamRole *string `field:"optional" json:"iamRole" yaml:"iamRole"`
	// The schedule for a one-time (at format) or recurring (cron format) scheduled action.
	//
	// Schedule invocations must be separated by at least one hour.
	//
	// Format of at expressions is " `at(yyyy-mm-ddThh:mm:ss)` ". For example, " `at(2016-03-04T17:27:00)` ".
	//
	// Format of cron expressions is " `cron(Minutes Hours Day-of-month Month Day-of-week Year)` ". For example, " `cron(0 10 ? * MON *)` ". For more information, see [Cron Expressions](https://docs.aws.amazon.com//AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions) in the *Amazon CloudWatch Events User Guide* .
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The description of the scheduled action.
	ScheduledActionDescription *string `field:"optional" json:"scheduledActionDescription" yaml:"scheduledActionDescription"`
	// The start time in UTC when the schedule is active.
	//
	// Before this time, the scheduled action does not trigger.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// A JSON format string of the Amazon Redshift API operation with input parameters.
	//
	// " `{\"ResizeCluster\":{\"NodeType\":\"ds2.8xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}` ".
	TargetAction interface{} `field:"optional" json:"targetAction" yaml:"targetAction"`
}

