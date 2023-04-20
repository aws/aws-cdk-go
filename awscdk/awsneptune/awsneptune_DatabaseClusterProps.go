package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for a new database cluster.
//
// Example:
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Vpc: Vpc,
//   	InstanceType: neptune.InstanceType_R5_LARGE(),
//   	Instances: jsii.Number(2),
//   })
//
// Experimental.
type DatabaseClusterProps struct {
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// What subnets to run the Neptune instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// A list of AWS Identity and Access Management (IAM) role that can be used by the cluster to access other AWS services.
	// Experimental.
	AssociatedRoles *[]awsiam.IRole `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// If set to true, Neptune will automatically update the engine of the entire cluster to the latest minor version after a stabilization window of 2 to 3 weeks.
	// Experimental.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// How many days to retain the backup.
	// Experimental.
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// Additional parameters to pass to the database engine.
	// Experimental.
	ClusterParameterGroup IClusterParameterGroup `field:"optional" json:"clusterParameterGroup" yaml:"clusterParameterGroup"`
	// An optional identifier for the cluster.
	// Experimental.
	DbClusterName *string `field:"optional" json:"dbClusterName" yaml:"dbClusterName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// What version of the database to start.
	// Experimental.
	EngineVersion EngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Map AWS Identity and Access Management (IAM) accounts to database accounts.
	// Experimental.
	IamAuthentication *bool `field:"optional" json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Number of Neptune compute instances.
	// Experimental.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The KMS key for storage encryption.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The port the Neptune cluster will listen on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'.
	// Experimental.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Whether to enable storage encryption.
	// Experimental.
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Existing subnet group for the cluster.
	// Experimental.
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

