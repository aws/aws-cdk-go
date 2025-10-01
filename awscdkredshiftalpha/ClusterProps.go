package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a new database cluster.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var vpc vpc
//
//
//   defaultRole := iam.NewRole(this, jsii.String("DefaultRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("redshift.amazonaws.com")),
//   })
//
//   awscdkredshiftalpha.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
//   	Roles: []iRole{
//   		defaultRole,
//   	},
//   	DefaultRole: defaultRole,
//   })
//
// Experimental.
type ClusterProps struct {
	// Username and password for the administrative user.
	// Experimental.
	MasterUser *Login `field:"required" json:"masterUser" yaml:"masterUser"`
	// The VPC to place the cluster in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster is created.
	// See: https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-recovery.html
	//
	// Default: - false.
	//
	// Experimental.
	AvailabilityZoneRelocation *bool `field:"optional" json:"availabilityZoneRelocation" yaml:"availabilityZoneRelocation"`
	// If this flag is set, the cluster resizing type will be set to classic.
	//
	// When resizing a cluster, classic resizing will always provision a new cluster and transfer the data there.
	//
	// Classic resize takes more time to complete, but it can be useful in cases where the change in node count or
	// the node type to migrate to doesn't fall within the bounds for elastic resize.
	// See: https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-operations.html#elastic-resize
	//
	// Default: - Elastic resize type.
	//
	// Experimental.
	ClassicResizing *bool `field:"optional" json:"classicResizing" yaml:"classicResizing"`
	// An optional identifier for the cluster.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Settings for the individual instances that are launched.
	// Default: `ClusterType.MULTI_NODE`
	//
	// Experimental.
	ClusterType ClusterType `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Name of a database which is automatically created inside the cluster.
	// Default: - default_db.
	//
	// Experimental.
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// A single AWS Identity and Access Management (IAM) role to be used as the default role for the cluster.
	//
	// The default role must be included in the roles list.
	// Default: - No default role is specified for the cluster.
	//
	// Experimental.
	DefaultRole awsiam.IRole `field:"optional" json:"defaultRole" yaml:"defaultRole"`
	// The Elastic IP (EIP) address for the cluster.
	// See: https://docs.aws.amazon.com/redshift/latest/mgmt/managing-clusters-vpc.html
	//
	// Default: - No Elastic IP.
	//
	// Experimental.
	ElasticIp *string `field:"optional" json:"elasticIp" yaml:"elasticIp"`
	// Whether to enable encryption of data at rest in the cluster.
	// Default: true.
	//
	// Experimental.
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The KMS key to use for encryption of data at rest.
	// Default: - AWS-managed key, if encryption at rest is enabled.
	//
	// Experimental.
	EncryptionKey awskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// If this flag is set, Amazon Redshift forces all COPY and UNLOAD traffic between your cluster and your data repositories through your virtual private cloud (VPC).
	// See: https://docs.aws.amazon.com/redshift/latest/mgmt/enhanced-vpc-routing.html
	//
	// Default: - false.
	//
	// Experimental.
	EnhancedVpcRouting *bool `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// Bucket details for log files to be sent to, including prefix.
	// Default: - No logging bucket is used.
	//
	// Experimental.
	LoggingProperties *LoggingProperties `field:"optional" json:"loggingProperties" yaml:"loggingProperties"`
	// The maintenance track name for the cluster.
	// See: https://docs.aws.amazon.com/redshift/latest/mgmt/managing-cluster-considerations.html#rs-mgmt-maintenance-tracks
	//
	// Default: undefined - Redshift default is current.
	//
	// Experimental.
	MaintenanceTrackName MaintenanceTrackName `field:"optional" json:"maintenanceTrackName" yaml:"maintenanceTrackName"`
	// Indicating whether Amazon Redshift should deploy the cluster in two Availability Zones.
	// Default: - false.
	//
	// Experimental.
	MultiAz *bool `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The node type to be provisioned for the cluster.
	// Default: `NodeType.DC2_LARGE`
	//
	// Experimental.
	NodeType NodeType `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Number of compute nodes in the cluster. Only specify this property for multi-node clusters.
	//
	// Value must be at least 2 and no more than 100.
	// Default: - 2 if `clusterType` is ClusterType.MULTI_NODE, undefined otherwise
	//
	// Experimental.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// Additional parameters to pass to the database engine https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html.
	// Default: - No parameter group.
	//
	// Experimental.
	ParameterGroup IClusterParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// What port to listen on.
	// Default: - The default for the engine is used.
	//
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Default: - 30-minute window selected at random from an 8-hour block of time for
	// each AWS Region, occurring on a random day of the week.
	//
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Whether to make cluster publicly accessible.
	// Default: false.
	//
	// Experimental.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// If this flag is set, the cluster will be rebooted when changes to the cluster's parameter group that require a restart to apply.
	// Default: false.
	//
	// Experimental.
	RebootForParameterChanges *bool `field:"optional" json:"rebootForParameterChanges" yaml:"rebootForParameterChanges"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Default: RemovalPolicy.RETAIN
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The Amazon Redshift operation to be performed.
	// Default: - no operation.
	//
	// Experimental.
	ResourceAction ResourceAction `field:"optional" json:"resourceAction" yaml:"resourceAction"`
	// A list of AWS Identity and Access Management (IAM) role that can be used by the cluster to access other AWS services.
	//
	// The maximum number of roles to attach to a cluster is subject to a quota.
	// Default: - No role is attached to the cluster.
	//
	// Experimental.
	Roles *[]awsiam.IRole `field:"optional" json:"roles" yaml:"roles"`
	// Security group.
	// Default: - a new security group is created.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A cluster subnet group to use with this cluster.
	// Default: - a new subnet group will be created.
	//
	// Experimental.
	SubnetGroup IClusterSubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	// Default: - private subnets.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

