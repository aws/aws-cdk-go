package awsappsync


// log-level for fields in AppSync.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	AuthorizationConfig: &AuthorizationConfig{
//   	},
//   	Name: jsii.String("myApi"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("myApi.graphql"))),
//   	LogConfig: &LogConfig{
//   		FieldLogLevel: appsync.FieldLogLevel_INFO,
//   		Retention: logs.RetentionDays_ONE_WEEK,
//   	},
//   })
//
type FieldLogLevel string

const (
	// Resolver logging is disabled.
	FieldLogLevel_NONE FieldLogLevel = "NONE"
	// Only Error messages appear in logs.
	FieldLogLevel_ERROR FieldLogLevel = "ERROR"
	// Info and Error messages appear in logs.
	FieldLogLevel_INFO FieldLogLevel = "INFO"
	// Debug, Info, and Error messages, appear in logs.
	FieldLogLevel_DEBUG FieldLogLevel = "DEBUG"
	// All messages (Debug, Error, Info, and Trace) appear in logs.
	FieldLogLevel_ALL FieldLogLevel = "ALL"
)

