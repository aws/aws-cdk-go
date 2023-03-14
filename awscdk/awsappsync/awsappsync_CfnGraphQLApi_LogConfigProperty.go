package awsappsync


// The `LogConfig` property type specifies the logging configuration when writing GraphQL operations and tracing to Amazon CloudWatch for an AWS AppSync GraphQL API.
//
// `LogConfig` is a property of the [AWS::AppSync::GraphQLApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigProperty := &logConfigProperty{
//   	cloudWatchLogsRoleArn: jsii.String("cloudWatchLogsRoleArn"),
//   	excludeVerboseContent: jsii.Boolean(false),
//   	fieldLogLevel: jsii.String("fieldLogLevel"),
//   }
//
type CfnGraphQLApi_LogConfigProperty struct {
	// The service role that AWS AppSync will assume to publish to Amazon CloudWatch Logs in your account.
	CloudWatchLogsRoleArn *string `field:"optional" json:"cloudWatchLogsRoleArn" yaml:"cloudWatchLogsRoleArn"`
	// Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging level.
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
	// The field logging level. Values can be NONE, ERROR, or ALL.
	//
	// - *NONE* : No field-level logs are captured.
	// - *ERROR* : Logs the following information only for the fields that are in error:
	//
	// - The error section in the server response.
	// - Field-level errors.
	// - The generated request/response functions that got resolved for error fields.
	// - *ALL* : The following information is logged for all fields in the query:
	//
	// - Field-level tracing information.
	// - The generated request/response functions that got resolved for each field.
	FieldLogLevel *string `field:"optional" json:"fieldLogLevel" yaml:"fieldLogLevel"`
}

