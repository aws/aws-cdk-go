package awsrds


// The database insights mode.
//
// Example:
//   var vpc vpc
//
//   rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AURORA(),
//   	Vpc: vpc,
//   	// If you enable the advanced mode of Database Insights,
//   	// Performance Insights is enabled and you must set the `performanceInsightRetention` to 465(15 months).
//   	DatabaseInsightsMode: rds.DatabaseInsightsMode_ADVANCED,
//   	PerformanceInsightRetention: rds.PerformanceInsightRetention_MONTHS_15,
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("Writer"), &ProvisionedClusterInstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R7G, ec2.InstanceSize_LARGE),
//   	}),
//   })
//
type DatabaseInsightsMode string

const (
	// Standard mode.
	DatabaseInsightsMode_STANDARD DatabaseInsightsMode = "STANDARD"
	// Advanced mode.
	DatabaseInsightsMode_ADVANCED DatabaseInsightsMode = "ADVANCED"
)

