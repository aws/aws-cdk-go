package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

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
	// - Can be any printable ASCII character (ASCII code 33-126) except `'` (single quote), `"` (double quote), `\` , `/` , or `@` .
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
	// This parameter is retired.
	//
	// It does not set the AQUA configuration status. Amazon Redshift automatically determines whether to use AQUA (Advanced Query Accelerator).
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
	// Constraints: The cluster must be provisioned in EC2-VPC and publicly-accessible through an Internet gateway. Don't specify the Elastic IP address for a publicly accessible cluster with availability zone relocation turned on. For more information about provisioning clusters in EC2-VPC, go to [Supported Platforms to Launch Your Cluster](https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html#cluster-platforms) in the Amazon Redshift Cluster Management Guide.
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
	// The name of the snapshot from which to create the new cluster.
	//
	// This parameter isn't case sensitive. You can specify this parameter or `snapshotArn` , but not both.
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

