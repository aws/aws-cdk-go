package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Options for creating a serverless v2 instance.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		CaCertificate: rds.CaCertificate_RDS_CA_RDS2048_G1(),
//   	}),
//   	Readers: []iClusterInstance{
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader"), &ServerlessV2ClusterInstanceProps{
//   			CaCertificate: rds.CaCertificate_Of(jsii.String("custom-ca")),
//   		}),
//   	},
//   	Vpc: Vpc,
//   })
//
type ServerlessV2ClusterInstanceProps struct {
	// Whether to allow upgrade of major version for the DB instance.
	// Default: - false.
	//
	AllowMajorVersionUpgrade *bool `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Whether to enable automatic upgrade of minor version for the DB instance.
	// Default: - true.
	//
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The identifier of the CA certificate for this DB cluster's instances.
	//
	// Specifying or updating this property triggers a reboot.
	//
	// For RDS DB engines:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html
	//
	// Default: - RDS will choose a certificate authority.
	//
	CaCertificate CaCertificate `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// Whether to enable Performance Insights for the DB instance.
	// Default: - false, unless ``performanceInsightRetention`` or ``performanceInsightEncryptionKey`` is set.
	//
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The identifier for the database instance.
	// Default: - CloudFormation generated identifier.
	//
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
	// Default: false.
	//
	IsFromLegacyInstanceProps *bool `field:"optional" json:"isFromLegacyInstanceProps" yaml:"isFromLegacyInstanceProps"`
	// The DB parameter group to associate with the instance.
	//
	// This is only needed if you need to configure different parameter
	// groups for each individual instance, otherwise you should not
	// provide this and just use the cluster parameter group.
	// Default: the cluster parameter group is used.
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
	// Indicates whether the DB instance is an internet-facing instance.
	//
	// If not specified,
	// the cluster's vpcSubnets will be used to determine if the instance is internet-facing
	// or not.
	// Default: - `true` if the cluster's `vpcSubnets` is `subnetType: SubnetType.PUBLIC`, `false` otherwise
	//
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Only applicable to reader instances.
	//
	// If this is true then the instance will be placed in promotion tier 1, otherwise
	// it will be placed in promotion tier 2.
	//
	// For serverless v2 instances this means:
	// - true: The serverless v2 reader will scale to match the writer instance (provisioned or serverless)
	// - false: The serverless v2 reader will scale with the read workfload on the instance.
	// Default: false.
	//
	ScaleWithWriter *bool `field:"optional" json:"scaleWithWriter" yaml:"scaleWithWriter"`
}

