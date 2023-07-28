package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Options for creating a provisioned instance.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_2_08_1(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R6G, ec2.InstanceSize_XLARGE4),
//   	}),
//   	ServerlessV2MinCapacity: jsii.Number(6.5),
//   	ServerlessV2MaxCapacity: jsii.Number(64),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader1"), &ServerlessV2ClusterInstanceProps{
//   			ScaleWithWriter: jsii.Boolean(true),
//   		}),
//   		rds.ClusterInstance_*ServerlessV2(jsii.String("reader2")),
//   	},
//   	Vpc: Vpc,
//   })
//
type ProvisionedClusterInstanceProps struct {
	// Whether to allow upgrade of major version for the DB instance.
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The identifier for the database instance.
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// Only used for migrating existing clusters from using `instanceProps` to `writer` and `readers`.
	//
	// Example:
	//   // existing cluster
	//   declare const vpc: ec2.Vpc;
	//   const cluster = new rds.DatabaseCluster(this, 'Database', {
	//     engine: rds.DatabaseClusterEngine.auroraMysql({
	//       version: rds.AuroraMysqlEngineVersion.VER_3_03_0,
	//     }),
	//     instances: 2,
	//     instanceProps: {
	//       instanceType: ec2.InstanceType.of(ec2.InstanceClass.BURSTABLE3, ec2.InstanceSize.SMALL),
	//       vpcSubnets: { subnetType: ec2.SubnetType.PUBLIC },
	//       vpc,
	//     },
	//   });
	//
	//   // migration
	//
	//   const instanceProps = {
	//     instanceType: ec2.InstanceType.of(ec2.InstanceClass.BURSTABLE3, ec2.InstanceSize.SMALL),
	//     isFromLegacyInstanceProps: true,
	//   };
	//
	//   const myCluster = new rds.DatabaseCluster(this, 'Database', {
	//     engine: rds.DatabaseClusterEngine.auroraMysql({
	//       version: rds.AuroraMysqlEngineVersion.VER_3_03_0,
	//     }),
	//     vpcSubnets: { subnetType: ec2.SubnetType.PUBLIC },
	//     vpc,
	//     writer: rds.ClusterInstance.provisioned('Instance1', {
	//       instanceType: instanceProps.instanceType,
	//       isFromLegacyInstanceProps: instanceProps.isFromLegacyInstanceProps,
	//     }),
	//     readers: [
	//       rds.ClusterInstance.provisioned('Instance2', {
	//         instanceType: instanceProps.instanceType,
	//         isFromLegacyInstanceProps: instanceProps.isFromLegacyInstanceProps,
	//       }),
	//     ],
	//   });
	//
	IsFromLegacyInstanceProps *bool `field:"optional" json:"isFromLegacyInstanceProps" yaml:"isFromLegacyInstanceProps"`
	// The DB parameter group to associate with the instance.
	//
	// This is only needed if you need to configure different parameter
	// groups for each individual instance, otherwise you should not
	// provide this and just use the cluster parameter group.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBParameterGroup.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The cluster instance type.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The promotion tier of the cluster instance.
	//
	// Can be between 0-15
	//
	// For provisioned instances this just determines the failover priority.
	// If multiple instances have the same priority then one will be picked at random.
	PromotionTier *float64 `field:"optional" json:"promotionTier" yaml:"promotionTier"`
}

