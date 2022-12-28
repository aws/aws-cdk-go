package awslogs


// How long, in days, the log contents will be retained.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var myLogsPublishingRole role
//   var vpc vpc
//
//
//   // Exporting logs from a cluster
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	engine: rds.databaseClusterEngine.aurora(&auroraClusterEngineProps{
//   		version: rds.auroraEngineVersion_VER_1_17_9(),
//   	}),
//   	instanceProps: &instanceProps{
//   		vpc: vpc,
//   	},
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("error"),
//   		jsii.String("general"),
//   		jsii.String("slowquery"),
//   		jsii.String("audit"),
//   	},
//   	 // Export all available MySQL-based logs
//   	cloudwatchLogsRetention: logs.retentionDays_THREE_MONTHS,
//   	 // Optional - default is to never expire logs
//   	cloudwatchLogsRetentionRole: myLogsPublishingRole,
//   })
//
//   // Exporting logs from an instance
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   		version: rds.postgresEngineVersion_VER_12_3(),
//   	}),
//   	vpc: vpc,
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("postgresql"),
//   	},
//   })
//
// Experimental.
type RetentionDays string

const (
	// 1 day.
	// Experimental.
	RetentionDays_ONE_DAY RetentionDays = "ONE_DAY"
	// 3 days.
	// Experimental.
	RetentionDays_THREE_DAYS RetentionDays = "THREE_DAYS"
	// 5 days.
	// Experimental.
	RetentionDays_FIVE_DAYS RetentionDays = "FIVE_DAYS"
	// 1 week.
	// Experimental.
	RetentionDays_ONE_WEEK RetentionDays = "ONE_WEEK"
	// 2 weeks.
	// Experimental.
	RetentionDays_TWO_WEEKS RetentionDays = "TWO_WEEKS"
	// 1 month.
	// Experimental.
	RetentionDays_ONE_MONTH RetentionDays = "ONE_MONTH"
	// 2 months.
	// Experimental.
	RetentionDays_TWO_MONTHS RetentionDays = "TWO_MONTHS"
	// 3 months.
	// Experimental.
	RetentionDays_THREE_MONTHS RetentionDays = "THREE_MONTHS"
	// 4 months.
	// Experimental.
	RetentionDays_FOUR_MONTHS RetentionDays = "FOUR_MONTHS"
	// 5 months.
	// Experimental.
	RetentionDays_FIVE_MONTHS RetentionDays = "FIVE_MONTHS"
	// 6 months.
	// Experimental.
	RetentionDays_SIX_MONTHS RetentionDays = "SIX_MONTHS"
	// 1 year.
	// Experimental.
	RetentionDays_ONE_YEAR RetentionDays = "ONE_YEAR"
	// 13 months.
	// Experimental.
	RetentionDays_THIRTEEN_MONTHS RetentionDays = "THIRTEEN_MONTHS"
	// 18 months.
	// Experimental.
	RetentionDays_EIGHTEEN_MONTHS RetentionDays = "EIGHTEEN_MONTHS"
	// 2 years.
	// Experimental.
	RetentionDays_TWO_YEARS RetentionDays = "TWO_YEARS"
	// 5 years.
	// Experimental.
	RetentionDays_FIVE_YEARS RetentionDays = "FIVE_YEARS"
	// 6 years.
	// Experimental.
	RetentionDays_SIX_YEARS RetentionDays = "SIX_YEARS"
	// 7 years.
	// Experimental.
	RetentionDays_SEVEN_YEARS RetentionDays = "SEVEN_YEARS"
	// 8 years.
	// Experimental.
	RetentionDays_EIGHT_YEARS RetentionDays = "EIGHT_YEARS"
	// 9 years.
	// Experimental.
	RetentionDays_NINE_YEARS RetentionDays = "NINE_YEARS"
	// 10 years.
	// Experimental.
	RetentionDays_TEN_YEARS RetentionDays = "TEN_YEARS"
	// Retain logs forever.
	// Experimental.
	RetentionDays_INFINITE RetentionDays = "INFINITE"
)

