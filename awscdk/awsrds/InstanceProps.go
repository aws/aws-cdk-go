package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Instance properties for database instances.
//
// Example:
//   cluster := rds.NewDatabaseCluster(stack, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA(),
//   	InstanceProps: &InstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   		Vpc: *Vpc,
//   	},
//   })
//
//   cluster.addRotationSingleUser()
//
//   clusterWithCustomRotationOptions := rds.NewDatabaseCluster(stack, jsii.String("CustomRotationOptions"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA(),
//   	InstanceProps: &InstanceProps{
//   		InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_BURSTABLE3, ec2.InstanceSize_SMALL),
//   		Vpc: *Vpc,
//   	},
//   })
//   clusterWithCustomRotationOptions.addRotationSingleUser(&RotationSingleUserOptions{
//   	AutomaticallyAfter: cdk.Duration_Days(jsii.Number(7)),
//   	ExcludeCharacters: jsii.String("!@#$%^&*"),
//   	SecurityGroup: SecurityGroup,
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	Endpoint: endpoint,
//   })
//
type InstanceProps struct {
	// What subnets to run the RDS instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Whether to allow upgrade of major version for the DB instance.
	// Default: - false.
	//
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	// Default: - true.
	//
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Whether to remove automated backups immediately after the DB instance is deleted for the DB instance.
	// Default: - true.
	//
	DeleteAutomatedBackups *bool `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Whether to enable Performance Insights for the DB instance.
	// Default: - false, unless ``performanceInsightRetention`` or ``performanceInsightEncryptionKey`` is set.
	//
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// What type of instance to start for the replicas.
	// Default: - t3.medium (or, more precisely, db.t3.medium)
	//
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The DB parameter group to associate with the instance.
	// Default: no parameter group.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	// Default: - None.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Default: - default master key.
	//
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	// Default: 7.
	//
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Default: - 30-minute window selected at random from an 8-hour block of time for
	// each AWS Region, occurring on a random day of the week.
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Indicates whether the DB instance is an internet-facing instance.
	// Default: - `true` if `vpcSubnets` is `subnetType: SubnetType.PUBLIC`, `false` otherwise
	//
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Security group.
	// Default: a new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Where to place the instances within the VPC.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

