package awsrds


// Aurora capacity units (ACUs).
//
// Each ACU is a combination of processing and memory capacity.
//
// Example:
//   var vpc vpc
//
//
//   cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &ServerlessClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA_POSTGRESQL(),
//   	ParameterGroup: rds.ParameterGroup_FromParameterGroupName(this, jsii.String("ParameterGroup"), jsii.String("default.aurora-postgresql10")),
//   	Vpc: Vpc,
//   	Scaling: &ServerlessScalingOptions{
//   		AutoPause: awscdk.Duration_Minutes(jsii.Number(10)),
//   		 // default is to pause after 5 minutes of idle time
//   		MinCapacity: rds.AuroraCapacityUnit_ACU_8,
//   		 // default is 2 Aurora capacity units (ACUs)
//   		MaxCapacity: rds.AuroraCapacityUnit_ACU_32,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.architecture
//
// Experimental.
type AuroraCapacityUnit string

const (
	// 1 Aurora Capacity Unit.
	// Experimental.
	AuroraCapacityUnit_ACU_1 AuroraCapacityUnit = "ACU_1"
	// 2 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_2 AuroraCapacityUnit = "ACU_2"
	// 4 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_4 AuroraCapacityUnit = "ACU_4"
	// 8 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_8 AuroraCapacityUnit = "ACU_8"
	// 16 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_16 AuroraCapacityUnit = "ACU_16"
	// 32 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_32 AuroraCapacityUnit = "ACU_32"
	// 64 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_64 AuroraCapacityUnit = "ACU_64"
	// 128 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_128 AuroraCapacityUnit = "ACU_128"
	// 192 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_192 AuroraCapacityUnit = "ACU_192"
	// 256 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_256 AuroraCapacityUnit = "ACU_256"
	// 384 Aurora Capacity Units.
	// Experimental.
	AuroraCapacityUnit_ACU_384 AuroraCapacityUnit = "ACU_384"
)

