package awscdkneptunealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for a new database cluster.
//
// Example:
//   cluster := neptune.NewDatabaseCluster(this, jsii.String("ServerlessDatabase"), &DatabaseClusterProps{
//   	Vpc: Vpc,
//   	InstanceType: neptune.InstanceType_SERVERLESS(),
//   	ServerlessScalingConfiguration: &ServerlessScalingConfiguration{
//   		MinCapacity: jsii.Number(1),
//   		MaxCapacity: jsii.Number(5),
//   	},
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
	// Default: - No role is attached to the cluster.
	//
	// Experimental.
	AssociatedRoles *[]awsiam.IRole `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// If set to true, Neptune will automatically update the engine of the entire cluster to the latest minor version after a stabilization window of 2 to 3 weeks.
	// Default: - false.
	//
	// Experimental.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// How many days to retain the backup.
	// Default: - cdk.Duration.days(1)
	//
	// Experimental.
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// See: https://docs.aws.amazon.com/neptune/latest/userguide/auditing.html#auditing-enable
	//
	// Default: - no log exports.
	//
	// Experimental.
	CloudwatchLogsExports *[]LogType `field:"optional" json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Default: - logs never expire.
	//
	// Experimental.
	CloudwatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Default: - a new role is created.
	//
	// Experimental.
	CloudwatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Additional parameters to pass to the database engine.
	// Default: - No parameter group.
	//
	// Experimental.
	ClusterParameterGroup IClusterParameterGroup `field:"optional" json:"clusterParameterGroup" yaml:"clusterParameterGroup"`
	// An optional identifier for the cluster.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	DbClusterName *string `field:"optional" json:"dbClusterName" yaml:"dbClusterName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Default: - true if ``removalPolicy`` is RETAIN, false otherwise.
	//
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// What version of the database to start.
	// Default: - The default engine version.
	//
	// Experimental.
	EngineVersion EngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Map AWS Identity and Access Management (IAM) accounts to database accounts.
	// Default: - `false`.
	//
	// Experimental.
	IamAuthentication *bool `field:"optional" json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Default: - `dbClusterName` is used with the word "Instance" appended. If `dbClusterName` is not provided, the
	// identifier is automatically generated.
	//
	// Experimental.
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Number of Neptune compute instances.
	// Default: 1.
	//
	// Experimental.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The KMS key for storage encryption.
	// Default: - default master key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	// Default: no parameter group.
	//
	// Experimental.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// A daily time range in 24-hours UTC format in which backups preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: '01:00-02:00'.
	// Default: - a 30-minute window selected at random from an 8-hour block of
	// time for each AWS Region. To see the time blocks available, see
	//
	// Experimental.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'.
	// Default: - 30-minute window selected at random from an 8-hour block of time for
	// each AWS Region, occurring on a random day of the week.
	//
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	// Default: - Retain cluster.
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	// Default: a new security group is created.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specify minimum and maximum NCUs capacity for a serverless cluster.
	//
	// See https://docs.aws.amazon.com/neptune/latest/userguide/neptune-serverless-capacity-scaling.html
	// Default: - required if instanceType is db.serverless
	//
	// Experimental.
	ServerlessScalingConfiguration *ServerlessScalingConfiguration `field:"optional" json:"serverlessScalingConfiguration" yaml:"serverlessScalingConfiguration"`
	// Whether to enable storage encryption.
	// Default: true.
	//
	// Experimental.
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Existing subnet group for the cluster.
	// Default: - a new subnet group will be created.
	//
	// Experimental.
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// Where to place the instances within the VPC.
	// Default: private subnets.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

