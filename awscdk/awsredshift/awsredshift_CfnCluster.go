package awsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
	// This parameter is retired.
	//
	// It does not set the AQUA configuration status. Amazon Redshift automatically determines whether to use AQUA (Advanced Query Accelerator).
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
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
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
	// Experimental.
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
	// Constraints: The cluster must be provisioned in EC2-VPC and publicly-accessible through an Internet gateway. Don't specify the Elastic IP address for a publicly accessible cluster with availability zone relocation turned on. For more information about provisioning clusters in EC2-VPC, go to [Supported Platforms to Launch Your Cluster](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms) in the Amazon Redshift Cluster Management Guide.
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
	// Experimental.
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
	// - Can be any printable ASCII character (ASCII code 33-126) except `'` (single quote), `"` (double quote), `\` , `/` , or `@` .
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
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
	// Experimental.
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
	// The name of the snapshot from which to create the new cluster.
	//
	// This parameter isn't case sensitive. You can specify this parameter or `snapshotArn` , but not both.
	//
	// Example: `my-snapshot-id`.
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A list of tag instances.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.
	//
	// Default: The default VPC security group is associated with the cluster.
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

func (j *jsiiProxy_CfnCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
func NewCfnCluster(scope awscdk.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	if err := validateNewCfnClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"monocdk.aws_redshift.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Redshift::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope awscdk.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_redshift.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster)SetAllowVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAquaConfigurationStatus(val *string) {
	_jsii_.Set(
		j,
		"aquaConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAutomatedSnapshotRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"automatedSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAvailabilityZoneRelocation(val interface{}) {
	if err := j.validateSetAvailabilityZoneRelocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRelocation",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetAvailabilityZoneRelocationStatus(val *string) {
	_jsii_.Set(
		j,
		"availabilityZoneRelocationStatus",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClassic(val interface{}) {
	if err := j.validateSetClassicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classic",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterParameterGroupName(val *string) {
	_jsii_.Set(
		j,
		"clusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterSecurityGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"clusterSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterSubnetGroupName(val *string) {
	_jsii_.Set(
		j,
		"clusterSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetClusterVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetDeferMaintenance(val interface{}) {
	if err := j.validateSetDeferMaintenanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deferMaintenance",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetDeferMaintenanceDuration(val *float64) {
	_jsii_.Set(
		j,
		"deferMaintenanceDuration",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetDeferMaintenanceEndTime(val *string) {
	_jsii_.Set(
		j,
		"deferMaintenanceEndTime",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetDeferMaintenanceStartTime(val *string) {
	_jsii_.Set(
		j,
		"deferMaintenanceStartTime",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetDestinationRegion(val *string) {
	_jsii_.Set(
		j,
		"destinationRegion",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetElasticIp(val *string) {
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetEnhancedVpcRouting(val interface{}) {
	if err := j.validateSetEnhancedVpcRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedVpcRouting",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetHsmClientCertificateIdentifier(val *string) {
	_jsii_.Set(
		j,
		"hsmClientCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetHsmConfigurationIdentifier(val *string) {
	_jsii_.Set(
		j,
		"hsmConfigurationIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetIamRoles(val *[]*string) {
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetLoggingProperties(val interface{}) {
	if err := j.validateSetLoggingPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetMaintenanceTrackName(val *string) {
	_jsii_.Set(
		j,
		"maintenanceTrackName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetManualSnapshotRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"manualSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetNumberOfNodes(val *float64) {
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetOwnerAccount(val *string) {
	_jsii_.Set(
		j,
		"ownerAccount",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetPreferredMaintenanceWindow(val *string) {
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetResourceAction(val *string) {
	_jsii_.Set(
		j,
		"resourceAction",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetRevisionTarget(val *string) {
	_jsii_.Set(
		j,
		"revisionTarget",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetRotateEncryptionKey(val interface{}) {
	if err := j.validateSetRotateEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rotateEncryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSnapshotClusterIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSnapshotCopyGrantName(val *string) {
	_jsii_.Set(
		j,
		"snapshotCopyGrantName",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSnapshotCopyManual(val interface{}) {
	if err := j.validateSetSnapshotCopyManualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotCopyManual",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSnapshotCopyRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"snapshotCopyRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetSnapshotIdentifier(val *string) {
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_CfnCluster)SetVpcSecurityGroupIds(val *[]*string) {
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshift.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshift.CfnCluster",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_redshift.CfnCluster",
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
		"monocdk.aws_redshift.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCluster) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

func (c *jsiiProxy_CfnCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

