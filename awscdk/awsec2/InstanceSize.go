package awsec2


// What size of instance to use.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
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
type InstanceSize string

const (
	// Instance size NANO (nano).
	InstanceSize_NANO InstanceSize = "NANO"
	// Instance size MICRO (micro).
	InstanceSize_MICRO InstanceSize = "MICRO"
	// Instance size SMALL (small).
	InstanceSize_SMALL InstanceSize = "SMALL"
	// Instance size MEDIUM (medium).
	InstanceSize_MEDIUM InstanceSize = "MEDIUM"
	// Instance size LARGE (large).
	InstanceSize_LARGE InstanceSize = "LARGE"
	// Instance size XLARGE (xlarge).
	InstanceSize_XLARGE InstanceSize = "XLARGE"
	// Instance size XLARGE2 (2xlarge).
	InstanceSize_XLARGE2 InstanceSize = "XLARGE2"
	// Instance size XLARGE3 (3xlarge).
	InstanceSize_XLARGE3 InstanceSize = "XLARGE3"
	// Instance size XLARGE4 (4xlarge).
	InstanceSize_XLARGE4 InstanceSize = "XLARGE4"
	// Instance size XLARGE6 (6xlarge).
	InstanceSize_XLARGE6 InstanceSize = "XLARGE6"
	// Instance size XLARGE8 (8xlarge).
	InstanceSize_XLARGE8 InstanceSize = "XLARGE8"
	// Instance size XLARGE9 (9xlarge).
	InstanceSize_XLARGE9 InstanceSize = "XLARGE9"
	// Instance size XLARGE10 (10xlarge).
	InstanceSize_XLARGE10 InstanceSize = "XLARGE10"
	// Instance size XLARGE12 (12xlarge).
	InstanceSize_XLARGE12 InstanceSize = "XLARGE12"
	// Instance size XLARGE16 (16xlarge).
	InstanceSize_XLARGE16 InstanceSize = "XLARGE16"
	// Instance size XLARGE18 (18xlarge).
	InstanceSize_XLARGE18 InstanceSize = "XLARGE18"
	// Instance size XLARGE24 (24xlarge).
	InstanceSize_XLARGE24 InstanceSize = "XLARGE24"
	// Instance size XLARGE32 (32xlarge).
	InstanceSize_XLARGE32 InstanceSize = "XLARGE32"
	// Instance size XLARGE48 (48xlarge).
	InstanceSize_XLARGE48 InstanceSize = "XLARGE48"
	// Instance size XLARGE56 (56xlarge).
	InstanceSize_XLARGE56 InstanceSize = "XLARGE56"
	// Instance size XLARGE96 (96xlarge).
	InstanceSize_XLARGE96 InstanceSize = "XLARGE96"
	// Instance size XLARGE112 (112xlarge).
	InstanceSize_XLARGE112 InstanceSize = "XLARGE112"
	// Instance size METAL (metal).
	InstanceSize_METAL InstanceSize = "METAL"
	// Instance size XLARGE16METAL (metal-16xl).
	InstanceSize_XLARGE16METAL InstanceSize = "XLARGE16METAL"
	// Instance size XLARGE24METAL (metal-24xl).
	InstanceSize_XLARGE24METAL InstanceSize = "XLARGE24METAL"
	// Instance size XLARGE32METAL (metal-32xl).
	InstanceSize_XLARGE32METAL InstanceSize = "XLARGE32METAL"
	// Instance size XLARGE48METAL (metal-48xl).
	InstanceSize_XLARGE48METAL InstanceSize = "XLARGE48METAL"
)

