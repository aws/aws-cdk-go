package awslogs


// How long, in days, the log contents will be retained.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   apiKeyProvider := &AppSyncAuthProvider{
//   	AuthorizationType: appsync.AppSyncAuthorizationType_API_KEY,
//   }
//
//   api := appsync.NewEventApi(this, jsii.String("api"), &EventApiProps{
//   	ApiName: jsii.String("Api"),
//   	OwnerContact: jsii.String("OwnerContact"),
//   	AuthorizationConfig: &EventApiAuthConfig{
//   		AuthProviders: []appSyncAuthProvider{
//   			apiKeyProvider,
//   		},
//   		ConnectionAuthModeTypes: []appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultPublishAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   		DefaultSubscribeAuthModeTypes: []*appSyncAuthorizationType{
//   			appsync.*appSyncAuthorizationType_API_KEY,
//   		},
//   	},
//   	LogConfig: &AppSyncLogConfig{
//   		FieldLogLevel: appsync.AppSyncFieldLogLevel_INFO,
//   		Retention: logs.RetentionDays_ONE_WEEK,
//   	},
//   })
//
//   api.AddChannelNamespace(jsii.String("default"))
//
type RetentionDays string

const (
	// 1 day.
	RetentionDays_ONE_DAY RetentionDays = "ONE_DAY"
	// 3 days.
	RetentionDays_THREE_DAYS RetentionDays = "THREE_DAYS"
	// 5 days.
	RetentionDays_FIVE_DAYS RetentionDays = "FIVE_DAYS"
	// 1 week.
	RetentionDays_ONE_WEEK RetentionDays = "ONE_WEEK"
	// 2 weeks.
	RetentionDays_TWO_WEEKS RetentionDays = "TWO_WEEKS"
	// 1 month.
	RetentionDays_ONE_MONTH RetentionDays = "ONE_MONTH"
	// 2 months.
	RetentionDays_TWO_MONTHS RetentionDays = "TWO_MONTHS"
	// 3 months.
	RetentionDays_THREE_MONTHS RetentionDays = "THREE_MONTHS"
	// 4 months.
	RetentionDays_FOUR_MONTHS RetentionDays = "FOUR_MONTHS"
	// 5 months.
	RetentionDays_FIVE_MONTHS RetentionDays = "FIVE_MONTHS"
	// 6 months.
	RetentionDays_SIX_MONTHS RetentionDays = "SIX_MONTHS"
	// 1 year.
	RetentionDays_ONE_YEAR RetentionDays = "ONE_YEAR"
	// 13 months.
	RetentionDays_THIRTEEN_MONTHS RetentionDays = "THIRTEEN_MONTHS"
	// 18 months.
	RetentionDays_EIGHTEEN_MONTHS RetentionDays = "EIGHTEEN_MONTHS"
	// 2 years.
	RetentionDays_TWO_YEARS RetentionDays = "TWO_YEARS"
	// 3 years.
	RetentionDays_THREE_YEARS RetentionDays = "THREE_YEARS"
	// 5 years.
	RetentionDays_FIVE_YEARS RetentionDays = "FIVE_YEARS"
	// 6 years.
	RetentionDays_SIX_YEARS RetentionDays = "SIX_YEARS"
	// 7 years.
	RetentionDays_SEVEN_YEARS RetentionDays = "SEVEN_YEARS"
	// 8 years.
	RetentionDays_EIGHT_YEARS RetentionDays = "EIGHT_YEARS"
	// 9 years.
	RetentionDays_NINE_YEARS RetentionDays = "NINE_YEARS"
	// 10 years.
	RetentionDays_TEN_YEARS RetentionDays = "TEN_YEARS"
	// Retain logs forever.
	RetentionDays_INFINITE RetentionDays = "INFINITE"
)

