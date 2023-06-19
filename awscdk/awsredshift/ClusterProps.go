package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Properties for a new database cluster.
//
// Example:
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//   cluster := awscdk.NewCluster(this, jsii.String("Redshift"), &ClusterProps{
//   	MasterUser: &Login{
//   		MasterUsername: jsii.String("admin"),
//   	},
//   	Vpc: Vpc,
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
	// An optional identifier for the cluster.
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Settings for the individual instances that are launched.
	// Experimental.
	ClusterType ClusterType `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Name of a database which is automatically created inside the cluster.
	// Experimental.
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Whether to enable encryption of data at rest in the cluster.
	// Experimental.
	Encrypted *bool `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The KMS key to use for encryption of data at rest.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Bucket to send logs to.
	//
	// Logging information includes queries and connection attempts, for the specified Amazon Redshift cluster.
	// Experimental.
	LoggingBucket awss3.IBucket `field:"optional" json:"loggingBucket" yaml:"loggingBucket"`
	// Prefix used for logging.
	// Experimental.
	LoggingKeyPrefix *string `field:"optional" json:"loggingKeyPrefix" yaml:"loggingKeyPrefix"`
	// The node type to be provisioned for the cluster.
	// Experimental.
	NodeType NodeType `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Number of compute nodes in the cluster. Only specify this property for multi-node clusters.
	//
	// Value must be at least 2 and no more than 100.
	// Experimental.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// Additional parameters to pass to the database engine https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html.
	// Experimental.
	ParameterGroup IClusterParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// What port to listen on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Whether to make cluster publicly accessible.
	// Experimental.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// A list of AWS Identity and Access Management (IAM) role that can be used by the cluster to access other AWS services.
	//
	// Specify a maximum of 10 roles.
	// Experimental.
	Roles *[]awsiam.IRole `field:"optional" json:"roles" yaml:"roles"`
	// Security group.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A cluster subnet group to use with this cluster.
	// Experimental.
	SubnetGroup IClusterSubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

